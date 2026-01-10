package main

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
	hueapi "github.com/Snansidansi/hue-api-go"
)

type appData struct {
	hueClient *hueapi.Client

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
		a.Lights.Append(l.ID, MapToAppLight(&l))
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
	GetName() binding.String
	GetOn() binding.Bool
	GetBrightness() binding.Float
}

type baseGroup struct {
	Name       binding.String
	On         binding.Bool
	Brightness binding.Float
}

func NewBaseGroup(name string, on bool, brightness float64) baseGroup {
	group := baseGroup{
		Name:       binding.NewString(),
		On:         binding.NewBool(),
		Brightness: binding.NewFloat(),
	}

	group.Name.Set(name)
	group.On.Set(on)
	group.Brightness.Set(brightness)

	return group
}

func (b *baseGroup) GetName() binding.String {
	return b.Name
}

func (b *baseGroup) GetOn() binding.Bool {
	return b.On
}

func (b *baseGroup) GetBrightness() binding.Float {
	return b.Brightness
}

type baseGroupData[T Groupable] struct {
	ByID    map[string]T
	GuiList binding.UntypedList
}

func NewBaseGroupData[T Groupable]() *baseGroupData[T] {
	return &baseGroupData[T]{
		ByID:    map[string]T{},
		GuiList: binding.NewUntypedList(),
	}
}

func (b *baseGroupData[T]) Append(id string, data T) {
	b.GuiList.Append(data)
	b.ByID[id] = data
}

func (b *baseGroupData[T]) Remove(id string) {
	items, _ := b.GuiList.Get()
	itemPtr := b.ByID[id]

	index := -1
	for i, item := range items {
		if any(itemPtr) == item {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	newRooms := append(items[:index], items[index+1:]...)
	b.GuiList.Set(newRooms)

	delete(b.ByID, id)
}
