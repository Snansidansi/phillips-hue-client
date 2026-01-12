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
	"github.com/Snansidansi/phillips-hue-client/assets"
)

type tabListEntryWidget struct {
	widget.BaseWidget

	NameLabel        *widget.Label
	BrightnessSlider *widget.Slider
	Background       *canvas.Rectangle
	BulbIcon         *widget.Icon
	FavouriteIcon    *ActionIcon

	Container *fyne.Container
}

func (t *tabListEntryWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.Container)
}

func NewTabListEntry() *tabListEntryWidget {
	r := &tabListEntryWidget{
		NameLabel:        widget.NewLabel(""),
		BrightnessSlider: widget.NewSlider(0, 100),
		Background:       canvas.NewRectangle(color.Transparent),
	}

	padding := canvas.NewRectangle(color.Transparent)
	padding.SetMinSize(fyne.NewSize(5, 0))

	r.BulbIcon = widget.NewIcon(nil)
	r.FavouriteIcon = NewActionIcon(assets.StarBorderIcon, nil)

	icons := container.NewHBox(
		r.BulbIcon,
		r.FavouriteIcon,
		padding,
	)

	header := container.NewBorder(nil, nil, r.NameLabel, icons, nil)
	content := container.NewVBox(header, r.BrightnessSlider)
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
	return NewTabListEntry()
}

func (v *tabListEntryView) UpdateItem(item binding.DataItem, obj fyne.CanvasObject) {
	untyped := item.(binding.Untyped)
	val, _ := untyped.Get()
	data := val.(Groupable)

	w := obj.(*tabListEntryWidget)

	w.NameLabel.SetText(data.GetName())
	w.BrightnessSlider.SetValue(data.GetBrightness())
	w.updateBackground(data.GetBrightness(), data.GetOn(), nil)

	if data.GetOn() {
		w.BulbIcon.SetResource(assets.LightBulbOnIcon)
	} else {
		w.BulbIcon.SetResource(assets.LightBulbOffIcon)
	}

	w.FavouriteIcon.OnTapped = func() {
		if w.FavouriteIcon.Resource == assets.StarBorderIcon {
			w.FavouriteIcon.SetResource(assets.StarFilledIcon)
			return
		}
		w.FavouriteIcon.SetResource(assets.StarBorderIcon)
	}

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
