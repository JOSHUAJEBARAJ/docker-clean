package main

import (
	"fmt"
	"os"

	"github.com/JOSHUAJEBARAJ/docker-clean/dock"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	cli, err := dock.Init()
	if err != nil {
		exit(err)
	}
	images, err := dock.GetImages(cli)
	if err != nil {
		exit(err)
	}

	app := createApp(images)

	if err := app.Run(); err != nil {
		panic(err)
	}

}

func exit(e error) {
	fmt.Println(e)
	os.Exit(1)
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
			image.Delete(id)
			list.RemoveItem(index)
			// valling the delete function

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
