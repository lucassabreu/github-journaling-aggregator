// +build dev

package formatter

import (
	"go/build"
	"log"
	"net/http"
)

func importPathToDir(importPath string) string {
	p, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		log.Fatalln(err)
	}
	return p.Dir
}

// Assets contains project assets.
var Assets = http.Dir(importPathToDir("github.com/lucassabreu/github-journaling-aggregator/formatter/assets"))
