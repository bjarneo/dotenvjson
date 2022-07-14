package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bjarneo/dotenvjson/core"
)

func outputStrategy(args core.Args) string {
	if args.Yaml {
		return "yaml"
	}

	return "json"
}

func inputStrategy(args core.Args) string {
	if args.Filename != "" {
		return core.FileContent(args.Filename)
	}

	return core.PipeInput()
}

func main() {
	args := core.Arg()
	data := inputStrategy(args)
	kind := outputStrategy(args)

	generatedData, AST := core.Compiler(data, kind)

	result := ""
	if args.Pretty && kind == "json" {
		prettyPrint, err := json.MarshalIndent(AST, "", "  ")

		if err != nil {
			log.Fatalf(" [!] Something happened while creating the pretty print JSON \n %s", err)
		}

		result = string(prettyPrint)
	} else {
		result = generatedData
	}

	if args.Output != "" {
		core.WriteFile(args.Output, result)
	}

	if args.PrintTerminal {
		fmt.Print(result)
	}
}
