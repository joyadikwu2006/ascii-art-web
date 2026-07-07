package asciiart

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// LoadBanner reads a banner .txt file and parses it into a
// map of rune to 8 rows of ASCII art strings.
func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file")
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")
	charMap := make(map[rune][]string)
	j := 0
	for i := ' '; i <= 126; i++ {
		for j < len(lines) && lines[j] == "" {
			j++
		}
		if j+8 > len(lines) {
			return nil, fmt.Errorf("invalid banner content")
		}
		charMap[i] = lines[j : j+8]
		j += 8
	}
	return charMap, nil
}

// GenerateArt takes a text string and a loaded banner map,
// and returns the full ASCII art representation as a string.
func GenerateArt(input string, banner map[rune][]string) string {
	var Art strings.Builder
	input = strings.ReplaceAll(input, "\n", "\\n")
	if input == "\\n" {
		return "\n"
	}

	if len(input) == 0 {
		return ""
	}

	char, err := ValidateInput(input)
	if char != 0 && err != nil {
		return "unsupported char"
	}
	line := Split(input)
	for i, ch := range line {
		if ch != "" {
			Art.WriteString(strings.Join(RenderLine(ch, banner), "\n"))
			Art.WriteString("\n")
		} else if i < len(line)-1 {
			Art.WriteRune('\n')
		}
	}

	return Art.String()
}

// RenderLine takes a single line of text and renders it
// as ASCII art using the provided charMap.
func RenderLine(input string, banner map[rune][]string) []string {
	style := make([]string, 8)
	for row := range 8 {
		var styleBuilder strings.Builder
		for _, i := range input {
			styleBuilder.WriteString(banner[i][row])
		}
		style[row] = styleBuilder.String()
	}
	return style
}

// Split breaks the input string on literal "\n" sequences.
func Split(input string) []string {
	return strings.Split(input, `\n`)
}

// ValidateInput checks that every character in the input
// is within the printable ASCII range (32–126).
// Returns the offending rune and an error if invalid.
func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char < rune(32) || char > 126 {
			return char, errors.New("unsupported input")
		}
	}
	return 0, nil
}
