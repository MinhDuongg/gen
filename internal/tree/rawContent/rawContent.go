package rawContent

import "context"

type RawContent struct {
	Content []byte
}

func NewRawContent(content []byte) RawContent {
	return RawContent{content}
}

func (r RawContent) ContentAvailbility(ctx context.Context) bool {
	if len(r.Content) == 0 {
		return false
	}

	return true
}

func (r RawContent) ContentWriter(ctx context.Context) ([]byte, error) {
	return r.Content, nil
}


