package generator

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Name    string
	Content string
}

// GoProjectStructureGenerator represents the necessary data types to generate standard go project structure
type GoProjectStructureGenerator struct {
	ModuleName            string
	Directories           map[string][]string
	AdditionalDirectories map[string][]File
	ModuleFiles           []string
	Permission            os.FileMode
}

const (
	SrcDirectory      = "src"
	AssetsDirectory   = "assets"
	ReadmeMD          = "README.md"
	Module            = "module"
	ProjectModuleName = "github.com/xkmsoft/go-project-structure-generator"
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
	var additionalDirectories = map[string][]File{
		"vendor": {File{
			Name:    ReadmeMD,
			Content: "# `/vendor`\n\nApplication dependencies (managed manually or by your favorite dependency " +
				"management tool like the new built-in, but still experimental, " +
				"[`modules`](https://github.com/golang/go/wiki/Modules) feature)." +
				"\n\nDon't commit your application dependencies if you are building a library." +
				"\n\nNote that since [`1.13`](https://golang.org/doc/go1.13#modules) " +
				"Go also enabled the module proxy feature (using `https://proxy.golang.org` as their module " +
				"proxy server by default). Read more about it [`here`](https://blog.golang.org/module-mirror-launch) " +
				"to see if it fits all of your requirements and constraints. " +
				"If it does, then you won't need the 'vendor' directory at all.\n",
		}},
	}
	return &GoProjectStructureGenerator{
		ModuleName:            moduleName,
		Directories:           directories,
		AdditionalDirectories: additionalDirectories,
		ModuleFiles:           moduleFiles,
		Permission:            0755,
	}
}

// Generate simply generates the project structure on the following path GOPATH/src/<provided-project-name>
// And copies all the README.md files from the assets directory into the project directory
func (g *GoProjectStructureGenerator) Generate() error {
	assetsDirectory, err := g.GetAssetDirectory()
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

	// Creating sub-directories and copying necessary files into the sub-directories from the assets directory
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

	// Creating additional directories and their corresponding files
	for dir, files := range g.AdditionalDirectories {
		path := filepath.Join(projectDir, dir)
		if err := os.Mkdir(path, g.Permission); err != nil {
			return err
		}
		for _, file := range files {
			destination := filepath.Join(path, file.Name)
			if _, err := CreateFile(destination, file.Content); err != nil {
				return err
			}
		}
	}

	fmt.Printf("Project with name %s created successfully on %s\n", g.ModuleName, projectDir)
	return nil
}

// GetAssetDirectory returns the asset directory of the project or a non-nil error if anything goes wrong.
func (g *GoProjectStructureGenerator) GetAssetDirectory() (string, error) {
	currentModule, err := GetCurrentModulePath()
	if err != nil {
		return "", err
	}
	goPath := GetGOPATH()
	if currentModule == ProjectModuleName {
		// Current project has the module path: github.com/xkmsoft/go-project-structure-generator
		current, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return filepath.Join(current, AssetsDirectory), nil
	} else {
		// Current project has different module path
		module, err := GetGoModule(ProjectModuleName)
		if err != nil {
			return "", err
		}
		parts := strings.Split(module.Path, "/")
		if len(parts) == 3 {
			full := filepath.Join(goPath, "pkg", "mod", parts[0], parts[1], fmt.Sprintf("%s@%s", parts[2], module.Version), AssetsDirectory)
			return full, nil
		}
	}
	return "", fmt.Errorf("asset directory could not be found")
}
