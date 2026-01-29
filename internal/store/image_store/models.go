package imagestore

import (
	"context"
	"io"
)

type ImageStore interface {
	Upload(ctx context.Context, key string, body io.ReadSeeker, size int64, contentType string) (string, error)
	Delete(ctx context.Context, key string) error
}
