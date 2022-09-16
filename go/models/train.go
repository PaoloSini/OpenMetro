package models

import (
	"math"
	"math/rand"
)

type Train struct {
	CurrentStation Station
	CurrentLine    Line
	Direction      bool
	PosX           float64
	PosY           float64
	Speed          float64
	Stopped        int
}

func (t *Train) getNextStation() (*Station, bool) {

	offset := 1
	if !t.Direction {
		offset = -1
	}

	for index, station := range t.CurrentLine.Stations {
		if station.Name == t.CurrentStation.Name {
			if (index+1 > len(t.CurrentLine.Stations)-1) && (t.Direction) {
				return &t.CurrentLine.Stations[index-offset], true
			}
			if (index == 0) && (!t.Direction) {
				return &t.CurrentLine.Stations[index-offset], true
			}
			return &t.CurrentLine.Stations[index+offset], false
		}
	}

	return nil, false
}

func (t *Train) Update() {

	if t.Stopped > 0 {
		t.Stopped -= 1
		return
	}

	nextStation, changeDirection := t.getNextStation()
	distance := math.Sqrt(
		math.Pow(t.PosX-nextStation.PosX, 2) + math.Pow(t.PosY-nextStation.PosY, 2),
	)

	deltaX := t.Speed * math.Sin((nextStation.PosX-t.PosX)/distance)
	deltaY := t.Speed * math.Sin((nextStation.PosY-t.PosY)/distance)
	t.PosX += deltaX
	t.PosY += deltaY

	if distance < 5 {
		t.Stopped = 50 + rand.Int()%100
		t.CurrentStation = *nextStation
		if changeDirection {
			t.Direction = !t.Direction
		}
	}

}
