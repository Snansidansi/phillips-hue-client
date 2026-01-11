package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type zone struct {
	baseGroup
}

func NewZone(id, name string, on bool, brightness float64) *zone {
	return &zone{
		baseGroup: *NewBaseGroup(id, name, on, brightness),
	}
}

func CreateZonePage(appData *appData) fyne.CanvasObject {
	view := tabListEntryView{Data: appData}

	return widget.NewListWithData(
		appData.Zones.GuiList,
		view.CreateItem,
		view.UpdateItem,
	)
}
