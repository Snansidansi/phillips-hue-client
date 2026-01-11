package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"
	"github.com/Snansidansi/phillips-hue-client/assets"
)

func main() {
	setupTranslations()

	a := app.New()
	w := a.NewWindow("Hue Client")

	appData := NewAppData()
	err := appData.LoadInitialData()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	startListenToEventstream(appData)

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(lang.X("menu.favorites", "Favorites"), assets.StarIcon, widget.NewLabel("...")),
		container.NewTabItemWithIcon(lang.X("menu.rooms", "Rooms"), assets.DoorIcon, CreateRoomPage(appData)),
		container.NewTabItemWithIcon(lang.X("menu.zones", "Zones"), assets.ZoneIcon, CreateZonePage(appData)),
		container.NewTabItemWithIcon(lang.X("menu.lights", "Lights"), assets.LightBulbIcon, CreateLightPage(appData)),
		container.NewTabItemWithIcon("", assets.SettingsIcon, widget.NewLabel("...")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun()
}
