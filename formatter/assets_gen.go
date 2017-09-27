// +build ignore

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lucassabreu/github-journaling-aggregator/formatter"
	"github.com/shurcooL/vfsgen"
)

type modTimeFS struct {
	fs http.FileSystem
}

func (fs modTimeFS) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return f, err
	}

	return modTimeFile{f}, nil
}

type modTimeFile struct {
	http.File
}

func (f modTimeFile) Stat() (os.FileInfo, error) {
	fi, err := f.File.Stat()
	if err != nil {
		return nil, err
	}

	return modTimeFileInfo{fi}, nil
}

type modTimeFileInfo struct {
	os.FileInfo
}

func (fi modTimeFileInfo) ModTime() time.Time {
	return time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
}

func main() {
	err := vfsgen.Generate(modTimeFS{formatter.Assets}, vfsgen.Options{
		PackageName:  "formatter",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
