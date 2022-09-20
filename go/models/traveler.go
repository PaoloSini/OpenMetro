package models

import (
	"math/rand"

	"github.com/google/uuid"
)

type Traveler struct {
	PosX        float64
	PosY        float64
	InTrain     bool
	Waiting     float64
	Id          uuid.UUID
	Destination Station
}

func (t *Traveler) Init(PosX float64, PosY float64) {
	t.PosX = PosX
	t.PosY = PosY
	t.InTrain = true
	t.Waiting = 0
	t.Id = uuid.New()

}

func (t *Traveler) GetPos() (float64, float64) {
	return t.PosX, t.PosY
}

func (t *Traveler) Wander() {

	t.PosX += rand.Float64() - 0.5
	t.PosY += rand.Float64() - 0.5
	if t.Waiting > 0 {
		t.Waiting -= 1
	}
}
