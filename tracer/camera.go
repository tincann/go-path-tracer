package tracer

import (
	"image"
	"math"
	"math/rand"
)

type Camera struct {
	Eye, Direction, Up Vector
	right              Vector //vector pointing right

	phi, theta float64 //latitude and longitude of look direction

	imagePlaneWidth, imagePlaneHeight float64 //width and height of the view plane
	distance                          float64 //camera distance from view plane
	imagePlane                        *imagePlane

	moveSpeed, lookSpeed float64
}

//Represents the upper and left border of the image plane
type imagePlane struct {
	TopLeft             Vector
	UpperEdge, LeftEdge Vector
}

type CameraMoveDirection int

func NewCamera(eye, direction, up Vector, dImagePlane, imagePlaneWidth, imagePlaneHeight float64, moveSpeed, lookSpeed float64) *Camera {
	direction = direction.Normalize()
	phi, theta := directionToAngles(direction)
	c := &Camera{
		Eye: eye, Direction: direction, Up: up.Normalize(),
		distance: dImagePlane, imagePlaneWidth: imagePlaneWidth, imagePlaneHeight: imagePlaneHeight,
		moveSpeed: moveSpeed, lookSpeed: lookSpeed,
		phi: phi, theta: theta,
	}
	c.updateImagePlane()
	return c
}

func (c *Camera) GenerateRays(pixelDimensions, area image.Rectangle, aaFactor float64) []*RayInfo {
	maxX, maxY := pixelDimensions.Dx(), pixelDimensions.Dy()

	topLeft, upper, left := c.imagePlane.TopLeft, c.imagePlane.UpperEdge, c.imagePlane.LeftEdge
	rayInfos := make([]*RayInfo, area.Dx()*area.Dy())
	i := 0

	for x := area.Min.X; x < area.Max.X; x++ {
		for y := area.Min.Y; y < area.Max.Y; y++ {
			vx := (float64(x) + 0.5 + rand.Float64()*aaFactor - aaFactor/2) / float64(maxX)
			vy := (float64(y) + 0.5 + rand.Float64()*aaFactor - aaFactor/2) / float64(maxY)

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

func (c *Camera) Rotate(deltaPhi, deltaTheta float64) {

	c.phi += deltaPhi * c.lookSpeed
	c.theta += deltaTheta * c.lookSpeed

	if c.theta < 0 {
		c.theta = 0
	} else if c.theta > math.Pi {
		c.theta = math.Pi
	}

	//calculate new direction vector
	c.Direction = anglesToDirection(c.phi, c.theta)
	c.updateImagePlane()
}

func anglesToDirection(phi, theta float64) Vector {
	return NewVector(
		math.Cos(phi)*math.Sin(theta),
		math.Sin(phi)*math.Sin(theta),
		math.Cos(theta),
	)
}

func directionToAngles(direction Vector) (phi, theta float64) {
	phi = math.Atan2(direction.Y, direction.X)
	theta = math.Acos(direction.Z)
	return phi, theta
}

func (c *Camera) updateImagePlane() {
	center := c.Eye.Add(c.Direction.Multiply(c.distance))
	left := c.Direction.Cross(Vector{0, 0, 1})
	c.right = left.Multiply(-1)
	c.Up = left.Cross(c.Direction)

	c.imagePlane = &imagePlane{
		TopLeft:   center.Add(left.Multiply(c.imagePlaneWidth / 2)).Add(c.Up.Multiply(c.imagePlaneHeight / 2)),
		UpperEdge: left.Multiply(-c.imagePlaneWidth),
		LeftEdge:  c.Up.Multiply(-c.imagePlaneHeight),
	}
}
