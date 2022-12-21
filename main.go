package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	specLocation := os.Args[1]

	loader := openapi3.NewLoader()
	spec, err := loader.LoadFromFile(specLocation)
	if err != nil {
		log.Fatalf("failed to load the OpenAPI specification file from \"%s\": %s", specLocation, err)
		return
	}

	path := spec.Paths.Find("/take")
	if path == nil {
		log.Fatalf("path \"/take\" not found")
		return
	}

	tpl := `
def {{.Name}}(self, {{ if eq .Schema.Value.Type "array" }}values: List[{{ .Schema.Value.Items.Value.Type }}]{{ else }}value{{ end }}): 	
    {{ if eq .Schema.Value.Type "array" }}self.options['{{.Name}}'] = values{{ else }}self.options['{{.Name}}'] = value{{ end }}

	return self
`

	t, err := template.New("python_parameter_function").Parse(tpl)
	if err != nil {
		log.Fatalf("failed to parse template: %s", err)
		return
	}

	buf := new(bytes.Buffer)

	for _, parameter := range path.Get.Parameters {
		t.Execute(buf, parameter.Value)		
	}

	fmt.Println(buf.String())
}
