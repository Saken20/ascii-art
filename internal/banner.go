package internal

import (
	"bufio"
	"fmt"
	"os"
)

const (
	asciiStart = 32  // ' '
	asciiEnd   = 126 // '~'
	charHeight = 8
)

type Banner map[rune][]string

func LoadBanner(path string) (Banner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open banner file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	banner := make(Banner)
	currentChar := rune(asciiStart)

	for {
		// Пропускаем разделительную строку (если есть)
		if !scanner.Scan() {
			break
		}
		lines := make([]string, 0, charHeight)
		for i := 0; i < charHeight; i++ {
			if !scanner.Scan() {
				if currentChar <= asciiEnd {
					return nil, fmt.Errorf("unexpected EOF at character %q", currentChar)
				}
				return banner, nil
			}
			lines = append(lines, scanner.Text())
		}
		if currentChar > asciiEnd {
			return nil, fmt.Errorf("excessive characters in banner file")
		}
		banner[currentChar] = lines
		currentChar++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}
	if currentChar <= asciiEnd {
		return nil, fmt.Errorf("missing characters from %q to ~", currentChar)
	}
	return banner, nil
}
