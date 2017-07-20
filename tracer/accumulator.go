package tracer

import (
	"image"
	"image/color"

	wde "github.com/skelterjohn/go.wde"
)

type Accumulator struct {
	Data   []Color
	Bounds image.Rectangle
	frames int
}

func NewAccumulator(bounds image.Rectangle) *Accumulator {
	data := make([]Color, bounds.Dx()*bounds.Dy())
	return &Accumulator{
		Data:   data,
		Bounds: bounds,
	}
}

func (a *Accumulator) DrawContents(screen wde.Image) {
	for y := 0; y < a.Bounds.Dy(); y++ {
		for x := 0; x < a.Bounds.Dx(); x++ {
			i := a.Bounds.Dx()*y + x
			screen.Set(a.Bounds.Min.X+x, a.Bounds.Min.Y+y, toSystemColor(a.Data[i]))
		}
	}
}

func toSystemColor(c Color) color.RGBA {
	return color.RGBA{
		R: uint8(Clamp(c.R, 0, 1) * 255),
		G: uint8(Clamp(c.G, 0, 1) * 255),
		B: uint8(Clamp(c.B, 0, 1) * 255),
	}
}

func (a *Accumulator) SetPixel(x, y int, color Color) Color {
	//normalize x and y
	xx := x - a.Bounds.Min.X
	yy := y - a.Bounds.Min.Y

	i := a.Bounds.Dx()*yy + xx

	c := a.Data[i]
	c.R = (c.R*float32(a.frames) + color.R) / float32(a.frames+1)
	c.G = (c.G*float32(a.frames) + color.G) / float32(a.frames+1)
	c.B = (c.B*float32(a.frames) + color.B) / float32(a.frames+1)

	a.Data[i] = c
	return c
}

func (a *Accumulator) NextFrame() {
	a.frames++
}

func (a *Accumulator) Reset() {
	a.frames = 0
}
