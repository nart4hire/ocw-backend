package storage

import (
	"context"
	"net/url"
)

type Storage interface {
	CreatePutSignedLink(ctx context.Context, path string) (string, error)
	CreateGetSignedLink(ctx context.Context, path string, reqParam url.Values) (string, error)
	Delete(ctx context.Context, path string) error
}
