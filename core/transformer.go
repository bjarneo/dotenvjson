package core

func Transformer(tokens []Token) map[string]interface{} {
	tokenData := make(map[string]interface{})

	prevKey := ""
	for _, token := range tokens {
		if token.Kind == "Key" {
			prevKey = token.Value
		}

		if token.Kind == "Value" {
			tokenData[prevKey] = token.Value
		}
	}

	return tokenData
}
