package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/wire20022/pkg/messages"
)

func TestProcessDirectoryWithoutRecursiveDoesNotPanic(t *testing.T) {
	dir := t.TempDir()
	nested := filepath.Join(dir, "nested")
	if err := os.Mkdir(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "foo.xml"), []byte(`<x/>`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(nested, "bar.xml"), []byte(`<x/>`), 0o644); err != nil {
		t.Fatal(err)
	}

	origPattern, origRecursive := pattern, recursive
	t.Cleanup(func() {
		pattern, recursive = origPattern, origRecursive
	})
	pattern = "*.xml"
	recursive = false

	result := processDirectory(messages.NewUniversalReader(), dir)
	if result.TotalFiles != 1 {
		t.Fatalf("processed %d files, want 1 (nested files should be skipped without -r)", result.TotalFiles)
	}
	if got := filepath.Base(result.Results[0].File); got != "foo.xml" {
		t.Fatalf("processed %q, want foo.xml", got)
	}
}

func TestProcessDirectoryRecursiveIncludesNestedFiles(t *testing.T) {
	dir := t.TempDir()
	nested := filepath.Join(dir, "nested")
	if err := os.Mkdir(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "foo.xml"), []byte(`<x/>`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(nested, "bar.xml"), []byte(`<x/>`), 0o644); err != nil {
		t.Fatal(err)
	}

	origPattern, origRecursive := pattern, recursive
	t.Cleanup(func() {
		pattern, recursive = origPattern, origRecursive
	})
	pattern = "*.xml"
	recursive = true

	result := processDirectory(messages.NewUniversalReader(), dir)
	if result.TotalFiles != 2 {
		t.Fatalf("processed %d files, want 2", result.TotalFiles)
	}
}
