package main

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("Hello, world!")
	a := app.New()
	defer a.Quit()
	w := a.NewWindow("Hello World")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	content := container.NewVBox(
		input, widget.NewButton("click me", func() {
			log.Println(strings.Contains(input.Text, "a"))
		}), widget.NewLabel("Hello World!"),
	)

	w.SetContent(content)
	w.ShowAndRun()

}
