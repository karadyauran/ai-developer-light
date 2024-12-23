// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        pgtype.UUID        `json:"id"`
	GithubID  int64              `json:"github_id"`
	Username  string             `json:"username"`
	AvatarUrl pgtype.Text        `json:"avatar_url"`
	Email     pgtype.Text        `json:"email"`
	Token     string             `json:"token"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
