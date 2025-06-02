package service

import (
	"encoding/json"
	"net/http"

	smodel "github.com/jp/authentication/internal/service/model"
)

func (s *AuthService) SaveGoogleToken(googleToken smodel.GoogleToken) (string, error) {
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
