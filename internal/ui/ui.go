package ui

import (
	"github.com/bytearena/ecs"
	"github.com/gopxl/pixel"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
	"github.com/timsims1717/pixel-go-utils/typeface"
	"github.com/timsims1717/pixel-go-utils/viewport"
)

type ElementType int

const (
	ButtonElement = iota
	CheckboxElement
	ContainerElement
	InputElement
	ScrollElement
	SpriteElement
	TextElement
	CustomElement
)

type ElementConstructor struct {
	Key         string
	Width       float64
	Height      float64
	SprKey      string
	SprKey2     string
	Batch       string
	Text        string
	HelpText    string
	Color       pixel.RGBA
	Position    pixel.Vec
	CanFocus    bool
	Left        string
	Right       string
	Up          string
	Down        string
	ElementType ElementType
	SubElements []ElementConstructor
}

type Element struct {
	Key      string
	Sprite   *img.Sprite
	Sprite2  *img.Sprite
	Delay    float64
	HelpText string
	Object   *object.Object
	Entity   *ecs.Entity
	Action   func()
	OnClick  func()
	OnHold   func()
	OnHover  func(bool)
	Left     string
	Right    string
	Up       string
	Down     string

	ElementType ElementType

	Checked    bool
	Value      string
	Focused    bool
	Text       *typeface.Text
	CaretIndex int
	CaretObj   *object.Object
	InputType  InputType
	MultiLine  bool

	Border       *Border
	BorderVP     *viewport.ViewPort
	BorderObject *object.Object
	BorderEntity *ecs.Entity
	ViewPort     *viewport.ViewPort
	Layer        int
	Elements     []*Element

	Bar          *Element
	ButtonHeight float64
	YTop         float64
	YBot         float64
}

type InputType int

const (
	AlphaNumeric = iota
	Numeric
	Special
	Any
)

func (e *Element) Get(key string) *Element {
	for _, e1 := range e.Elements {
		if e1.Key == key {
			return e1
		}
	}
	return nil
}
