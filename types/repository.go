package types

import (
	"io"
	"context"
)

//Repository provides access to the resource storage
type Repository interface {
	// Get io.Reader from the resource storage by CID
	Get(ctx context.Context, scid string) (io.Reader, error)
	// Remove and existing resource using CID
	Remove(ctx context.Context, scid string) error
	// Add a resource with the CID as name generated from the content of io.Reader
	Add(ctx context.Context, reader io.Reader) (string, error)
}