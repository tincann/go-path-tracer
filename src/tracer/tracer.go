package tracer

import "image/color"

type Tracer struct {
	viewplaneWidth, viewplaneHeight float32 //width and height of the view plane
	distance                        float32 //camera distance from view plane
}

func NewTracer(viewplaneWidth, viewplaneHeight, distance float32) *Tracer {
	return &Tracer{viewplaneWidth: viewplaneWidth, viewplaneHeight: viewplaneHeight, distance: distance}
}

func (t *Tracer) TraceRay(r Ray) color.Color {
	return color.RGBA{B: 128}
}
