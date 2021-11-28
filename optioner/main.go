package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	FilePermissions = 0755
)

type data struct {
	Type     string
	Name     string
	Package  string
	Dir      string
	Filename string
	Path     string
}

func main() {
	var d data
	flag.StringVar(&d.Type, "type", "", "The type used for the option being generated")
	flag.StringVar(&d.Name, "name", "", "The name used for the option being generated. This should start with a capital letter so that it is exported.")
	flag.StringVar(&d.Package, "package", "", "The package used for the option being generated.")
	flag.StringVar(&d.Dir, "dir", "", "The output directory used for the option being generated.")
	flag.StringVar(&d.Filename, "filename", "", "The filename used for the option being generated.")
	flag.Parse()

	if err := processType(&d); err != nil {
		handleErr(err)
	}

	if err := processPath(&d); err != nil {
		handleErr(err)
	}

	bs, err := processTemplate(d)

	if err != nil {
		handleErr(err)
	}

	if err := processWrite(d, bs); err != nil {
		handleErr(err)
	}
}

func handleErr(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func processType(args *data) error {
	if args.Type == "" {
		return errors.New("must provide 'type' parameter")
	}

	if args.Name == "" {
		args.Name = strings.Title(args.Type)
	}

	if args.Package == "" {
		args.Package = "optioner"
	}

	return nil
}

func processPath(args *data) error {
	if args.Dir == "" {
		cwd, err := os.Getwd()

		if err != nil {
			return err
		}

		base := filepath.Base(cwd)

		if base != args.Package {
			args.Dir = filepath.Join(".", args.Package)
		}
	}

	if args.Filename == "" {
		args.Filename = fmt.Sprintf("%v.optioner.go", args.Name)
	}

	args.Path = filepath.Join(args.Dir, args.Filename)

	return nil
}

func processWrite(data data, bs []byte) error {
	err := os.MkdirAll(data.Dir, FilePermissions)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(data.Path, bs, FilePermissions); err != nil {
		return err
	}

	return nil
}

func processTemplate(args data) ([]byte, error) {
	funcs := template.FuncMap{
		"title": func(s string) string {
			return strings.Title(s)
		},
	}

	var t *template.Template
	var err error

	t, err = template.New(args.Filename).Funcs(funcs).Parse(tmpl)

	if err != nil {
		return []byte{}, nil
	}

	var b bytes.Buffer

	if err := t.Execute(&b, args); err != nil {
		return []byte{}, nil
	}

	return format.Source(b.Bytes())
}

var tmpl = `
package {{.Package}}

import (
	"fmt"
	"encoding/json"
)

// {{.Name}}Option contains a initialized or empty copy of the {{.Type}} type.
type {{.Name}}Option struct {
	v *{{.Type}}
}

// New{{.Name}}Value creates a new {{.Name}}Option type with an initialized value.
func New{{.Name}}Value(v {{.Type}}) {{.Name}}Option {
	return {{.Name}}Option{
		v: &v,
	}
}

// New{{.Name}}Value creates a new {{.Name}}Option type with the given value.
func New{{.Name}}Ptr(v *{{.Type}}) {{.Name}}Option {
	return {{.Name}}Option{
		v: v,
	}
}

// {{.Name}}2Ptr create a pointer from the given type.
func {{.Name}}2Ptr(v {{.Type}}) *{{.Type}} {
	return &v
}

// New{{.Name}}Value creates a new {{.Name}}Option type with an empty value.
func New{{.Name}}Empty() {{.Name}}Option {
	return {{.Name}}Option{}
}

// New{{.Name}}EmptyIfZeroValue creates a new initialized {{.Name}}Option type if the input {{.Type}} doesn't equal the {{.Type}}'s zero value, or an empty {{.Name}}Option otherwise.
func New{{.Name}}EmptyIfZeroValue(v {{.Type}}) {{.Name}}Option {
	var e {{.Type}}

	if v == e {
		return New{{.Name}}Empty()
	}

	return New{{.Name}}Value(v)
}

//IsEmpty returns true if the contained {{.Type}} value is empty, false otherwise.
func (t {{.Name}}Option) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained {{.Type}} value is initialized, false otherwise.
func (t {{.Name}}Option) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained {{.Type}} value.
//NOTE: If the value is empty, this will return the {{.Type}} zero value.
func (t {{.Name}}Option) Get() {{.Type}} {
	var v {{.Type}}

	if t.IsEmpty() {
		return v
	}
	
	return *t.v
}

//Get returns the contained {{.Type}} value if the contained value is initialized or the input value is the value is not initialized.
func (t {{.Name}}Option) GetOrElse(v {{.Type}}) {{.Type}} {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the {{.Name}}Option type into the JSON representation.
func (t {{.Name}}Option) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the {{.Name}}Option type.
func (t *{{.Name}}Option) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v {{.Type}}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the {{.Name}}Option type into the YAML representation.
func (t {{.Name}}Option) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}


//UnmarshalYAML unmarshals the YAML representation to the {{.Name}}Option type.
func (t *{{.Name}}Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v {{.Type}}

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *{{.Name}}Option) String() string {
	return fmt.Sprintf("%v", t.Get())
}
`
