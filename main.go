package main

import (
	"fmt"
	"os"

	"github.com/lucassabreu/github-journaling-aggregator/github"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <username>", os.Args[0])
		return
	}

	api := github.NewAPI(os.Args[1], os.Getenv("GITHUB_TOKEN"))
	_, err := api.GetUser()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
}
