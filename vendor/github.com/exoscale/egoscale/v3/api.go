package v3

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httputil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UUID string

func (u UUID) String() string {
	return string(u)
}

func ParseUUID(s string) (UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return "", err
	}

	return UUID(id.String()), nil
}

// Wait is a helper that waits for async operation to reach the final state.
// Final states are one of: failure, success, timeout.
// If states argument are given, returns an error if the final state not match on of those.
func (c Client) Wait(ctx context.Context, op *Operation, states ...OperationState) (*Operation, error) {
	const abortErrorsCount = 5

	if op == nil {
		return nil, fmt.Errorf("operation is nil")
	}

	startTime := time.Now()

	ticker := time.NewTicker(pollInterval(0))
	defer ticker.Stop()

	if op.State != OperationStatePending {
		return op, nil
	}

	var subsequentErrors int
	var operation *Operation
polling:
	for {
		select {
		case <-ticker.C:
			runTime := time.Since(startTime)

			if c.waitTimeout != 0 && runTime > c.waitTimeout {
				return nil, fmt.Errorf("operation: %q: max wait timeout reached", op.ID)
			}

			newInterval := pollInterval(runTime)
			ticker.Reset(newInterval)

			o, err := c.GetOperation(ctx, op.ID)
			if err != nil {
				subsequentErrors++
				if subsequentErrors >= abortErrorsCount {
					return nil, err
				}
				continue
			}
			subsequentErrors = 0

			if o.State == OperationStatePending {
				continue
			}

			operation = o
			break polling
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	if len(states) == 0 {
		return operation, nil
	}

	for _, st := range states {
		if operation.State == st {
			return operation, nil
		}
	}

	var ref OperationReference
	if operation.Reference != nil {
		ref = *operation.Reference
	}

	return nil,
		fmt.Errorf("operation: %q %v, state: %s, reason: %q, message: %q",
			operation.ID,
			ref,
			operation.State,
			operation.Reason,
			operation.Message,
		)
}

func String(s string) *string {
	return &s
}

func Int64(i int64) *int64 {
	return &i
}

func Bool(b bool) *bool {
	return &b
}

func Ptr[T any](v T) *T {
	return &v
}

// Validate any struct from schema or request
func (c Client) Validate(s any) error {
	err := c.validate.Struct(s)
	if err == nil {
		return nil
	}

	// Print better error messages
	validationErrors := err.(validator.ValidationErrors)

	if len(validationErrors) > 0 {
		e := validationErrors[0]
		errorString := fmt.Sprintf(
			"request validation error: '%s' = '%v' does not validate ",
			e.StructNamespace(),
			e.Value(),
		)
		if e.Param() == "" {
			errorString += fmt.Sprintf("'%s'", e.ActualTag())
		} else {
			errorString += fmt.Sprintf("'%s=%v'", e.ActualTag(), e.Param())
		}
		return errors.New(errorString)
	}

	return err
}

// pollInterval returns the wait interval (as a time.Duration) before the next poll, based on the current runtime of a job.
// The polling frequency is:
//   - every 3 seconds for the first 30 seconds
//   - then increases linearly to reach 1 minute at 15 minutes of runtime
//   - after 15 minutes, it stays at 1 minute intervals
func pollInterval(runTime time.Duration) time.Duration {
	runTimeSeconds := runTime.Seconds()

	// Coefficients for the linear equation y = a * x + b
	a := 57.0 / 870.0
	b := 3.0 - 30.0*a

	minWait := 3.0
	maxWait := 60.0

	interval := a*runTimeSeconds + b
	interval = math.Max(minWait, interval)
	interval = math.Min(maxWait, interval)

	return time.Duration(interval) * time.Second
}

func prepareJSONBody(body any) (*bytes.Reader, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buf), nil
}

func prepareJSONResponse(resp *http.Response, v any) error {
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(buf, v); err != nil {
		return err
	}

	return nil
}

func (c Client) signRequest(req *http.Request) error {
	var (
		sigParts    []string
		headerParts []string
	)

	var expiration = time.Now().UTC().Add(time.Minute * 10)

	// Request method/URL path
	sigParts = append(sigParts, fmt.Sprintf("%s %s", req.Method, req.URL.EscapedPath()))
	headerParts = append(headerParts, "EXO2-HMAC-SHA256 credential="+c.apiKey)

	// Request body if present
	body := ""
	if req.Body != nil {
		data, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}
		err = req.Body.Close()
		if err != nil {
			return err
		}
		body = string(data)
		req.Body = io.NopCloser(bytes.NewReader(data))
	}
	sigParts = append(sigParts, body)

	// Request query string parameters
	// Important: this is order-sensitive, we have to have to sort parameters alphabetically to ensure signed
	// values match the names listed in the "signed-query-args=" signature pragma.
	signedParams, paramsValues := extractRequestParameters(req)
	sigParts = append(sigParts, paramsValues)
	if len(signedParams) > 0 {
		headerParts = append(headerParts, "signed-query-args="+strings.Join(signedParams, ";"))
	}

	// Request headers -- none at the moment
	// Note: the same order-sensitive caution for query string parameters applies to headers.
	sigParts = append(sigParts, "")

	// Request expiration date (UNIX timestamp, no line return)
	sigParts = append(sigParts, fmt.Sprint(expiration.Unix()))
	headerParts = append(headerParts, "expires="+fmt.Sprint(expiration.Unix()))

	h := hmac.New(sha256.New, []byte(c.apiSecret))
	if _, err := h.Write([]byte(strings.Join(sigParts, "\n"))); err != nil {
		return err
	}
	headerParts = append(headerParts, "signature="+base64.StdEncoding.EncodeToString(h.Sum(nil)))

	req.Header.Set("Authorization", strings.Join(headerParts, ","))

	return nil
}

