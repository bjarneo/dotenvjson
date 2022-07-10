package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/bjarneo/dotenvjson/core"
)

func main() {
	pretty := flag.Bool("p", false, "Pretty print the output. Default 'false'")
	printTerminal := flag.Bool("pt", true, "Print to terminal. Default 'true'.")
	output := flag.String("o", "", "Save the output to file. -o=file.json")

	flag.Parse()

	// Get the dotenv file
	filename := flag.Arg(0)

	if filename == "" {
		log.Fatal(" [!] The file is empty")
	}

	JSON, AST := core.Compiler(core.FileContent(filename))

	jsonOutput := ""

	if *pretty {
		prettyPrint, err := json.MarshalIndent(AST, "", "  ")

		if err != nil {
			log.Fatalf(" [!] Something happened while creating the pretty print JSON \n %s", err)
		}

		jsonOutput = string(prettyPrint)
	} else {
		jsonOutput = string(JSON)
	}

	if *output != "" {
		core.WriteFile(*output, jsonOutput)
	}

	if *printTerminal {
		fmt.Print(jsonOutput)
	}
}
