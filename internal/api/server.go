package api

import (
	"context"
	"fmt"
	"github.com/iden3/go-merkletree-sql/v2"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-schema-processor/verifiable"
	"github.com/iden3/iden3comm"
	"github.com/iden3/iden3comm/packers"
	"github.com/iden3/iden3comm/protocol"
	"github.com/pkg/errors"

	"github.com/polygonid/sh-id-platform/internal/common"
	"github.com/polygonid/sh-id-platform/internal/config"
	"github.com/polygonid/sh-id-platform/internal/core/domain"
	"github.com/polygonid/sh-id-platform/internal/core/ports"
	"github.com/polygonid/sh-id-platform/internal/core/services"
	"github.com/polygonid/sh-id-platform/internal/gateways"
	"github.com/polygonid/sh-id-platform/internal/health"
	"github.com/polygonid/sh-id-platform/internal/log"
	"github.com/polygonid/sh-id-platform/internal/repositories"
	"github.com/polygonid/sh-id-platform/pkg/schema"
)

// Server implements StrictServerInterface and holds the implementation of all API controllers
// This is the glue to the API autogenerated code
type Server struct {
	cfg              *config.Configuration
	identityService  ports.IdentityService
	claimService     ports.ClaimsService
	usersService     ports.UsersService
	publisherGateway ports.Publisher
	packageManager   *iden3comm.PackageManager
	health           *health.Status
}

// NewServer is a Server constructor
func NewServer(cfg *config.Configuration, identityService ports.IdentityService, claimsService ports.ClaimsService, usersService ports.UsersService, publisherGateway ports.Publisher, packageManager *iden3comm.PackageManager, health *health.Status) *Server {
	return &Server{
		cfg:              cfg,
		identityService:  identityService,
		claimService:     claimsService,
		usersService:     usersService,
		publisherGateway: publisherGateway,
		packageManager:   packageManager,
		health:           health,
	}
}

// Health is a method
func (s *Server) Health(_ context.Context, _ HealthRequestObject) (HealthResponseObject, error) {
	var resp Health200JSONResponse = s.health.Status()

	return resp, nil
}

// GetDocumentation this method will be overridden in the main function
func (s *Server) GetDocumentation(_ context.Context, _ GetDocumentationRequestObject) (GetDocumentationResponseObject, error) {
	return nil, nil
}

// GetFavicon this method will be overridden in the main function
func (s *Server) GetFavicon(_ context.Context, _ GetFaviconRequestObject) (GetFaviconResponseObject, error) {
	return nil, nil
}

// GetYaml this method will be overridden in the main function
func (s *Server) GetYaml(_ context.Context, _ GetYamlRequestObject) (GetYamlResponseObject, error) {
	return nil, nil
}

// CreateIdentity is created identity controller
func (s *Server) CreateIdentity(ctx context.Context, request CreateIdentityRequestObject) (CreateIdentityResponseObject, error) {
	method := request.Body.DidMetadata.Method
	blockchain := request.Body.DidMetadata.Blockchain
	network := request.Body.DidMetadata.Network

	identity, err := s.identityService.Create(ctx, method, blockchain, network, s.cfg.ServerUrl)
	if err != nil {
		if errors.Is(err, services.ErrWrongDIDMetada) {
			return CreateIdentity400JSONResponse{
				N400JSONResponse{
					Message: err.Error(),
				},
			}, nil
		}
		return nil, err
	}

	return CreateIdentity201JSONResponse{
		Identifier: &identity.Identifier,
		State: &IdentityState{
			BlockNumber:        identity.State.BlockNumber,
			BlockTimestamp:     identity.State.BlockTimestamp,
			ClaimsTreeRoot:     identity.State.ClaimsTreeRoot,
			CreatedAt:          identity.State.CreatedAt,
			ModifiedAt:         identity.State.ModifiedAt,
			PreviousState:      identity.State.PreviousState,
			RevocationTreeRoot: identity.State.RevocationTreeRoot,
			RootOfRoots:        identity.State.RootOfRoots,
			State:              identity.State.State,
			Status:             string(identity.State.Status),
			TxID:               identity.State.TxID,
		},
	}, nil
}

