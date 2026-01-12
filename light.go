package main

import (
	"fyne.io/fyne/v2/widget"
)

type light struct {
	baseGroup
}

func NewLight(id, name string, on bool, brightness float64) *light {
	return &light{
		baseGroup: *NewBaseGroup(id, name, on, brightness),
	}
}

func CreateLightPage(appData *appData) *widget.List {
	view := tabListEntryView{Data: appData}

	view.SliderOnChanged = func(id string, val float64) {
		appData.hueClient.Lights.SetBrightness(id, val)
	}

	list := widget.NewListWithData(
		appData.Lights.GuiList,
		view.CreateItem,
		view.UpdateItem,
	)

	list.OnSelected = func(id widget.ListItemID) {
		list.Unselect(id)

		val, _ := view.Data.Lights.GuiList.GetValue(id)
		light, _ := val.(*light)
		newState := !light.On

		view.Data.hueClient.Lights.SetOnOff(light.ID, newState)
		light.On = newState
	}

	return list
}
