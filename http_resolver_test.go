package core

import (
	"testing"
)

func TestHTTPResolverImplementsResolverInterface(t *testing.T) {
	var resolver Resolver = &HTTPResolver{}
	_, ok := resolver.(*HTTPResolver)
	if !ok {
		t.Error("HTTPResolver not implements Resolver")
	}
}

func TestHTTPResolverShouldBeAbleToFetchAURL(t *testing.T) {
	locator := "https://gist.githubusercontent.com/Kelvur/896265085d32db2c3ab065ea0995b0a3/raw/62df0ad25eed20cdfc27235e8c39cbbbdf967ed3/2020_content.md"
	resource, err := ResourceFactory(locator)
	if err != nil {
		t.Errorf("ResourceFactory fails with locator %s", locator)
		t.Error(err)
	}
	var resolver HTTPResolver = HTTPResolver{}
	readCloser, err := resolver.FetchData(&resource)
	if err != nil {
		t.Error(err)
	}
	readCloser.Close()
}
