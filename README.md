# Go Project Structure Generator

## Overview

This is a simple CLI application to create the basic layout for Go application projects mentioned in repository [Golang standards - project layout](https://github.com/golang-standards/project-layout) 

## Basic Usage
The application basically checks the global GOPATH environment and creates the desired project structure in the %GOPATH/src directory 
```
chasank@macbookpro16 go-project-structure-generator % go run cmd/generator/main.go -project=example
Project with name example created successfully on /Users/chasank/go/src/example
chasank@macbookpro16 go-project-structure-generator % tree /Users/chasank/go/src/example
/Users/chasank/go/src/example
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
chasank@macbookpro16 go-project-structure-generator % 
```
