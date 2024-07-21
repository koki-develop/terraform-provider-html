package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

//go:embed elements.yaml
var elementsYAML []byte

//go:embed resource.go.tmpl
var resourceGoTmpl []byte

type Schema struct {
	Elements         []Element   `yaml:"elements"`
	GlobalAttributes []Attribute `yaml:"global_attributes"`
}

type Element struct {
	Name        string      `yaml:"name"`
	URL         string      `yaml:"url"`
	Description string      `yaml:"description"`
	Attributes  []Attribute `yaml:"attributes"`
}

type Attribute struct {
	Name        string `yaml:"name"`
	URL         string `yaml:"url"`
	Description string `yaml:"description"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var schema Schema
	if err := yaml.Unmarshal(elementsYAML, &schema); err != nil {
		return fmt.Errorf("failed to unmarshal elements.yaml: %w", err)
	}

	for _, elm := range schema.Elements {
		p := fmt.Sprintf("./internal/resources/%s.generated.go", elm.Name)
		f, err := os.Create(p)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", p, err)
		}
		defer f.Close()

		funcMap := map[string]interface{}{
			"title": cases.Title(language.Dutch).String,
			"escape": func(s string) string {
				return strings.ReplaceAll(s, "\"", "\\\"")
			},
		}

		t, err := template.New("resource").Funcs(funcMap).Parse(string(resourceGoTmpl))
		if err != nil {
			return fmt.Errorf("failed to parse template: %w", err)
		}
		if err := t.Execute(f, map[string]interface{}{
			"Element":          elm,
			"GlobalAttributes": schema.GlobalAttributes,
		}); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
	}

	return nil
}
