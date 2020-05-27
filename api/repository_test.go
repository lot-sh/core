package api

import (
	"io/ioutil"
	"testing"
	"bytes"
	"context"
	"bufio"
	"github.com/lot-sh/core/api"
)
 
func TestRepositoryAdd(t *testing.T) {
	ctx := context.Background()
	repo := api.NewFileRepository()
	buf := bytes.NewBufferString("This is a test")
	str, err := repo.Add(ctx, bufio.NewReader(buf));
	if err != nil {
		t.Fatalf("Test failed in create resource; %s\n", err)
	}

	t.Logf("CID: %s", str)
}

func TestRepositoryGet(t *testing.T) {
	ctx := context.Background()
	repo := api.NewFileRepository()
	reader, err := repo.Get(ctx, "zb2rhoF5PNBBjjGXU86EhU4Lt5su47dfqeBd87YvGXh9qcDDx");
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Logf("Test failed in get resource; %s\n", err)
	}
	
	t.Logf("Content: %x", buf)
}

func TestRepositoryRemove(t *testing.T) {
	ctx := context.Background()
	repo := api.NewFileRepository()
	err := repo.Remove(ctx, "zb2rhoF5PNBBjjGXU86EhU4Lt5su47dfqeBd87YvGXh9qcDDx");
	if err != nil {
		t.Fatalf("Test failed in remove resource; %s\n", err)
	}
}