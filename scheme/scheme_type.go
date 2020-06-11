package scheme

import "regexp"

// SchemeType type
type SchemeType int

// Scheme types supported by this program. Also
// UNKNOWN is defined to be used when the scheme is unknown
// https://tools.ietf.org/html/rfc3986#section-3
const (
	FTP SchemeType = iota
	GIST
	HTTP
	HTTPS
	IPFS
	LOT
	UNKNOWN
)

// SEPARATOR defines where the scheme definition
// ends in a formated string
const SEPARATOR = ":"

// ListNames string naming mapping for each SchemeType
var ListNames = []string{
	"ftp",
	"gist",
	"http",
	"https",
	"ipfs",
	"lot",
}

func (st SchemeType) String() string {
	if st < 0 || st >= SchemeType(len(ListNames)) {
		return "unknown"
	}
	return ListNames[st]
}

// GetSchemeTypeFrom try to identify the scheme of the locator
// and returns the SchemeType which correspond, can return the
// SchemeType.UNKNOWN in case it cannot identify the scheme
func GetSchemeTypeFrom(locator string) SchemeType {
	re := regexp.MustCompile(`^[a-z]*`)
	switch string(re.Find([]byte(locator))) {
	case "ftp":
		return FTP
	case "gist":
		return GIST
	case "http":
		return HTTP
	case "https":
		return HTTPS
	case "ipfs":
		return IPFS
	case "lot":
		return LOT
	default:
		return UNKNOWN
	}
}
