package genbase

import (
	"errors"
	"go/ast"
	"path/filepath"
	"strings"
)

func pathJoinAll(directory string, names ...string) []string {
	if directory == "." {
		return names
	}
	ret := make([]string, len(names))
	for i, name := range names {
		ret[i] = filepath.Join(directory, name)
	}
	return ret
}

func findAnnotation(doc *ast.CommentGroup, directive string) *ast.Comment {
	if doc == nil {
		return nil
	}

	for _, c := range doc.List {
		l := c.Text
		t := strings.TrimLeft(l, "/ ")
		if !strings.HasPrefix(t, directive) {
			continue
		}

		t = strings.TrimPrefix(t, directive)

		if len(t) > 0 && t[0] != ' ' {
			continue
		}

		return c
	}

	return nil
}

func IsReferenceToOtherPackage(expr ast.Expr) (bool, string) {
	switch t := expr.(type) {
	case *ast.Ident:
		return false, ""
	case *ast.StarExpr:
		return IsReferenceToOtherPackage(t.X)
	case *ast.SelectorExpr:
		name, err := ExprToTypeName(t.X)
		if err != nil {
			return false, ""
		}
		return true, name
	default:
		return false, ""
	}
}

func ExprToTypeName(expr ast.Expr) (string, error) {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name, nil
	}
	if star, ok := expr.(*ast.StarExpr); ok {
		x, err := ExprToTypeName(star.X)
		if err != nil {
			return "", nil
		}
		return "*" + x, nil
	}
	if selector, ok := expr.(*ast.SelectorExpr); ok {
		x, err := ExprToTypeName(selector.X)
		if err != nil {
			return "", nil
		}
		sel, err := ExprToTypeName(selector.Sel)
		if err != nil {
			return "", nil
		}
		return x + "." + sel, nil
	}
	if array, ok := expr.(*ast.ArrayType); ok {
		x, err := ExprToTypeName(array.Elt)
		if err != nil {
			return "", nil
		}
		return "[]" + x, nil
	}
	return "", errors.New("can't detect type name")
}

func ExprToBaseTypeName(expr ast.Expr) (string, error) {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name, nil
	}
	if star, ok := expr.(*ast.StarExpr); ok {
		x, err := ExprToBaseTypeName(star.X)
		if err != nil {
			return "", nil
		}
		return x, nil
	}
	if selector, ok := expr.(*ast.SelectorExpr); ok {
		x, err := ExprToBaseTypeName(selector.X)
		if err != nil {
			return "", nil
		}
		sel, err := ExprToBaseTypeName(selector.Sel)
		if err != nil {
			return "", nil
		}
		return x + "." + sel, nil
	}
	if array, ok := expr.(*ast.ArrayType); ok {
		x, err := ExprToBaseTypeName(array.Elt)
		if err != nil {
			return "", nil
		}
		return x, nil
	}
	return "", errors.New("can't detect type name")
}

func GetKeys(tag string) []string {
	result := []string{}

	// from reflect.StructTag.Get(string)

	for tag != "" {
		// skip leading space
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		// scan to colon.
		// a space or a quote is a syntax error
		i = 0
		for i < len(tag) && tag[i] != ' ' && tag[i] != ':' && tag[i] != '"' {
			i++
		}
		if i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := string(tag[:i])
		result = append(result, name)
		tag = tag[i+1:]

		// scan quoted string to find value
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		tag = tag[i+1:]
	}
	return result
}
