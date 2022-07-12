package core

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func YAMLGenerator(ast map[string]interface{}) string {
	yamlData, err := yaml.Marshal(ast)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	return string(yamlData)
}
