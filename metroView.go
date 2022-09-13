package main

import (
	"image/color"

	"github.com/PaoloSini/OpenMetro/models"
	"github.com/PaoloSini/OpenMetro/primitives"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const STATION_RADIUS = 20

type MetroView struct {
	MetroMap     models.MetroMap
	StationImage *ebiten.Image
}

func (mv *MetroView) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIN_WIDTH, WIN_HEIGHT
}

func (mv *MetroView) Update() error {
	return nil
}

func (mv *MetroView) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	mv.drawMap(screen)
}

func (mv *MetroView) drawMap(screen *ebiten.Image) {

	for _, station := range mv.MetroMap.Stations {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(station.PosX), float64(station.PosY))
		screen.DrawImage(primitives.DrawCircle(STATION_RADIUS, color.Black), op)
	}

	for _, line := range mv.MetroMap.Lines {
		for i := 0; i == len(line.Stations)-2; i++ {
			ebitenutil.DrawLine(
				screen,
				float64(line.Stations[i].PosX+STATION_RADIUS), float64(line.Stations[i].PosY+STATION_RADIUS),
				float64(line.Stations[i+1].PosX+STATION_RADIUS), float64(line.Stations[i+1].PosY+STATION_RADIUS),
				line.Color,
			)
		}
	}
}
