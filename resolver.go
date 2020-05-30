package core

import "io"

// Resolver is an interface representing the abitily
// to fetch data without specify the source
type Resolver interface {
	FetchData(*Resource) (io.ReadCloser, error)
}
