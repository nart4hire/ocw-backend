package storage

import (
	"bufio"
	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

func (s S3Storage) Delete(ctx context.Context, path string) error {
	return s.minio.RemoveObject(ctx, s.env.BucketName, path, minio.RemoveObjectOptions{})
}

func (s S3Storage) Get(ctx context.Context, path string) ([]byte, error) {
	result := []byte{}
	obj, err := s.minio.GetObject(ctx, s.env.BucketName, path, minio.GetObjectOptions{})

	if err != nil {
		return result, err
	}

	var buffer bytes.Buffer
	writter := bufio.NewWriter(&buffer)

	_, err = io.Copy(writter, obj)
	if err != nil {
		return result, err
	}

	result = buffer.Bytes()

	return result, err
}
