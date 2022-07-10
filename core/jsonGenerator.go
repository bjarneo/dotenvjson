package core

import (
	"encoding/json"
	"log"
)

func JSONGenerator(ast map[string]interface{}) string {
	json, err := json.Marshal(ast)

	if err != nil {
		log.Fatal(err)
	}

	// return the json as a string
	return string(json)
}
