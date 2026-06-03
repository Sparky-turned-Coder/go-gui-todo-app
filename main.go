// Building a native desktop software app in Go (A To-Do list)

package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 1. Initialize the native desktop application window
	myApp := app.New()
	myWindow := myApp.NewWindow("Go To-Do App")
	myWindow.Resize(fyne.NewSize(400, 500))

	// 2. Setup our simple string slice data store
	todoItems := []string{
		"Learn Go fundamentals",
		"Build a standard REST API",
		"Build a pure Go desktop GUI app",
	}

	// 3. Create a graphical UI list component
	list := widget.NewList(
		// Count the rows in our data slice
		func() int {
			return len(todoItems)
		},
		// Create a dynamic visual Checkbox template for each row
		func() fyne.CanvasObject {
			return widget.NewCheck("Template Text", func(checked bool) {
				// We leave this template callback empty
			})
		},
		// Map our actual task data and click actions to the checkbox row
		func(i widget.ListItemID, o fyne.CanvasObject) {
			checkBox := o.(*widget.Check)

			// Set the visual text to match the current task string
			checkBox.Text = todoItems[i]

			// Set up what happens when the user clicks the checkbox
			checkBox.OnChanged = func(checked bool) {
				if checked {
					fmt.Printf("Task completed: %s\n", todoItems[i])
					// optional: You could modify your data store array state here
				} else {
					fmt.Printf("Task unchecked: %s\n", todoItems[i])
				}
			}
			// Refresh the single checkbox widget state to render the text update properly
			checkBox.Refresh()
		},
	)

	// 4. Create an input box and submit button to add items
	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Type a new task here...")

	addButton := widget.NewButton("Add Task", func() {
		// If the input box is empty, ignore the button click
		if inputEntry.Text == "" {
			return
		}

		// Append the typed text to our internal slice
		todoItems = append(todoItems, inputEntry.Text)

		// Clear out the text inside the input field for the next item
		inputEntry.SetText("")

		// Tell the UI list components to redraw itself with the new item data
		list.Refresh()
	})

	// Structure the interface layout design
	// Place the input layout items at the bottom, and the list taking up the rest of the space
	inputContainer := container.NewBorder(nil, nil, nil, addButton, inputEntry)
	mainLayout := container.NewBorder(nil, inputContainer, nil, nil, list)

	// Inject layout into the window framework and activate the server thread loop
	myWindow.SetContent(mainLayout)
	myWindow.ShowAndRun()
}
