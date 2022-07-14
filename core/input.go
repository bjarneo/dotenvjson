package core

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func PipeInput() string {
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

type inputs struct {
	Pretty        bool
	PrintTerminal bool
	Yaml          bool
	Output        string
	Filename      string
}

func Input() inputs {
	pretty := flag.Bool("p", false, "Pretty print the output. Default 'false'")
	printTerminal := flag.Bool("pt", true, "Print to terminal. Default 'true'.")
	output := flag.String("o", "", "Save the output to file. -o=file.json")
	yaml := flag.Bool("y", false, "Transform it to YAML")

	flag.Parse()

	// Get the dotenv file
	filename := flag.Arg(0)

	return inputs{
		Pretty:        *pretty,
		PrintTerminal: *printTerminal,
		Yaml:          *yaml,
		Output:        *output,
		Filename:      filename,
	}
}
