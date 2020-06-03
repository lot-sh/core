package core

import (
	"fmt"
	"io"
	"github.com/lot-sh/core/scheme"
)

// Resource struct is the data of the principal resource of
// this application, which is pieces of code and they origin
type Resource struct {
	Locator string
	Scheme  scheme.SchemeType
}

func (r *Resource) String() string {
	return fmt.Sprintf("Locator: %v\nScheme: %v", r.Locator, r.Scheme)
}

// FetchData return the data which this resource points to
func (r *Resource) FetchData() (io.ReadCloser, error) {
	resolver, err := ResolverFactory(r.Scheme)
	var readCloser io.ReadCloser
	if err != nil {
		return readCloser, err
	}
	return resolver.FetchData(r)
}
