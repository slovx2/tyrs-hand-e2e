package greeting

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("Codex"); got != "Hello, Codex!" {
		t.Fatalf("Hello() = %q", got)
	}
}

func TestFriendlyHello(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "name", input: "Codex", want: "Hello, Codex!"},
		{name: "surrounding whitespace", input: "  Codex\t", want: "Hello, Codex!"},
		{name: "empty", input: "", want: "Hello, friend!"},
		{name: "whitespace only", input: " \t\n", want: "Hello, friend!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FriendlyHello(tt.input); got != tt.want {
				t.Fatalf("FriendlyHello(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
