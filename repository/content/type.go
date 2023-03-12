package content

import "context"

type ContentRepository interface {
	GenerateNewLink(ctx context.Context, path string) (string, error)
	GetLink(ctx context.Context, path string) (string, error)
	Delete(ctx context.Context, path string) error
}
