package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type room struct {
	baseGroup
}

func NewRoom(id, name string, on bool, brightness float64) *room {
	return &room{
		baseGroup: *NewBaseGroup(id, name, on, brightness),
	}
}

func CreateRoomPage(appData *appData) fyne.CanvasObject {
	view := tabListEntryView{Data: appData}

	return widget.NewListWithData(
		appData.Rooms.GuiList,
		view.CreateItem,
		view.UpdateItem,
	)
}
