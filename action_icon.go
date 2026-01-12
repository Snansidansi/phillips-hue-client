package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ActionIcon struct {
	widget.Icon
	OnTapped func()
}

func NewActionIcon(res fyne.Resource, tapped func()) *ActionIcon {
	i := &ActionIcon{OnTapped: tapped}
	i.SetResource(res)
	i.ExtendBaseWidget(i)
	return i
}

func (i *ActionIcon) Tapped(_ *fyne.PointEvent) {
	if i.OnTapped != nil {
		i.OnTapped()
	}
}
