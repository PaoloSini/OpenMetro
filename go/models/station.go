package models

type Station struct {
	Name string
	PosX float64
	PosY float64
}

func (s *Station) GetPos() (float64, float64) {
	return s.PosX, s.PosY
}
