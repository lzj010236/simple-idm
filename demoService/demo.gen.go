// Package demoService provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version v0.3.0 DO NOT EDIT.
package demoService

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/discord-gophers/goapi-gen/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Greeting defines model for Greeting.
type Greeting struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GreetingRequest defines model for GreetingRequest.
type GreetingRequest struct {
	// The name to include in the greeting.
	Name *string `json:"name,omitempty"`
}

// GetHelloParams defines parameters for GetHello.
type GetHelloParams struct {
	// The name to include in the greeting.
	Name string `json:"name"`
}

// PostHelloJSONBody defines parameters for PostHello.
type PostHelloJSONBody GreetingRequest

// GetHelloIDParams defines parameters for GetHelloID.
type GetHelloIDParams struct {
	// The name to include in the greeting.
	Name string `json:"name"`
}

// PostHelloJSONRequestBody defines body for PostHello for application/json ContentType.
type PostHelloJSONRequestBody PostHelloJSONBody

// Bind implements render.Binder.
func (PostHelloJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// Response is a common response struct for all the API calls.
// A Response object may be instantiated via functions for specific operation responses.
// It may also be instantiated directly, for the purpose of responding with a single status code.
type Response struct {
	body        interface{}
	Code        int
	contentType string
}

// Render implements the render.Renderer interface. It sets the Content-Type header
// and status code based on the response definition.
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", resp.contentType)
	render.Status(r, resp.Code)
	return nil
}

// Status is a builder method to override the default status code for a response.
func (resp *Response) Status(code int) *Response {
	resp.Code = code
	return resp
}

// ContentType is a builder method to override the default content type for a response.
func (resp *Response) ContentType(contentType string) *Response {
	resp.contentType = contentType
	return resp
}

// MarshalJSON implements the json.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(resp.body)
}

// MarshalXML implements the xml.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(resp.body)
}

// GetHelloJSON200Response is a constructor method for a GetHello response.
// A *Response is returned with the configured status code and content type from the spec.
func GetHelloJSON200Response(body Greeting) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostHelloJSON200Response is a constructor method for a PostHello response.
// A *Response is returned with the configured status code and content type from the spec.
func PostHelloJSON200Response(body Greeting) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetHelloIDJSON200Response is a constructor method for a GetHelloID response.
// A *Response is returned with the configured status code and content type from the spec.
func GetHelloIDJSON200Response(body Greeting) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a greeting
	// (GET /hello)
	GetHello(w http.ResponseWriter, r *http.Request, params GetHelloParams) *Response
	// Post a greeting
	// (POST /hello)
	PostHello(w http.ResponseWriter, r *http.Request) *Response
	// Get a greeting
	// (GET /hello/{id})
	GetHelloID(w http.ResponseWriter, r *http.Request, id int, params GetHelloIDParams) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// GetHello operation middleware
