package soap

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type SOAPEncoder interface {
	Encode(v interface{}) error
	Flush() error
}

type SOAPDecoder interface {
	Decode(v interface{}) error
}

type SOAPEnvelopeResponse struct {
	XMLName     xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header      *SOAPHeaderResponse
	Body        SOAPBodyResponse
	Attachments []MIMEMultipartAttachment `xml:"attachments,omitempty"`
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XmlNS   string   `xml:"xmlns:soap,attr"`

	Header *SOAPHeader
	Body   SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"soap:Header"`

	Headers []interface{}
}
type SOAPHeaderResponse struct {
	XMLName xml.Name `xml:"Header"`

	Headers []interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"soap:Body"`

	Content interface{} `xml:",omitempty"`

	// faultOccurred indicates whether the XML body included a fault;
	// we cannot simply store SOAPFault as a pointer to indicate this, since
	// fault is initialized to non-nil with user-provided detail type.
	faultOccurred bool
	Fault         *SOAPFault `xml:",omitempty"`
}

type SOAPBodyResponse struct {
	XMLName xml.Name `xml:"Body"`

	Content interface{} `xml:",omitempty"`

	// faultOccurred indicates whether the XML body included a fault;
	// we cannot simply store SOAPFault as a pointer to indicate this, since
	// fault is initialized to non-nil with user-provided detail type.
	faultOccurred bool
	Fault         *SOAPFault `xml:",omitempty"`
}

type MIMEMultipartAttachment struct {
	Name string
	Data []byte
}

// UnmarshalXML unmarshals SOAPBody xml
func (b *SOAPBodyResponse) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Content = nil

				b.faultOccurred = true
				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (b *SOAPBody) ErrorFromFault() error {
	if b.faultOccurred {
		return b.Fault
	}
	b.Fault = nil
	return nil
}

func (b *SOAPBodyResponse) ErrorFromFault() error {
	if b.faultOccurred {
		return b.Fault
	}
	b.Fault = nil
	return nil
}

type DetailContainer struct {
	Detail interface{}
}

type FaultError interface {
	// ErrorString should return a short version of the detail as a string,
	// which will be used in place of <faultstring> for the error message.
	// Set "HasData()" to always return false if <faultstring> error
	// message is preferred.
	ErrorString() string
	// HasData indicates whether the composite fault contains any data.
	HasData() bool
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string     `xml:"faultcode,omitempty"`
	String string     `xml:"faultstring,omitempty"`
	Actor  string     `xml:"faultactor,omitempty"`
	Detail FaultError `xml:"detail,omitempty"`
}

func (f *SOAPFault) Error() string {
	if f.Detail != nil && f.Detail.HasData() {
		return f.Detail.ErrorString()
	}
	return f.String
}

// HTTPError is returned whenever the HTTP request to the server fails
type HTTPError struct {
	//StatusCode is the status code returned in the HTTP response
	StatusCode int
	//ResponseBody contains the body returned in the HTTP response
	ResponseBody []byte
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP Status %d: %s", e.StatusCode, string(e.ResponseBody))
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE       string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU        string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType       string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
	mtomContentType string = `multipart/related; start-info="application/soap+xml"; type="application/xop+xml"; boundary="%s"`
	XmlNsSoapEnv    string = "http://schemas.xmlsoap.org/soap/envelope/"
)

type WSSSecurityHeader struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Id string `xml:"wsu:Id,attr,omitempty"`

	Username *WSSUsername `xml:",omitempty"`
	Password *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Data string `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`

	Data string `xml:",chardata"`
}

// NewWSSSecurityHeader creates WSSSecurityHeader instance
func NewWSSSecurityHeader(user, pass, tokenID, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: tokenID}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
}

type basicAuth struct {
	Login    string
	Password string
}

type options struct {
	tlsCfg           *tls.Config
	auth             *basicAuth
	timeout          time.Duration
	contimeout       time.Duration
	tlshshaketimeout time.Duration
	client           HTTPClient
	httpHeaders      map[string]string
	mtom             bool
	mma              bool
}

var defaultOptions = options{
	timeout:          time.Duration(30 * time.Second),
	contimeout:       time.Duration(90 * time.Second),
	tlshshaketimeout: time.Duration(15 * time.Second),
}

