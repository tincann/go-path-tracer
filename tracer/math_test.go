package tracer_test

import (
	"testing"

	"fmt"

	. "github.com/franela/goblin"
	"github.com/tincann/go-path-tracer/tracer"
)

const epsilon = 1e-8

func TestVector(t *testing.T) {
	g := Goblin(t)
	g.Describe("Length", func() {
		g.It("Should have length 1", func() {
			v := tracer.NewVector(0, 1, 0)
			assertFloat(g, v.Length(), 1)

			tracer.NewVector(0, 1, 0)
		})
	})

	g.Describe("Cross product", func() {
		g.It("Example 1", func() {
			v1 := tracer.NewVector(3, -3, 1)
			v2 := tracer.NewVector(4, 9, 2)
			result := v1.Cross(v2)

			assertVector(g, result, tracer.NewVector(-15, -2, 39))
		})

		g.It("Example 2 - UnitVectors", func() {
			v1 := tracer.NewVector(0.688247, -0.688247, 0.229416)
			v2 := tracer.NewVector(0.398015, 0.895533, 0.199007)
			result := v1.Cross(v2)

			assertVector(g, result, tracer.NewVector(-0.342415569457, -0.04565496148899999, 0.8902805303560002))
		})
	})
}

//http://www.wolframalpha.com/input/?i=intersect+line+((0,0,0),+(0,1,0))+with+triangle+(+(0,+1,+0),+(1,+1,+1),+(1,+1,+0))
func TestIntersection(t *testing.T) {
	g := Goblin(t)
	g.Describe("Intersection", func() {
		g.It("Ray with triangle", func() {
			ray := tracer.Ray{
				Origin:    tracer.NewVector(0, 0, 0),
				Direction: tracer.NewVector(0, 1, 0),
			}
			triangle := tracer.Triangle{
				P1: tracer.NewVector(-0.5, -0.5, -0.5),
				P2: tracer.NewVector(0.5, -0.5, -0.5),
				P3: tracer.NewVector(0, 0, 0.5),
			}
			yes, _ := triangle.Intersect(ray)

			g.Assert(yes).IsTrue()
		})
	})
}

func assertVector(g *G, actual, expected tracer.Vector) {
	if equalsFloat(actual.X, expected.X) &&
		equalsFloat(actual.X, expected.X) &&
		equalsFloat(actual.X, expected.X) {
		return
	}

	g.Fail(fmt.Sprintf("Actual value %+v is not equal to %+v", actual, expected))
}

func assertFloat(g *G, actual, expected float64) {
	if equalsFloat(actual, expected) {
		return
	}
	g.Fail(fmt.Sprintf("Actual value %f is not equal to %f", actual, expected))
}

func equalsFloat(val1, val2 float64) bool {
	return val1-val2 < epsilon && val2-val1 < epsilon
}
