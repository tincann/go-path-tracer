package tracer

import "image"

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
