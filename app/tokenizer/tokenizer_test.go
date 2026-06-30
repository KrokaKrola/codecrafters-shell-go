package tokenizer_test

import (
	"slices"
	"testing"

	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		// base test cases
		{name: "basic case", input: "echo hello world", want: []string{"echo", "hello", "world"}},
		{name: "consequtive spaces are collapsed unless quoted", input: "echo hello    world", want: []string{"echo", "hello", "world"}},
		{name: "sanitizes space in the end of the input", input: "echo hello ", want: []string{"echo", "hello"}},
		{name: "sanitizes space in the begining of the input", input: " echo hello ", want: []string{"echo", "hello"}},

		{name: "spaces are preserved within single quotes", input: "echo 'hello    world'", want: []string{"echo", "hello    world"}},
		{name: "adjacent single quoted strings are concatenated", input: "echo 'hello''world'", want: []string{"echo", "helloworld"}},
		{name: "empty single quotes are ignored", input: "echo hello''world", want: []string{"echo", "helloworld"}},
		{name: "adjacent single quoted words concatenates into one arg", input: "echo hello'world'", want: []string{"echo", "helloworld"}},

		{name: "spaces are preserved within double quotes", input: "echo \"hello    world\"", want: []string{"echo", "hello    world"}},
		{name: "adjacent double quoted strings are concatenated", input: "echo \"hello\"\"world\"", want: []string{"echo", "helloworld"}},
		{name: "empty double quotes are ignored", input: "echo hello\"\"world", want: []string{"echo", "helloworld"}},
		{name: "adjacent double quoted words concatenates into one arg", input: "echo hello\"world\"", want: []string{"echo", "helloworld"}},

		{name: "each \\ creates a literal space as a part of one argument", input: "echo three\\ \\ \\ spaces", want: []string{"echo", "three   spaces"}},
		{name: "backslah preserves the first space literaly, but the shell collapses the subsequent unescaped spaces", input: "echo before\\     after", want: []string{"echo", "before ", "after"}},
		{name: "\\n becomes just n", input: "echo test\\nexample", want: []string{"echo", "testnexample"}},
		{name: "first backslash escapes the second, and the result is a single literal backslash", input: "echo hello\\\\world", want: []string{"echo", "hello\\world"}},
		{name: "\\' makes the single quotes literal characters", input: "echo \\'hello\\'", want: []string{"echo", "'hello'"}},

		{name: "backslah within double qoutes", input: "echo \"hello'example'\\\\'shell\"", want: []string{"echo", "hello'example'\\'shell"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, gotErr := tokenizer.Tokenize(tt.input)

			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Tokenize() failed: %#v", gotErr)
				}
				return
			}

			if tt.wantErr {
				t.Fatal("Tokenize() succeeded unexpectedly")
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("Tokenize() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
