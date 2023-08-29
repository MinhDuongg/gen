package tree

import "context"

type ContentI interface {
	ContentAvailbility(ctx context.Context) bool
	ContentWriter(ctx context.Context) ([]byte, error)
}
