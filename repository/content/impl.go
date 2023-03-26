package content

import (
	"context"
	"net/url"

	"gitlab.informatika.org/ocw/ocw-backend/provider/storage"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

type ContentRepositoryImpl struct {
	storage.Storage
	*env.Environment
}

func (c ContentRepositoryImpl) GenerateNewLink(ctx context.Context, path string) (string, error) {
	return c.Storage.CreatePutSignedLink(ctx, path)
}

func (c ContentRepositoryImpl) GetLink(ctx context.Context, path string) (string, error) {
	res, err := c.Storage.CreateGetSignedLink(ctx, path, url.Values{})
	return res, err
}

func (c ContentRepositoryImpl) Delete(ctx context.Context, path string) error {
	return c.Storage.Delete(ctx, path)
}
