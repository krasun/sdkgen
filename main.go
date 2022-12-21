package main

import (
	"fmt"
	"log"
	"os"


	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	specLocation := os.Args[1]

	loader := openapi3.NewLoader()
	spec, err := loader.LoadFromFile(specLocation)
	if err != nil {
		log.Fatalf("failed to load the OpenAPI specification file from \"%s\"", specLocation)
		return
	}

	path := spec.Paths.Find("/take")
	if path == nil {
		log.Fatalf("path \"/take\" not found")
		return
	}

	for _, parameter := range path.Get.Parameters {
		fmt.Println(parameter.Value.Name)		
	}
}
