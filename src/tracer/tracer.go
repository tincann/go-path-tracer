package tracer

import "image/color"
import "math"
import "math/rand"

type Tracer struct {
	ViewplaneWidth, ViewplaneHeight float64 //width and height of the view plane
	Distance                        float64 //camera distance from view plane
}

func NewTracer(viewplaneWidth, viewplaneHeight, distance float64) *Tracer {
	return &Tracer{ViewplaneWidth: viewplaneWidth, ViewplaneHeight: viewplaneHeight, Distance: distance}
}

func (t *Tracer) TraceRay(r Ray, s Scene) color.Color {
	minT := math.MaxFloat64
	var closestObject Intersectable
	for _, object := range s.Objects {
		if yes, t := object.Intersect(r); yes && t < minT {
			minT = t
			closestObject = object
		}
	}

	if closestObject != nil {
		return color.RGBA{R: 128}
	}

	return color.RGBA{
		R: uint8(rand.Int31n(64)),
		G: uint8(rand.Int31n(64)),
		B: uint8(rand.Int31n(64)),
	}
}
