package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	db "karadyaur.io/ai-dev-light/ad-oauth/internal/db/sqlc"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/utils"
)

type Service struct {
	OAuthService *OAuthService
}

func NewService(pool *pgxpool.Pool, gitHubAuth *utils.GitHubOAuth) *Service {
	queries := db.New(pool)
	return &Service{
		OAuthService: NewOAuthService(queries, gitHubAuth),
	}
}
