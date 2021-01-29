package oauth2

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// ClientCredentials represent the config type for the authorization code grant
// type
type ClientCredentials struct {
	B []byte
	c *clientcredentials.Config
}

// Token retrieves an Oauth2 token based on the credentials retrieved from a file
func (a *ClientCredentials) Token() (*oauth2.Token, error) {
	var err error
	a.c, err = configClientCredentials(a.B)
	if err != nil {
		return nil, err
	}
	return a.c.Token(context.TODO())
}

// Client returns an HTTP client using the provided token. The token will
// auto-refresh as necessary. The returned client should not be modified
func (a *ClientCredentials) Client(token *oauth2.Token) *http.Client {
	return a.c.Client(context.Background())
}
func configClientCredentials(b []byte) (*clientcredentials.Config, error) {
	type cred struct {
		ClientID     string   `json:"client_id"`
		ClientSecret string   `json:"client_secret"`
		TokenURI     string   `json:"token_uri"`
		Scopes       []string `json:"scopes"`
	}
	var c cred
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return &clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     c.TokenURI,
		Scopes:       c.Scopes,
	}, nil
}
