// +build dev

package formatter

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("assets")
