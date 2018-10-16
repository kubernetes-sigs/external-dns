// A toolbox to check/generate CloudStack API commands from the JSON description

package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/exoscale/egoscale"
)

// must be sorted
var ignoredFields = []string{
	"aclid",
	"acltype",
	"forvpc",
	"policyid",
	"project",
	"projectid",
	"vpc",
	"vpcavailable",
	"vpcid",
	"vpclimit",
	"vpctotal",
}

var cmd = flag.String("cmd", "", "CloudStack command name")
var source = flag.String("apis", "", "listApis response in JSON")
var rtype = flag.String("type", "", "Actual type to check against the cmd (need cmd)")

var apiTypes = map[string]string{
	"short":   "int16",
	"integer": "int",
	"long":    "int64",
	"map":     "map[string]string",
	"list":    "[]struct{}",
	"set":     "[]struct{}",
	"uuid":    "*UUID",
	"boolean": "*bool",
	"date":    "string",
}

// fieldInfo represents the inner details of a field
type fieldInfo struct {
	Var       *types.Var
	OmitEmpty bool
	Doc       string
}

// command represents a struct within the source code
type command struct {
	name        string
	description string
	sync        string
	s           *types.Struct
	position    token.Pos
	fields      map[string]fieldInfo
	errors      map[string][]error
}

