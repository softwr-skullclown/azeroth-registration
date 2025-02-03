package email

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed templates
var f embed.FS

func EmbedTemplates(useOS bool) fs.FS {
	if useOS {
		return fs.FS(os.DirFS("templates"))
	}

	fsys, err := fs.Sub(f, "templates")
	if err != nil {
		panic(err)
	}

	return fs.FS(fsys)
}
