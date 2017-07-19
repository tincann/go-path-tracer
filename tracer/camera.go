package tracer

import (
	"image"
)

type Camera struct {
	Eye, Direction, Up Vector
	right              Vector //vector pointing right

	imagePlaneWidth, imagePlaneHeight float64 //width and height of the view plane
	distance                          float64 //camera distance from view plane
	imagePlane                        *imagePlane

	moveSpeed float64
}

//Represents the upper and left border of the image plane
type imagePlane struct {
	TopLeft             Vector
	UpperEdge, LeftEdge Vector
}

type CameraMoveDirection int

func NewCamera(eye, direction, up Vector, dImagePlane, imagePlaneWidth, imagePlaneHeight float64, moveSpeed float64) *Camera {
	c := &Camera{
		Eye: eye, Direction: direction.Normalize(), Up: up.Normalize(),
		distance: dImagePlane, imagePlaneWidth: imagePlaneWidth, imagePlaneHeight: imagePlaneHeight,
		moveSpeed: moveSpeed,
	}
	c.updateImagePlane()
	return c
}

func (c *Camera) GenerateRays(pixelDimensions image.Point, area image.Rectangle) []*RayInfo {
	maxX, maxY := pixelDimensions.X, pixelDimensions.Y

	topLeft, upper, left := c.imagePlane.TopLeft, c.imagePlane.UpperEdge, c.imagePlane.LeftEdge
	rayInfos := make([]*RayInfo, area.Dx()*area.Dy())
	i := 0
	for x := area.Min.X; x < area.Max.X; x++ {
		for y := area.Min.Y; y < area.Max.Y; y++ {
			vx := float64(x) / float64(maxX)
			vy := float64(y) / float64(maxY)

			p := topLeft.Add(upper.Multiply(vx)).Add(left.Multiply(vy))
			direction := p.Subtract(c.Eye).Normalize()

			rayInfo := &RayInfo{
				Ray: Ray{Origin: c.Eye, Direction: direction},
				X:   x, Y: y,
			}
			rayInfos[i] = rayInfo
			i++
		}
	}

	return rayInfos
}

//x = strafe, y = forward/backward, z = up/down
func (c *Camera) Move(direction Vector) {
	strafe := c.right.Multiply(direction.X * c.moveSpeed)
	forward := c.Direction.Multiply(direction.Y * c.moveSpeed)
	up := c.Up.Multiply(direction.Z * c.moveSpeed)

	c.Eye = c.Eye.Add(strafe).Add(forward).Add(up)
	c.updateImagePlane()
}

func (c *Camera) Rotate(deltaTheta, deltaPhi float64) {
	//rotate round up vector
	// a1 := c.Direction.Multiply(math.Cos(deltaTheta))
	// a2 := c.right.Multiply(math.Sin(deltaTheta))
	// a3 := c.Direction.Multiply(c.Up.Dot(c.Direction) * (1 - math.Cos(deltaTheta)))
	// c.Direction = a1.Add(a2).Add(a3)
	// c.updateImagePlane()
}

func (c *Camera) updateImagePlane() {
	center := c.Eye.Add(c.Direction.Multiply(c.distance))
	left := c.Direction.Cross(c.Up)
	c.right = left.Multiply(-1)

	c.imagePlane = &imagePlane{
		TopLeft:   center.Add(left.Multiply(c.imagePlaneWidth / 2)).Add(c.Up.Multiply(c.imagePlaneHeight / 2)),
		UpperEdge: left.Multiply(-c.imagePlaneWidth),
		LeftEdge:  c.Up.Multiply(-c.imagePlaneHeight),
	}
}
