package core

import (
	"encoding/json"
	"log"
	"strings"
)

func JSONGenerator(AST any) string {
	json, err := json.Marshal(AST)

	if err != nil {
		log.Fatal(err)
	}

	// return the json as a string
	return string(json)
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

func PrettyPrint(AST any) string {
	prettyPrint, err := json.MarshalIndent(AST, "", "  ")

	if err != nil {
		log.Fatalf(" [!] Something happened while creating the pretty print JSON \n %s", err)
	}

	return string(prettyPrint)
}
