package core

import "errors"

// ResolverFactory returns the implementation of Resolver
// which can handle the SchemeType passed as argument
//
// • when there is not a Resolver associated to the SchemeType
// passed as argument it will return a nil Resolver and an error
func ResolverFactory(st SchemeType) (Resolver, error) {
	var resolver Resolver
	switch st {
	case HTTP, HTTPS:
		resolver = &HTTPResolver{}
		return resolver, nil
	}
	return resolver, errors.New("There is not implementation known which supports the given SchemeType")
}
