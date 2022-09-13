package main

import (
	"log"

	"github.com/PaoloSini/OpenMetro/models"
	"github.com/hajimehoshi/ebiten/v2"
)

const WIN_WIDTH = 1000
const WIN_HEIGHT = 1000

func main() {

	metroView := MetroView{
		MetroMap: initMetro(),
	}

	if err := ebiten.RunGame(&metroView); err != nil {
		log.Fatal(err)
	}

}

func initMetro() models.MetroMap {
	return initMap()
}

func initMap() models.MetroMap {
	metroMap := new(models.MetroMap)
	loadMap("paris.json", metroMap)
	return *metroMap
}

func initWindow() {
	ebiten.SetWindowSize(WIN_WIDTH, WIN_HEIGHT)
	ebiten.SetWindowTitle("OpenMetro")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