// CreateClaim is claim creation controller
func (s *Server) CreateClaim(ctx context.Context, request CreateClaimRequestObject) (CreateClaimResponseObject, error) {
	did, err := core.ParseDID(request.Identifier)
	if err != nil {
		return CreateClaim400JSONResponse{N400JSONResponse{Message: err.Error()}}, nil
	}
	var expiration *time.Time
	if request.Body.Expiration != nil {
		expiration = common.ToPointer(time.Unix(*request.Body.Expiration, 0))
	}

	req := ports.NewCreateClaimRequest(did, request.Body.CredentialSchema, request.Body.CredentialSubject, expiration, request.Body.Type, request.Body.Version, request.Body.SubjectPosition, request.Body.MerklizedRootPosition, common.ToPointer(true), common.ToPointer(true), nil, false)

	resp, err := s.claimService.Save(ctx, req)
	if err != nil {
		if errors.Is(err, services.ErrJSONLdContext) {
			return CreateClaim400JSONResponse{N400JSONResponse{Message: err.Error()}}, nil
		}
		if errors.Is(err, services.ErrProcessSchema) {
			return CreateClaim400JSONResponse{N400JSONResponse{Message: err.Error()}}, nil
		}
		if errors.Is(err, services.ErrLoadingSchema) {
			return CreateClaim422JSONResponse{N422JSONResponse{Message: err.Error()}}, nil
		}
		if errors.Is(err, services.ErrMalformedURL) {
			return CreateClaim400JSONResponse{N400JSONResponse{Message: err.Error()}}, nil
		}
		if errors.Is(err, services.ErrParseClaim) {
			return CreateClaim400JSONResponse{N400JSONResponse{Message: err.Error()}}, nil
		}
		if errors.Is(err, services.ErrInvalidCredentialSubject) {
			return CreateClaim400JSONResponse{N400JSONResponse{Message: err.Error()}}, nil
		}
		if errors.Is(err, services.ErrLoadingSchema) {
			return CreateClaim400JSONResponse{N400JSONResponse{Message: err.Error()}}, nil
		}
		return CreateClaim500JSONResponse{N500JSONResponse{Message: err.Error()}}, nil
	}
	return CreateClaim201JSONResponse{Id: resp.ID.String()}, nil
}

// RevokeClaim is the revocation claim controller
func (s *Server) RevokeClaim(ctx context.Context, request RevokeClaimRequestObject) (RevokeClaimResponseObject, error) {
	did, err := core.ParseDID(request.Identifier)
	if err != nil {
		log.Warn(ctx, "revoke claim invalid did", "err", err, "req", request)
		return RevokeClaim400JSONResponse{N400JSONResponse{err.Error()}}, nil
	}

	if err := s.claimService.Revoke(ctx, *did, uint64(request.Nonce), ""); err != nil {
		if errors.Is(err, repositories.ErrClaimDoesNotExist) {
			return RevokeClaim404JSONResponse{N404JSONResponse{
				Message: "the claim does not exist",
			}}, nil
		}

		return RevokeClaim500JSONResponse{N500JSONResponse{Message: err.Error()}}, nil
	}
	return RevokeClaim202JSONResponse{
		Message: "claim revocation request sent",
	}, nil
}

// GetRevocationStatus is the controller to get revocation status
func (s *Server) GetRevocationStatus(ctx context.Context, request GetRevocationStatusRequestObject) (GetRevocationStatusResponseObject, error) {
	issuerDID, err := core.ParseDID(request.Identifier)
	if err != nil {
		return GetRevocationStatus500JSONResponse{N500JSONResponse{
			Message: err.Error(),
		}}, nil
	}

	stateHash := ""
	if request.Params.StateHash != nil {
		stateHash = *request.Params.StateHash
	}

	rs, err := s.claimService.GetRevocationStatus(ctx, *issuerDID, uint64(request.Nonce), stateHash)
	if err != nil {
		return GetRevocationStatus500JSONResponse{N500JSONResponse{
			Message: err.Error(),
		}}, nil
	}

	response := GetRevocationStatus200JSONResponse{}
	response.Issuer.State = rs.Issuer.State
	response.Issuer.RevocationTreeRoot = rs.Issuer.RevocationTreeRoot
	response.Issuer.RootOfRoots = rs.Issuer.RootOfRoots
	response.Issuer.ClaimsTreeRoot = rs.Issuer.ClaimsTreeRoot
	response.Mtp.Existence = rs.MTP.Existence

	if rs.MTP.NodeAux != nil {
		key := rs.MTP.NodeAux.Key
		decodedKey := key.BigInt().String()
		value := rs.MTP.NodeAux.Value
		decodedValue := value.BigInt().String()
		response.Mtp.NodeAux = &struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		}{
			Key:   &decodedKey,
			Value: &decodedValue,
		}
	}

	response.Mtp.Existence = rs.MTP.Existence
	siblings := make([]string, 0)
	for _, s := range rs.MTP.AllSiblings() {
		siblings = append(siblings, s.BigInt().String())
	}
	response.Mtp.Siblings = &siblings

	return response, err
}

