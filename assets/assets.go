package assets

import (
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func themedIcon(name string, data []byte) fyne.Resource {
	return theme.NewThemedResource(
		fyne.NewStaticResource(name, data),
	)
}

//go:embed icons/door.svg
var door []byte
var DoorIcon = themedIcon("door.svg", door)

//go:embed icons/star.svg
var star []byte
var StarIcon = themedIcon("star.svg", star)

//go:embed icons/zone.svg
var zone []byte
var ZoneIcon = themedIcon("zone.svg", zone)

//go:embed icons/lightbulb.svg
var lightbulb []byte
var LightBulbIcon = themedIcon("lightBulb.svg", lightbulb)

//go:embed icons/settings.svg
var settings []byte
var SettingsIcon = themedIcon("settings.svg", settings)
