package model

type AuthenticateUserRequestSwagger struct {
	Code string `json:"code" binding:"required"`
}

type UserResponseSwagger struct {
	ID        string `json:"id"`
	GitHubID  int64  `json:"github_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	AvatarURL string `json:"avatar_url"`
}