// extractRequestParameters returns the list of request URL parameters names
// and a strings concatenating the values of the parameters.
func extractRequestParameters(req *http.Request) ([]string, string) {
	var (
		names  []string
		values string
	)

	for param, values := range req.URL.Query() {
		// Keep only parameters that hold exactly 1 value (i.e. no empty or multi-valued parameters)
		if len(values) == 1 {
			names = append(names, param)
		}
	}
	sort.Strings(names)

	for _, param := range names {
		values += req.URL.Query().Get(param)
	}

	return names, values
}

func dumpRequest(req *http.Request, operationID string) {
	if req != nil {
		if dump, err := httputil.DumpRequest(req, true); err == nil {
			fmt.Fprintf(os.Stderr, ">>> Operation: %s\n%s\n", operationID, dump)
		}
	}
}

func dumpResponse(resp *http.Response) {
	if resp != nil {
		if dump, err := httputil.DumpResponse(resp, true); err == nil {
			fmt.Fprintf(os.Stderr, "<<< %s\n", dump)
			fmt.Fprintln(os.Stderr, "----------------------------------------------------------------------")
		}
	}
}

// FindInstanceType attempts to find an InstanceType by id, or by a string or the form FAMILY.SIZE or SIZE,
// where family defaults to "standard"
func (l ListInstanceTypesResponse) FindInstanceTypeByIdOrFamilyAndSize(familyAndSizeOrId string) (InstanceType, error) {
	var result []InstanceType

	var typeFamily, typeSize string
	parts := strings.SplitN(familyAndSizeOrId, ".", 2)
	if l := len(parts); l > 0 {
		if l == 1 {
			typeFamily, typeSize = "standard", strings.ToLower(parts[0])
		} else {
			typeFamily, typeSize = strings.ToLower(parts[0]), strings.ToLower(parts[1])
		}
	}

	for i, elem := range l.InstanceTypes {
		if string(elem.ID) == familyAndSizeOrId || (string(elem.Size) == typeSize && string(elem.Family) == typeFamily) {
			result = append(result, l.InstanceTypes[i])
		}
	}
	if len(result) == 1 {
		return result[0], nil
	}

	if len(result) > 1 {
		return InstanceType{}, fmt.Errorf("%q too many found in ListInstanceTypesResponse: %w", familyAndSizeOrId, ErrConflict)
	}

	return InstanceType{}, fmt.Errorf("%q not found in ListInstanceTypesResponse: %w", familyAndSizeOrId, ErrNotFound)
}
