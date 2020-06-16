package core

import (
	"errors"

	"github.com/lot-sh/core/scheme"
)

// ResolverFactory returns the implementation of Resolver
// which can handle the scheme.Type passed as argument
//
// • when there is not a Resolver associated to the scheme.Type
// passed as argument it will return a nil Resolver and an error
func ResolverFactory(st scheme.Type) (Resolver, error) {
	var resolver Resolver
	switch st {
	case scheme.HTTP, scheme.HTTPS:
		resolver = &HTTPResolver{}
		return resolver, nil
	}
	return resolver, errors.New("There is not implementation known which supports the given scheme.Type")
}
