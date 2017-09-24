// +build ignore

package main

import (
	"log"

	"github.com/lucassabreu/github-journaling-aggregator/formatter"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(formatter.Assets, vfsgen.Options{
		PackageName:  "formatter",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
