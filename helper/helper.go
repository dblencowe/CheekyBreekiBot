package helper

import (
	"fmt"
	"os"
	"strings"
)

func GenerateImageFromHtml(html string) []byte {
	c := WKHTMLImageOptions{Input: "-", Format: "png", Html: html, BinaryPath: "/usr/local/bin/wkhtmltoimage"}
	out, err := GenerateImage(&c)
	fmt.Printf("Err: %+v\n", err)

	CreateFile("/tmp/example.png", out)

	return out
}

func Levenshtein(s1 string, s2 string) (distance int) {
	target := []rune(strings.ToLower(s1))
	compareTo := []rune(strings.ToLower(s2))

	rows := len(target) + 1
	cols := len(compareTo) + 1

	var d1, d2, d3, i, j int
	dist := make([]int, rows*cols)

	for i = 0; i < rows; i++ {
		dist[i*cols] = i
	}

	for j = 0; j < cols; j++ {
		dist[j] = j
	}

	for j = 1; j < cols; j++ {
		for i = 1; i < rows; i++ {
			if target[i-1] == compareTo[j-1] {
				dist[(i*cols)+j] = dist[((i-1)*cols)+(j-1)]
			} else {
				d1 = dist[((i-1)*cols)+j] + 1
				d2 = dist[(i*cols)+(j-1)] + 1
				d3 = dist[((i-1)*cols)+(j-1)] + 1

				dist[(i*cols)+j] = min(d1, min(d2, d3))
			}
		}
	}

	distance = dist[(cols*rows)-1]

	return
}

func min(a int, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}

func Contains(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}

	return -1, false
}

func CreateFile(path string, content []byte) {
	f, _ := os.Create(path)
	defer f.Close()
	f.Write(content)
}
