package srv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	googleOAUTHUrl = "https://oauth2.googleapis.com/tokeninfo?id_token="
)

var (
	errAuthFailed = errors.New("auth failed")
)

type Token interface {
	GetEmail() string
}

func errUnknownUserToken(t Token) error {
	return fmt.Errorf("%w: %s", errNotFound, t.GetEmail())
}

func authenticate(headers http.Header) (Token, error) {
	val := headers.Get(headerAuthorization)

	tokens := strings.Split(val, " ")
	if len(tokens) != 2 {
		return nil, errAuthFailed
	}
	// TODO: check bearer?
	token := strings.TrimSpace(tokens[1])
	if token == "" {
		return nil, fmt.Errorf("%w: empty token", errAuthFailed)
	}

	res, err := doDebugAuth(token)
	if err != nil {
		return nil, err
	} else if res != nil {
		return res, nil
	}

	res, err = doGoogleAuth(token)
	if err != nil {
		return nil, err
	}

	if res.GetEmail() == "" {
		return nil, fmt.Errorf("oauth bad response: empty email")
	}
	return res, nil

}

func doDebugAuth(token string) (Token, error) {
	if token == "123-test-1" {
		return &googleToken{
			Email: "test-1@localhost.io",
		}, nil
	}
	if token == "123-test-2" {
		return &googleToken{
			Email: "test@localhost.io",
		}, nil
	}
	return nil, nil
}

func doGoogleAuth(token string) (Token, error) {
	resp, err := http.Get(googleOAUTHUrl + token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var gtok googleToken
	if resp.StatusCode != 200 {
		if resp.StatusCode < 500 {
			if err := json.NewDecoder(resp.Body).Decode(&gtok); err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("%w: google.oauth: status: %d; error: '%s'; description: '%s';",
				errAuthFailed, resp.StatusCode, gtok.Error, gtok.ErrorDescription)
		}
		raw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		// 503?
		return nil, fmt.Errorf("google.oauth status: %d; raw response: %s", resp.StatusCode, string(raw))
	}

	if err := json.NewDecoder(resp.Body).Decode(&gtok); err != nil {
		return nil, err
	}

	// TODO: implement all checks for security: expiration, etc.

	return &gtok, nil
}

type googleToken struct {
	// error fields
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`

	// normal fields
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

func (t *googleToken) GetEmail() string {
	return t.Email
}
