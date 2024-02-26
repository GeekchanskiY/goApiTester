package main

import (
	"fmt"

	//"log"
	//"strings"
	"strconv"

	// validators "ApiTester/Validators"
	utils "ApiTester/Utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	var selectedMethod string = "GET"

	a := app.New()
	defer a.Quit()
	w := a.NewWindow("ApiTester")

	// URL Input
	urlInput := widget.NewEntry()
	urlInput.SetPlaceHolder("Enter URL")

	// Request Methods Choice
	methods := []string{"GET", "POST", "PUT", "DELETE"}
	methodChoice := widget.NewSelect(methods, func(selected string) {
		// Handle selected method
	})

	// Output Labels
	outputLabel1 := widget.NewLabel("Status code")
	outputLabel2 := widget.NewLabel("Execution time")
	outputLabel3 := widget.NewLabel("Params")

	sidebar := widget.NewList(
		func() int {
			return 10 // Example: Number of previous queries
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Query")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText("Query " + strconv.Itoa(id))
		},
	)

	sidebar.OnSelected = func(id widget.ListItemID) {
		// Handle click on specific item
		// id here represents the index or identifier of the clicked item
		fmt.Println("Clicked on item:", id)

		if selectedMethod == "GET" {
			utils.ReadFile(strconv.Itoa(id) + "_query.txt")
		}
	}

	content := container.New(layout.NewVBoxLayout(),
		container.NewGridWithColumns(2,
			widget.NewLabel("URL:"),
			urlInput,
			widget.NewLabel("Method:"),
			methodChoice,
		),
		container.NewGridWithColumns(1,
			outputLabel1,
			outputLabel2,
			outputLabel3,
		),
	)

	sidebarContainer := container.NewBorder(nil, nil, nil, nil, sidebar)

	split := container.NewHSplit(content, sidebarContainer)

	w.Resize(fyne.NewSize(600, 400))

	w.SetContent(split)
	w.ShowAndRun()

}
