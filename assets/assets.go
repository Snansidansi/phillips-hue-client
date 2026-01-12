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

//go:embed icons/star_border.svg
var starBorder []byte
var StarBorderIcon = themedIcon("star_border.svg", starBorder)

//go:embed icons/star_filled.svg
var starFilled []byte
var StarFilledIcon = themedIcon("star_filled.svg", starFilled)

//go:embed icons/zone.svg
var zone []byte
var ZoneIcon = themedIcon("zone.svg", zone)

//go:embed icons/lightbulb_on.svg
var lightbulbOn []byte
var LightBulbOnIcon = themedIcon("lightBulb_on.svg", lightbulbOn)

//go:embed icons/lightbulb_off.svg
var lightbulbOff []byte
var LightBulbOffIcon = themedIcon("lightBulb_off.svg", lightbulbOff)

//go:embed icons/settings.svg
var settings []byte
var SettingsIcon = themedIcon("settings.svg", settings)
