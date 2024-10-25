package server

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/generated"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/service"
)

type AuthServer struct {
	generated.UnimplementedOAuthServiceServer
	authService *service.OAuthService
}

func NewAuthServer(service *service.Service) *AuthServer {
	return &AuthServer{
		authService: service.OAuthService,
	}
}

func (s *AuthServer) Authenticate(ctx context.Context, request *generated.AuthenticateUserRequest) (*generated.UserResponse, error) {
	authenticate, err := s.authService.Authenticate(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	idStr, err := UUIDToString(authenticate.ID)
	if err != nil {
		return nil, err
	}

	emailValue := TextToStringValue(authenticate.Email)

	response := &generated.UserResponse{
		Id:       idStr,
		GithubId: authenticate.GitHubID,
		Username: authenticate.Username,
		Email:    emailValue,
		Token:    authenticate.Token,
	}

	return response, nil
}

func UUIDToString(u pgtype.UUID) (string, error) {
	if !u.Valid {
		return "", fmt.Errorf("UUID is not valid")
	}
	uuidValue, err := uuid.FromBytes(u.Bytes[:])
	if err != nil {
		return "", err
	}
	return uuidValue.String(), nil
}

func TextToStringValue(t pgtype.Text) *wrapperspb.StringValue {
	if t.Valid {
		return &wrapperspb.StringValue{Value: t.String}
	}
	return nil
}
