package core

import "testing"

func TestResolverFactoryShouldWorksWhenPassedHTTPSchemeType(t *testing.T) {
	_, err := ResolverFactory(HTTP)
	if err != nil {
		t.Error("Error should be no returned by ResolverFactory when passed a HTTP SchemeType")
	}
}

func TestResolverFactoryShouldReturnErrorWhenPassedUNKNOWNSchemeType(t *testing.T) {
	_, err := ResolverFactory(UNKNOWN)
	if err == nil {
		t.Error("Error should be returned by ResolverFactory when passed a UNKNONW SchemeType")
	}
}
