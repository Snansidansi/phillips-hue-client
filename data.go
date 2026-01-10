package main

type appData struct {
	Rooms *roomData
}

func NewAppData() *appData {
	return &appData{
		Rooms: NewRoomData(),
	}
}
