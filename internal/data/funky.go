package data

import (
	"github.com/gopxl/pixel"
	pxginput "github.com/timsims1717/pixel-go-input"
	"github.com/timsims1717/pixel-go-utils/timing"
	"github.com/timsims1717/pixel-go-utils/viewport"
)

type Funky struct {
	Fn func()
}

func NewFn(fn func()) *Funky {
	return &Funky{Fn: fn}
}

type HoverClick struct {
	Input     *pxginput.Input
	View      *viewport.ViewPort
	Func      func(*HoverClick)
	Hover     bool
	ViewHover bool
	Pos       pixel.Vec
}

func NewHoverClickFn(in *pxginput.Input, vp *viewport.ViewPort, fn func(*HoverClick)) *HoverClick {
	return &HoverClick{
		Input: in,
		View:  vp,
		Func:  fn,
	}
}

type TimerFunc struct {
	Timer *timing.Timer
	Func  func() bool
}

func NewTimerFunc(fn func() bool, dur float64) *TimerFunc {
	return &TimerFunc{
		Timer: timing.New(dur),
		Func:  fn,
	}
}

type FrameFunc struct {
	Func func() bool
}

func NewFrameFunc(fn func() bool) *FrameFunc {
	return &FrameFunc{Func: fn}
}
