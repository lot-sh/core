package core

import "testing"

func TestGetSchemeTypeFromShouldWorkAsExpected(t *testing.T) {
	locatorFTP := "ftp://username:pass@example.com"
	if GetSchemeTypeFrom(locatorFTP) != FTP {
		t.Errorf("The SchemeType of %s should be FTP, instead got %s", locatorFTP, GetSchemeTypeFrom(locatorFTP))
	}
	locatorGist := "gist:username/path/to/file.sh"
	if GetSchemeTypeFrom(locatorGist) != GIST {
		t.Errorf("The SchemeType of %s should be FTP, instead got %s", locatorGist, GetSchemeTypeFrom(locatorGist))
	}
	locatorHTTP := "http://thelifeofbrian.com/the/best/film.sh"
	if GetSchemeTypeFrom(locatorHTTP) != HTTP {
		t.Errorf("The SchemeType of %s should be FTP, instead got %s", locatorHTTP, GetSchemeTypeFrom(locatorHTTP))
	}
	locatorHTTPS := "https://gist.githubusercontent.com/Kelvur/896265085d32db2c3ab065ea0995b0a3/raw/62df0ad25eed20cdfc27235e8c39cbbbdf967ed3/2020_content.md"
	if GetSchemeTypeFrom(locatorHTTPS) != HTTPS {
		t.Errorf("The SchemeType of %s should be FTP, instead got %s", locatorHTTPS, GetSchemeTypeFrom(locatorHTTPS))
	}
	locatorIPFS := "ipfs:QmT5NvUtoM5nWFfrQdVrFtvGfKFmG7AHE8P34isapyhCxX"
	if GetSchemeTypeFrom(locatorIPFS) != IPFS {
		t.Errorf("The SchemeType of %s should be FTP, instead got %s", locatorIPFS, GetSchemeTypeFrom(locatorIPFS))
	}
	locatorLot := "lot:QmT5NvUtoM5nWFfrQdVrFtvGfKFmG7AHE8P34isapyhCxX"
	if GetSchemeTypeFrom(locatorLot) != LOT {
		t.Errorf("The SchemeType of %s should be FTP, instead got %s", locatorLot, GetSchemeTypeFrom(locatorLot))
	}
	locatorUnknown := "monty://john.cleese/dead/parrot"
	if GetSchemeTypeFrom(locatorUnknown) != UNKNOWN {
		t.Errorf("The SchemeType of %s should be FTP, instead got %s", locatorUnknown, GetSchemeTypeFrom(locatorUnknown))
	}
}
