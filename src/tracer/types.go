package tracer

type Ray struct {
	Start     Vector
	Direction Vector
}

type Vector struct {
	X, Y, Z float32
}

func NewVector(x, y, z float32) Vector {
	return Vector{X: x, Y: y, Z: z}
}
