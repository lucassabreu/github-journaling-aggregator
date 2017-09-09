package githubclient

import (
	"net/http"

	"github.com/google/go-github/github"
)

func NewGithubClient(username string, token string, tr http.RoundTripper) *github.Client {
	if tr == nil {
		tr = http.DefaultTransport
	}

	return github.NewClient(&http.Client{Transport: &transportBasicAuth{
		tr:       tr,
		username: username,
		token:    token,
	}})
}

type transportBasicAuth struct {
	tr       http.RoundTripper
	username string
	token    string
}

func (ba *transportBasicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(ba.username, ba.token)
	return ba.tr.RoundTrip(req)
}
