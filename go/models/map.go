package models

import "encoding/json"

type MetroMap struct {
	Stations map[string]Station
	Lines    map[string]Line
}

func (m *MetroMap) ToJSON() []byte {
	jsonString, _ := json.Marshal(m)
	return jsonString
}
