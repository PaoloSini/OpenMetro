package models

import "image/color"

type Line struct {
	Stations []Station
	Name     string
	Trains   int
	Color    color.Color
}
