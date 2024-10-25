package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/config"
	"net/http"
	"net/url"
)

type GitHubOAuth struct {
	ClientID     string
	ClientSecret string
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type GitHubUser struct {
	ID        int64  `json:"id"`
	Username  string `json:"login"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

func NewGitHubOAuth(cfg *config.Config) *GitHubOAuth {
	return &GitHubOAuth{
		ClientID:     cfg.GitHubClientID,
		ClientSecret: cfg.GitHubClientSecret,
	}
}

func (gh *GitHubOAuth) ExchangeCode(ctx context.Context, code string) (*Token, error) {
	params := url.Values{}
	params.Add("client_id", gh.ClientID)
	params.Add("client_secret", gh.ClientSecret)
	params.Add("code", code)

	req, err := http.NewRequestWithContext(ctx, "POST", "https://github.com/login/oauth/access_token", nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var token Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (gh *GitHubOAuth) GetUser(ctx context.Context, token *Token) (*GitHubUser, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	if user.Email == "" {
		email, err := gh.getPrimaryEmail(ctx, token)
		if err != nil {
			return nil, err
		}
		user.Email = email
	}

	return &user, nil
}

func (gh *GitHubOAuth) getPrimaryEmail(ctx context.Context, token *Token) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var emails []struct {
		Email   string `json:"email"`
		Primary bool   `json:"primary"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
		return "", err
	}

	for _, e := range emails {
		if e.Primary {
			return e.Email, nil
		}
	}

	return "", fmt.Errorf("primary email not found")
}
