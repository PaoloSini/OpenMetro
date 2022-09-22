package models

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/dominikbraun/graph"
	"github.com/google/uuid"
)

type MetroMap struct {
	Stations     map[uuid.UUID]*Station
	Lines        map[uuid.UUID]*Line
	Trains       []*Train
	Travelers    map[uuid.UUID]*Traveler
	Graph        graph.Graph[string, Station]
	travelerLock *sync.RWMutex
}

func (mm *MetroMap) Init() {
	mm.travelerLock = &sync.RWMutex{}
	mm.Travelers = make(map[uuid.UUID]*Traveler)
}

func (mm *MetroMap) ToJSON() []byte {
	jsonString, err := json.Marshal(mm)

	if err != nil {
		fmt.Println(err)
	}
	return jsonString
}

func (mm *MetroMap) DispatchTravelers() {
	travelersPerTrain := len(mm.Travelers) / len(mm.Trains)
	for _, train := range mm.Trains {
		travelersToPickUp := make([]*Traveler, 0)
		for travelerUUID, traveler := range mm.Travelers {
			if len(travelersToPickUp) == travelersPerTrain {
				break
			}
			delete(mm.Travelers, travelerUUID)
			travelersToPickUp = append(travelersToPickUp, traveler)
			travelersPerTrain -= 1
		}
		train.PickupTravelers(travelersToPickUp)
	}
}

func (mm *MetroMap) Update() {
	for _, train := range mm.Trains {
		droppedTravelers := train.Update()
		for _, droppedTraveler := range droppedTravelers {
			mm.travelerLock.Lock()
			mm.Travelers[droppedTraveler.Id] = droppedTraveler
			mm.travelerLock.Unlock()
		}
		if train.Stopped > 0 {
			closeTravelers := mm.getCloseTravelers(train.CurrentStation)
			if len(closeTravelers) > 0 {
				pickedUpTravelers := train.PickupTravelers(closeTravelers)
				mm.removeTravelers(pickedUpTravelers)
			}
		}
	}
	mm.travelerLock.RLock()
	for _, traveler := range mm.Travelers {
		traveler.Wander()
	}
	mm.travelerLock.RUnlock()
}

func (mm *MetroMap) getCloseTravelers(station Station) []*Traveler {

	closeTravelers := make([]*Traveler, 0)

	mm.travelerLock.RLock()
	for _, traveler := range mm.Travelers {
		if GetDistance(traveler, &station) < 50 {
			closeTravelers = append(
				closeTravelers,
				traveler,
			)
		}
	}
	mm.travelerLock.RUnlock()

	return closeTravelers
}

func (mm *MetroMap) removeTravelers(travelers []*Traveler) {

	for _, removedTraveler := range travelers {
		mm.travelerLock.Lock()
		delete(mm.Travelers, removedTraveler.Id)
		mm.travelerLock.Unlock()
	}

}
