package repository

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/lot-sh/core/types"
)

var (
	repo        types.Repository
	directory   = path.Join(os.TempDir(), "lotsh/lib")
	filecontent = "This is a test"
	hash        = "zb2rhk5zU8kdMPiPKgdxjhXTxs7KBxaXFqDhK2sBr5xQckX5f"
	err         error
)

func TestErrorWhenInitialize(t *testing.T) {
	prohibitedDir := "/test"
	repo, err = NewFileSystemRepository(prohibitedDir)
	if err == nil {
		t.Errorf("Error must be raised but got nil")
	}
	actual := err.Error()
	expected := fmt.Sprintf("Failed to initialize repository the directory %s can't be created", prohibitedDir)
	if !strings.Contains(actual, expected) {
		t.Errorf("Missmatch content \n-> Actual: %s \n-> Expected: %s", actual, expected)
	}
}

func TestInitializeWithoutErrors(t *testing.T) {
	repo, err = NewFileSystemRepository(directory)
	if err != nil {
		t.Error(err)
	}
}

func TestRepositoryAdd(t *testing.T) {
	buf := bytes.NewBufferString(filecontent)
	actual, err := repo.Add(bufio.NewReader(buf))
	expected := hash
	if err != nil {
		t.Fatalf("Test failed in create resource; %s\n", err)
	}

	if !strings.Contains(string(actual), expected) {
		t.Errorf("Missmatch content \n-> Actual: %s \n-> Expected: %s", actual, expected)
	}
}

func TestRepositoryGet(t *testing.T) {
	reader, err := repo.Get(hash)
	actual, err := ioutil.ReadAll(reader)
	expected := filecontent
	if err != nil {
		t.Logf("Test failed in get resource; %s\n", err)
	}

	if !strings.Contains(string(actual), expected) {
		t.Errorf("Missmatch content \n-> Actual: %s \n-> Expected: %s", actual, expected)
	}
}

func TestRepositoryRemove(t *testing.T) {
	if err := repo.Remove(hash); err != nil {
		t.Fatalf("Test failed when remove resource: %s\n", err)
	}
}
