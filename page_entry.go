package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type tabListEntryWidget struct {
	widget.BaseWidget

	NameLabel        *widget.Label
	BrightnessSlider *widget.Slider
	Container        *fyne.Container
}

func (t *tabListEntryWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.Container)
}

func NewTabListEntry(name string, brightness float64) *tabListEntryWidget {
	r := &tabListEntryWidget{
		NameLabel:        widget.NewLabel(name),
		BrightnessSlider: widget.NewSlider(0, 100),
	}

	r.BrightnessSlider.SetValue(brightness)
	r.Container = container.NewVBox(r.NameLabel, r.BrightnessSlider)

	r.ExtendBaseWidget(r)
	return r
}

type tabListEntryView struct {
	Data *appData
}

func (v *tabListEntryView) CreateItem() fyne.CanvasObject {
	return NewTabListEntry("", 0)
}

func (v *tabListEntryView) UpdateItem(item binding.DataItem, obj fyne.CanvasObject) {
	untyped := item.(binding.Untyped)
	val, _ := untyped.Get()
	data := val.(Groupable)

	w := obj.(*tabListEntryWidget)

	w.NameLabel.Bind(data.GetName())
	w.BrightnessSlider.Bind(data.GetBrightness())
}
