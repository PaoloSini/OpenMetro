package primitives

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawCircle(r float64, c color.Color) *ebiten.Image {

	dc := gg.NewContext(int(r*2), int(r*2))
	dc.DrawCircle(r, r, r)
	dc.SetColor(c)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}
