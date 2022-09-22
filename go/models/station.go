package models

import (
	"github.com/google/uuid"
)

type Station struct {
	Name string
	PosX float64
	PosY float64
	Id   uuid.UUID
}

func (s *Station) String() string {
	return s.Name
}

func (s *Station) Int() int {
	return s.Id.ClockSequence()
}

func (s *Station) Init(name string, PosX float64, PosY float64) {

	stationUUID, _ := uuid.NewUUID()

	s.Name = name
	s.PosX = PosX
	s.PosY = PosY
	s.Id = stationUUID

}

func (s *Station) GetPos() (float64, float64) {
	return s.PosX, s.PosY
}
