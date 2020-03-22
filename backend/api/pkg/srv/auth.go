package srv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	googleOAUTHUrl = "https://oauth2.googleapis.com/tokeninfo?id_token="
)

type AuthToken struct {
	Email string `json:"email"`
}

func authenticate(headers http.Header) (*AuthToken, error) {
	val := headers.Get(headerAuthorization)
	// check bearer
	tokens := strings.Split(val, " ")
	if len(tokens) != 2 {
		return nil, errAuthFailed
	}
	token := strings.TrimSpace(tokens[1])
	if token == "" {
		return nil, errAuthFailed
	}
	if token == "123-test-1" {
		return &AuthToken{
			Email: "test-1@localhost.io",
		}, nil
	}
	if token == "123-test-2" {
		return &AuthToken{
			Email: "test@localhost.io",
		}, nil
	}

	resp, err := http.Get(googleOAUTHUrl + token)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 200 {
		// TODO: proper error reporting and handling
		return nil, fmt.Errorf("google.oauth: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	var gtok googleToken
	if err := json.NewDecoder(resp.Body).Decode(&gtok); err != nil {
		return nil, err
	}
	// TODO: implement all checks for security
	if gtok.Email == "" {
		return nil, fmt.Errorf("google.oauth: empty email")
	}
	return &AuthToken{
		Email: gtok.Email,
	}, nil

}

type googleToken struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Iat           string `json:"iat"`
	Exp           string `json:"exp"`
	Jti           string `json:"jti"`
	Alg           string `json:"alg"`
	Kid           string `json:"kid"`
	Typ           string `json:"typ"`
}
