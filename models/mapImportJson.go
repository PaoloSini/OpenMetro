package models

type MetroMapJson struct {
	Data struct {
		Lines []struct {
			Name     string
			Trains   int
			Stations []string
			Color    struct {
				R uint8
				G uint8
				B uint8
			}
		}
		Stations []struct {
			Name string
			PosX int
			PosY int
		}
	}
}
