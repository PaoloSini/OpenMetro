package models

import "encoding/json"

type MetroMap struct {
	Stations  map[string]Station
	Lines     map[string]Line
	Trains    []*Train
	Travelers []*Traveler
}

func (mm *MetroMap) ToJSON() []byte {
	jsonString, _ := json.Marshal(mm)
	return jsonString
}

func (mm *MetroMap) Update() {
	for _, train := range mm.Trains {
		droppedTravelers := train.Update()
		mm.Travelers = append(mm.Travelers, droppedTravelers...)
		if train.Stopped > 0 {
			closeTravelers := mm.getCloseTravelers(train.CurrentStation)
			if len(closeTravelers) > 0 {
				pickedUpTravelers := train.PickupTravelers(closeTravelers)
				mm.removeTravelers(pickedUpTravelers)
			}
		}
	}

	for _, traveler := range mm.Travelers {
		traveler.Wander()
	}
}

func (mm *MetroMap) getCloseTravelers(station Station) []*Traveler {

	closeTravelers := make([]*Traveler, 0)

	for travelerIndex, traveler := range mm.Travelers {
		if GetDistance(traveler, &station) < 50 {
			closeTravelers = append(
				closeTravelers,
				mm.Travelers[travelerIndex],
			)
		}
	}

	return closeTravelers
}

func (mm *MetroMap) removeTravelers(travelers []*Traveler) {

	removedTravelerIndices := make([]int, 0)

	for travelerIndex, traveler := range mm.Travelers {
		for _, removedTraveler := range travelers {
			if traveler == removedTraveler {
				removedTravelerIndices = append(
					removedTravelerIndices,
					travelerIndex,
				)
			}
		}
	}

	if len(removedTravelerIndices) == 0 {
		return
	}

	for i := len(removedTravelerIndices) - 1; i >= 0; i-- {
		lastTraveler := mm.Travelers[len(mm.Travelers)-1]
		mm.Travelers[removedTravelerIndices[i]] = lastTraveler
		mm.Travelers = mm.Travelers[:len(mm.Travelers)-1]
	}

}
