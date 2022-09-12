package main

import (
	"github.com/PaoloSini/OpenMetro/models"
)

func main() {

	metroMap := new(models.MetroMap)
	loadMap("paris.json", metroMap)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
