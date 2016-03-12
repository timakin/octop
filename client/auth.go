package client

import (
	"net/http"

	"golang.org/x/oauth2"
)

func Authenticate(token string) *http.Client {
	TokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	TokenClient := oauth2.NewClient(oauth2.NoContext, TokenSource)
	return TokenClient
}