// GetClaim is the controller to get a client.
func (s *Server) GetClaim(ctx context.Context, request GetClaimRequestObject) (GetClaimResponseObject, error) {
	if request.Identifier == "" {
		return GetClaim400JSONResponse{N400JSONResponse{"invalid did, cannot be empty"}}, nil
	}

	did, err := core.ParseDID(request.Identifier)
	if err != nil {
		return GetClaim400JSONResponse{N400JSONResponse{"invalid did"}}, nil
	}

	if request.Id == "" {
		return GetClaim400JSONResponse{N400JSONResponse{"cannot proceed with an empty claim id"}}, nil
	}

	clID, err := uuid.Parse(request.Id)
	if err != nil {
		return GetClaim400JSONResponse{N400JSONResponse{"invalid claim id"}}, nil
	}

	claim, err := s.claimService.GetByID(ctx, did, clID)
	if err != nil {
		if errors.Is(err, services.ErrClaimNotFound) {
			return GetClaim404JSONResponse{N404JSONResponse{err.Error()}}, nil
		}
		return GetClaim500JSONResponse{N500JSONResponse{err.Error()}}, nil
	}

	w3c, err := schema.FromClaimModelToW3CCredential(*claim)
	if err != nil {
		return GetClaim500JSONResponse{N500JSONResponse{"invalid claim format"}}, nil
	}

	return GetClaim200JSONResponse(toGetClaim200Response(w3c)), nil
}

// GetClaims is the controller to get multiple claims of a determined identity
func (s *Server) GetClaims(ctx context.Context, request GetClaimsRequestObject) (GetClaimsResponseObject, error) {
	if request.Identifier == "" {
		return GetClaims400JSONResponse{N400JSONResponse{"invalid did, cannot be empty"}}, nil
	}

	did, err := core.ParseDID(request.Identifier)
	if err != nil {
		return GetClaims400JSONResponse{N400JSONResponse{"invalid did"}}, nil
	}

	filter, err := ports.NewClaimsFilter(
		request.Params.SchemaHash,
		request.Params.SchemaType,
		request.Params.Subject,
		request.Params.QueryField,
		request.Params.QueryValue,
		request.Params.Self,
		request.Params.Revoked)
	if err != nil {
		return GetClaims400JSONResponse{N400JSONResponse{err.Error()}}, nil
	}

	claims, err := s.claimService.GetAll(ctx, *did, filter)
	if err != nil && !errors.Is(err, services.ErrClaimNotFound) {
		return GetClaims500JSONResponse{N500JSONResponse{"there was an internal error trying to retrieve claims for the requested identifier"}}, nil
	}

	w3Claims, err := schema.FromClaimsModelToW3CCredential(claims)
	if err != nil {
		return GetClaims500JSONResponse{N500JSONResponse{"there was an internal error parsing the claims"}}, nil
	}

	return toGetClaims200Response(w3Claims), nil
}

