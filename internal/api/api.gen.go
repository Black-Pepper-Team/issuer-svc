// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-merkletree-sql"
)

// CreateClaimRequest defines model for CreateClaimRequest.
type CreateClaimRequest struct {
	CredentialSchema      string          `json:"credentialSchema"`
	CredentialSubject     json.RawMessage `json:"credentialSubject"`
	Expiration            int64           `json:"expiration"`
	MerklizedRootPosition string          `json:"merklizedRootPosition"`
	RevNonce              *uint64         `json:"revNonce,omitempty"`
	SubjectPosition       string          `json:"subjectPosition"`
	Type                  string          `json:"type"`
	Version               uint32          `json:"version"`
}

// CreateClaimResponse defines model for CreateClaimResponse.
type CreateClaimResponse struct {
	Id string `json:"id"`
}

// CreateIdentityResponse defines model for CreateIdentityResponse.
type CreateIdentityResponse struct {
	Identifier *string        `json:"identifier,omitempty"`
	Immutable  bool           `json:"immutable"`
	Relay      string         `json:"relay"`
	State      *IdentityState `json:"state,omitempty"`
}

// GenericErrorMessage defines model for GenericErrorMessage.
type GenericErrorMessage struct {
	Message *string `json:"message,omitempty"`
}

// Health defines model for Health.
type Health struct {
	Cache bool `json:"cache"`
	Db    bool `json:"db"`
}

// IdentityState defines model for IdentityState.
type IdentityState struct {
	BlockNumber        *int      `json:"blockNumber,omitempty"`
	BlockTimestamp     *int      `json:"blockTimestamp,omitempty"`
	ClaimsTreeRoot     *string   `json:"claimsTreeRoot,omitempty"`
	CreatedAt          time.Time `json:"createdAt"`
	Identifier         string    `json:"-"`
	ModifiedAt         time.Time `json:"modifiedAt"`
	PreviousState      *string   `json:"previousState,omitempty"`
	RevocationTreeRoot *string   `json:"revocationTreeRoot,omitempty"`
	RootOfRoots        *string   `json:"rootOfRoots,omitempty"`
	State              *string   `json:"state,omitempty"`
	StateID            int64     `json:"-"`
	Status             string    `json:"status"`
	TxID               *string   `json:"txID,omitempty"`
}

// Pong defines model for Pong.
type Pong struct {
	Response *string `json:"response,omitempty"`
}

// PublishStateResponse defines model for PublishStateResponse.
type PublishStateResponse struct {
	Hex *string `json:"hex,omitempty"`
}

// RevocationStatusResponse defines model for RevocationStatusResponse.
type RevocationStatusResponse struct {
	Issuer struct {
		ClaimsTreeRoot     *string `json:"claimsTreeRoot,omitempty"`
		RevocationTreeRoot *string `json:"revocationTreeRoot,omitempty"`
		RootOfRoots        *string `json:"rootOfRoots,omitempty"`
		State              *string `json:"state,omitempty"`
	} `json:"issuer"`
	Mtp struct {
		Existence bool `json:"existence"`
		NodeAux   *struct {
			Key   *merkletree.Hash `json:"key,omitempty"`
			Value *merkletree.Hash `json:"value,omitempty"`
		} `json:"nodeAux,omitempty"`
	} `json:"mtp"`
}

// RevokeClaimResponse defines model for RevokeClaimResponse.
type RevokeClaimResponse struct {
	Status string `json:"status"`
}

// PathIdentifier defines model for pathIdentifier.
type PathIdentifier = core.DID

// PathNonce defines model for pathNonce.
type PathNonce = int64

// N400 defines model for 400.
type N400 = GenericErrorMessage

// N401 defines model for 401.
type N401 = GenericErrorMessage

// N402 defines model for 402.
type N402 = GenericErrorMessage

// N407 defines model for 407.
type N407 = GenericErrorMessage

// N500 defines model for 500.
type N500 = GenericErrorMessage

