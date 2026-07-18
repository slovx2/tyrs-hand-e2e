package greeting

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("Codex"); got != "Hello, Codex!" {
		t.Fatalf("Hello() = %q", got)
	}
}

func TestWelcome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "name", input: "Codex", want: "Welcome, Codex!"},
		{name: "trim whitespace", input: "  Codex\t", want: "Welcome, Codex!"},
		{name: "empty", input: "", want: "Welcome!"},
		{name: "whitespace only", input: " \t\n", want: "Welcome!"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Welcome(test.input); got != test.want {
				t.Fatalf("Welcome(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}
