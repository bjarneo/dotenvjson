package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bjarneo/dotenvjson/core"
)

func pipeInput() string {
	data := ""

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		data += scanner.Text() + "\n"
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	pretty := flag.Bool("p", false, "Pretty print the output. Default 'false'")
	printTerminal := flag.Bool("pt", true, "Print to terminal. Default 'true'.")
	output := flag.String("o", "", "Save the output to file. -o=file.json")
	yaml := flag.Bool("y", false, "Transform it to YAML")

	flag.Parse()

	// Get the dotenv file
	filename := flag.Arg(0)

	data := ""
	if filename != "" {
		data = core.FileContent(filename)
	} else {
		data = pipeInput()
	}

	kind := ""
	if *yaml {
		kind = "yaml"
	} else {
		kind = "json"
	}

	generatedData, AST := core.Compiler(data, kind)

	result := ""
	if *pretty && kind == "json" {
		prettyPrint, err := json.MarshalIndent(AST, "", "  ")

		if err != nil {
			log.Fatalf(" [!] Something happened while creating the pretty print JSON \n %s", err)
		}

		result = string(prettyPrint)
	} else {
		result = generatedData
	}

	if *output != "" {
		core.WriteFile(*output, result)
	}

	if *printTerminal {
		fmt.Print(result)
	}
}
