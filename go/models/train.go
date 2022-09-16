package models

import (
	"math"
)

type Train struct {
	CurrentStation Station
	CurrentLine    Line
	Direction      bool
	PosX           float64
	PosY           float64
	Speed          float64
}

func (t *Train) getNextStation() *Station {
	for index, station := range t.CurrentLine.Stations {
		if (index == len(t.CurrentLine.Stations)-1) && (t.Direction) {
			return &t.CurrentLine.Stations[len(t.CurrentLine.Stations)-1]
		}
		if (index == 1) && (!t.Direction) {
			return &t.CurrentLine.Stations[0]
		}
		if station == t.CurrentStation {
			return &t.CurrentLine.Stations[index+1]
		}
	}

	return nil
}

func (t *Train) Update() {
	nextStation := t.getNextStation()
	distance := math.Sqrt(
		math.Pow(t.PosX-nextStation.PosX, 2) + math.Pow(t.PosY-nextStation.PosY, 2),
	)

	deltaX := t.Speed * math.Sin((nextStation.PosX-t.PosX)/distance)
	deltaY := t.Speed * math.Sin((nextStation.PosY-t.PosY)/distance)

	t.PosX += deltaX
	t.PosY += deltaY

}