func (siw *ServerInterfaceWrapper) GetHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params GetHelloParams

	// ------------- Required query parameter "name" -------------

	if err := runtime.BindQueryParameter("form", true, true, "name", r.URL.Query(), &params.Name); err != nil {
		err = fmt.Errorf("invalid format for parameter name: %w", err)
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "name"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetHello(w, r, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// PostHello operation middleware
func (siw *ServerInterfaceWrapper) PostHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostHello(w, r)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetHelloID operation middleware
func (siw *ServerInterfaceWrapper) GetHelloID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id int

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetHelloIDParams

	// ------------- Required query parameter "name" -------------

	if err := runtime.BindQueryParameter("form", true, true, "name", r.URL.Query(), &params.Name); err != nil {
		err = fmt.Errorf("invalid format for parameter name: %w", err)
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "name"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetHelloID(w, r, id, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter %s: %v", err.paramName, err.err)
}

func (err UnescapedCookieParamError) Unwrap() error { return err.err }

type UnmarshalingParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnmarshalingParamError) Error() string {
	return fmt.Sprintf("error unmarshaling parameter %s as JSON: %v", err.paramName, err.err)
}

func (err UnmarshalingParamError) Unwrap() error { return err.err }

type RequiredParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err RequiredParamError) Error() string {
	if err.err == nil {
		return fmt.Sprintf("query parameter %s is required, but not found", err.paramName)
	} else {
		return fmt.Sprintf("query parameter %s is required, but errored: %s", err.paramName, err.err)
	}
}

func (err RequiredParamError) Unwrap() error { return err.err }

type RequiredHeaderError struct {
	paramName string
}

// Error implements error.
func (err RequiredHeaderError) Error() string {
	return fmt.Sprintf("header parameter %s is required, but not found", err.paramName)
}

type InvalidParamFormatError struct {
	err       error
	paramName string
}

// Error implements error.
func (err InvalidParamFormatError) Error() string {
	return fmt.Sprintf("invalid format for parameter %s: %v", err.paramName, err.err)
}

func (err InvalidParamFormatError) Unwrap() error { return err.err }

type TooManyValuesForParamError struct {
	NumValues int
	paramName string
}

// Error implements error.
func (err TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("expected one value for %s, got %d", err.paramName, err.NumValues)
}

// ParameterName is an interface that is implemented by error types that are
// relevant to a specific parameter.
type ParameterError interface {
	error
	// ParamName is the name of the parameter that the error is referring to.
	ParamName() string
}

func (err UnescapedCookieParamError) ParamName() string  { return err.paramName }
func (err UnmarshalingParamError) ParamName() string     { return err.paramName }
func (err RequiredParamError) ParamName() string         { return err.paramName }
func (err RequiredHeaderError) ParamName() string        { return err.paramName }
func (err InvalidParamFormatError) ParamName() string    { return err.paramName }
func (err TooManyValuesForParamError) ParamName() string { return err.paramName }

type ServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

type ServerOption func(*ServerOptions)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface, opts ...ServerOption) http.Handler {
	options := &ServerOptions{
		BaseURL:    "/",
		BaseRouter: chi.NewRouter(),
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
	}

	for _, f := range opts {
		f(options)
	}

	r := options.BaseRouter
	wrapper := ServerInterfaceWrapper{
		Handler:          si,
		ErrorHandlerFunc: options.ErrorHandlerFunc,
	}

	r.Route(options.BaseURL, func(r chi.Router) {
		r.Get("/hello", wrapper.GetHello)
		r.Post("/hello", wrapper.PostHello)
		r.Get("/hello/{id}", wrapper.GetHelloID)
	})
	return r
}

func WithRouter(r chi.Router) ServerOption {
	return func(s *ServerOptions) {
		s.BaseRouter = r
	}
}

func WithServerBaseURL(url string) ServerOption {
	return func(s *ServerOptions) {
		s.BaseURL = url
	}
}

func WithErrorHandler(handler func(w http.ResponseWriter, r *http.Request, err error)) ServerOption {
	return func(s *ServerOptions) {
		s.ErrorHandlerFunc = handler
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xUO2/bMBD+K8S1oyA57catDyB1pyLxFmRgpbPEQHyYPNU1DP734ihbtuwADYoMHbLY",
	"Ank8fvc9uIfaGe8sWoog9xDrDo3Kn7cBkbRt+dsH5zGQxryjG/7F38r4HkHeFEA7jyBBW8IWA6QCrDI4",
	"q4LvrrMwlUYK3DulAgJuBh2wAfnArQ9nH6dS9/MJa+KmR0h3uBkw0jWy460NxjpoT9pZkLDqUPCOICe0",
	"rfuhQaGtoA5Fe+hYPovsAgAvabt2fEPtLKmaTpfCygWvWidWqAwUMIQeJHREPsqq2m63JY0FZe0MDzPH",
	"eK+ZJrFsjDCuGXpkQJoyc/faXG79whDHgzflolxwP+fRKq9Bwse8VIBX1GVaqg77PsNuka75uUMago1C",
	"TXQIgzGqFpkW5ldx5bIBCbdI33Izbh+UQcIQQT78I+eaazcDht1Rdzn+nduCwoDFwZqM/lKoRy6O3tk4",
	"uuDDYnGUCG2eV3nf6zpPUT1FBrg/6/c+4BokvKtOWagOQaimFGT15zN+Er2OJNx6GimW2TdxMEaF3UjX",
	"Ga0sk3fxGQm+BFSEL5Pgh4uTBmFMwmfX7F595GPK0jykrEb6fxlndmaUp+Lg/2qvm/SqIVh+vY5BdjQn",
	"72To/Kj91c7T45mKtzC9KEwppT8BAAD//3sM5wDABgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}