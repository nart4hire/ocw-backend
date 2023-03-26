package res

import "io/fs"

type Resource interface {
	GetBytesResource(path string) ([]byte, error)
	GetStringResource(path string) (string, error)
	GetFile(path string) (fs.File, error)
}