func main() {
	flag.Parse()

	sourceFile, _ := os.Open(*source)
	decoder := json.NewDecoder(sourceFile)
	apis := new(egoscale.ListAPIsResponse)
	if err := decoder.Decode(&apis); err != nil {
		panic(err)
	}

	files, err := filepath.Glob("*.go")
	if err != nil {
		panic(err)
	}
	fset := token.NewFileSet()
	astFiles := make([]*ast.File, len(files))
	for i, file := range files {
		f, er := parser.ParseFile(fset, file, nil, 0)
		if er != nil {
			panic(er)
		}
		astFiles[i] = f
	}

	config := types.Config{
		Importer: importer.For("source", nil),
	}

	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}

	_, err = config.Check("egoscale", fset, astFiles, info)
	if err != nil {
		_, e := fmt.Fprintf(os.Stderr, err.Error())
		if e != nil {
			panic(e)
		}
		os.Exit(1)
	}

	commands := make(map[string]*command)

	for id, obj := range info.Defs {
		if obj == nil || !obj.Exported() {
			continue
		}

		typ := obj.Type().Underlying()

		switch typ.(type) {
		case *types.Struct:
			commands[strings.ToLower(obj.Name())] = &command{
				name:     obj.Name(),
				s:        typ.(*types.Struct),
				position: id.Pos(),
			}
		}
	}

	re := regexp.MustCompile(`\bjson:"(?P<name>[^,"]+)(?P<omit>,omitempty)?"`)
	reDoc := regexp.MustCompile(`\bdoc:"(?P<doc>[^"]+)"`)

	for _, a := range apis.API {
		name := strings.ToLower(a.Name)
		params := a.Params

		if strings.ToLower(*cmd) == name && *rtype != "" {
			panic(fmt.Errorf("checking return type is temporary disabled"))
			/*
				name = strings.ToLower(*rtype)
				*cmd = name
				params = a.Response
				_, e := fmt.Fprintf(os.Stderr, "Checking return type of %sResult, using %q\n", a.Name, *rtype)
				if e != nil {
					panic(e)
				}
			*/
		}

		if command, ok := commands[name]; !ok {
			// too much information
			//fmt.Fprintf(os.Stderr, "Unknown command: %q\n", name)
		} else {
			command.description = strings.Trim(a.Description, " ")
			// mapping from name to field
			command.fields = make(map[string]fieldInfo)
			command.errors = make(map[string][]error)

			if a.IsAsync {
				command.sync = " (A)"
			}

			hasMeta := false

			for i := 0; i < command.s.NumFields(); i++ {
				f := command.s.Field(i)

				if !f.IsField() || !f.Exported() {
					if f.Name() != "_" {
						continue
					}

					tag := (reflect.StructTag)(command.s.Tag(i))
					name, nameOK := tag.Lookup("name")
					description, descriptionOK := tag.Lookup("description")
					if !nameOK || !descriptionOK {
						command.errors["_"] = append(command.errors["_"], fmt.Errorf("meta field incomplete, wanted\n\t_ bool `name:%q description:%q`", a.Name, command.description))
					} else {
						if name != a.Name || description != command.description {
							command.errors["_"] = append(command.errors["_"], fmt.Errorf("meta field incorrect, got %q %q, wanted\n\t_ bool `name:%q description:%q`", name, description, a.Name, command.description))
						}
					}

					hasMeta = true
					continue
				}

				tag := command.s.Tag(i)
				match := re.FindStringSubmatch(tag)
				if len(match) == 0 {
					n := f.Name()
					command.errors[n] = append(command.errors[n], errors.New("field error: no json annotation found"))
					continue
				}
				name := match[1]
				omitempty := len(match) == 3 && match[2] == ",omitempty"

				doc := ""
				match = reDoc.FindStringSubmatch(tag)
				if len(match) == 2 {
					doc = match[1]
				}

				command.fields[name] = fieldInfo{
					Var:       f,
					OmitEmpty: omitempty,
					Doc:       doc,
				}
			}

			for _, p := range params {
				n := p.Name
				index := sort.SearchStrings(ignoredFields, p.Name)
				ignored := index < len(ignoredFields) && ignoredFields[index] == p.Name
				if ignored {
					continue
				}
				field, ok := command.fields[p.Name]
				description := strings.Trim(p.Description, " ")

				omit := ""
				if !p.Required {
					omit = ",omitempty"
				}

				if !ok {
					doc := ""
					if description != "" {
						doc = fmt.Sprintf(" doc:%q", description)
					}

					apiType, ok := apiTypes[p.Type]
					if !ok {
						apiType = p.Type
					}

					command.errors[n] = append(command.errors[n], fmt.Errorf("missing field:\n\t%s %s `json:\"%s%s\"%s`", strings.Title(p.Name), apiType, p.Name, omit, doc))
					continue
				}
				delete(command.fields, p.Name)

				typename := field.Var.Type().String()

				if field.Doc != description {
					if field.Doc == "" {
						command.errors[n] = append(command.errors[n], fmt.Errorf("missing doc:\n\t\t`doc:%q`", description))
					} else {
						command.errors[n] = append(command.errors[n], fmt.Errorf("wrong doc want %q got %q", description, field.Doc))
					}
				}

				if p.Required == field.OmitEmpty {
					command.errors[n] = append(command.errors[n], fmt.Errorf("wrong omitempty, want `json:\"%s%s\"`", p.Name, omit))
					continue
				}

				expected := ""
				switch p.Type {
				case "short":
					if typename != "int16" {
						expected = "int16"
					}
				case "int":
				case "integer":
					// uint are used by port and icmp types
					if typename != "int" && typename != "uint16" && typename != "uint8" {
						expected = "int"
					}
				case "long":
					if typename != "int64" && typename != "uint64" {
						expected = "int64"
					}
				case "boolean":
					if typename != "bool" && typename != "*bool" {
						expected = "bool"
					}
				case "string":
				case "date":
				case "tzdate":
				case "imageformat":
					if typename != "string" {
						expected = "string"
					}
				case "uuid":
					if typename != "*egoscale.UUID" {
						expected = "*UUID"
					}
				case "list":
					if !strings.HasPrefix(typename, "[]") {
						expected = "[]string"
					}
				case "map":
				case "set":
					if !strings.HasPrefix(typename, "[]") {
						expected = "array"
					}
				default:
					command.errors[n] = append(command.errors[n], fmt.Errorf("unknown type %q <=> %q", p.Type, field.Var.Type().String()))
				}

				if expected != "" {
					command.errors[n] = append(command.errors[n], fmt.Errorf("expected to be a %s, got %q", expected, typename))
				}
			}

			if !hasMeta && *rtype == "" {
				command.errors["_"] = append(command.errors["_"], fmt.Errorf("meta field missing, wanted\n\t\t_ bool `name:%q description:%q`", a.Name, a.Description))
			}

			for name := range command.fields {
				command.errors[name] = append(command.errors[name], errors.New("extra field found"))
			}
		}
	}

	for name, c := range commands {
		pos := fset.Position(c.position)
		er := len(c.errors)

		if *cmd == "" {
			if er != 0 {
				fmt.Printf("%5d %s: %s%s\n", er, pos, c.name, c.sync)
			}
		} else if strings.ToLower(*cmd) == name {
			errs := make([]string, 0, len(c.errors))
			for k, es := range c.errors {
				var b strings.Builder
				for i, e := range es {
					if i > 0 {
						if _, err := fmt.Fprintln(&b, ""); err != nil {
							panic(e)
						}
					}
					if _, err := fmt.Fprintf(&b, "%s: %s", k, e.Error()); err != nil {
						panic(err)
					}
				}
				errs = append(errs, b.String())
			}
			sort.Strings(errs)
			for _, e := range errs {
				fmt.Println(e)
			}
			fmt.Printf("\n%s: %s%s has %d error(s)\n", pos, c.name, c.sync, er)
			os.Exit(er)
		}
	}

	if *cmd != "" {
		fmt.Printf("%s not found\n", *cmd)
		os.Exit(1)
	}
}