// A Option sets options such as credentials, tls, etc.
type Option func(*options)

// WithHTTPClient is an Option to set the HTTP client to use
// This cannot be used with WithTLSHandshakeTimeout, WithTLS,
// WithTimeout options
func WithHTTPClient(c HTTPClient) Option {
	return func(o *options) {
		o.client = c
	}
}

// WithTLSHandshakeTimeout is an Option to set default tls handshake timeout
// This option cannot be used with WithHTTPClient
func WithTLSHandshakeTimeout(t time.Duration) Option {
	return func(o *options) {
		o.tlshshaketimeout = t
	}
}

// WithRequestTimeout is an Option to set default end-end connection timeout
// This option cannot be used with WithHTTPClient
func WithRequestTimeout(t time.Duration) Option {
	return func(o *options) {
		o.contimeout = t
	}
}

// WithBasicAuth is an Option to set BasicAuth
func WithBasicAuth(login, password string) Option {
	return func(o *options) {
		o.auth = &basicAuth{Login: login, Password: password}
	}
}

// WithTLS is an Option to set tls config
// This option cannot be used with WithHTTPClient
func WithTLS(tls *tls.Config) Option {
	return func(o *options) {
		o.tlsCfg = tls
	}
}

// WithTimeout is an Option to set default HTTP dial timeout
func WithTimeout(t time.Duration) Option {
	return func(o *options) {
		o.timeout = t
	}
}

// WithHTTPHeaders is an Option to set global HTTP headers for all requests
func WithHTTPHeaders(headers map[string]string) Option {
	return func(o *options) {
		o.httpHeaders = headers
	}
}

// WithMTOM is an Option to set Message Transmission Optimization Mechanism
// MTOM encodes fields of type Binary using XOP.
func WithMTOM() Option {
	return func(o *options) {
		o.mtom = true
	}
}

// WithMIMEMultipartAttachments is an Option to set SOAP MIME Multipart attachment support.
// Use Client.AddMIMEMultipartAttachment to add attachments of type MIMEMultipartAttachment to your SOAP request.
func WithMIMEMultipartAttachments() Option {
	return func(o *options) {
		o.mma = true
	}
}

// Client is soap client
type Client struct {
	url         string
	opts        *options
	headers     []interface{}
	attachments []MIMEMultipartAttachment
}

// HTTPClient is a client which can make HTTP requests
// An example implementation is net/http.Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewClient creates new SOAP client instance
func NewClient(url string, opt ...Option) *Client {
	opts := defaultOptions
	for _, o := range opt {
		o(&opts)
	}
	return &Client{
		url:  url,
		opts: &opts,
	}
}

// AddHeader adds envelope header
// For correct behavior, every header must contain a `XMLName` field.  Refer to #121 for details
func (s *Client) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

// AddMIMEMultipartAttachment adds an attachment to the client that will be sent only if the
// WithMIMEMultipartAttachments option is used
func (s *Client) AddMIMEMultipartAttachment(attachment MIMEMultipartAttachment) {
	s.attachments = append(s.attachments, attachment)
}

// SetHeaders sets envelope headers, overwriting any existing headers.
// For correct behavior, every header must contain a `XMLName` field.  Refer to #121 for details
func (s *Client) SetHeaders(headers ...interface{}) {
	s.headers = headers
}

// CallContext performs HTTP POST request with a context
func (s *Client) CallContext(ctx context.Context, soapAction string, request, response interface{}) error {
	return s.call(ctx, soapAction, request, response, nil, nil)
}

// Call performs HTTP POST request.
// Note that if the server returns a status code >= 400, a HTTPError will be returned
func (s *Client) Call(soapAction string, request, response interface{}) error {
	return s.call(context.Background(), soapAction, request, response, nil, nil)
}

// CallContextWithAttachmentsAndFaultDetail performs HTTP POST request.
// Note that if SOAP fault is returned, it will be stored in the error.
// On top the attachments array will be filled with attachments returned from the SOAP request.
func (s *Client) CallContextWithAttachmentsAndFaultDetail(ctx context.Context, soapAction string, request,
	response interface{}, faultDetail FaultError, attachments *[]MIMEMultipartAttachment) error {
	return s.call(ctx, soapAction, request, response, faultDetail, attachments)
}

