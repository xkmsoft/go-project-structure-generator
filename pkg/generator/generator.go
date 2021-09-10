package generator

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)
// GoProjectStructureGenerator represents the necessary data types to generate standard go project structure
type GoProjectStructureGenerator struct {
	ModuleName  string
	Directories map[string][]string
	ModuleFiles []string
	Permission  os.FileMode
}

const (
	SrcDirectory    = "src"
	AssetsDirectory = "assets"
	ReadmeMD        = "README.md"
	Module          = "module"
)

// NewGoProjectStructureGenerator returns new *GoProjectStructureGenerator with provided project name
func NewGoProjectStructureGenerator(moduleName string) *GoProjectStructureGenerator {
	var directories = map[string][]string{
		"api":         {ReadmeMD},
		"assets":      {ReadmeMD},
		"build":       {ReadmeMD},
		"cmd":         {ReadmeMD},
		"configs":     {ReadmeMD},
		"deployments": {ReadmeMD},
		"docs":        {ReadmeMD},
		"examples":    {ReadmeMD},
		"githooks":    {ReadmeMD},
		"init":        {ReadmeMD},
		"internal":    {ReadmeMD},
		"pkg":         {ReadmeMD},
		"scripts":     {ReadmeMD},
		"test":        {ReadmeMD},
		"third_party": {ReadmeMD},
		"tools":       {ReadmeMD},
		"vendor":      {ReadmeMD},
		"web":         {ReadmeMD},
		"website":     {ReadmeMD},
	}
	var moduleFiles = []string{
		".editorconfig",
		".gitignore",
		"LICENSE.md",
		"Makefile",
		"README.md",
	}
	return &GoProjectStructureGenerator{
		ModuleName:  moduleName,
		Directories: directories,
		ModuleFiles: moduleFiles,
		Permission:  0755,
	}
}

// Generate simply generates the project structure on the following path GOPATH/src/<provided-project-name>
// And copies all the README.md files from the assets directory into the project directory
func (g *GoProjectStructureGenerator) Generate() error {
	assetsDirectory, err := GetAssetDirectory()
	if err != nil {
		return err
	}

	goPath := GetGOPATH()
	src := filepath.Join(goPath, SrcDirectory)
	if !IsDirExists(src) {
		return errors.New("src directory could not be found in the GOPATH\n")
	}

	projectDir := filepath.Join(src, g.ModuleName)
	if IsDirExists(projectDir) {
		return errors.New(fmt.Sprintf("module with name %s already exists in %s\n", g.ModuleName, projectDir))
	}

	if err := os.Mkdir(projectDir, g.Permission); err != nil {
		return err
	}

	// Copying module files into the root project directory
	for _, file := range g.ModuleFiles {
		source := filepath.Join(assetsDirectory, Module, file)
		destination := filepath.Join(projectDir, file)
		if _, err := CopyFile(source, destination); err != nil {
			return err
		}
	}

	// Creating sub-directories and copying necessary files into the sub-directories
	for dir, files := range g.Directories {
		path := filepath.Join(projectDir, dir)
		if err := os.Mkdir(path, g.Permission); err != nil {
			return err
		}
		for _, file := range files {
			source := filepath.Join(assetsDirectory, dir, file)
			destination := filepath.Join(path, file)
			if _, err := CopyFile(source, destination); err != nil {
				return err
			}
		}
	}
	fmt.Printf("Project with name %s created successfully on %s\n", g.ModuleName, projectDir)
	return nil
}
