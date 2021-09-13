package generator

import (
	"bufio"
	"fmt"
	"go/build"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
	"io"
	"io/ioutil"
	"os"
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

// CreateFile simply creates file from given filename and content. It returns the number of bytes written and error
func CreateFile(fileName string, content string) (int, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}(f)

	writer := bufio.NewWriter(f)
	n, err := writer.Write([]byte(content))
	if err != nil {
		return 0, err
	}

	err = writer.Flush()
	if err != nil {
		return 0, err
	}

	return n, nil
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

// GetGoModule reads and parses the go.mod file and returns the go module path and its version if the
// desired module name is found in the file
func GetGoModule(modulePath string) (*module.Version, error) {
	bytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return nil, err
	}
	mod, err := modfile.Parse("go.mod", bytes, nil)
	if err != nil {
		return nil, err
	}
	for _, requirement := range mod.Require {
		if requirement.Mod.Path == modulePath && !requirement.Indirect {
			return &requirement.Mod, nil
		}
	}
	return nil, fmt.Errorf("requirement %s could not be found in go.mod", modulePath)
}

// GetCurrentModulePath returns the current project's module path
func GetCurrentModulePath() (string, error) {
	bytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "", err
	}
	mod, err := modfile.Parse("go.mod", bytes, nil)
	if err != nil {
		return "", err
	}
	return mod.Module.Mod.Path, nil
}
