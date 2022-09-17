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
	}

	for _, traveler := range mm.Travelers {
		traveler.Wander()
	}
}

func (mm *MetroMap) getCloseTravelers(station Station) []*Traveler {

	for (Traveler)

}
