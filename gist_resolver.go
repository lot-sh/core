package core

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

// GistResolver is able to understand gist locators and fetch
// them returning a stream
//
// • It implements the infertace Resolver
type GistResolver struct {
}

// UserGistFile represents a gist file in a Gisthub response
type UserGistFile struct {
	Filename string
	Raw_URL  string
}

// UserGistsResponse represents the response of Github when
// fetching all the gist of a user
type UserGistsResponse struct {
	ID    string
	Files map[string]UserGistFile
}

// GithubAPIURL constant reference the base url of Github API
var GithubAPIURL string = "https://api.github.com"

// GistLocatorRegexp regexp to validate and match the parts of the gist locator
var GistLocatorRegexp string = `^[a-zA-Z0-9]+://(?P<username>[a-zA-Z0-9_]+[a-zA-Z0-9])/(?P<resource>[^/<>|:&]+)$`

// ObtainUsernameAndResourceFromGistLocator return an array
// with the username and resource of a gist locator
func ObtainUsernameAndResourceFromGistLocator(locator string) ([2]string, error) {
	re := regexp.MustCompile(GistLocatorRegexp)
	if !re.MatchString(locator) {
		return [2]string{}, fmt.Errorf("The locator %s doesn't match a well formed gist locator", locator)
	}
	matchList := re.FindAllStringSubmatch(locator, -1)[0]
	return [2]string{matchList[1], matchList[2]}, nil
}

// GetUserGistURL builds a URL which points to all the gist of
// a user
func GetUserGistURL(username string) string {
	return fmt.Sprintf("%s/users/%s/gists", GithubAPIURL, username)
}

// FetchData function fetch the code where the resource Locator
// points to, in this case the resource points to a gist
func (resolver *GistResolver) FetchData(resource *Resource) (io.ReadCloser, error) {
	values, err := ObtainUsernameAndResourceFromGistLocator(resource.Locator)
	if err != nil {
		return nil, err
	}
	url := GetUserGistURL(values[0])
	// Fetch all the gists of the user
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch the url %s %v", url, err)
	}
	// Defer the close of the gists stream
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	// Decode the response
	var gistList []UserGistsResponse

	if err = json.Unmarshal(body, &gistList); err != nil {
		return nil, fmt.Errorf("Unable to decode the request to %s %v", url, err)
	}
	// Look for the resource in all the gist
	var found UserGistFile = UserGistFile{"", ""}
	for idxGist := range gistList {
		for filename, fileList := range gistList[idxGist].Files {
			// Check if the filename is equals to the resource ol the locator
			if filename == values[1] {
				found = fileList
				break
			}
		}
	}
	if found.Filename == "" {
		return nil, fmt.Errorf("%s not found for the user %s", values[1], values[0])
	}
	// Leave the fetch to HTTPResolver
	httpResolver := &HTTPResolver{}
	return httpResolver.FetchData(&Resource{
		found.Raw_URL,
		HTTPS,
	})
}
