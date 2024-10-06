package data

import "github.com/gopxl/pixel"

var (
	RoomBorder = pixel.R(-300., -240., 300., 240.)
	Layers     = true
)

func GetRandomX() float64 {
	return GlobalRand.Float64()*300. - 150.
}

func GetRandomY() float64 {
	return GlobalRand.Float64()*200. - 100.
}