// CallContextWithFault performs HTTP POST request.
// Note that if SOAP fault is returned, it will be stored in the error.
func (s *Client) CallContextWithFaultDetail(ctx context.Context, soapAction string, request, response interface{}, faultDetail FaultError) error {
	return s.call(ctx, soapAction, request, response, faultDetail, nil)
}

// CallWithFaultDetail performs HTTP POST request.
// Note that if SOAP fault is returned, it will be stored in the error.
// the passed in fault detail is expected to implement FaultError interface,
// which allows to condense the detail into a short error message.
func (s *Client) CallWithFaultDetail(soapAction string, request, response interface{}, faultDetail FaultError) error {
	return s.call(context.Background(), soapAction, request, response, faultDetail, nil)
}

func (s *Client) call(ctx context.Context, soapAction string, request, response interface{}, faultDetail FaultError,
	retAttachments *[]MIMEMultipartAttachment) error {
	// SOAP envelope capable of namespace prefixes
	envelope := SOAPEnvelope{
		XmlNS: XmlNsSoapEnv,
	}

	if s.headers != nil && len(s.headers) > 0 {
		envelope.Header = &SOAPHeader{
			Headers: s.headers,
		}
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)
	var encoder SOAPEncoder
	if s.opts.mtom && s.opts.mma {
		return fmt.Errorf("cannot use MTOM (XOP) and MMA (MIME Multipart Attachments) option at the same time")
	} else if s.opts.mtom {
		encoder = newMtomEncoder(buffer)
	} else if s.opts.mma {
		encoder = newMmaEncoder(buffer, s.attachments)
	} else {
		encoder = xml.NewEncoder(buffer)
	}

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.opts.auth != nil {
		req.SetBasicAuth(s.opts.auth.Login, s.opts.auth.Password)
	}

	req = req.WithContext(ctx)

	if s.opts.mtom {
		req.Header.Add("Content-Type", fmt.Sprintf(mtomContentType, encoder.(*mtomEncoder).Boundary()))
	} else if s.opts.mma {
		req.Header.Add("Content-Type", fmt.Sprintf(mmaContentType, encoder.(*mmaEncoder).Boundary()))
	} else {
		req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	}
	req.Header.Add("SOAPAction", soapAction)
	req.Header.Set("User-Agent", "gowsdl/0.1")
	if s.opts.httpHeaders != nil {
		for k, v := range s.opts.httpHeaders {
			req.Header.Set(k, v)
		}
	}
	req.Close = true

	client := s.opts.client
	if client == nil {
		tr := &http.Transport{
			TLSClientConfig: s.opts.tlsCfg,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				d := net.Dialer{Timeout: s.opts.timeout}
				return d.DialContext(ctx, network, addr)
			},
			TLSHandshakeTimeout: s.opts.tlshshaketimeout,
		}
		client = &http.Client{Timeout: s.opts.contimeout, Transport: tr}
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(res.Body)
		return &HTTPError{
			StatusCode:   res.StatusCode,
			ResponseBody: body,
		}
	}

	// xml Decoder (used with and without MTOM) cannot handle namespace prefixes (yet),
	// so we have to use a namespace-less response envelope
	respEnvelope := new(SOAPEnvelopeResponse)
	respEnvelope.Body = SOAPBodyResponse{
		Content: response,
		Fault: &SOAPFault{
			Detail: faultDetail,
		},
	}

	mtomBoundary, err := getMtomHeader(res.Header.Get("Content-Type"))
	if err != nil {
		return err
	}

	var mmaBoundary string
	if s.opts.mma{
		mmaBoundary, err = getMmaHeader(res.Header.Get("Content-Type"))
		if err != nil {
			return err
		}
	}

	var dec SOAPDecoder
	if mtomBoundary != "" {
		dec = newMtomDecoder(res.Body, mtomBoundary)
	} else if mmaBoundary != "" {
		dec = newMmaDecoder(res.Body, mmaBoundary)
	} else {
		dec = xml.NewDecoder(res.Body)
	}

	if err := dec.Decode(respEnvelope); err != nil {
		return err
	}

	if respEnvelope.Attachments != nil {
		*retAttachments = respEnvelope.Attachments
	}
	return respEnvelope.Body.ErrorFromFault()
}
