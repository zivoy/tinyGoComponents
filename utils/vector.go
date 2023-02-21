package utils

import (
	"math"
)

type Vector3 struct {
	X, Y, Z float64
}

type Plane struct {
	Normal, Up Vector3
}

func (p Plane) GetRotation(vector Vector3) (angle, magnitude float64) {
	// Step 1: Project vector onto plane
	projV := p.Project(vector)

	// Step 2: Find 2D coordinates of projected vector
	xBasis, yBasis := p.GetBasis()
	projX := projV.Dot(xBasis)
	projY := projV.Dot(yBasis)

	return math.Atan2(projY, projX), projV.Magnitude()
}

func (p Plane) Project(v Vector3) Vector3 {
	perp := p.Normal.Dot(v) / p.Normal.LengthSquared()
	perpVector := Vector3{p.Normal.X * perp, p.Normal.Y * perp, p.Normal.Z * perp}
	return Vector3{v.X - perpVector.X, v.Y - perpVector.Y, v.Z - perpVector.Z}
}

// GetBasis gets the two basis vectors that span a plane
func (p Plane) GetBasis() (Vector3, Vector3) {
	xBasis := p.Normal.Cross(p.Up)
	yBasis := p.Normal.Cross(xBasis)
	return xBasis, yBasis
}

func (v Vector3) Dot(o Vector3) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vector3) Cross(o Vector3) Vector3 {
	return Vector3{
		v.Y*o.Z - v.Z*o.Y,
		v.Z*o.X - v.X*o.Z,
		v.X*o.Y - v.Y*o.X,
	}
}

func (v Vector3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector3) Magnitude() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vector3) Normalized() Vector3 {
	mag := v.Magnitude()
	return Vector3{v.X / mag, v.Y / mag, v.Z / mag}
}
