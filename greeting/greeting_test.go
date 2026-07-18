package greeting

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("Codex"); got != "Hello, Codex!" {
		t.Fatalf("Hello() = %q", got)
	}
}
