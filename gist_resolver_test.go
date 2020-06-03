package core

import (
	"regexp"
	"testing"
)

func TestGistLocatorRegexpShouldBeWellFormed(t *testing.T) {
	re := regexp.MustCompile(GistLocatorRegexp)
	locator := "gist://Kelvur/multiply.py"
	username := "Kelvur"
	resource := "multiply.py"
	if !re.MatchString(locator) {
		t.Errorf("GistLocatorRegexp doesn't match the locator: %s", locator)
	}
	listMatches := re.FindAllStringSubmatch("gist://Kelvur/multiply.py", -1)[0]
	if listMatches[1] != "Kelvur" {
		t.Errorf("GistLocatorRegexp doesn't match correctly the username, expected %s get %s", username, listMatches[1])
	}
	if listMatches[2] != "multiply.py" {
		t.Errorf("GistLocatorRegexp doesn't match correctly the resource, expected %s get %s", resource, listMatches[2])
	}
}

func TestObtainUsernameAndResourceFromGistLocatorShouldWorkAsExpected(t *testing.T) {
	gistLocator := "gist://Kelvur/multiply.py"
	username := "Kelvur"
	resource := "multiply.py"
	values, err := ObtainUsernameAndResourceFromGistLocator(gistLocator)
	if err != nil {
		t.Errorf("ObtainUsernameAndResourceFromGistLocator should work with a valid gist locator as %s", gistLocator)
	}
	if values[0] != username {
		t.Errorf("after calling ObtainUsernameAndResourceFromGistLocator with %s the expected fist value is %s, get %s", gistLocator, username, values[0])
	}
	if values[1] != resource {
		t.Errorf("after calling ObtainUsernameAndResourceFromGistLocator with %s the expected fist value is %s, get %s", gistLocator, resource, values[1])
	}
}

func TestObtainUsernameAndResourceFromGistLocatorShouldFailAsExpected(t *testing.T) {
	incorrectGistLocator := "gist:/-asd/<>"
	_, err := ObtainUsernameAndResourceFromGistLocator(incorrectGistLocator)
	if err == nil {
		t.Errorf("ObtainUsernameAndResourceFromGistLocator should fail when called with %s", incorrectGistLocator)
	}
}

func TestGistResolverImplementsResolverInterface(t *testing.T) {
	var resolver Resolver = &GistResolver{}
	_, ok := resolver.(*GistResolver)
	if !ok {
		t.Error("GistResolver not implements Resolver")
	}
}

func TestGistResolverShouldBeAbleToFetchAGist(t *testing.T) {
	gistLocator := "gist://Kelvur/multiplication.py"
	resource, err := ResourceFactory(gistLocator)
	if err != nil {
		t.Errorf("ResourceFactory fails with locator %s", gistLocator)
		t.Error(err)
	}
	var resolver GistResolver = GistResolver{}
	readCloser, err := resolver.FetchData(&resource)
	defer readCloser.Close()
	if err != nil {
		t.Error(err)
	}
}
