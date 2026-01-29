package imagestore

import (
	"context"
	"errors"
	"fmt"
	"io"
	"main/internal/exceptions"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2Store struct {
	client    *s3.Client
	bucket    string
	publicURL string
}

type R2Config struct {
	AccountID       string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	PublicBaseURL   string
}

func NewR2Store(ctx context.Context, cfg R2Config) (ImageStore, error) {
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.AccountID)

	awsCfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("auto"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, "")),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	})

	return &R2Store{
		client:    client,
		bucket:    cfg.Bucket,
		publicURL: strings.TrimRight(cfg.PublicBaseURL, "/"),
	}, nil
}

func (s *R2Store) Upload(ctx context.Context, key string, body io.ReadSeeker, size int64, contentType string) (string, error) {
    if _, err := body.Seek(0, 0); err != nil {
        return "", err
    }

    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    _, err := s.client.PutObject(ctx, &s3.PutObjectInput{
        Bucket:        aws.String(s.bucket),
        Key:           aws.String(key),
        Body:          body,
        ContentType:   aws.String(contentType),
        ContentLength: aws.Int64(size),
    })
    if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return "", exceptions.ErrTimeoutExceeded
		}
        return "", err
    }

    if s.publicURL == "" {
        return "", nil
    }
    return s.publicURL + "/" + key, nil
}

func (s *R2Store) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}
