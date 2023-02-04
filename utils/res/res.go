package res

type Resource interface {
	GetBytesResource(path string) ([]byte, error)
	GetStringResource(path string) (string, error)
}
