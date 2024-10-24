package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	db "karadyaur.io/ai-dev-light/ad-oauth/internal/db/sqlc"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/model"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/utils"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id pgtype.UUID) error
	GetUserByGitHubID(ctx context.Context, githubID int64) (db.User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (db.User, error)
	UpdateUserToken(ctx context.Context, arg db.UpdateUserTokenParams) error
}

type UserService struct {
	userRepository IUserRepository
	GitHubAuth     *utils.GitHubOAuth
}

func NewUserService(userRepository IUserRepository, gitHubAuth *utils.GitHubOAuth) *UserService {
	return &UserService{
		userRepository: userRepository,
		GitHubAuth:     gitHubAuth,
	}
}

func (s *UserService) Authenticate(ctx context.Context, code string) (*model.User, error) {
	token, err := s.GitHubAuth.ExchangeCode(ctx, code)
	if err != nil {
		return nil, err
	}

	ghUser, err := s.GitHubAuth.GetUser(ctx, token)
	if err != nil {
		return nil, err
	}

	user, err := s.GetUserByGitHubID(ctx, ghUser.ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user = &model.User{
			GitHubID: ghUser.ID,
			Username: ghUser.Username,
			Email: pgtype.Text{
				String: ghUser.Email,
				Valid:  ghUser.Email != "",
			},
			Token: token.AccessToken,
		}
		err = s.CreateUser(ctx, user)
		if err != nil {
			return nil, err
		}
	} else {
		err = s.userRepository.UpdateUserToken(ctx, db.UpdateUserTokenParams{
			ID:    user.ID,
			Token: token.AccessToken,
		})
		if err != nil {
			return nil, err
		}
		user.Token = token.AccessToken
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id pgtype.UUID) (*model.User, error) {
	dbUser, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user := &model.User{
		ID:       dbUser.ID,
		GitHubID: dbUser.GithubID,
		Username: dbUser.Username,
		Email:    dbUser.Email,
		Token:    dbUser.Token,
	}

	return user, nil
}

func (s *UserService) GetUserByGitHubID(ctx context.Context, githubID int64) (*model.User, error) {
	dbUser, err := s.userRepository.GetUserByGitHubID(ctx, githubID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user := &model.User{
		ID:       dbUser.ID,
		GitHubID: dbUser.GithubID,
		Username: dbUser.Username,
		Email:    dbUser.Email,
		Token:    dbUser.Token,
	}

	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	createdUser, err := s.userRepository.CreateUser(ctx, db.CreateUserParams{
		GithubID: user.GitHubID,
		Username: user.Username,
		Email:    user.Email,
		Token:    user.Token,
	})
	if err != nil {
		return err
	}
	user.ID = createdUser.ID
	return nil
}
