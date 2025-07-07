package importgraph

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func StartJohnnyWalker() {
	dir := ""
	if len(os.Args) < 2 {
		dir = "./"
	} else {
		dir = os.Args[1]
	}
	dir = "/Users/izik.golan/git/platform/model-analysis-service"

	dependencies := make(map[string][]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".go" {
			fset := token.NewFileSet()
			node, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
			if err != nil {
				return err
			}
			var imports []string
			for _, imp := range node.Imports {
				if strings.HasPrefix(imp.Path.Value, "\"model-analysis-service") {
					parts := strings.Split(imp.Path.Value, "/")
					lastPart := parts[len(parts)-1]
					imports = append(imports, strings.Trim(lastPart, "\""))
					//imports = append(imports, strings.Trim(imp.Path.Value, "\""))
				}
			}
			dependencies[node.Name.Name] = imports
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Edge list:")
	for pkg, imports := range dependencies {
		for _, imp := range imports {
			fmt.Printf("%s>%s\n", pkg, imp)
		}
	}
}

func StartJohnnyWalker_() {
	//dir := "" // Folder to scan
	dir := "/Users/izik.golan/git/platform/model-analysis-service"
	if len(os.Args) < 2 {
		dir = "./"
	} else {
		dir = os.Args[1]
	}

	dependencies := make(map[string][]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".go" {
			fset := token.NewFileSet()
			node, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
			if err != nil {
				return err
			}
			var imports []string
			for _, imp := range node.Imports {
				imports = append(imports, imp.Path.Value)
			}
			dependencies[node.Name.Name] = imports
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Package Dependencies:")
	for pkg, imports := range dependencies {
		fmt.Printf("%s -> %v\n", pkg, imports)
	}
}
