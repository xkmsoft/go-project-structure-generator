package generator

import (
	"fmt"
	"go/build"
	"io"
	"os"
	"path/filepath"
)

// CopyFile copies the file from the src into des.
// It returns the number of bytes copied and error if anything goes wrong
func CopyFile(src string, dest string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer func(source *os.File) {
		if err := source.Close(); err != nil {
			fmt.Printf("Error closing source file: %s\n", err.Error())
		}
	}(source)

	destination, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer func(destination *os.File) {
		if err := destination.Close(); err != nil {
			fmt.Printf("Error closing destination file: %s\n", err.Error())
		}

	}(destination)
	n, err := io.Copy(destination, source)
	return n, err
}

// GetGOPATH returns the GOPATH environment variable
func GetGOPATH() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}

// IsDirExists simply checks whether a directory exists or not
func IsDirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// GetAssetDirectory returns the asset directory of the project or a non-nil error if anything goes wrong.
func GetAssetDirectory() (string, error) {
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(current, AssetsDirectory), nil
}
