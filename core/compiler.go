package core

func Compiler(input string) (string, map[string]interface{}) {
	// Create tokens
	tokens := Tokenizer(input)
	// Transform the token into an "AST" (not really an AST though)
	AST := Transformer(tokens)
	// Generate JSON from the "AST"
	JSON := JSONGenerator(AST)

	// return the JSON string
	return JSON, AST
}
