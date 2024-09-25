package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/styles"
)

func main() {
	text := `
## Hello World

This is some text.

 + [link](https://google.com/1)
 + https://google.com/2
 + [https://google.com/3](https://google.com/3)

	`
	renderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(styles.DarkStyleConfig),
		glamour.WithWordWrap(78),
	)
	if err != nil {
		log.Fatalln()
	}

	buf, err := renderer.Render(text)
	if err != nil {
		log.Fatalln()
	}

	fmt.Println(buf)
}
