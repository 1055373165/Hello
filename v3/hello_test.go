package Hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello smy this is v3"
	if got := Hello("smy"); got != want {
		t.Errorf("Hello(\"smy\") = %q, want %q", got, want)
	}
}