// N500CreateIdentity defines model for 500-CreateIdentity.
type N500CreateIdentity struct {
	Code      *int    `json:"code,omitempty"`
	Error     *string `json:"error,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
}

// CreateClaimJSONRequestBody defines body for CreateClaim for application/json ContentType.
type CreateClaimJSONRequestBody = CreateClaimRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get the documentation
	// (GET /)
	Get(w http.ResponseWriter, r *http.Request)
	// Play Ping Pong
	// (GET /ping)
	Ping(w http.ResponseWriter, r *http.Request)
	// Return random responses and status codes
	// (GET /random)
	Random(w http.ResponseWriter, r *http.Request)
	// Healthcheck
	// (GET /status)
	Health(w http.ResponseWriter, r *http.Request)
	// Create Identity
	// (POST /v1/identities)
	CreateIdentity(w http.ResponseWriter, r *http.Request)
	// Publish State On-Chain
	// (POST /v1/identities/state)
	PublishState(w http.ResponseWriter, r *http.Request)
	// Create Claim
	// (POST /v1/{identifier}/claims)
	CreateClaim(w http.ResponseWriter, r *http.Request, identifier PathIdentifier)
	// Get Revocation Status
	// (GET /v1/{identifier}/claims/revocation/status/{nonce})
	GetRevocationStatus(w http.ResponseWriter, r *http.Request, identifier PathIdentifier, nonce PathNonce)
	// Revoke Claim
	// (POST /v1/{identifier}/claims/revoke/{nonce})
	RevokeClaim(w http.ResponseWriter, r *http.Request, identifier PathIdentifier, nonce PathNonce)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// Get operation middleware
func (siw *ServerInterfaceWrapper) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Get(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// Ping operation middleware
func (siw *ServerInterfaceWrapper) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Ping(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// Random operation middleware
func (siw *ServerInterfaceWrapper) Random(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Random(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// Health operation middleware
func (siw *ServerInterfaceWrapper) Health(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Health(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateIdentity operation middleware
func (siw *ServerInterfaceWrapper) CreateIdentity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateIdentity(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PublishState operation middleware
func (siw *ServerInterfaceWrapper) PublishState(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PublishState(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateClaim operation middleware
func (siw *ServerInterfaceWrapper) CreateClaim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "identifier" -------------
	var identifier PathIdentifier

	err = runtime.BindStyledParameterWithLocation("simple", false, "identifier", runtime.ParamLocationPath, chi.URLParam(r, "identifier"), &identifier)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "identifier", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateClaim(w, r, identifier)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetRevocationStatus operation middleware
func (siw *ServerInterfaceWrapper) GetRevocationStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "identifier" -------------
	var identifier PathIdentifier

	err = runtime.BindStyledParameterWithLocation("simple", false, "identifier", runtime.ParamLocationPath, chi.URLParam(r, "identifier"), &identifier)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "identifier", Err: err})
		return
	}

	// ------------- Path parameter "nonce" -------------
	var nonce PathNonce

	err = runtime.BindStyledParameterWithLocation("simple", false, "nonce", runtime.ParamLocationPath, chi.URLParam(r, "nonce"), &nonce)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "nonce", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRevocationStatus(w, r, identifier, nonce)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// RevokeClaim operation middleware
func (siw *ServerInterfaceWrapper) RevokeClaim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "identifier" -------------
	var identifier PathIdentifier

	err = runtime.BindStyledParameterWithLocation("simple", false, "identifier", runtime.ParamLocationPath, chi.URLParam(r, "identifier"), &identifier)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "identifier", Err: err})
		return
	}

	// ------------- Path parameter "nonce" -------------
	var nonce PathNonce

	err = runtime.BindStyledParameterWithLocation("simple", false, "nonce", runtime.ParamLocationPath, chi.URLParam(r, "nonce"), &nonce)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "nonce", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RevokeClaim(w, r, identifier, nonce)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/", wrapper.Get)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/ping", wrapper.Ping)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/random", wrapper.Random)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/status", wrapper.Health)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/identities", wrapper.CreateIdentity)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/identities/state", wrapper.PublishState)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/{identifier}/claims", wrapper.CreateClaim)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/{identifier}/claims/revocation/status/{nonce}", wrapper.GetRevocationStatus)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/{identifier}/claims/revoke/{nonce}", wrapper.RevokeClaim)
	})

	return r
}

type N400JSONResponse GenericErrorMessage

type N401JSONResponse GenericErrorMessage

type N402JSONResponse GenericErrorMessage

type N407JSONResponse GenericErrorMessage

type N500JSONResponse GenericErrorMessage

type N500CreateIdentityJSONResponse struct {
	Code      *int    `json:"code,omitempty"`
	Error     *string `json:"error,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
}

