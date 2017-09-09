package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/lucassabreu/github-journaling-aggregator/githubclient"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s username <since>", os.Args[0])
		return
	}

	username := os.Args[1]
	token := os.Getenv("GITHUB_TOKEN")

	client := githubclient.NewGithubClient(username, os.Getenv("GITHUB_TOKEN"))

	// opt := &github.RepositoryListOptions{
	// 	Sort:        "updated",
	// 	Direction:   "desc",
	// 	ListOptions: github.ListOptions{PerPage: 10},
	// }

	opt := &github.ListOptions{
		PerPage: 10,
	}
	events, _, err := client.Activity.ListEventsPerformedByUser(context.Background(), username, false, opt)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(os.Stdout).Encode(&events)

	// var allRepos []*github.Repository
	// for {
	// 	repos, resp, err := client.Repositories.List(context.Background(), "", opt)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(len(repos))

	// 	allRepos = append(allRepos, repos...)
	// 	if resp.NextPage == 0 {
	// 		break
	// 	}
	// 	opt.Page = resp.NextPage
	// }
	// json.NewEncoder(os.Stdout).Encode(&allRepos)
}
