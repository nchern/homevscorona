package restapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
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
	return fmt.Errorf("%w: '%s' not found; signup required", errAuthFailed, t.GetEmail())
}

func authenticate(headers http.Header) (Token, error) {
	val := headers.Get(headerAuthorization)

	token := ""
	if _, err := fmt.Sscanf(val, "Bearer %s", &token); err != nil {
		return nil, err
	}

	token = strings.TrimSpace(token)

	if token == "" {
		return nil, fmt.Errorf("%w: empty token", errValidationFailed)
	}

	res, err := verify(token)
	if err != nil {
		return nil, err
	}

	if res.GetEmail() == "" {
		return nil, fmt.Errorf("oauth bad response: empty email")
	}
	return res, nil

}

func verify(token string) (Token, error) {
	segments := strings.Split(token, ".")
	if len(segments) != 3 {
		return nil, fmt.Errorf("%w: bad jwt format", errValidationFailed)
	}

	var tok jwtToken
	if err := tok.FromB64String(segments[1]); err != nil {
		return nil, err
	}

	if tok.Iss == "test" {
		// TOOD: add signature using HMAC
		return &tok, nil
	}

	res, err := doGoogleAuth(token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func doGoogleAuth(token string) (Token, error) {
	resp, err := http.Get(googleOAUTHUrl + token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var gtok jwtToken
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

type jwtToken struct {

	// Google error fields
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`

	// public claims
	Exp string `json:"exp"`
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Iat string `json:"iat"`
	Jti string `json:"jti"`
	Alg string `json:"alg"`

	// normal fields
	Email string `json:"email"`

	Azp           string `json:"azp"`
	EmailVerified string `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Kid           string `json:"kid"`
	Typ           string `json:"typ"`
}

func (t *jwtToken) FromB64String(s string) error {
	b, err := jwt.DecodeSegment(s)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &t)
}

func (t *jwtToken) GetEmail() string {
	return t.Email
}
