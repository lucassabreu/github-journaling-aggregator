package main

import (
	"context"
	"encoding/json"
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

	opt := &github.RepositoryListOptions{
		Sort:        "updated",
		Direction:   "desc",
		ListOptions: github.ListOptions{PerPage: 10},
	}
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(context.Background(), "", opt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(len(repos))

		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	json.NewEncoder(os.Stdout).Encode(&allRepos)

	// fmt.Println(repos)
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
