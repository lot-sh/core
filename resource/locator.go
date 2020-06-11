package resource

import (
	"fmt"
	"strings"

	"github.com/lot-sh/core/scheme"
)

// Locator represents the minimal locator functionality
// For more details visit the RFC
// https://github.com/lot-sh/docs/blob/master/rfcs/002-resource-locator.md
type Locator struct {
	scheme scheme.SchemeType
	path   string
}

// NewLocator instance a locator from a given string which parse
// and ensure if the string is well formated
func NewLocator(strloc string) (*Locator, error) {
	res := Locator{}
	parts := strings.Split(strloc, scheme.SEPARATOR)
	res.scheme = scheme.GetSchemeTypeFrom(parts[0])

	if res.scheme == scheme.UNKNOWN {
		return nil, fmt.Errorf("Failure schema detection when parsing %s", strloc)
	}

	partsLen := len(parts)
	if partsLen != 2 {
		if partsLen == 1 {
			return nil, fmt.Errorf("Malformatted locator missing path when parsing %s", strloc)
		}

		return nil, fmt.Errorf("Malformatted locator when parsing %s", strloc)
	}

	res.path = parts[1]

	return &res, nil
}

// Scheme gets the scheme type with string format
func (l *Locator) Scheme() string {
	return l.scheme.String()
}

// Path returns the part of the locator wich contains
// the identification of a resource under the specified scheme
func (l *Locator) Path() string {
	return l.path
}

// Tag returns the name of version of a resource
func (l *Locator) Tag() string {
	parts := strings.Split(l.path, "@")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

// String format representation
func (l *Locator) String() string {
	s := []string{l.Scheme(), l.Path()}
	return strings.Join(s, scheme.SEPARATOR)
}