type GetRequestObject struct {
}

type GetResponseObject interface {
	VisitGetResponse(w http.ResponseWriter) error
}

type PingRequestObject struct {
}

type PingResponseObject interface {
	VisitPingResponse(w http.ResponseWriter) error
}

type Ping201JSONResponse Pong

func (response Ping201JSONResponse) VisitPingResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type Ping500JSONResponse struct{ N500JSONResponse }

func (response Ping500JSONResponse) VisitPingResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type RandomRequestObject struct {
}

type RandomResponseObject interface {
	VisitRandomResponse(w http.ResponseWriter) error
}

type Random400JSONResponse struct{ N400JSONResponse }

func (response Random400JSONResponse) VisitRandomResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type Random401JSONResponse struct{ N401JSONResponse }

func (response Random401JSONResponse) VisitRandomResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type Random402JSONResponse struct{ N402JSONResponse }

func (response Random402JSONResponse) VisitRandomResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(402)

	return json.NewEncoder(w).Encode(response)
}

type Random407JSONResponse struct{ N407JSONResponse }

func (response Random407JSONResponse) VisitRandomResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(407)

	return json.NewEncoder(w).Encode(response)
}

type Random500JSONResponse struct{ N500JSONResponse }

func (response Random500JSONResponse) VisitRandomResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type HealthRequestObject struct {
}

type HealthResponseObject interface {
	VisitHealthResponse(w http.ResponseWriter) error
}

type Health200JSONResponse Health

func (response Health200JSONResponse) VisitHealthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type Health500JSONResponse struct{ N500JSONResponse }

func (response Health500JSONResponse) VisitHealthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type CreateIdentityRequestObject struct {
}

type CreateIdentityResponseObject interface {
	VisitCreateIdentityResponse(w http.ResponseWriter) error
}

type CreateIdentity201JSONResponse CreateIdentityResponse

func (response CreateIdentity201JSONResponse) VisitCreateIdentityResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateIdentity500JSONResponse struct{ N500CreateIdentityJSONResponse }

func (response CreateIdentity500JSONResponse) VisitCreateIdentityResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PublishStateRequestObject struct {
}

type PublishStateResponseObject interface {
	VisitPublishStateResponse(w http.ResponseWriter) error
}

type PublishState200JSONResponse PublishStateResponse

func (response PublishState200JSONResponse) VisitPublishStateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PublishState500JSONResponse struct{ N500JSONResponse }

func (response PublishState500JSONResponse) VisitPublishStateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type CreateClaimRequestObject struct {
	Identifier PathIdentifier `json:"identifier"`
	Body       *CreateClaimJSONRequestBody
}

type CreateClaimResponseObject interface {
	VisitCreateClaimResponse(w http.ResponseWriter) error
}

type CreateClaim201JSONResponse CreateClaimResponse

func (response CreateClaim201JSONResponse) VisitCreateClaimResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateClaim400JSONResponse struct{ N400JSONResponse }

func (response CreateClaim400JSONResponse) VisitCreateClaimResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type CreateClaim500JSONResponse struct{ N500JSONResponse }

