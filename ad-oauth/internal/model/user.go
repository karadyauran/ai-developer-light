package model

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	ID       pgtype.UUID `json:"id"`
	GitHubID int64       `json:"github_id"`
	Username string      `json:"username"`
	Email    pgtype.Text `json:"email"`
	Token    string      `json:"token"`
}
