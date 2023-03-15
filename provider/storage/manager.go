package storage

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (s S3Storage) Delete(ctx context.Context, path string) error {
	return s.minio.RemoveObject(ctx, s.env.BucketName, path, minio.RemoveObjectOptions{})
}
