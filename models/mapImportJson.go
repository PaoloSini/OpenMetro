package models

type MetroMapJson struct {
	Data struct {
		Lines []struct {
			Name     string
			Trains   int
			Stations []string
		}
		Stations []struct {
			Name string
			PosX int
			PosY int
		}
	}
}
