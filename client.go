package ftapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nikunicke/ftapi/oauth2"
)

// Service includes an http client with authorization to communicate with the
// 42 API.
type Service struct {
	c       *http.Client
	baseURL string
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

// NewService returns an Service or an error. The grant type must be specified:
// 	- ClientCredentials
// 	- AuthorizationCode: default
func NewService(gt GrantType) (*Service, error) {
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

// Me gets the current users userdata. Exactly one of *UsersItem or error will
// be non-nil. Any non 2xx status code is an error
func (s *Service) Me() (*UsersItem, error) {
	urls := s.baseURL + "/me"
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := s.Do(req)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	ret := &UsersItem{}
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Do commits a request to the URL specified in the request
func (s *Service) Do(req *http.Request) (*http.Response, error) {
	resp, err := s.c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func client(config oauth2.Config) (*Service, error) {
	token, err := config.Token()
	if err != nil {
		return nil, err
	}
	return &Service{c: config.Client(token), baseURL: "https://api.intra.42.fr/v2"}, nil
}
