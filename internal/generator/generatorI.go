package generator

import "context"

type Generator interface {
	Generate(ctx context.Context, destination string) error
}