func (response CreateClaim500JSONResponse) VisitCreateClaimResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetRevocationStatusRequestObject struct {
	Identifier PathIdentifier `json:"identifier"`
	Nonce      PathNonce      `json:"nonce"`
}

type GetRevocationStatusResponseObject interface {
	VisitGetRevocationStatusResponse(w http.ResponseWriter) error
}

type GetRevocationStatus200JSONResponse RevocationStatusResponse

func (response GetRevocationStatus200JSONResponse) VisitGetRevocationStatusResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetRevocationStatus400JSONResponse struct{ N400JSONResponse }

func (response GetRevocationStatus400JSONResponse) VisitGetRevocationStatusResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetRevocationStatus500JSONResponse struct{ N500JSONResponse }

func (response GetRevocationStatus500JSONResponse) VisitGetRevocationStatusResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type RevokeClaimRequestObject struct {
	Identifier PathIdentifier `json:"identifier"`
	Nonce      PathNonce      `json:"nonce"`
}

type RevokeClaimResponseObject interface {
	VisitRevokeClaimResponse(w http.ResponseWriter) error
}

type RevokeClaim202JSONResponse RevokeClaimResponse

func (response RevokeClaim202JSONResponse) VisitRevokeClaimResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)

	return json.NewEncoder(w).Encode(response)
}

type RevokeClaim400JSONResponse struct{ N400JSONResponse }

func (response RevokeClaim400JSONResponse) VisitRevokeClaimResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type RevokeClaim500JSONResponse struct{ N500JSONResponse }

func (response RevokeClaim500JSONResponse) VisitRevokeClaimResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get the documentation
	// (GET /)
	Get(ctx context.Context, request GetRequestObject) (GetResponseObject, error)
	// Play Ping Pong
	// (GET /ping)
	Ping(ctx context.Context, request PingRequestObject) (PingResponseObject, error)
	// Return random responses and status codes
	// (GET /random)
	Random(ctx context.Context, request RandomRequestObject) (RandomResponseObject, error)
	// Healthcheck
	// (GET /status)
	Health(ctx context.Context, request HealthRequestObject) (HealthResponseObject, error)
	// Create Identity
	// (POST /v1/identities)
	CreateIdentity(ctx context.Context, request CreateIdentityRequestObject) (CreateIdentityResponseObject, error)
	// Publish State On-Chain
	// (POST /v1/identities/state)
	PublishState(ctx context.Context, request PublishStateRequestObject) (PublishStateResponseObject, error)
	// Create Claim
	// (POST /v1/{identifier}/claims)
	CreateClaim(ctx context.Context, request CreateClaimRequestObject) (CreateClaimResponseObject, error)
	// Get Revocation Status
	// (GET /v1/{identifier}/claims/revocation/status/{nonce})
	GetRevocationStatus(ctx context.Context, request GetRevocationStatusRequestObject) (GetRevocationStatusResponseObject, error)
	// Revoke Claim
	// (POST /v1/{identifier}/claims/revoke/{nonce})
	RevokeClaim(ctx context.Context, request RevokeClaimRequestObject) (RevokeClaimResponseObject, error)
}

type StrictHandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// Get operation middleware
func (sh *strictHandler) Get(w http.ResponseWriter, r *http.Request) {
	var request GetRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.Get(ctx, request.(GetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Get")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetResponseObject); ok {
		if err := validResponse.VisitGetResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// Ping operation middleware
func (sh *strictHandler) Ping(w http.ResponseWriter, r *http.Request) {
	var request PingRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.Ping(ctx, request.(PingRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Ping")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PingResponseObject); ok {
		if err := validResponse.VisitPingResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// Random operation middleware
func (sh *strictHandler) Random(w http.ResponseWriter, r *http.Request) {
	var request RandomRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.Random(ctx, request.(RandomRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Random")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(RandomResponseObject); ok {
		if err := validResponse.VisitRandomResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// Health operation middleware
func (sh *strictHandler) Health(w http.ResponseWriter, r *http.Request) {
	var request HealthRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.Health(ctx, request.(HealthRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Health")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(HealthResponseObject); ok {
		if err := validResponse.VisitHealthResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// CreateIdentity operation middleware
func (sh *strictHandler) CreateIdentity(w http.ResponseWriter, r *http.Request) {
	var request CreateIdentityRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CreateIdentity(ctx, request.(CreateIdentityRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateIdentity")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CreateIdentityResponseObject); ok {
		if err := validResponse.VisitCreateIdentityResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// PublishState operation middleware
func (sh *strictHandler) PublishState(w http.ResponseWriter, r *http.Request) {
	var request PublishStateRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PublishState(ctx, request.(PublishStateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PublishState")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PublishStateResponseObject); ok {
		if err := validResponse.VisitPublishStateResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// CreateClaim operation middleware
func (sh *strictHandler) CreateClaim(w http.ResponseWriter, r *http.Request, identifier PathIdentifier) {
	var request CreateClaimRequestObject

	request.Identifier = identifier

	var body CreateClaimJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CreateClaim(ctx, request.(CreateClaimRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateClaim")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CreateClaimResponseObject); ok {
		if err := validResponse.VisitCreateClaimResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// GetRevocationStatus operation middleware
func (sh *strictHandler) GetRevocationStatus(w http.ResponseWriter, r *http.Request, identifier PathIdentifier, nonce PathNonce) {
	var request GetRevocationStatusRequestObject

	request.Identifier = identifier
	request.Nonce = nonce

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetRevocationStatus(ctx, request.(GetRevocationStatusRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetRevocationStatus")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetRevocationStatusResponseObject); ok {
		if err := validResponse.VisitGetRevocationStatusResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// RevokeClaim operation middleware
func (sh *strictHandler) RevokeClaim(w http.ResponseWriter, r *http.Request, identifier PathIdentifier, nonce PathNonce) {
	var request RevokeClaimRequestObject

	request.Identifier = identifier
	request.Nonce = nonce

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.RevokeClaim(ctx, request.(RevokeClaimRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "RevokeClaim")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(RevokeClaimResponseObject); ok {
		if err := validResponse.VisitRevokeClaimResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RZTXPbOBL9KyjsHinRVrKbKt0cO5W4ajdR2ZlTxgeIbImISAABQI80Lv73KXzwS4Qo",
	"KbGdy0wsNIDuh9cP3eATTnghOAOmFZ4/YUEkKUCD9H/p7DYFpumKgjS/pKASSYWmnOE5dmN6h2hrFGFq",
	"hsxUHGFGCsBz3BuX8KOkElI817KECKskg4KY1fVOGGulJWVrHOHtZM0n/seES5je3N50f57QQnCpzVS/",
	"UVnSFEdu9zleU52Vy2nCi9h48CZe84n9x8SshquqcqafOUtgGN51TmiBmB0MhlUPHY5oxWVBtEGA6f++",
	"xVEdImUa1iCdCxKU4EyBxfztxYX5X8KZBmZDI0LkNCHGqfi7Mp49dXb4t4QVnuN/xe1Bxm5UxR+BgaTJ",
	"Bym5/D8oRdY+6H6c70mK7uBHCUrjKsJvLy5f24M/GCl1xiX9G1Lnwuy1XViQXQFMWyDsWVo33r26G5Jv",
	"d+iq1JnJGLdTz6f/vD4/bpkGyUiO7kE+gkRg7L0vk2sJREOtBGe5JiQXIDV1xE94Ch0RaDIkwm6/gT5U",
	"LvFA6dubwGjVJBtffodEnxFZVaewdcxFaMWgzpKh8xIsAiS/H1EzXlANhTAwrUiuoIq6E0vnZlc1LGbR",
	"aSvBVlBJXGAnCE+EC5Cb3GTcHed6wRWt5wZgfmwkslm4PLyycqGMrul+OBGlR5BqPzCz/5tZSFK7gvxt",
	"eDR+Rgj7HorttsOIDqH3MCBd1OePU/ohgWh6Ihh70dF0ZM86K8e27d7ug0OiRVFqssy7J7XkPAfCHC9y",
	"sgtOVJpoOCY/tXv31ngYmZeUep+uO26HUgWDD8naIPKiHYAtKYSJEd/zAnRG2RplRAgY5l5AVSL8CUhu",
	"6o2BKJAk62/gyoMhkOnyBLM9fNKlobDdIgRCH9yBb8ucJ5vPZbHsnXwnh63BV1qA0qQQYZvEcFp9lQAm",
	"CYJESCwR06u+rqVEw0TTAoYAR6Oc9MWf0cUJXTNTxVm0jJzx1Ew6aysh4ZHyUjUghaSPuwtsNEzJuf6y",
	"MsNqPB/CI+4CO6bZI8H7bAgK7fbQ9dhlU+1G1K/V/bo9dLuHGmLegrP1kHCyo0En5NSiXOZUZfZgDstX",
	"BtsT17trzvHehjQiiUqVjnp72Xyc7C9MllBchRZDV2FLlQZ/YQ/VhvEUrsrtcN4GdsfaMHvtgZYA009E",
	"ZePdmLcY6cba5SbqR25DfCR5Cb/bjUDx2M+XFuKHY6aeUO6sHg5wc3OsQGgzvL2uBLDUYXN+4XDw/jSG",
	"lK34sB++4UlpeiTXlKy4RDoDdA/5Cn3iSkOKbm/QIifaCNmf9vKk2t2rYZtOiTXHF9PL6YWBgwtgRFA8",
	"x2/sT+7kbOix+c8a7OEacKwntyme44+gca+TZmWem7qtKIjcOQPrbtoNwp5VLKgTLL9wP+ivGWEbhTRH",
	"OqMKAUsFp0yjHS9RQhgSOdkhswIS3ELf92vhzqPX4s+escG2YhvqqIW5+NJpp2UMLdP4FRsj2/g0iC1M",
	"ZMZ/tHCRabJWhjr3kFgGXOIHi58kLOXFCIJd3CToUjKF3CTUOID+ojpDjCMFhv/7ON65PcKPJeORGaP2",
	"WeOY7WXn/eGY7azzSHDM9t0vnMSdxWwIGWEpcnmMTPesRs6olY9g+vgSdkDU53tr8DsEqHqV50iBfKSJ",
	"CUkCkiVj/ir8ScDcZkkGycal+ONl7PsJL6eCqwBVP9Qs1Ry5IgcRhjqdSB+1vbePF0zzA/1c6EmjfpP1",
	"Rdo5KO4/5vRBdYOoE25NtuanhyHYcVPNHIdcuKLPchoQZ5MkI5QNNbVTG74kZYM1aABya1A7fx7i+5Lr",
	"AXArfmGTaw/AYaSf2nq9il2Neh6/kS08DnC7Hut+F/gWjqs1ife+G1QPzUvde57unjkpeg9zVb/M8d3R",
	"C6dlv3ILEMR9SOgk5Dn31k9SyWdrfYA1gdzfY+yJ2y7G3xrxk/3QUR284bu0Wvs6q10FNU3koGTb78p+",
	"mWnRSTPcY6aj5Qtpx8F+M/zmz1evRAtTBre+oQb2M/mxgS4pjouNm4MISoJi0+mAfjcFZs9Kgc1xcbhK",
	"EhCvpwvOq8O6YIztxxAHft/X//GE5DjCpczxHGdai3kcX87eTS+mF9NLC6dfcH9mU482nYBqv6O2xao5",
	"vfDE2djEWWDiNc9zP8xX7WQkITcybFjZKWX8gm3p8zPrXbvrt1nNgVo9VP8EAAD//0cAnBRnHwAA",
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
