package models

type MetroMap struct {
	Stations map[string]Station
	Lines    map[string]Line
}
