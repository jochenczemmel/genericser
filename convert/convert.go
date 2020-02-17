package convert

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

// ConvertFile changes the types in the go source and writes it
// to the file targetFile.
func ConvertFile(sourceFile, targetFile string,
	replacer Replacer) error {

	src, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return fmt.Errorf("can not read file %s: %w", sourceFile, err)
	}

	fset := token.NewFileSet()
	srcTree, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("can not parse file %s: %w", sourceFile, err)
	}

	replacer.filterTypeAlias(srcTree)

	ast.Walk(replacer, srcTree)

	fh, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("can not open file %s: %w", targetFile, err)
	}
	defer fh.Close()

	format.Node(fh, fset, srcTree)

	return nil
}

// Replacer contains the replacement rules (key=from, value=to)
type Replacer map[string]string

// Visit replaces the type names according to the definition
// in the Replacer.
func (r Replacer) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	switch x := node.(type) {
	case *ast.Ident:
		for key, value := range r {
			if x.Name == key {
				x.Name = value
			} else if strings.HasSuffix(x.Name, key) {
				idx := strings.LastIndex(x.Name, key)
				x.Name = x.Name[:idx] + value
			}
		}
	}

	return r
}

// filterTypeAlias removes type alias definitions for the
// defined replacement types.
func (r Replacer) filterTypeAlias(srcTree *ast.File) {

	newDecl := []ast.Decl{}

	for _, decl := range srcTree.Decls {
		keep := true
		switch x := decl.(type) {
		case *ast.GenDecl:
			if x.Tok == token.TYPE {
				for _, spec := range x.Specs {
					switch y := spec.(type) {
					case *ast.TypeSpec:
						for key := range r {
							if y.Name.Name == key && y.Assign > 0 {
								keep = false
							}
						}
					}
				}
			}
		}
		if keep {
			newDecl = append(newDecl, decl)
		}
	}

	srcTree.Decls = newDecl
}
