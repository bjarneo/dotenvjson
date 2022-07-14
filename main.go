package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bjarneo/dotenvjson/core"
)

func main() {
	input := core.Input()

	data := ""
	if input.Filename != "" {
		data = core.FileContent(input.Filename)
	} else {
		data = core.PipeInput()
	}

	kind := ""
	if input.Yaml {
		kind = "yaml"
	} else {
		kind = "json"
	}

	generatedData, AST := core.Compiler(data, kind)

	result := ""
	if input.Pretty && kind == "json" {
		prettyPrint, err := json.MarshalIndent(AST, "", "  ")

		if err != nil {
			log.Fatalf(" [!] Something happened while creating the pretty print JSON \n %s", err)
		}

		result = string(prettyPrint)
	} else {
		result = generatedData
	}

	if input.Output != "" {
		core.WriteFile(input.Output, result)
	}

	if input.PrintTerminal {
		fmt.Print(result)
	}
}
