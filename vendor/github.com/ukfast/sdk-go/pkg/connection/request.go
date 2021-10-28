package connection

import (
	"errors"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

type Validatable interface {
	Validate() *ValidationError
}

type ValidationError struct {
	Message string
}

func NewValidationError(msg string) *ValidationError {
	return &ValidationError{Message: msg}
}

func (e *ValidationError) Error() string {
	return e.Message
}

// APIRequestBodyDefaultValidator provides convenience method Validate() for
// validating a struct/item, utilising the validator.v9 package
type APIRequestBodyDefaultValidator struct {
}

// Validate performs validation on item v using the validator.v9 package
func (a *APIRequestBodyDefaultValidator) Validate(v interface{}) *ValidationError {
	err := validator.New().Struct(v)
	if err != nil {
		return NewValidationError(err.Error())
	}

	return nil
}

type APIRequest struct {
	Method     string
	Resource   string
	Body       interface{}
	Parameters APIRequestParameters
}

// APIRequestParameters holds a collection of supported API request parameters
type APIRequestParameters struct {
	Pagination APIRequestPagination
	Sorting    APIRequestSorting
	Filtering  []APIRequestFiltering
}

func NewAPIRequestParameters() *APIRequestParameters {
	return &APIRequestParameters{}
}

// WithPagination is a fluent method for adding pagination to request parameters
func (p *APIRequestParameters) WithPagination(pagination APIRequestPagination) *APIRequestParameters {
	p.Pagination = pagination
	return p
}

// WithSorting is a fluent method for adding sorting to request parameters
func (p *APIRequestParameters) WithSorting(sorting APIRequestSorting) *APIRequestParameters {
	p.Sorting = sorting
	return p
}

// WithFilter is a fluent method for adding a set of filters to request parameters
func (p *APIRequestParameters) WithFilter(filters ...APIRequestFiltering) *APIRequestParameters {
	p.Filtering = append(p.Filtering, filters...)

	return p
}

func (p *APIRequestParameters) Copy() APIRequestParameters {
	newParameters := APIRequestParameters{}
	newParameters.Sorting = p.Sorting
	newParameters.Pagination = p.Pagination
	newParameters.Filtering = p.Filtering

	return newParameters
}

type APIRequestPagination struct {
	PerPage int
	Page    int
}

type APIRequestSorting struct {
	Property   string
	Descending bool
}

type APIRequestFiltering struct {
	Property string
	Operator APIRequestFilteringOperator
	Value    []string
}

func NewAPIRequestFiltering(property string, operator APIRequestFilteringOperator, value []string) *APIRequestFiltering {
	return &APIRequestFiltering{
		Property: property,
		Operator: operator,
		Value:    value,
	}
}

type APIRequestFilteringOperator int

const (
	// EQOperator - equals
	EQOperator APIRequestFilteringOperator = iota
	// LKOperator - like
	LKOperator
	// GTOperator - greater than
	GTOperator
	// LTOperator - less than
	LTOperator
	// INOperator - in set
	INOperator
	// NEQOperator - not equal
	NEQOperator
	// NINOperator - not in set
	NINOperator
	// NLKOperator - not like
	NLKOperator
)

func (o APIRequestFilteringOperator) String() string {
	operators := [...]string{
		"eq",
		"lk",
		"gt",
		"lt",
		"in",
		"neq",
		"nin",
		"nlk",
	}

	if o < EQOperator || o > NLKOperator {
		return "Unknown"
	}

	return operators[o]
}

// ParseOperator attempts to parse an operator from string
func ParseOperator(o string) (APIRequestFilteringOperator, error) {
	switch strings.ToUpper(o) {
	case "EQ":
		return EQOperator, nil
	case "LK":
		return LKOperator, nil
	case "GT":
		return GTOperator, nil
	case "LT":
		return LTOperator, nil
	case "IN":
		return INOperator, nil
	case "NEQ":
		return NEQOperator, nil
	case "NIN":
		return NINOperator, nil
	case "NLK":
		return NLKOperator, nil
	}

	return 0, errors.New("Invalid filtering operator")
}

// PaginatedGetFunc represents a function which can be called for returning an implementation of Paginated
type PaginatedGetFunc func(parameters APIRequestParameters) (Paginated, error)

// Paginated represents a paginated API response
type Paginated interface {
	TotalPages() int
	CurrentPage() int
	Total() int
	First() (Paginated, error)
	Previous() (Paginated, error)
	Next() (Paginated, error)
	Last() (Paginated, error)
}

// PaginatedBase implements the Paginated interface, and can be embedded to represent different paginated types
type PaginatedBase struct {
	parameters APIRequestParameters
	pagination APIResponseMetadataPagination
	getFunc    PaginatedGetFunc
}

// NewPaginatedBase returns a pointer to an initialised PaginatedBase
func NewPaginatedBase(parameters APIRequestParameters, pagination APIResponseMetadataPagination, getFunc PaginatedGetFunc) *PaginatedBase {
	return &PaginatedBase{
		parameters: parameters,
		pagination: pagination,
		getFunc:    getFunc,
	}
}

// TotalPages returns the total number of pages
func (p *PaginatedBase) TotalPages() int {
	return p.pagination.TotalPages
}

// CurrentPage returns the current page
func (p *PaginatedBase) CurrentPage() int {
	if p.parameters.Pagination.Page > 0 {
		return p.parameters.Pagination.Page
	}

	return 1
}

// Total returns the total number of items for current page
func (p *PaginatedBase) Total() int {
	return p.pagination.Total
}

// First returns the the first page
func (p *PaginatedBase) First() (Paginated, error) {
	newParameters := p.parameters.Copy()
	newParameters.Pagination.Page = 1
	return p.getFunc(newParameters)
}

// Previous returns the the previous page, or nil if current page is the first page
func (p *PaginatedBase) Previous() (Paginated, error) {
	if p.CurrentPage() <= 1 {
		return nil, nil
	}

	newParameters := p.parameters.Copy()
	newParameters.Pagination.Page = p.CurrentPage() - 1
	return p.getFunc(newParameters)
}

// Next returns the the next page, or nil if current page is the last page
func (p *PaginatedBase) Next() (Paginated, error) {
	if p.CurrentPage() >= p.TotalPages() {
		return nil, nil
	}

	newParameters := p.parameters.Copy()
	newParameters.Pagination.Page = p.CurrentPage() + 1
	return p.getFunc(newParameters)
}

// Last returns the the last page
func (p *PaginatedBase) Last() (Paginated, error) {
	newParameters := p.parameters.Copy()
	newParameters.Pagination.Page = p.TotalPages()
	return p.getFunc(newParameters)
}

// RequestAllResponseFunc represents a function which is called after every request within Invoke()
type RequestAllResponseFunc func(Paginated)

// RequestAll contains a single method Invoke(), which is a convenience method for iterating over paginated API responses
type RequestAll struct {
	getFunc      PaginatedGetFunc
	responseFunc RequestAllResponseFunc
}

// NewRequestAll returns a pointer to an initialised RequestAll struct
func NewRequestAll(getFunc PaginatedGetFunc, responseFunc RequestAllResponseFunc) *RequestAll {
	return &RequestAll{
		getFunc:      getFunc,
		responseFunc: responseFunc,
	}
}

// Invoke iterates over all pages using provided parameters and the getFunc which RequestAll was intialised with
func (r *RequestAll) Invoke(parameters APIRequestParameters) error {
	totalPages := 1
	for currentPage := 1; currentPage <= totalPages; currentPage++ {
		parameters.Pagination.Page = currentPage
		paginated, err := r.getFunc(parameters)
		if err != nil {
			return err
		}

		r.responseFunc(paginated)
		totalPages = paginated.TotalPages()
	}

	return nil
}

// InvokeRequestAll is a convenience method for initialising RequestAll and calling Invoke()
func InvokeRequestAll(getFunc PaginatedGetFunc, responseFunc RequestAllResponseFunc, parameters APIRequestParameters) error {
	r := NewRequestAll(getFunc, responseFunc)

	return r.Invoke(parameters)
}
