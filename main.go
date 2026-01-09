package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"
)

func main() {
	setupTranslations()

	a := app.New()
	w := a.NewWindow("Hue Client")

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(lang.X("menu.favorites", "Favorites"), StarIcon, widget.NewLabel("...")),
		container.NewTabItemWithIcon(lang.X("menu.rooms", "Rooms"), DoorIcon, widget.NewLabel("...")),
		container.NewTabItemWithIcon(lang.X("menu.zones", "Zones"), ZoneIcon, widget.NewLabel("...")),
		container.NewTabItemWithIcon(lang.X("menu.lights", "Lights"), LightBulbIcon, widget.NewLabel("...")),
		container.NewTabItemWithIcon("", SettingsIcon, widget.NewLabel("...")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun()
}
