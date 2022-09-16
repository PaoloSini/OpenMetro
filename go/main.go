package main

import (
	"log"
	"net/http"

	"github.com/PaoloSini/OpenMetro/models"
)

const WIN_WIDTH = 1000
const WIN_HEIGHT = 1000

func main() {
	MetroMap := initMap()
	SetupRoutes(MetroMap)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initMap() *models.MetroMap {
	metroMap := new(models.MetroMap)
	loadMap("paris.json", metroMap)
	return metroMap
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
