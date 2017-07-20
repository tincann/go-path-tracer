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
	for y := a.Bounds.Min.Y; y < a.Bounds.Max.Y; y++ {
		for x := a.Bounds.Min.X; x < a.Bounds.Max.X; x++ {
			i := a.Bounds.Dx()*y + x
			screen.Set(x, y, toSystemColor(a.Data[i]))
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
	i := a.Bounds.Dx()*y + x

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
