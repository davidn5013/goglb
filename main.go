// Lists all global variables (and constants) in a module path
package main

/*

TODO show on file with variable and consts
TODO add file and line number

TODO always show variables and constans
TODO argument for show variables or constans

*/

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	modulePath := flag.String("path", ".", "path to module")
	variableAndConst := flag.Bool("varconst", false, "list global variabel and constrants")
	flag.Usage = func() {
		fmt.Println("Go tool for listing global variables (and constans) in go module path")
		flag.PrintDefaults()
	}
	flag.Parse()

	files := listGoFiles(*modulePath)

	if *variableAndConst == false {
		showGlbVariables(files)
	} else {
		showGlbVariablesAndConst(files)
	}

}

func showGlbVariables(files []string) {
	for _, file := range files {

		variables := listGlobalVariables(file)
		fmt.Printf("Global variables in %s:\n", file)
		for _, variable := range variables {
			fmt.Println(variable)
		}
	}

}

func showGlbVariablesAndConst(files []string) {
	for _, file := range files {
		variables, constants := listGlobals(file)
		fmt.Printf("Globals in %s:\n", file)
		fmt.Println("Variables:")
		for _, variable := range variables {
			fmt.Println(variable)
		}
		fmt.Println("Constants:")
		for _, constant := range constants {
			fmt.Println(constant)
		}
	}
}

func listGoFiles(modulePath string) []string {
	var goFiles []string

	err := filepath.Walk(modulePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	return goFiles
}

func listGlobalVariables(filename string) []string {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var variables []string

	for _, decl := range node.Decls {
		if gd, ok := decl.(*ast.GenDecl); ok && gd.Tok == token.VAR {
			for _, spec := range gd.Specs {
				if vs, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range vs.Names {
						variables = append(variables, name.Name)
					}
				}
			}
		}
	}

	return variables
}

func listGlobals(filename string) ([]string, []string) {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}

	var variables []string
	var constants []string

	for _, decl := range node.Decls {
		if gd, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range gd.Specs {
				if vs, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range vs.Names {
						if gd.Tok == token.VAR {
							variables = append(variables, name.Name)
						} else if gd.Tok == token.CONST {
							constants = append(constants, name.Name)
						}
					}
				}
			}
		}
	}

	return variables, constants
}
