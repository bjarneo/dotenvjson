package main

import (
	"fmt"

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

	AST := core.Compiler(data, kind)

	result := ""

	if kind == "json" && args.Kv != "" {
		result = core.JSONGenerator(core.TransformKV(AST, args.Kv), args.Pretty)
	}

	if kind == "json" && args.Kv == "" {
		result = core.JSONGenerator(AST, args.Pretty)
	}

	if kind == "yaml" {
		result = core.YAMLGenerator(AST)
	}

	if args.Output != "" {
		core.WriteFile(args.Output, result)
	}

	if args.PrintTerminal {
		fmt.Print(result)
	}
}
