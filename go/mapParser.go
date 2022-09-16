package main

import (
	"encoding/json"
	"os"

	"github.com/PaoloSini/OpenMetro/models"
)

func loadMap(path string, metroMap *models.MetroMap) {
	data, err := os.ReadFile(path)
	check(err)

	result := new(models.MetroMapJson)
	json.Unmarshal([]byte(data), &result)

	//Create Stations
	metroMap.Stations = make(map[string]models.Station)
	for _, station := range result.Data.Stations {
		newStation := models.Station{
			Name: station.Name,
			PosX: station.PosX,
			PosY: station.PosY,
		}
		metroMap.Stations[station.Name] = newStation
	}

	//Create Lines
	metroMap.Lines = make(map[string]models.Line)
	for _, line := range result.Data.Lines {

		lineStations := (*new([]models.Station))
		for _, stationName := range line.Stations {
			lineStations = append(lineStations, metroMap.Stations[stationName])
		}
		newLine := models.Line{
			Name:     line.Name,
			Stations: lineStations,
			Trains:   line.Trains,
			Color:    line.Color,
		}

		//Create Trains
		for trainNb := 0; trainNb < newLine.Trains; trainNb++ {
			metroMap.Trains = append(
				metroMap.Trains,
				&models.Train{
					CurrentStation: newLine.Stations[trainNb],
					CurrentLine:    newLine,
					Direction:      true,
					PosX:           newLine.Stations[trainNb].PosX,
					PosY:           newLine.Stations[trainNb].PosY,
					Speed:          0.5,
				},
			)
		}

		metroMap.Lines[line.Name] = newLine
	}

}
