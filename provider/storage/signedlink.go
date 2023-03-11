package storage

import (
	"context"
	"net/url"
	"time"
)

func (s S3Storage) CreatePutSignedLink(ctx context.Context, path string) (string, error) {
	url, err := s.minio.PresignedPutObject(
		ctx,
		s.env.BucketName,
		path,
		time.Duration(s.env.BucketSignedPutDuration)*time.Second,
	)

	if err != nil {
		return "", err
	}

	return url.String(), nil
}

func (s S3Storage) CreateGetSignedLink(ctx context.Context, path string, reqParam url.Values) (string, error) {
	url, err := s.minio.PresignedGetObject(
		ctx,
		s.env.BucketName,
		path,
		time.Duration(s.env.BucketSignedGetDuration)*time.Second,
		reqParam,
	)

	if err != nil {
		return "", err
	}

	return url.String(), nil
}
