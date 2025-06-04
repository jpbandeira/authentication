package domain

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleToken struct {
	ID           string
	UserEmail    string
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (s *domain) SaveGoogleToken(googleToken GoogleToken) (string, error) {
	return "", nil
}

func FetchUserEmail(client *http.Client) (string, error) {
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Email string `json:"email"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result.Email, err
}

func (d *domain) SaveToken(ctx context.Context, code string) (User, error) {
	var googleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://local.fidelity.com:8083/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/calendar.readonly",
			"https://www.googleapis.com/auth/userinfo.email",
			"openid",
		},
		Endpoint: google.Endpoint,
	}

	token, err := googleOAuthConfig.Exchange(ctx, code)
	if err != nil {
		return User{}, err
	}

	client := googleOAuthConfig.Client(ctx, token)
	email, err := FetchUserEmail(client)
	if err != nil {
		return User{}, err
	}

	user, err := d.db.FindByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	err = d.db.SaveGoogleToken(ctx, GoogleToken{
		ID:           user.ID,
		UserEmail:    user.Email,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	})
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (d *domain) IsFirstLogin(ctx context.Context, email string) bool {
	return d.db.GetGoogleToken(ctx, email) != nil
}
