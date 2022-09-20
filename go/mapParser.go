package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/PaoloSini/OpenMetro/models"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/google/uuid"
)

func loadMap(path string, metroMap *models.MetroMap) {
	data, err := os.ReadFile(path)
	check(err)

	result := new(models.MetroMapJson)
	json.Unmarshal([]byte(data), &result)

	//Create Stations
	metroMap.Stations = make(map[uuid.UUID]*models.Station)
	for _, station := range result.Data.Stations {
		newStation := new(models.Station)
		newStation.Init(
			station.Name,
			station.PosX,
			station.PosY,
		)
		metroMap.Stations[newStation.Id] = newStation
	}

	//Create Lines
	metroMap.Lines = make(map[uuid.UUID]*models.Line)
	for _, line := range result.Data.Lines {

		lineStations := make(map[uuid.UUID]*models.Station)
		lineStationsOrder := make([]*models.Station, 0)
		for _, stationName := range line.Stations {
			for stationUUID, station := range metroMap.Stations {
				if stationName == station.Name {
					lineStations[stationUUID] = station
					lineStationsOrder = append(lineStationsOrder, station)
				}
			}
		}

		newLine := new(models.Line)
		newLine.Init(
			line.Name,
			lineStations,
			lineStationsOrder,
			line.Trains,
			line.Color,
		)

		//Create Trains
		for trainNb := 0; trainNb < newLine.TrainsNb; trainNb++ {

			newTrain := new(models.Train)
			newTrain.Init(*newLine, trainNb)

			metroMap.Trains = append(
				metroMap.Trains,
				newTrain,
			)
		}

		metroMap.Lines[newLine.Id] = newLine

		for _, line := range metroMap.Lines {
			for stationUUID, station := range line.Stations {
				fmt.Println(stationUUID.String())
				fmt.Println(station.Name)
			}
		}
	}

	metroMapGraph := graphFromMap(*metroMap)
	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(metroMapGraph, file)
}

func graphFromMap(mm models.MetroMap) graph.Graph[string, models.Station] {
	metroGraphHash := func(s models.Station) string {
		return s.Name
	}
	metroGraph := graph.New(metroGraphHash)

	for _, station := range mm.Stations {
		metroGraph.AddVertex(*station)
	}

	for _, line := range mm.Lines {
		for i:=0; i < len(line.StationOrder) -1; i++ {
			metroGraph.AddEdge(
				line.StationOrder[i].Name,
				line.StationOrder[i+1].Name,
				graph.EdgeWeight(int(models.GetDistance(line.StationOrder[i], line.StationOrder[i+1]))),
			)
		}
	}
	return metroGraph
}
