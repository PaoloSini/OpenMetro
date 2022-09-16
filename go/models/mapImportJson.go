package models

type MetroMapJson struct {
	Data struct {
		Lines []struct {
			Name     string
			Trains   int
			Stations []string
			Color    string
		}
		Stations []struct {
			Name string
			PosX float64
			PosY float64
		}
	}
}
