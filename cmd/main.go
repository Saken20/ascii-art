package main

import (
	"ascii-art/internal"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [STRING]")
		return
	}

	input := os.Args[1]
	bannerPath := "./banners/standard.txt"

	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "shadow":
			bannerPath = "../banners/shadow.txt"
		}
		case "thinkertoy":
            bannerPath = "../banners/thinkertoy.txt"
        case "standard":
            bannerPath = "../banners/standard.txt"
        default:
            fmt.Println("Unknown banner, using standard.")
	}

	renderer, err := internal.NewRenderer(bannerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner: %v\n", err)
		os.Exit(1)
	}

	output, err := renderer.Render(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Render error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(output)
}
