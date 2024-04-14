package songs

import (
	"io/fs"
)

type Song struct {
	Info fs.FileInfo
	Path string
}
