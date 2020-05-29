package core

import "testing"

func TestResourceFactoryShouldWorkAsExpected(t *testing.T) {
	locator := "lot:QmT5NvUtoM5nWFfrQdVrFtvGfKFmG7AHE8P34isapyhCxX"
	resource, err := ResourceFactory(locator)
	if err != nil {
		t.Error(err)
	}
	if resource.Scheme != LOT {
		t.Errorf("the scheme of %s should be %s, found %s", locator, LOT, resource.Scheme)
	}
	if resource.Locator != locator {
		t.Errorf("the locator should be %s, found %s", locator, resource.Locator)
	}
}
