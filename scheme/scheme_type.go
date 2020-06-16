package scheme

import "regexp"

// Type type
type Type int

// Scheme types supported by this program. Also
// UNKNOWN is defined to be used when the scheme is unknown
// https://tools.ietf.org/html/rfc3986#section-3
const (
	FTP Type = iota
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

// ListNames string naming mapping for each Type
var ListNames = []string{
	"ftp",
	"gist",
	"http",
	"https",
	"ipfs",
	"lot",
}

func (st Type) String() string {
	if st < 0 || st >= Type(len(ListNames)) {
		return "unknown"
	}
	return ListNames[st]
}

// GetTypeFrom try to identify the scheme of the locator
// and returns the Type which correspond, can return the
// Type.UNKNOWN in case it cannot identify the scheme
func GetTypeFrom(locator string) Type {
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
