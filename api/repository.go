package api

import (
	"os"
	"io"
	"context"
	"io/ioutil"
	"path"
	cid "github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
	mbase "github.com/multiformats/go-multibase"
)

type FileRepository struct {
	FileMode os.FileMode
}

func NewFileRepository() *FileRepository {
	return &FileRepository{
		FileMode: 0644,
	}
}

func (repo *FileRepository) Add(ctx context.Context, reader io.Reader) (string, error) {

	hash := cid.Prefix{
		Version: 1,
		Codec: cid.Raw,
		MhType: mh.SHA2_256,
		MhLength: -1, // default length
	}

	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err;
	}

	cid, err := hash.Sum(buf)
	if  err != nil {
		return "", err;
	}

	scid, err := cid.StringOfBase(mbase.Base58BTC);
	if  err != nil {
		return "", err;
	}

	fullpath := getFullpathFilename(ctx, scid)
	if err := ioutil.WriteFile(fullpath, buf, repo.FileMode); err != nil {
		return "", err;
	}

	return scid, nil
}

func (rp *FileRepository) Get(ctx context.Context, scid string) (io.Reader, error) { 
	return os.Open(getFullpathFilename(ctx, scid))
}

func (rp *FileRepository) Remove(ctx context.Context, scid string) error { 
	return os.Remove(getFullpathFilename(ctx, scid))
}

func getFullpathFilename(ctx context.Context, scid string) string {
	return path.Join(getLibDirFromConfig(ctx), scid)
}

func getLibDirFromConfig(ctx context.Context) string {
	// TODO TDB 
	// config := ctx.GetValue("config");
	libdir := "/usr/local/opt/lotsh/lib"
	return libdir
}