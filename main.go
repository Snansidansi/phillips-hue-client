package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"
)

func main() {
	setupTranslations()

	a := app.New()
	w := a.NewWindow("Hue Client")

	w.SetContent(widget.NewLabel(lang.X("test", "Hue lights")))
	w.ShowAndRun()
}