// GetClaimQrCode returns a GetClaimQrCodeResponseObject that can be used with any QR generator to create a QR and
// scan it with polygon wallet to accept the claim
func (s *Server) GetClaimQrCode(ctx context.Context, request GetClaimQrCodeRequestObject) (GetClaimQrCodeResponseObject, error) {
	if request.Identifier == "" {
		return GetClaimQrCode400JSONResponse{N400JSONResponse{"invalid did, cannot be empty"}}, nil
	}

	did, err := core.ParseDID(request.Identifier)
	if err != nil {
		return GetClaimQrCode400JSONResponse{N400JSONResponse{"invalid did"}}, nil
	}

	if request.Id == "" {
		return GetClaimQrCode400JSONResponse{N400JSONResponse{"cannot proceed with an empty claim id"}}, nil
	}

	claimID, err := uuid.Parse(request.Id)
	if err != nil {
		return GetClaimQrCode400JSONResponse{N400JSONResponse{"invalid claim id"}}, nil
	}

	claim, err := s.claimService.GetByID(ctx, did, claimID)
	if err != nil {
		if errors.Is(err, services.ErrClaimNotFound) {
			return GetClaimQrCode404JSONResponse{N404JSONResponse{err.Error()}}, nil
		}
		return GetClaimQrCode500JSONResponse{N500JSONResponse{err.Error()}}, nil
	}
	return toGetClaimQrCode200JSONResponse(claim, s.cfg.ServerUrl), nil
}

func (s *Server) GetClaimMTP(ctx context.Context, request GetClaimMTPRequestObject) (GetClaimMTPResponseObject, error) {
	if request.Id == "" {
		return GetClaimMTP400JSONResponse{N400JSONResponse{"cannot proceed with an empty claim id"}}, nil
	}

	claimID, err := uuid.Parse(request.Id)
	if err != nil {
		return GetClaimMTP400JSONResponse{N400JSONResponse{"invalid claim id"}}, nil
	}

	claim, err := s.claimService.GetBySingleID(ctx, claimID)
	if err != nil {
		if errors.Is(err, services.ErrClaimNotFound) {
			return GetClaimMTP404JSONResponse{N404JSONResponse{err.Error()}}, nil
		}
		return GetClaimMTP500JSONResponse{N500JSONResponse{}}, nil
	}

	state := new(domain.IdentityState)
	if request.Params.StateHash == nil {
		issuerDID, err := core.ParseDID(claim.Issuer)
		if err != nil {
			log.Error(ctx, "failed to parse DID", err)
			return GetClaimMTP500JSONResponse{N500JSONResponse{}}, nil
		}
		state, err = s.identityService.GetLatestStateByID(context.Background(), *issuerDID)
		if err != nil {
			log.Error(ctx, "failed to get latest state by ID", err)
			return GetClaimMTP500JSONResponse{N500JSONResponse{}}, nil
		}
	} else {
		state, err = s.identityService.GetStateByHash(ctx, *request.Params.StateHash)
		if err != nil {
			log.Error(ctx, "failed to get state by hash", err)
			return GetClaimMTP500JSONResponse{N500JSONResponse{}}, nil
		}
	}

	leaf, err := claim.CoreClaim.Get().HIndex()
	if err != nil {
		log.Error(ctx, "failed to get HIndex", err)
		return GetClaimMTP500JSONResponse{N500JSONResponse{}}, nil
	}

	root, err := merkletree.NewHashFromHex(*state.ClaimsTreeRoot)
	if err != nil {
		log.Error(ctx, "failed to get new hash from hex", err)
		return GetClaimMTP500JSONResponse{N500JSONResponse{}}, nil
	}

	iMT, err := s.claimService.GetMTByKey(ctx, *state.ClaimsTreeRoot)
	if err != nil {
		log.Error(ctx, "failed to get MT proof", err)
		return GetClaimMTP500JSONResponse{N500JSONResponse{}}, nil
	}

	proof, err := s.claimService.GetMTProof(ctx, leaf, root, iMT.MTID)
	if err != nil {
		log.Error(ctx, "failed to get MT proof", err)
		return GetClaimMTP500JSONResponse{N500JSONResponse{}}, nil
	}

	return toGetClaimMTP200JSONResponse(state, proof), nil
}

// GetIdentities is the controller to get identities
func (s *Server) GetIdentities(ctx context.Context, request GetIdentitiesRequestObject) (GetIdentitiesResponseObject, error) {
	var response GetIdentities200JSONResponse
	var err error
	response, err = s.identityService.Get(ctx)
	if err != nil {
		return GetIdentities500JSONResponse{N500JSONResponse{
			Message: err.Error(),
		}}, nil
	}

	return response, nil
}

