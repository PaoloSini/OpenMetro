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
	travelers      []*Traveler
}

func (t *Train) GetPos() (float64, float64) {
	return t.PosX, t.PosY
}

func (t *Train) GenerateTravelers(travelersNb int) {

	for i := 0; i < travelersNb; i++ {
		t.travelers = append(
			t.travelers,
			&Traveler{
				t.PosX,
				t.PosY,
				true,
				0,
			},
		)
	}
}

func (t *Train) DropTravelers(travelersNb int) []*Traveler {

	droppedTravelers := make([]*Traveler, 0)

	for i := 0; i < travelersNb; i++ {
		droppedTraveler := t.travelers[0]
		t.travelers = t.travelers[1:]
		droppedTraveler.PosX, droppedTraveler.PosY = t.PosX+2, t.PosY+2
		droppedTraveler.Waiting = 100
		droppedTravelers = append(droppedTravelers, droppedTraveler)
	}

	return droppedTravelers
}

func (t *Train) PickupTravelers(travelers []*Traveler) []*Traveler {

	pickedUpTravelers := make([]*Traveler, 0)

	for _, traveler := range travelers {
		if traveler.Waiting == 0 {
			t.travelers = append(t.travelers, traveler)
			pickedUpTravelers = append(pickedUpTravelers, traveler)
		}
	}

	return pickedUpTravelers

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

func (t *Train) Update() []*Traveler {

	if t.Stopped > 0 {
		droppedTravelers := t.handleDropTravelers()
		t.Stopped -= 1
		return droppedTravelers
	}

	nextStation, changeDirection := t.getNextStation()
	distance := GetDistance(t, nextStation)

	deltaX := t.Speed * math.Sin((nextStation.PosX-t.PosX)/distance)
	deltaY := t.Speed * math.Sin((nextStation.PosY-t.PosY)/distance)
	t.PosX += deltaX
	t.PosY += deltaY

	if distance < 5 {
		t.handleStop(nextStation, changeDirection)
	}

	return nil

}

func (t *Train) handleDropTravelers() []*Traveler {

	droppedTravelers := make([]*Traveler, 0)

	if len(t.travelers) > 0 {
		if rand.Intn(1000) >= 900 {
			droppedTravelers = append(
				droppedTravelers,
				t.DropTravelers(1)...,
			)
		}

	}

	return droppedTravelers
}

func (t *Train) handleStop(nextStation *Station, changeDirection bool) {

	t.Stopped = 50 + rand.Intn(100)
	t.CurrentStation = *nextStation
	if changeDirection {
		t.Direction = !t.Direction
	}
}
