package model

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	ID        pgtype.UUID `json:"id"`
	GitHubID  int64       `json:"github_id"`
	AvatarURL pgtype.Text `json:"avatar_url"`
	Username  string      `json:"username"`
	Email     pgtype.Text `json:"email"`
	Token     string      `json:"token"`
}
