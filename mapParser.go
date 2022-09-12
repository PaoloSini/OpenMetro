package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/PaoloSini/OpenMetro/models"
)

func loadMap(path string, metroMap *models.MetroMap) {
	data, err := os.ReadFile(path)
	check(err)

	result := new(models.MetroMapJson)
	json.Unmarshal([]byte(data), &result)

	metroMap.Stations = make(map[string]models.Station)
	for _, station := range result.Data.Stations {
		newStation := models.Station{
			Name: station.Name,
			PosX: station.PosX,
			PosY: station.PosY,
		}
		metroMap.Stations[station.Name] = newStation
	}

	metroMap.Lines = make(map[string]models.Line)
	for _, line := range result.Data.Lines {

		lineStations := (*new([]models.Station))
		for _, stationName := range line.Stations {
			lineStations = append(lineStations, metroMap.Stations[stationName])
		}
		newLine := models.Line{
			Name:     line.Name,
			Stations: lineStations,
		}

		metroMap.Lines[line.Name] = newLine
	}

	fmt.Println(metroMap)
}
