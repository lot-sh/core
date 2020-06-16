package resource

import (
	"fmt"
	"strings"
	"testing"
)

var suitePassMsg = `
PASS: %s
-> Input: %s`

var suiteFailMsg = `
FAIL: %s
-> Input: %s
-> Actual %s
-> Expected: %s`

func TestLocatorShouldParseAndFormatWithoutErrors(t *testing.T) {
	type suiteExpected struct {
		scheme string
		path   string
		tag    string
	}
	suites := []struct {
		message  string
		input    string
		expected suiteExpected
	}{
		{
			"Regular http resource",
			"http://example.com/testfile",
			suiteExpected{
				"http", "//example.com/testfile", "",
			},
		}, {
			"Regular ftp resource",
			"ftp://example.com/testfile?test#foo",
			suiteExpected{
				"ftp", "//example.com/testfile?test#foo", "",
			},
		}, {
			"With tag",
			"lot:rubeniskov/semver-patcher@v1.0.1",
			suiteExpected{
				"lot", "rubeniskov/semver-patcher@v1.0.1", "v1.0.1",
			},
		}, {
			"With ID",
			"lot:QmT5NvUtoM5nWFfrQdVrFtvGfKFmG7AHE8P34isapyhCxX",
			suiteExpected{
				"lot", "QmT5NvUtoM5nWFfrQdVrFtvGfKFmG7AHE8P34isapyhCxX", "",
			},
		},
	}

	for _, suite := range suites {
		loc, err := NewLocator(suite.input)
		if err != nil {
			t.Error(err)
		}
		if loc.Scheme() != suite.expected.scheme {
			t.Errorf(
				suiteFailMsg,
				fmt.Sprintf("%s: Wrong expected scheme", suite.message),
				suite.input,
				loc.Scheme(),
				suite.expected.scheme,
			)
		}
		if loc.Path() != suite.expected.path {
			t.Errorf(
				suiteFailMsg,
				fmt.Sprintf("%s: Wrong expected path", suite.message),
				loc.Path(),
				suite.expected.path,
			)
		}
		if loc.Tag() != suite.expected.tag {
			t.Errorf(
				suiteFailMsg,
				fmt.Sprintf("%s: Wrong expected tag", suite.message),
				loc.Path(),
				suite.expected.tag,
			)
		}
		t.Logf(
			suitePassMsg,
			fmt.Sprintf(
				"%s should parse without errors (%s)",
				suite.message,
				loc,
			),
			suite.input,
		)
	}
}

func TestLocatorShouldRaiseErrorScheme(t *testing.T) {
	suites := []struct {
		message  string
		input    string
		expected string
	}{
		{
			"Missing scheme",
			"//example.com/testfile",
			"Failure schema detection when parsing //example.com/testfile",
		}, {
			"Wrong scheme",
			"scheme:",
			"Failure schema detection when parsing scheme:",
		}, {
			"Missing path",
			"ftp",
			"Malformatted locator missing path when parsing ftp",
		}, {
			"Wrong path",
			"ftp:test:",
			"Malformatted locator when parsing ftp:test:",
		},
	}

	for _, suite := range suites {
		_, err := NewLocator(suite.input)
		if err == nil {
			t.Error(err)
		}
		if !strings.Contains(err.Error(), suite.expected) {
			t.Errorf(suiteFailMsg, suite.message, suite.input, err.Error(), suite.expected)
		}
		t.Logf(
			suitePassMsg,
			fmt.Sprintf(
				"%s should raise error with the espected message (%s)",
				suite.message,
				err.Error(),
			),
			suite.input,
		)
	}
}
