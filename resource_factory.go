package core

import "errors"

// ResourceFactory function returns a Resource given a locator
func ResourceFactory(locator string) (Resource, error) {
	var scheme SchemeType = GetSchemeTypeFrom(locator)
	var resource Resource = Resource{
		locator,
		scheme,
	}
	if scheme == UNKNOWN {
		return resource, errors.New("Unknown scheme, the locator may be a invalid one")
	}
	return resource, nil
}
