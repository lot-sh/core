package repository

import (
	"os"
	"io"
	"io/ioutil"
	"fmt"
	"log"
	"path"
	cid "github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
	mbase "github.com/multiformats/go-multibase"
)

// FileSystemRepository is a repository implementation which uses a 
// file system to gets and store the data
type FileSystemRepository struct {
	FileMode os.FileMode
	Directory string
}

// NewFileSystemRepository constructor
func NewFileSystemRepository(dir string) (*FileSystemRepository, error) {
	if err := makeDirectoryIfDoesntExist(dir); err != nil {
		return nil, fmt.Errorf("Failed to initialize repository the directory %s can't be created", dir)
	}
	
	return &FileSystemRepository{
		FileMode: 0644,
		Directory: dir,
	}, nil
}

// Add create resource and returns a unique identifier
func (repo *FileSystemRepository) Add(reader io.Reader) (string, error) {

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

	fullpath := getFullpathFilename(repo.Directory, scid)
	if err := ioutil.WriteFile(fullpath, buf, repo.FileMode); err != nil {
		return "", err;
	}

	return scid, nil
}


func (repo *FileSystemRepository) Get(scid string) (io.Reader, error) { 
	return os.Open(getFullpathFilename(repo.Directory, scid))
}

func (repo *FileSystemRepository) Remove(scid string) error { 
	return os.Remove(getFullpathFilename(repo.Directory, scid))
}

func getFullpathFilename(dir string, scid string) string {
	return path.Join(dir, scid)
}

// Checks if directory exists and if not then 
// create it recusrsively like mkdir -p
func makeDirectoryIfDoesntExist(dir string) error {
	_, err := os.Stat(dir)
	if err == nil {
		log.Printf("FileSystemRepository the directory \"%s\" already exists ignoring creation", dir)
		return nil
	}
	err = os.MkdirAll(dir, os.ModePerm)
	if err == nil {
		log.Printf("FileSystemRepository root directory created! %s", dir)
	}
	return err
}