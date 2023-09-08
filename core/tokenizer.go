package core

import (
	"bufio"
	"strings"
)

type Token struct {
	Kind  string
	Value string
}

func Tokenizer(content string) []Token {
	// The tokens
	_tokens := make([]Token, 0)

	// Tracking our position on each line
	current := 0

	scanner := bufio.NewScanner(strings.NewReader(content))

	// read it line by line
	for scanner.Scan() {
		line := scanner.Text()

		// This handles the comments that are on their own lines
		if strings.HasPrefix(line, "#") {
			_tokens = append(_tokens, Token{Kind: "Comment", Value: line})

			continue
		}

		lineLength := len(line)
		for current < lineLength {
			char := string(line[current])

			// Find the key by starting from the first letter of the line
			if current == 0 {
				key := ""

				// Go on until we find the variable assignment
				for string(line[current]) != "=" {
					key += char

					current++

					char = string(line[current])
				}

				_tokens = append(_tokens, Token{Kind: "Key", Value: strings.TrimSpace(key)})
			}

			char = string(line[current])

			// Find the variable value
			if char == "=" {
				// The variable value
				value := ""

				// If we find a line comment
				comment := ""

				hashAllowed := false

				// Skip the "="
				current++

				// If the env variable is empty
				if len(line) <= current {
					_tokens = append(_tokens, Token{Kind: "Value", Value: ""})

					break
				}

				char = string(line[current])

				for current < lineLength {
					// If the variable value is wrapped in quotes
					// and we hit a "#", then "#" is allowed
					if char == "\"" || char == "'" {
						hashAllowed = true
					}

					// I.e we hit a space and we are not in a quote
					// and this is a password with the hash
					// then we allow the hash
					if !hashAllowed &&
						string(line[current]) == "#" &&
						current + 1 <= len(line) &&
						string(line[current + 1]) != " " &&
						string(line[current - 1]) != " " {
						hashAllowed = true
					}

					// Skip everything if we hit a comment
					if !hashAllowed && char == "#" {
						comment += string(line[current])

						current++

						// Add a comment to the tokens if it exists
						// and we are at the end of line
						if current == lineLength && comment != "" {
							_tokens = append(_tokens, Token{Kind: "Comment", Value: strings.TrimSpace(comment)})
						}

						continue
					}

					value += char

					current++

					// Important to break here so we do not break the line length index
					if current == lineLength {
						break
					}

					char = string(line[current])
				}

				_tokens = append(_tokens, Token{Kind: "Value", Value: strings.TrimSpace(value)})
			}

			if current == lineLength {
				break
			}
		}

		// reset the tracker for each line, this has to be the last thing we do
		current = 0
	}

	return _tokens
}
