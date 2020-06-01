package core

import (
	"io"
	"net/http"
)

// HTTPResolver is a HTTP/S client whose purpose is fetch data
// using the HTTP protocol with a simple API
//
// • It implements the infertace Resolver
type HTTPResolver struct {
}

// FetchData function fetch the code where the resource Locator
// points to
func (resolver *HTTPResolver) FetchData(resource *Resource) (io.ReadCloser, error) {
	res, err := http.Get(resource.Locator)
	return res.Body, err
}
