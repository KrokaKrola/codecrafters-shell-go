package tokenizer

import (
	"strings"
	"unicode"
)

type tokenType string

const (
	word       tokenType = "word"
	Whitespace tokenType = "whitespace"
)

type Token struct {
	Type  tokenType
	Value string
}

func tokenize(input string) ([]Token, error) {
	var tokens []Token

	pos := 0

	for pos < len(input) {
		ch := input[pos]
		switch {
		case unicode.IsSpace(rune(ch)):
			if len(tokens) > 0 && tokens[len(tokens)-1].Type == Whitespace {
				pos += 1
				continue
			}

			tokens = append(tokens, Token{Type: Whitespace})
			pos += 1
		case ch == '\'':
			sb := strings.Builder{}

			// starting with a quote, so we are moving to the inner content
			pos += 1
			ch = input[pos]

			for pos < len(input) && ch != '\'' {
				if err := sb.WriteByte(ch); err != nil {
					return nil, err
				}

				pos += 1

				if pos < len(input) {
					ch = input[pos]
				}
			}

			// found closing quote, moving to the next char
			pos += 1

			if sb.Len() == 0 {
				continue
			}

			tokens = append(tokens, Token{Type: word, Value: sb.String()})
		default:
			sb := strings.Builder{}

			for pos < len(input) && input[pos] != '\'' && !unicode.IsSpace(rune(input[pos])) {
				if err := sb.WriteByte(input[pos]); err != nil {
					return nil, err
				}
				pos += 1
			}

			tokens = append(tokens, Token{Type: word, Value: sb.String()})
		}
	}

	if len(tokens) > 0 && tokens[len(tokens)-1].Type == Whitespace {
		return tokens[:len(tokens)-1], nil
	}

	return tokens, nil
}

func process(tokens []Token) ([]Token, error) {
	var result []Token

	pos := 0

	for pos < len(tokens) {
		token := tokens[pos]

		switch token.Type {
		case word:
			sb := strings.Builder{}

			sb.WriteString(token.Value)

			for pos+1 < len(tokens) && tokens[pos+1].Type == word {
				sb.WriteString(tokens[pos+1].Value)
				pos += 1
			}

			pos += 1

			if len(result) > 0 && result[len(result)-1].Type == word {
				result[len(result)-1].Value += sb.String()
			} else {
				result = append(result, Token{Type: word, Value: sb.String()})
			}
		case Whitespace:
			pos += 1

			result = append(result, token)
		}

	}

	return result, nil
}

func Tokenize(input string) ([]string, error) {
	tokens, err := tokenize(input)
	if err != nil {
		return nil, err
	}

	tokens, err = process(tokens)
	if err != nil {
		return nil, err
	}

	var result []string

	for _, token := range tokens {
		if token.Type != Whitespace {
			result = append(result, token.Value)
		}
	}

	return result, nil
}
