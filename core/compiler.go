package core

func Compiler(input string, kind string) map[string]interface{} {
	// Create tokens
	tokens := Tokenizer(input)
	// Transform the token into an "AST" (not really an AST though)
	AST := Transformer(tokens)

	// return the JSON string
	return AST
}
