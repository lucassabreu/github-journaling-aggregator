package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <pemfile>", os.Args[0])
		return
	}

	client := NewGithubClient(os.Args[1], os.Getenv("GITHUB_TOKEN"))

	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
	repos, _, err := client.Repositories.List(context.Background(), "", opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repos)
}

func NewGithubClient(username string, token string) *github.Client {
	return github.NewClient(&http.Client{Transport: &TransportBasicAuth{
		tr:       http.DefaultTransport,
		username: username,
		token:    token,
	}})
}

type TransportBasicAuth struct {
	tr       http.RoundTripper
	username string
	token    string
}

func (ba *TransportBasicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(ba.username, ba.token)
	return ba.tr.RoundTrip(req)
}
