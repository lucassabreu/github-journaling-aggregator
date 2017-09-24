package main

//go:generate go run -tags=dev ./generate_assets/assets_gen.go

import (
	"github.com/lucassabreu/github-journaling-aggregator/cmd"
)

func main() {
	cmd.Execute()
}
