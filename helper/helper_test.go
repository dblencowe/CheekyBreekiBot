package helper

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerateImage(t *testing.T) {
	dat, _ := os.ReadFile("../test/full_graph_html.html")
	html := string(dat)
	GenerateImageFromHtml(html)
	fi, err := os.Stat("/tmp/example.png")
	if err != nil {
		t.Fatalf("GenerateImage(...) returned an error: %v, error\n", err)
	}

	if fi.Size() <= 0 {
		t.Fatalf("GenerateImage().Size() = %d, want X > 0\n", fi.Size())
	}
	fmt.Printf("GenerateImage().Size() = %d\n", fi.Size())
}
