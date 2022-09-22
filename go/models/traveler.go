package models

import (
	"github.com/dominikbraun/graph"
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
	MetroGraph  *graph.Graph[string, Station]
}

func (t *Traveler) Init(
	PosX float64,
	PosY float64,
	DesiredDest Station,
	MetroGraph *graph.Graph[string, Station],
) {
	t.PosX = PosX
	t.PosY = PosY
	t.InTrain = true
	t.Waiting = 0
	t.Id = uuid.New()
	t.Destination = DesiredDest
	t.MetroGraph = MetroGraph

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

func (t *Traveler) wantToBoard(train *Train) bool {
	path, _ := graph.ShortestPath(*t.MetroGraph, train.CurrentStation.Name, t.Destination.Name)

	if len(path) <= 1 {
		return false
	}

	trainNextStation, _ := train.getNextStation()

	if path[1] == trainNextStation.Name {
		return true
	}
	return false
}
