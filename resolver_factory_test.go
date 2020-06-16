package core

import (
	"testing"
	"github.com/lot-sh/core/scheme"
)

func TestResolverFactoryShouldWorksWhenPassedHTTPType(t *testing.T) {
	_, err := ResolverFactory(scheme.HTTP)
	if err != nil {
		t.Error("Error should be no returned by ResolverFactory when passed a HTTP Type")
	}
}

func TestResolverFactoryShouldReturnErrorWhenPassedUNKNOWNType(t *testing.T) {
	_, err := ResolverFactory(scheme.UNKNOWN)
	if err == nil {
		t.Error("Error should be returned by ResolverFactory when passed a UNKNONW Type")
	}
}
