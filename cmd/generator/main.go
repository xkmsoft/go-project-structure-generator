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
