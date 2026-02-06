package main

import (
	"os"
	"path/filepath"
	"testing"
)

func assertEqualFileContents(t *testing.T, wantFile, gotFile string) {
	t.Helper()
	wantBytes, err := os.ReadFile(wantFile)
	if err != nil {
		t.Fatalf("failed to read want file: %s", err)
		return
	}
	want := string(wantBytes)
	gotBytes, err := os.ReadFile(gotFile)
	if err != nil {
		t.Fatalf("failed to read got file: %s", err)
		return
	}
	got := string(gotBytes)
	if got != want {
		t.Errorf("unexpected contents:\nwant=%q\ngot=%s", want, got)
	}
}

func TestEncode(t *testing.T) {
	in, err := os.Open("testdata/test1.original")
	if err != nil {
		t.Fatalf("failed to open: %s", err)
	}
	defer in.Close()

	tmpdir := t.TempDir()
	gotFile := filepath.Join(tmpdir, "encoded.txt")
	out, err := os.Create(gotFile)
	if err != nil {
		t.Fatalf("failed to create: %s", err)
	}

	err = encode(out, in)
	out.Close()
	if err != nil {
		t.Fatalf("failed to encode: %s", err)
	}

	const wantFile = "testdata/test1.encoded"
	assertEqualFileContents(t, wantFile, gotFile)
}

func TestDecode(t *testing.T) {
	in, err := os.Open("testdata/test1.encoded")
	if err != nil {
		t.Fatalf("failed to open: %s", err)
	}
	defer in.Close()

	tmpdir := t.TempDir()
	gotFile := filepath.Join(tmpdir, "decoded.txt")
	out, err := os.Create(gotFile)
	if err != nil {
		t.Fatalf("failed to create: %s", err)
	}

	err = decode(out, in)
	out.Close()
	if err != nil {
		t.Fatalf("failed to encode: %s", err)
	}

	const wantFile = "testdata/test1.original"
	assertEqualFileContents(t, wantFile, gotFile)
}
