package models

import (
	"math"
	"math/rand"
	"sync"

	"github.com/google/uuid"
)

type Train struct {
	CurrentStation Station
	CurrentLine    Line
	Direction      bool
	PosX           float64
	PosY           float64
	Speed          float64
	Stopped        int
	Travelers      map[uuid.UUID]*Traveler
	travelerLock   *sync.RWMutex
}

func (t *Train) Init(line Line, trainNb int) {

	t.CurrentStation = *line.StationOrder[trainNb]
	t.CurrentLine = line
	t.PosX = line.StationOrder[trainNb].PosX - 8
	t.PosY = line.StationOrder[trainNb].PosY - 8
	t.Direction = true
	t.Speed = 0.5
	t.Travelers = make(map[uuid.UUID]*Traveler, 0)
	t.travelerLock = &sync.RWMutex{}

}

func (t *Train) GetPos() (float64, float64) {
	return t.PosX, t.PosY
}

func (t *Train) DropTravelers(travelersNb int) []*Traveler {

	droppedTravelers := make([]*Traveler, 0)

	t.travelerLock.RLock()
	for k, v := range t.Travelers {
		if travelersNb == 0 {
			break
		}
		droppedTraveler := v
		delete(t.Travelers, k)
		droppedTraveler.PosX, droppedTraveler.PosY = t.PosX+2, t.PosY+2
		droppedTraveler.Waiting = 100
		droppedTravelers = append(droppedTravelers, droppedTraveler)
		travelersNb -= 1
	}
	t.travelerLock.RUnlock()
	return droppedTravelers
}

func (t *Train) PickupTravelers(travelers []*Traveler) []*Traveler {

	pickedUpTravelers := make([]*Traveler, 0)

	for _, traveler := range travelers {
		if traveler.Waiting == 0 {
			t.travelerLock.Lock()
			t.Travelers[traveler.Id] = traveler
			traveler.InTrain = true
			t.travelerLock.Unlock()
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

	for index, station := range t.CurrentLine.StationOrder {
		if station.Name == t.CurrentStation.Name {
			if (index+1 > len(t.CurrentLine.Stations)-1) && (t.Direction) {
				return t.CurrentLine.StationOrder[index-offset], true
			}
			if (index == 0) && (!t.Direction) {
				return t.CurrentLine.StationOrder[index-offset], true
			}
			return t.CurrentLine.StationOrder[index+offset], false
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

	t.travelerLock.RLock()
	if len(t.Travelers) > 0 {
		if rand.Intn(1000) >= 900 {
			droppedTravelers = append(
				droppedTravelers,
				t.DropTravelers(1)...,
			)
		}
	}
	t.travelerLock.RUnlock()

	return droppedTravelers
}

func (t *Train) handleStop(nextStation *Station, changeDirection bool) {

	t.Stopped = 50 + rand.Intn(100)
	t.CurrentStation = *nextStation
	if changeDirection {
		t.Direction = !t.Direction
	}
}
