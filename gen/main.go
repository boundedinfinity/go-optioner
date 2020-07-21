package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

type data struct {
	Type    string
	Name    string
	Package string
	Path    string
}

func main() {
	var d data
	flag.StringVar(&d.Type, "type", "", "The type used for the optional being generated")
	flag.StringVar(&d.Name, "name", "", "The name used for the optinoal being generated. This should start with a capital letter so that it is exported.")
	flag.StringVar(&d.Package, "package", "", "The package used for the optinoal being generated.")
	flag.StringVar(&d.Path, "path", "", "The output path used for the optinoal being generated.")
	flag.Parse()

	if d.Package == "" {
		d.Package = "optional"
	}

	if d.Path == "" {
		d.Path = "."
	}

	n := fmt.Sprintf("optional_%v", d.Type)
	t := template.Must(template.New(n).Parse(queueTemplate))
	t.Execute(os.Stdout, d)
}

var queueTemplate = `
package {{.Package}}

import "encoding/json"

func New{{.Name}}(v *{{.Type}}) {{.Name}}Opt {
	return {{.Name}}Opt {
		v: v,
	}
}

type {{.Name}}Opt struct {
	v *{{.Type}}
}

func (t {{.Name}})Opt GetOrElse(v {{.Type}}) {{.Type}} {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

func (t {{.Name}}Opt) Get() {{.Type}} {
	var v {{.Type}}

	if t.IsDefined() {
		v = *t.v
	}

	return v
}

func (t {{.Name}}Opt) IsDefined() bool {
	var defined bool

	if t.v != nil {
		defined = true
	}

	return defined
}

func (t {{.Name}}Opt) IsEmpty() bool {
	return !t.IsDefined()
}

func (t {{.Name}}Opt) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

func (t *{{.Name}}Opt) UnmarshalJSON(data []byte) error {
	var v {{.Type}}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}

		t.v = &v
	}

	return nil
}
`
