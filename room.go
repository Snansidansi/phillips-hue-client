package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type room struct {
	Name       binding.String
	On         binding.Bool
	Brightness binding.Float
}

type roomData struct {
	RoomByID  map[string]*room
	RoomIndex map[string]int
	RoomList  binding.UntypedList
}

type roomListView struct {
	Data *appData
}

type roomWidget struct {
	widget.BaseWidget

	NameLabel        *widget.Label
	BrightnessSlider *widget.Slider
	Container        *fyne.Container
}

func (r *roomWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(r.Container)
}

func NewRoomWidget(name string, brightness float64) *roomWidget {
	r := &roomWidget{
		NameLabel:        widget.NewLabel(name),
		BrightnessSlider: widget.NewSlider(0, 100),
	}

	r.BrightnessSlider.SetValue(brightness)
	r.Container = container.NewVBox(r.NameLabel, r.BrightnessSlider)

	r.ExtendBaseWidget(r)
	return r
}

func NewRoomData() *roomData {
	return &roomData{
		RoomByID:  map[string]*room{},
		RoomList:  binding.NewUntypedList(),
		RoomIndex: map[string]int{},
	}
}

func (r *roomData) AddRoom(id string, room *room) {
	r.RoomList.Append(room)
	r.RoomByID[id] = room
	r.RoomIndex[id] = r.RoomList.Length() - 1
}

func (r *roomData) DeleteRoom(id string) {
	items, _ := r.RoomList.Get()
	i := r.RoomIndex[id]
	if i < 0 || i >= len(items) {
		return
	}

	newItems := append(items[:i], items[i+1:]...)
	r.RoomList.Set(newItems)

	delete(r.RoomByID, id)
	delete(r.RoomIndex, id)
}

func CreateRoomPage(appData *appData) fyne.CanvasObject {
	view := roomListView{Data: appData}

	return widget.NewListWithData(
		appData.Rooms.RoomList,
		view.CreateItem,
		view.UpdateItem,
	)
}

func (v *roomListView) CreateItem() fyne.CanvasObject {
	return NewRoomWidget("", 0)
}

func (v *roomListView) UpdateItem(item binding.DataItem, obj fyne.CanvasObject) {
	untyped := item.(binding.Untyped)
	val, _ := untyped.Get()
	room := val.(*room)

	roomWidget := obj.(*roomWidget)

	roomWidget.NameLabel.Bind(room.Name)
	roomWidget.BrightnessSlider.Bind(room.Brightness)
}
