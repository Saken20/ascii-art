package internal

import (
	"strings"
	"testing"
)

func TestRenderHello(t *testing.T) {
	renderer, err := NewRenderer("banners/standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}
	got, err := renderer.Render("hello")
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(got), "\n")
	if len(lines) != charHeight {
		t.Errorf("expected %d lines, got %d", charHeight, len(lines))
	}
	if len(strings.TrimSpace(got)) == 0 {
		t.Error("output is empty")
	}
}

func TestRenderMultiLine(t *testing.T) {
	renderer, err := NewRenderer("banners/standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}
	got, err := renderer.Render("Hi\\nThere")
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	blocks := strings.Split(got, "\n\n")
	if len(blocks) < 2 {
		t.Error("expected at least two blocks for multi-line input")
	}
}

func TestRenderUnknownChar(t *testing.T) {
	renderer, err := NewRenderer("banners/standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}
	_, err = renderer.Render("hello☺")
	if err == nil {
		t.Error("expected error for unknown character, got nil")
	}
}
