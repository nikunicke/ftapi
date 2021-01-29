package oauth2

import (
	"net/http"

	"golang.org/x/oauth2"
)

// Config generalizes oauth2 config types. Currently supporting:
// 	- oauth2.Config
// 	- oauth2/clientcredentials.Config
type Config interface {
	Token() (*oauth2.Token, error)
	Client(token *oauth2.Token) *http.Client
}
