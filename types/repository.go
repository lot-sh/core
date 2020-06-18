package types

import (
	"io"
)

//Repository provides access to the resource storage
type Repository interface {
	// Get io.Reader from the resource storage by CID
	Get(scid string) (io.Reader, error)
	// Remove and existing resource using CID
	Remove(scid string) error
	// Add a resource with the CID as name generated from the content of io.Reader
	Add(reader io.Reader) (string, error)
}
