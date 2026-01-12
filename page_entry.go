package main

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type tabListEntryWidget struct {
	widget.BaseWidget

	NameLabel        *widget.Label
	BrightnessSlider *widget.Slider
	Background       *canvas.Rectangle
	Container        *fyne.Container
}

func (t *tabListEntryWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.Container)
}

func NewTabListEntry(name string, brightness float64) *tabListEntryWidget {
	r := &tabListEntryWidget{
		NameLabel:        widget.NewLabel(name),
		BrightnessSlider: widget.NewSlider(0, 100),
		Background:       canvas.NewRectangle(color.Transparent),
	}

	r.BrightnessSlider.SetValue(brightness)
	content := container.NewVBox(r.NameLabel, r.BrightnessSlider)
	r.Container = container.NewStack(r.Background, content)

	r.ExtendBaseWidget(r)
	return r
}

func (t *tabListEntryWidget) updateBackground(brightness float64, isOn bool, customColor *color.NRGBA) {
	if !isOn || brightness <= 0 {
		t.Background.FillColor = color.Transparent
		return
	}

	if customColor == nil {
		customColor = &color.NRGBA{R: 255, G: 200, B: 0, A: 0}
	}

	alpha := uint8(brightness)
	customColor.A = alpha

	t.Background.FillColor = customColor
	t.Background.Refresh()
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

	w.NameLabel.SetText(data.GetName())
	w.BrightnessSlider.SetValue(data.GetBrightness())
	w.updateBackground(data.GetBrightness(), data.GetOn(), nil)

	w.BrightnessSlider.OnChanged = func(val float64) {
		w.updateBackground(val, data.GetOn(), nil)

		if time.Since(v.Data.lastSliderUpdate) < sliderChangeIntervallLimit {
			return
		}

		data.SetBrightness(val)
		v.Data.lastSliderUpdate = time.Now()
		v.SliderOnChanged(data.GetID(), val)
	}

	w.BrightnessSlider.OnChangeEnded = func(val float64) {
		currentBrightness := data.GetBrightness()
		if math.Abs(val-currentBrightness) < 1.0 {
			return
		}

		data.SetBrightness(val)
		v.SliderOnChanged(data.GetID(), val)
	}
}
