package exporter_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/kieranajp/langoustine/pkg/domain"
	"github.com/kieranajp/langoustine/pkg/exporter"
)

func TestImageGeneration(t *testing.T) {
	c := exporter.NewCoverImage()
	recipe := domain.Recipe{
		Name: "Test Recipe",
	}

	err := c.Generate(&recipe)
	if err != nil {
		t.Fatalf("failed to generate recipe image: %v", err)
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	_, err = c.Write(w)
	if err != nil {
		t.Fatalf("failed to get recipe image bytes: %v", err)
	}
}
