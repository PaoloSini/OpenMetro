package models

import "math"

type Positionnable interface {
	GetPos() (float64, float64)
}

func GetDistance(p1 Positionnable, p2 Positionnable) float64 {

	p1PosX, p1PosY := p1.GetPos()
	p2PosX, p2PosY := p2.GetPos()

	return math.Sqrt(
		math.Pow(p1PosX-p2PosX, 2) + math.Pow(p1PosY-p2PosY, 2),
	)

}
