package models

import "encoding/json"

type MetroMap struct {
	Stations map[string]Station
	Lines    map[string]Line
	Trains   []*Train
}

func (mm *MetroMap) ToJSON() []byte {
	jsonString, _ := json.Marshal(mm)
	return jsonString
}

func (mm *MetroMap) Update() {
	for _, train := range mm.Trains {
		train.Update()
	}
}
