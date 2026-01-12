package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	hueapi "github.com/Snansidansi/hue-api-go"
)

const sliderChangeIntervallLimit = 400 * time.Millisecond

type appData struct {
	hueClient        *hueapi.Client
	lastSliderUpdate time.Time

	Rooms  *baseGroupData[*room]
	Zones  *baseGroupData[*zone]
	Lights *baseGroupData[*light]
}

func NewAppData() *appData {
	return &appData{
		Rooms:  NewBaseGroupData[*room](),
		Zones:  NewBaseGroupData[*zone](),
		Lights: NewBaseGroupData[*light](),
	}
}

func (a *appData) LoadInitialData() error {
	a.hueClient = getHueClient()

	err := a.LoadRooms()
	if err != nil {
		return err
	}

	err = a.LoadZones()
	if err != nil {
		return err
	}

	err = a.LoadLights()
	if err != nil {
		return err
	}

	return nil
}

func (a *appData) LoadLights() error {
	hueResponse, err := a.hueClient.Lights.GetAllLights()
	if err != nil {
		return err
	}
	if len(hueResponse.Errors) > 0 {
		return fmt.Errorf("Hue error: %+v", hueResponse.Errors)
	}
	if len(hueResponse.Data) == 0 {
		return nil
	}

	for _, l := range hueResponse.Data {
		appLight := MapToAppLight(&l)
		a.Lights.Append(l.ID, appLight)
	}
	return nil
}

func (a *appData) LoadZones() error {
	hueResponse, err := a.hueClient.Zones.GetZones()
	if err != nil {
		return err
	}
	if len(hueResponse.Errors) > 0 {
		return fmt.Errorf("Hue error: %+v", hueResponse.Errors)
	}
	if len(hueResponse.Data) == 0 {
		return nil
	}

	for _, z := range hueResponse.Data {
		a.Zones.Append(z.ID, MapToAppZone(&z))
	}
	return nil
}

func (a *appData) LoadRooms() error {
	hueResponse, err := a.hueClient.Rooms.GetAllRooms()
	if err != nil {
		return err
	}
	if len(hueResponse.Errors) > 0 {
		return fmt.Errorf("Hue error: %+v", hueResponse.Errors)
	}
	if len(hueResponse.Data) == 0 {
		return nil
	}

	for _, r := range hueResponse.Data {
		a.Rooms.Append(r.ID, MapToAppRoom(&r))
	}
	return nil
}

type Groupable interface {
	GetID() string
	GetName() string
	GetOn() bool
	SetOn(state bool)
	GetBrightness() float64
	SetBrightness(val float64)
}

type baseGroup struct {
	ID         string
	Name       string
	On         bool
	Brightness float64
}

func NewBaseGroup(id, name string, on bool, brightness float64) *baseGroup {
	group := baseGroup{
		ID:         id,
		Name:       name,
		On:         on,
		Brightness: brightness,
	}

	return &group
}

func (b *baseGroup) GetID() string {
	return b.ID
}

func (b *baseGroup) GetName() string {
	return b.Name
}

func (b *baseGroup) GetOn() bool {
	return b.On
}

func (b *baseGroup) SetOn(state bool) {
	b.On = state
}

func (b *baseGroup) GetBrightness() float64 {
	return b.Brightness
}

func (b *baseGroup) SetBrightness(val float64) {
	b.Brightness = val
}

type baseGroupData[T Groupable] struct {
	ByID      map[string]T
	GuiList   binding.UntypedList
	GuiListId map[string]int
	List      *widget.List
}

func NewBaseGroupData[T Groupable]() *baseGroupData[T] {
	return &baseGroupData[T]{
		ByID:      map[string]T{},
		GuiList:   binding.NewUntypedList(),
		GuiListId: map[string]int{},
		List:      nil,
	}
}

func (b *baseGroupData[T]) Append(id string, data T) {
	b.GuiList.Append(data)
	b.ByID[id] = data
	b.GuiListId[id] = b.GuiList.Length()
}

func (b *baseGroupData[T]) Remove(id string) {
	items, _ := b.GuiList.Get()

	i, ok := b.GuiListId[id]
	if !ok {
		return
	}

	newData := append(items[:i], items[i+1:]...)
	b.GuiList.Set(newData)

	delete(b.ByID, id)

	newGuiListID := make(map[string]int, len(b.GuiListId)-1)
	for i := range b.GuiList.Length() - 1 {
		entry, _ := b.GuiList.GetValue(i)
		newGuiListID[entry.(T).GetID()] = i
	}
}
