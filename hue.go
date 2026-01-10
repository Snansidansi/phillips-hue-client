package main

import (
	"fmt"
	"os"

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
	return hueapi.NewClient(bridge, apiKey, nil, false)
}

func MapToAppRoom(hueRoom *models.Room) *room {
	return NewRoom(hueRoom.Metadata.Name, false, 0)
}

func MapToAppZone(hueZone *models.Zone) *zone {
	return NewZone(hueZone.Metadata.Name, false, 0)
}

func MapToAppLight(hueLight *models.Light) *light {
	return NewLight(*hueLight.Metadata.Name, *hueLight.On.On, *hueLight.Dimming.Brightness)
}
