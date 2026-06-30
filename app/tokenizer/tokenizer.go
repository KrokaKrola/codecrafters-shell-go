package tokenizer

import (
	"fmt"
	"strings"
	"unicode"
)

type tokenType string

const (
	Word           tokenType = "word"
	whitespace     tokenType = "whitespace"
	StdoutRedirect tokenType = "stdoutRedirect"
	StderrRedirect tokenType = "stderrRedirect"
)

const (
	singleQuote  byte = '\''
	doubleQuote  byte = '"'
	escapeChar   byte = '\\'
	dollarChar   byte = '$'
	backtickChar byte = '`'
	newLineChar  byte = '\n'
)

type Token struct {
	Type  tokenType
	Value string
}

func isDoubleQuoteEscape(ch byte) bool {
	return ch == doubleQuote || ch == escapeChar || ch == dollarChar || ch == backtickChar || ch == newLineChar
}

func isChar(input string, pos int) bool {
	return input[pos] != singleQuote && input[pos] != doubleQuote && !unicode.IsSpace(rune(input[pos]))
}

func tokenize(input string) ([]Token, error) {
	var tokens []Token

	pos := 0

	for pos < len(input) {
		ch := input[pos]

		switch {
		case unicode.IsSpace(rune(ch)):
			if len(tokens) > 0 && tokens[len(tokens)-1].Type == whitespace {
				pos += 1
				continue
			}

			tokens = append(tokens, Token{Type: whitespace})
			pos += 1
		case ch == singleQuote, ch == doubleQuote:
			sb := strings.Builder{}

			quoteType := singleQuote

			if ch == doubleQuote {
				quoteType = doubleQuote
			}

			// starting with a quote, so we are moving to the inner content
			pos += 1
			ch = input[pos]

			for pos < len(input) && ch != quoteType {
				if ch == escapeChar && quoteType == doubleQuote {
					next := pos + 1
					if next < len(input) && isDoubleQuoteEscape(input[next]) {
						pos += 1
						ch = input[pos]
					}
				}

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

			tokens = append(tokens, Token{Type: Word, Value: sb.String()})
		default:
			sb := strings.Builder{}

			isBackslashed := false

			for pos < len(input) {
				if !isChar(input, pos) && !isBackslashed {
					break
				}

				if input[pos] == escapeChar && !isBackslashed {
					isBackslashed = true
					pos += 1
					continue
				}

				if err := sb.WriteByte(input[pos]); err != nil {
					return nil, err
				}

				if isBackslashed {
					isBackslashed = false
				}

				pos += 1
			}

			tokens = append(tokens, Token{Type: Word, Value: sb.String()})
		}
	}

	if len(tokens) > 0 && tokens[len(tokens)-1].Type == whitespace {
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
		case Word:
			sb := strings.Builder{}

			sb.WriteString(token.Value)

			// handle case, when there is no whitespace between current word and next word
			for pos+1 < len(tokens) && tokens[pos+1].Type == Word {
				sb.WriteString(tokens[pos+1].Value)
				pos += 1
			}

			pos += 1

			value := sb.String()

			switch value {
			case ">", "1>", "2>":
				if pos+1 >= len(tokens) || tokens[pos+1].Type != Word {
					return nil, fmt.Errorf("Expected a string, but found end of the input")
				}

				t := StdoutRedirect
				if value == "2>" {
					t = StderrRedirect
				}

				result = append(result, Token{Type: t, Value: tokens[pos+1].Value})
				pos += 2
			default:
				result = append(result, Token{Type: Word, Value: sb.String()})
			}
		case whitespace:
			pos += 1

			result = append(result, token)
		default:
			return nil, fmt.Errorf("Unknown token type: %s, value=%q", token.Type, token.Value)
		}
	}

	return result, nil
}

func Tokenize(input string) ([]string, *Token, *Token, error) {
	var stdoutRedirect *Token
	var stderrRedirect *Token
	tokens, err := tokenize(input)
	if err != nil {
		return nil, nil, nil, err
	}

	tokens, err = process(tokens)
	if err != nil {
		return nil, nil, nil, err
	}

	var result []string

	for idx, token := range tokens {
		if idx == 0 && token.Type != Word {
			return nil, nil, nil, fmt.Errorf("Expected a string, but found a: %s", token.Type)
		}

		if token.Type == StdoutRedirect {
			stdoutRedirect = &Token{Type: token.Type, Value: token.Value}
			continue
		}

		if token.Type == StderrRedirect {
			stderrRedirect = &Token{Type: token.Type, Value: token.Value}
			continue
		}

		if token.Type != whitespace {
			result = append(result, token.Value)
		}
	}

	return result, stdoutRedirect, stderrRedirect, nil
}
