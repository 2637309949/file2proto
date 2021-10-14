package main

import (
	"errors"
	"fmt"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

type go2proto struct {
}

func (to *go2proto) Transform(uri string) ([]*message, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error Getwd: %s", err)
		return []*message{}, err
	}
	pkgs, err := loadPackages(pwd, []string{uri})
	if err != nil {
		log.Fatalf("error fetching packages: %s", err)
		return []*message{}, err
	}
	msgs := packagesToMessages(pkgs)
	return msgs, nil
}

func (to *go2proto) Check(uri string) bool {
	if strings.Contains(uri, ".go") {
		return true
	}
	s, err := os.Stat(uri)
	if err == nil && s.IsDir() {
		files, _ := ioutil.ReadDir(uri)
		for _, f := range files {
			if strings.Contains(f.Name(), ".go") {
				return true
			}
		}
	}
	return false
}

func loadPackages(pwd string, pkgs []string) ([]*packages.Package, error) {
	fset := token.NewFileSet()
	cfg := &packages.Config{
		Dir:  pwd,
		Mode: packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax | packages.NeedTypesInfo,
		Fset: fset,
	}
	packages, err := packages.Load(cfg, pkgs...)
	if err != nil {
		return nil, err
	}
	var errs = ""
	for _, p := range packages {
		if len(p.Errors) > 0 {
			errs += fmt.Sprintf("error fetching package %s: ", p.String())
			for _, e := range p.Errors {
				errs += e.Error()
			}
			errs += "; "
		}
	}
	if errs != "" {
		return nil, errors.New(errs)
	}
	return packages, nil
}

func packagesToMessages(pkgs []*packages.Package) []*message {
	var out []*message
	seen := map[string]struct{}{}
	for _, p := range pkgs {
		for _, t := range p.TypesInfo.Defs {
			if t == nil {
				continue
			}
			if !t.Exported() {
				continue
			}
			if _, ok := seen[t.Name()]; ok {
				continue
			}
			if _, ok := t.(*types.TypeName); !ok {
				continue
			}
			if s, ok := t.Type().Underlying().(*types.Struct); ok {
				seen[t.Name()] = struct{}{}
				out = appendPackagesMessage(out, t, s)
			}
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func appendPackagesMessage(out []*message, t types.Object, s *types.Struct) []*message {
	msg := &message{
		Name:   t.Name(),
		Fields: []*field{},
	}

	for i := 0; i < s.NumFields(); i++ {
		f := s.Field(i)
		if !f.Exported() {
			continue
		}
		newField := &field{
			Name:       toProtoFieldName(f.Name()),
			TypeName:   toProtoFieldTypeNameByVar(f),
			IsRepeated: isRepeated(f),
			Order:      i + 1,
		}
		msg.Fields = append(msg.Fields, newField)
	}
	out = append(out, msg)
	return out
}

func toProtoFieldTypeNameByVar(f *types.Var) string {
	switch f.Type().Underlying().(type) {
	case *types.Basic:
		name := f.Type().String()
		return normalizeType(name)
	case *types.Slice:
		name := splitNameHelper(f)
		return normalizeType(strings.TrimLeft(name, "[]"))

	case *types.Pointer, *types.Struct:
		name := splitNameHelper(f)
		return normalizeType(name)
	}
	return f.Type().String()
}

func splitNameHelper(f *types.Var) string {
	parts := strings.Split(f.Type().String(), ".")
	name := parts[len(parts)-1]
	if name[0] == '*' {
		name = name[1:]
	}
	return name
}

func isRepeated(f *types.Var) bool {
	_, ok := f.Type().Underlying().(*types.Slice)
	return ok
}

func init() {
	register("go2proto", new(go2proto))
}
