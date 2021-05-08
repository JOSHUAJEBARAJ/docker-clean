package main

import (
	"github.com/JOSHUAJEBARAJ/docker-clean/dock"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	images := dock.GetImages()

	app := createApp(images)

	if err := app.Run(); err != nil {
		panic(err)
	}

}

func createApp(images []dock.Image) (app *tview.Application) {

	app = tview.NewApplication()
	// creating command
	list := tview.NewList()
	list.SetBorder(true).SetTitle("Docker Images")
	list.ShowSecondaryText(false)
	for i, image := range images {

		name := image.Name
		list.AddItem(name, "", rune(i), nil)

	}

	// variable to delete
	var image dock.Image

	var id string

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {
		case tcell.KeyEnter:

			index := list.GetCurrentItem()
			name, _ := list.GetItemText(index)
			// getting the id
			for _, v := range images {
				if v.Name == name {
					id = v.Id
				}
			}
			list.RemoveItem(index)
			// valling the delete function

			image.Delete(id)

			return nil
		case tcell.KeyEsc:
			// Exit the application
			app.Stop()
			return nil
		}

		return event
	})

	pages := tview.NewPages()

	layout := createLayout(list)
	pages.AddPage("root", layout, true, true)

	app.SetRoot(pages, true)
	return app
}

func createLayout(commandList tview.Primitive) (layout *tview.Flex) {

	mainLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(commandList, 150, 1, true)

	info := tview.NewTextView()
	info.SetBorder(true)
	info.SetText("Created by Joshua")
	info.SetTextAlign(tview.AlignCenter)

	layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainLayout, 0, 20, true).
		AddItem(info, 3, 1, false)

	return layout
}
