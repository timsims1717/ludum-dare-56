package ui

import (
	"github.com/gopxl/pixel"
)

type Border struct {
	Width  int
	Height int
	Rect   pixel.Rect
	Empty  bool
	Style  BorderStyle
}

type BorderStyle int

const (
	FancyBorder = iota
	ThinBorder
)
