package ui

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
)

//go:embed dist
var f embed.FS

func New(useOS bool) http.FileSystem {
	if useOS {
		return http.FS(os.DirFS("ui/dist"))
	}

	fsys, err := fs.Sub(f, "dist")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
