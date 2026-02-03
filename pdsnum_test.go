package pdsnum_test

import (
	"testing"

	"github.com/koron-go/pdsnum"
)

func testEncode(t *testing.T, src, want string) {
	t.Helper()
	got, err := pdsnum.Encode(src)
	if err != nil {
		t.Errorf("failed: %s", err)
		return
	}
	if got != want {
		t.Errorf("unmatch: want=%s got=%s", want, got)
	}
}

func TestEncode(t *testing.T) {
	testEncode(t, "123", "_ 1 03 2 02 3 01 _")
}

func testDecode(t *testing.T, src, want string) {
	t.Helper()
	got, err := pdsnum.Decode(src)
	if err != nil {
		t.Errorf("failed: %s", err)
		return
	}
	if got != want {
		t.Errorf("unmatch: want=%s got=%s", want, got)
	}
}

func TestDecode(t *testing.T) {
	testDecode(t, "_ 1 03 2 02 3 01 _", "123")
}
