package core

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
)

func JSONGenerator(AST any, indent bool) string {
	var buf bytes.Buffer

	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)

	if indent {
		encoder.SetIndent("", "  ")
	}

	if err := encoder.Encode(AST); err != nil {
		log.Fatal(err)
	}

	// return the json as a string
	return buf.String()
}

type M map[string]interface{}

func TransformKV(AST map[string]interface{}, kv string) []M {
	var output []M

	kvData := strings.Split(kv, ",")

	for key, value := range AST {
		output = append(output, M{kvData[0]: key, kvData[1]: value})
	}

	// return the json as a string
	return output
}

