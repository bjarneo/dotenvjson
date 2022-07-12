package core

func Compiler(input string, kind string) (string, map[string]interface{}) {
	// predefine our output string
	output := ""

	// Create tokens
	tokens := Tokenizer(input)
	// Transform the token into an "AST" (not really an AST though)
	AST := Transformer(tokens)

	if kind == "json" {
		// Generate JSON from the "AST"
		output = JSONGenerator(AST)
	}

	if kind == "yaml" {
		output = YAMLGenerator(AST)
	}

	// return the JSON string
	return output, AST
}
