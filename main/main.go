package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/styles"
)

func main() {
	text, err := os.ReadFile("text.md")
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(styles.DarkStyleConfig),
		glamour.WithWordWrap(0),
	)
	if err != nil {
		log.Fatalln()
	}

	buf, err := renderer.Render(string(text))
	if err != nil {
		log.Fatalln()
	}

	fmt.Println(buf)
}
