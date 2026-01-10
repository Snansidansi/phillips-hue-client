package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type light struct {
	baseGroup
}

func NewLight(name string, on bool, brightness float64) *light {
	return &light{
		baseGroup: NewBaseGroup(name, on, brightness),
	}
}

func CreateLightPage(appData *appData) fyne.CanvasObject {
	view := tabListEntryView{Data: appData}

	return widget.NewListWithData(
		appData.Lights.GuiList,
		view.CreateItem,
		view.UpdateItem,
	)
}
