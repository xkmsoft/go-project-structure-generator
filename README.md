# Go Project Structure Generator

## Overview

This is a simple CLI application to create the basic layout for Go application projects mentioned in repository [Golang standards - project layout](https://github.com/golang-standards/project-layout) 

## Basic Usage
The application basically checks the global GOPATH environment and creates the desired project structure in the %GOPATH/src directory.

You can use the [cmd/generator/main.go](https://github.com/xkmsoft/go-project-structure-generator/blob/master/cmd/generator/main.go) as a starting point
```go
package main

import (
	"flag"
	"fmt"
	"github.com/xkmsoft/go-project-structure-generator/pkg/generator"
	"log"
	"os"
)

func main() {
	project := flag.String("project", "", "Project name")
	flag.Parse()

	if *project == "" {
		_, _ = fmt.Fprintf(os.Stderr, "missing required -project argument/flag\n")
		os.Exit(2)
	}

	g := generator.NewGoProjectStructureGenerator(*project)
	if err := g.Generate(); err != nil {
		log.Fatal(err)
	}
}
```

And then simply execute the following commands to init and tidy the go modules.
```
chasank@macbookpro16 generator % go mod tidy
go: finding module for package github.com/xkmsoft/go-project-structure-generator/pkg/generator
go: found github.com/xkmsoft/go-project-structure-generator/pkg/generator in github.com/xkmsoft/go-project-structure-generator v0.1.0
go: finding module for package golang.org/x/mod/module
go: finding module for package golang.org/x/mod/modfile
go: found golang.org/x/mod/modfile in golang.org/x/mod v0.5.0
go: found golang.org/x/mod/module in golang.org/x/mod v0.5.0
chasank@macbookpro16 generator % go run main.go -project=test
Project with name test created successfully on /Users/chasank/go/src/test
chasank@macbookpro16 generator % tree ~/go/src/test 
/Users/chasank/go/src/test
├── LICENSE.md
├── Makefile
├── README.md
├── api
│   └── README.md
├── assets
│   └── README.md
├── build
│   └── README.md
├── cmd
│   └── README.md
├── configs
│   └── README.md
├── deployments
│   └── README.md
├── docs
│   └── README.md
├── examples
│   └── README.md
├── githooks
│   └── README.md
├── init
│   └── README.md
├── internal
│   └── README.md
├── pkg
│   └── README.md
├── scripts
│   └── README.md
├── test
│   └── README.md
├── third_party
│   └── README.md
├── tools
│   └── README.md
├── vendor
│   └── README.md
├── web
│   └── README.md
└── website
    └── README.md

19 directories, 22 files
```