// Agent is the controller to fetch credentials from mobile
func (s *Server) Agent(ctx context.Context, request AgentRequestObject) (AgentResponseObject, error) {
	if request.Body == nil || *request.Body == "" {
		log.Debug(ctx, "agent empty request")
		return Agent400JSONResponse{N400JSONResponse{"cannot proceed with an empty request"}}, nil
	}
	basicMessage, err := s.packageManager.UnpackWithType(packers.MediaTypeZKPMessage, []byte(*request.Body))
	if err != nil {
		log.Debug(ctx, "agent bad request", "err", err, "body", *request.Body)
		return Agent400JSONResponse{N400JSONResponse{"cannot proceed with the given request"}}, nil
	}

	req, err := ports.NewAgentRequest(basicMessage)
	if err != nil {
		log.Error(ctx, "agent parsing request", "err", err)
		return Agent400JSONResponse{N400JSONResponse{err.Error()}}, nil
	}

	agent, err := s.claimService.Agent(ctx, req)
	if err != nil {
		log.Error(ctx, "agent error", "err", err)
		return Agent400JSONResponse{N400JSONResponse{err.Error()}}, nil
	}
	return Agent200JSONResponse{
		Body:     agent.Body,
		From:     agent.From,
		Id:       agent.ID,
		ThreadID: agent.ThreadID,
		To:       agent.To,
		Typ:      string(agent.Typ),
		Type:     string(agent.Type),
	}, nil
}

// PublishIdentityState - publish identity state on chain
func (s *Server) PublishIdentityState(ctx context.Context, request PublishIdentityStateRequestObject) (PublishIdentityStateResponseObject, error) {
	did, err := core.ParseDID(request.Identifier)
	if err != nil {
		return PublishIdentityState400JSONResponse{N400JSONResponse{"invalid did"}}, nil
	}

	publishedState, err := s.publisherGateway.PublishState(ctx, did)
	if err != nil {
		if errors.Is(err, gateways.ErrNoStatesToProcess) || errors.Is(err, gateways.ErrStateIsBeingProcessed) {
			return PublishIdentityState200JSONResponse{Message: err.Error()}, nil
		}
		return PublishIdentityState500JSONResponse{N500JSONResponse{err.Error()}}, nil
	}

	return PublishIdentityState202JSONResponse{
		ClaimsTreeRoot:     publishedState.ClaimsTreeRoot,
		RevocationTreeRoot: publishedState.RevocationTreeRoot,
		RootOfRoots:        publishedState.RootOfRoots,
		State:              publishedState.State,
		TxID:               publishedState.TxID,
	}, nil
}

// RetryPublishState - retry to publish the current state if it failed previously.
func (s *Server) RetryPublishState(ctx context.Context, request RetryPublishStateRequestObject) (RetryPublishStateResponseObject, error) {
	did, err := core.ParseDID(request.Identifier)
	if err != nil {
		return RetryPublishState400JSONResponse{N400JSONResponse{"invalid did"}}, nil
	}

	publishedState, err := s.publisherGateway.RetryPublishState(ctx, did)
	if err != nil {
		log.Error(ctx, "error retrying the publishing the state", "err", err)
		if errors.Is(err, gateways.ErrStateIsBeingProcessed) || errors.Is(err, gateways.ErrNoFailedStatesToProcess) {
			return RetryPublishState400JSONResponse{N400JSONResponse{Message: err.Error()}}, nil
		}
		return RetryPublishState500JSONResponse{N500JSONResponse{Message: err.Error()}}, nil
	}
	return RetryPublishState202JSONResponse{
		ClaimsTreeRoot:     publishedState.ClaimsTreeRoot,
		RevocationTreeRoot: publishedState.RevocationTreeRoot,
		RootOfRoots:        publishedState.RootOfRoots,
		State:              publishedState.State,
		TxID:               publishedState.TxID,
	}, nil
}

// AddUser - add login and password to the database
func (s *Server) AddUser(ctx context.Context, request AddUserRequestObject) (AddUserResponseObject, error) {
	did, err := core.ParseDID(request.Body.Did)
	if err != nil {
		log.Error(context.Background(), "failed to parse did", err)
		return AddUser500JSONResponse{N500JSONResponse{"invalid did"}}, nil
	}
	if err := s.usersService.AddUser(ctx, request.Body.Login, request.Body.Password, *did); err != nil {
		log.Error(context.Background(), "failed to add user", err)
		return AddUser500JSONResponse{N500JSONResponse{"failed to add user"}}, nil
	}
	return AddUser200JSONResponse{}, nil
}

