package core

import "fmt"

func Transformer(tokens []Token) map[string]interface{} {
	tokenData := make(map[string]interface{})

	prevKey := ""
	for _, token := range tokens {
		if token.Kind == "Key" {
			prevKey = token.Value
		}

		if token.Kind == "Value" {
			intToken := 0
			_, err := fmt.Sscan(token.Value, &intToken)

			// if the error is nil, then it is converted to an integer
			// to be used for i.e. port numbers
			if err == nil {
				tokenData[prevKey] = intToken

				continue
			}

			tokenData[prevKey] = token.Value
		}
	}

	return tokenData
}
