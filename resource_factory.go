package core

import (
	"errors"
	"github.com/lot-sh/core/scheme"
)

// ResourceFactory function returns a Resource given a locator
func ResourceFactory(locator string) (Resource, error) {
	var sch scheme.SchemeType = scheme.GetSchemeTypeFrom(locator)
	var resource Resource = Resource{
		locator,
		sch,
	}
	if sch == scheme.UNKNOWN {
		return resource, errors.New("Unknown scheme, the locator may be a invalid one")
	}
	return resource, nil
}
