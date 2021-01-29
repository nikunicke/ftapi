package oauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/xerrors"
)

// AuthorizationCode represent the config type for the authorization code grant
// type
type AuthorizationCode struct {
	B []byte
	c *oauth2.Config
}

// Token generates a *oauth2.Token. The token will be cached in token.json.
// Refreshtoken might not work properly.
func (a *AuthorizationCode) Token() (*oauth2.Token, error) {
	tokenFile := "token.json"
	token, err := getTokenFromFile(tokenFile)
	if err != nil {
		a.c, err = configAuthorizationCode(a.B)
		if err != nil {
			return nil, err
		}
		token, err = getTokenFromWeb(a.c)
		if err != nil {
			return nil, err
		}
		if err := saveToken(tokenFile, token); err != nil {
			return nil, err
		}
	}
	return token, nil
}

// Client returns an HTTP client using the provided token. The token will
// auto-refresh as necessary. The returned client should not be modified.
func (a *AuthorizationCode) Client(token *oauth2.Token) *http.Client {
	return a.c.Client(context.Background(), token)
}
func configAuthorizationCode(b []byte) (*oauth2.Config, error) {
	type cred struct {
		ClientID     string   `json:"client_id"`
		ClientSecret string   `json:"client_secret"`
		RedirectURIs []string `json:"redirect_uris"`
		AuthURI      string   `json:"auth_uri"`
		TokenURI     string   `json:"token_uri"`
		Scopes       []string `json:"scopes"`
	}
	var c cred
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	if len(c.RedirectURIs) < 1 {
		return nil, xerrors.New("Missing redirect URL.")
	}
	return &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURIs[0],
		Scopes:       c.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.AuthURI,
			TokenURL: c.TokenURI,
		},
	}, nil
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link and type in the "+
		"authorization code: \n%v\n", authURL)
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, err
	}
	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func getTokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

func saveToken(path string, token *oauth2.Token) error {
	fmt.Println("Saving credential file to:", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}