// RegisterStatic add method to the mux that are not documented in the API.
func RegisterStatic(mux *chi.Mux) {
	mux.Get("/", documentation)
	mux.Get("/static/docs/api/api.yaml", swagger)
	mux.Get("/favicon.ico", favicon)
}

func toGetClaims200Response(claims []*verifiable.W3CCredential) GetClaims200JSONResponse {
	response := make(GetClaims200JSONResponse, len(claims))
	for i := range claims {
		response[i] = toGetClaim200Response(claims[i])
	}

	return response
}

func toGetClaim200Response(claim *verifiable.W3CCredential) GetClaimResponse {
	return GetClaimResponse{
		Context: claim.Context,
		CredentialSchema: CredentialSchema{
			claim.CredentialSchema.ID,
			claim.CredentialSchema.Type,
		},
		CredentialStatus:  claim.CredentialStatus,
		CredentialSubject: claim.CredentialSubject,
		Expiration:        claim.Expiration,
		Id:                claim.ID,
		IssuanceDate:      claim.IssuanceDate,
		Issuer:            claim.Issuer,
		Proof:             claim.Proof,
		Type:              claim.Type,
	}
}

func toGetClaimQrCode200JSONResponse(claim *domain.Claim, hostURL string) *GetClaimQrCode200JSONResponse {
	id := uuid.New()
	return &GetClaimQrCode200JSONResponse{
		Body: struct {
			Credentials []struct {
				Description string `json:"description"`
				Id          string `json:"id"`
			} `json:"credentials"`
			Url string `json:"url"`
		}{
			Credentials: []struct {
				Description string `json:"description"`
				Id          string `json:"id"`
			}{
				{
					Description: claim.SchemaType,
					Id:          claim.ID.String(),
				},
			},
			Url: fmt.Sprintf("%s/v1/agent", strings.TrimSuffix(hostURL, "/")),
		},
		From: claim.Issuer,
		Id:   id.String(),
		Thid: id.String(),
		To:   claim.OtherIdentifier,
		Typ:  string(packers.MediaTypePlainMessage),
		Type: string(protocol.CredentialOfferMessageType),
	}
}

func toGetClaimMTP200JSONResponse(state *domain.IdentityState, proof *merkletree.Proof) *GetClaimMTP200JSONResponse {
	response := GetClaimMTP200JSONResponse{
		Issuer: struct {
			ClaimTreeRoot      *string `json:"claimTreeRoot,omitempty"`
			RevocationTreeRoot *string `json:"revocationTreeRoot,omitempty"`
			RootOfRoots        *string `json:"rootOfRoots,omitempty"`
			State              *string `json:"state,omitempty"`
		}{
			ClaimTreeRoot:      state.ClaimsTreeRoot,
			RevocationTreeRoot: state.RevocationTreeRoot,
			RootOfRoots:        state.RootOfRoots,
			State:              state.State,
		},
	}

	response.Mtp.Existence = proof.Existence

	if proof.NodeAux != nil {
		key := proof.NodeAux.Key
		decodedKey := key.BigInt().String()
		value := proof.NodeAux.Value
		decodedValue := value.BigInt().String()
		response.Mtp.NodeAux = &struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		}{
			Key:   &decodedKey,
			Value: &decodedValue,
		}
	}

	siblings := make([]string, 0)
	for _, s := range proof.AllSiblings() {
		siblings = append(siblings, s.BigInt().String())
	}
	response.Mtp.Siblings = &siblings

	return &response
}

func documentation(w http.ResponseWriter, _ *http.Request) {
	writeFile("api/spec.html", "text/html; charset=UTF-8", w)
}

func favicon(w http.ResponseWriter, _ *http.Request) {
	writeFile("api/polygon.png", "image/png", w)
}

func swagger(w http.ResponseWriter, _ *http.Request) {
	writeFile("api/api.yaml", "text/html; charset=UTF-8", w)
}

func writeFile(path string, mimeType string, w http.ResponseWriter) {
	f, err := os.ReadFile(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not found"))
	}
	w.Header().Set("Content-Type", mimeType)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(f)
}
