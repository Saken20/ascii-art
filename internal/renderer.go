package internal

import (
	"fmt"
	"strings"
)

type Renderer struct {
	banner Banner
}

func NewRenderer(bannerPath string) (*Renderer, error) {
	banner, err := LoadBanner(bannerPath)
	if err != nil {
		return nil, err
	}
	return &Renderer{banner: banner}, nil
}

func (r *Renderer) Render(input string) (string, error) {
	// Обработка \n как перевода строки (Go-стиль)
	lines := strings.Split(input, "\\n")
	var sb strings.Builder

	for _, line := range lines {
		// if idx > 0 {
		// 	sb.WriteString("\n")
		// }
		rendered, err := r.renderLine(line)
		if err != nil {
			return "", err
		}
		for _, l := range rendered {
			sb.WriteString(l)
			sb.WriteString("\n")
		}
	}
	return sb.String(), nil
}

func (r *Renderer) renderLine(line string) ([]string, error) {
	if line == "" {
		return make([]string, 1), nil
	}
	output := make([]string, charHeight)
	for _, ch := range line {
		art, ok := r.banner[ch]
		if !ok {
			return nil, fmt.Errorf("character %q not found in banner", ch)
		}
		for i := 0; i < charHeight; i++ {
			output[i] += art[i]
		}
	}
	return output, nil
}
