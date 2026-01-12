package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	hueapi "github.com/Snansidansi/hue-api-go"
	"github.com/Snansidansi/hue-api-go/models"
	"github.com/joho/godotenv"
)

func getHueClient() *hueapi.Client {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error while loading .env file: %v", err)
		os.Exit(1)
	}

	bridge := models.Bridge{
		Id:       os.Getenv("HUE_BRIDGE_ID"),
		IPAdress: os.Getenv("HUE_BRIDGE_IP"),
	}

	apiKey := os.Getenv("HUE_BRIDGE_USERNAME")
	return hueapi.NewClient(bridge, apiKey, nil, true)
}

func startListenToEventstream(appData *appData) {
	go func() {
		eventStream := appData.hueClient.EventStream.GetEventStream(20)
		errorStream := appData.hueClient.EventStream.GetErrorStream(5)
		appData.hueClient.EventStream.Start()
		defer appData.hueClient.EventStream.Stop()

		for {
			select {
			case event := <-eventStream:
				handleStreamEvent(appData, event)
			case event := <-errorStream:
				handleStreamError(event)
			}
		}
	}()
}

func checkResponseInvalid[T any](hueResponse *models.HueResponse[T], err error) bool {
	if err != nil {
		fmt.Printf("Error while reaction on eventstream: %v", err)
		return true
	}
	if len(hueResponse.Errors) > 0 {
		fmt.Printf("Hue error: %+v", hueResponse.Errors)
		return true
	}
	if len(hueResponse.Data) == 0 {
		return true
	}
	return false
}

func handleStreamEvent(appData *appData, event any) {
	switch e := event.(type) {
	case *models.LightChangeEvent:
		switch e.EventType {
		case models.EventTypeAdd:
			hueResponse, err := appData.hueClient.Lights.GetLightByID(e.ID)
			if checkResponseInvalid(hueResponse, err) {
				return
			}
			appData.Lights.Append(e.ID, MapToAppLight(&hueResponse.Data[0]))

		case models.EventTypeDelete:
			appData.Lights.Remove(e.ID)

		case models.EventTypeUpdate:
			light := appData.Lights.ByID[e.ID]
			defer fyne.DoAndWait(func() {
				listEntryID := appData.Lights.GuiListId[e.ID]
				appData.Lights.List.RefreshItem(widget.ListItemID(listEntryID))
			})

			if !e.StateChanges {
				hueResponse, err := appData.hueClient.Lights.GetLightByID(e.ID)
				if checkResponseInvalid(hueResponse, err) {
					return
				}
				light.Name = *hueResponse.Data[0].Metadata.Name
				return
			}
			if e.On != nil {
				light.On = *e.On.On
			}
			if e.Dimming != nil {
				light.Brightness = *e.Dimming.Brightness
			}
		}
	}
}

func handleStreamError(err error) {
	fmt.Println("Hue eventstream error: ", err)
}

func MapToAppRoom(hueRoom *models.Room) *room {
	return NewRoom(hueRoom.ID, hueRoom.Metadata.Name, false, 0)
}

func MapToAppZone(hueZone *models.Zone) *zone {
	return NewZone(hueZone.ID, hueZone.Metadata.Name, false, 0)
}

func MapToAppLight(hueLight *models.Light) *light {
	return NewLight(hueLight.ID, *hueLight.Metadata.Name, *hueLight.On.On, *hueLight.Dimming.Brightness)
}
