package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

type S3Storage struct {
	env   *env.Environment
	minio *minio.Client
}

func NewS3(
	env *env.Environment,
) (*S3Storage, error) {
	if !env.UseBucket {
		return nil, nil
	}

	settings := &minio.Options{
		Creds:  credentials.NewStaticV4(env.BucketAccessKey, env.BucketSecretKey, env.BucketTokenKey),
		Secure: env.BucketUseSSL,
	}
	client, err := minio.New(env.BucketEndpoint, settings)

	if err != nil {
		return nil, err
	}

	return &S3Storage{
		minio: client,
		env:   env,
	}, nil
}
