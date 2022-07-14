package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bjarneo/dotenvjson/core"
)

func outputStrategy(input core.Inputs) string {
	if input.Yaml {
		return "yaml"
	}

	return "json"
}

func inputStrategy(input core.Inputs) string {
	if input.Filename != "" {
		return core.FileContent(input.Filename)
	}

	return core.PipeInput()
}

func main() {
	input := core.Input()
	data := inputStrategy(input)
	kind := outputStrategy(input)

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
