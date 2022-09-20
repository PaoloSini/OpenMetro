package models

import "github.com/google/uuid"

type Line struct {
	Stations     map[uuid.UUID]*Station
	StationOrder []*Station
	Name         string
	TrainsNb     int
	Color        string
	Id           uuid.UUID
}

func (l *Line) Init(
	name string,
	lineStations map[uuid.UUID]*Station,
	lineStationsOrder []*Station,
	trainsNb int,
	color string,
) {

	lineUUID, _ := uuid.NewUUID()

	l.Name = name
	l.Stations = lineStations
	l.StationOrder = lineStationsOrder
	l.TrainsNb = trainsNb
	l.Color = color
	l.Id = lineUUID
}
