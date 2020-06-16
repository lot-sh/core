package scheme

import "testing"

func TestGetTypeFromShouldWorkAsExpected(t *testing.T) {
	locatorFTP := "ftp://username:pass@example.com"
	if GetTypeFrom(locatorFTP) != FTP {
		t.Errorf("The Type of %s should be FTP, instead got %s", locatorFTP, GetTypeFrom(locatorFTP))
	}
	locatorGist := "gist:username/path/to/file.sh"
	if GetTypeFrom(locatorGist) != GIST {
		t.Errorf("The Type of %s should be FTP, instead got %s", locatorGist, GetTypeFrom(locatorGist))
	}
	locatorHTTP := "http://thelifeofbrian.com/the/best/film.sh"
	if GetTypeFrom(locatorHTTP) != HTTP {
		t.Errorf("The Type of %s should be FTP, instead got %s", locatorHTTP, GetTypeFrom(locatorHTTP))
	}
	locatorHTTPS := "https://gist.githubusercontent.com/Kelvur/896265085d32db2c3ab065ea0995b0a3/raw/62df0ad25eed20cdfc27235e8c39cbbbdf967ed3/2020_content.md"
	if GetTypeFrom(locatorHTTPS) != HTTPS {
		t.Errorf("The Type of %s should be FTP, instead got %s", locatorHTTPS, GetTypeFrom(locatorHTTPS))
	}
	locatorIPFS := "ipfs:QmT5NvUtoM5nWFfrQdVrFtvGfKFmG7AHE8P34isapyhCxX"
	if GetTypeFrom(locatorIPFS) != IPFS {
		t.Errorf("The Type of %s should be FTP, instead got %s", locatorIPFS, GetTypeFrom(locatorIPFS))
	}
	locatorLot := "lot:QmT5NvUtoM5nWFfrQdVrFtvGfKFmG7AHE8P34isapyhCxX"
	if GetTypeFrom(locatorLot) != LOT {
		t.Errorf("The Type of %s should be FTP, instead got %s", locatorLot, GetTypeFrom(locatorLot))
	}
	locatorUnknown := "monty://john.cleese/dead/parrot"
	if GetTypeFrom(locatorUnknown) != UNKNOWN {
		t.Errorf("The Type of %s should be FTP, instead got %s", locatorUnknown, GetTypeFrom(locatorUnknown))
	}
}

func TestCastingString(t *testing.T) {
	actual := FTP.String()
	expected := "ftp"
	if actual != "ftp" {
		t.Errorf("The Type %s should be casted as string with the value of %s, instead got %s", FTP, expected, actual)
	}
	actual = HTTP.String()
	expected = "http"
	if actual != "http" {
		t.Errorf("The Type %s should be casted as string with the value of %s, instead got %s", HTTP, expected, actual)
	}
}
