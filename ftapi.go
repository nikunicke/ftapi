package ftapi

import (
	"io/ioutil"
	"net/http"

	"github.com/nikunicke/ftapi/oauth2"
)

// FtClient includes an http client with authorization to communicate with the
// 42 API.
type FtClient struct {
	c *http.Client
}

// GrantType specifies the gran type to be used for authentication
type GrantType int

const (
	// ClientCredentials grant type is used by client to obtain an access token
	// outside of the context of a user. This is typically used by clients to
	// access resources about themselves rather than to access a users's
	// resources.
	ClientCredentials GrantType = iota
	// AuthorizationCode grant type is used by confidential and public client to
	// exchange an authorization code for an access token. After the user
	// returns to the client vie the redirect URL, the application will get the
	// authorization code from the URL and use it to request an access token.
	AuthorizationCode
)

// Client returns an FtClient or an error. The grant type must be specified:
// 	- ClientCredentials
// 	- AuthorizationCode: default
func Client(gt GrantType) (*FtClient, error) {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}
	switch gt {
	case ClientCredentials:
		return client(&oauth2.ClientCredentials{B: b})
	default:
		return client(&oauth2.AuthorizationCode{B: b})
	}
}

// Do commits a request to the URL specified in the request
func (c *FtClient) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func client(config oauth2.Config) (*FtClient, error) {
	token, err := config.Token()
	if err != nil {
		return nil, err
	}
	return &FtClient{c: config.Client(token)}, nil
}
