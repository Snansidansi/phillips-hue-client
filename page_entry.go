package main

import (
	"time"

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
	Data            *appData
	SliderOnChanged func(id string, val float64)
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
	w.BrightnessSlider.OnChanged = func(val float64) {
		if data.GetValueIsUpdating().Load() {
			return
		}

		if time.Since(v.Data.lastSliderUpdate) < sliderChangeIntervallLimit {
			return
		}

		data.GetBrightness().Set(val)
		v.Data.lastSliderUpdate = time.Now()
		v.SliderOnChanged(data.GetID(), val)
	}
	w.BrightnessSlider.OnChangeEnded = func(val float64) {
		if data.GetValueIsUpdating().Load() {
			return
		}

		currentBrightness, _ := data.GetBrightness().Get()
		if currentBrightness == val {
			return
		}

		data.GetBrightness().Set(val)
		v.SliderOnChanged(data.GetID(), val)
	}
}
