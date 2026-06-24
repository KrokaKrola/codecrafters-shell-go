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
		{name: "basic case", input: "echo hello world", want: []string{"echo", "hello", "world"}},
		{name: "consequtive spaces are collapsed unless quoted", input: "echo hello    world", want: []string{"echo", "hello", "world"}},
		{name: "spaces are preserved within quotes", input: "echo 'hello    world'", want: []string{"echo", "hello    world"}},
		{name: "adjacent quoted strings are concatenated", input: "echo 'hello''world'", want: []string{"echo", "helloworld"}},
		{name: "empty quotes are ignored", input: "echo hello''world", want: []string{"echo", "helloworld"}},
		{name: "adjacent quoted words concatenates into one arg", input: "echo hello'world'", want: []string{"echo", "helloworld"}},
		{name: "sanitizes space in the end of the input", input: "echo hello ", want: []string{"echo", "hello"}},
		{name: "sanitizes space in the begining of the input", input: " echo hello ", want: []string{"echo", "hello"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tokenizer.Tokenize(tt.input)

			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Tokenize() failed: %v", gotErr)
				}
				return
			}

			if tt.wantErr {
				t.Fatal("Tokenize() succeeded unexpectedly")
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
