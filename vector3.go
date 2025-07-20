package vectors

import (
	"math"
)

// IVector3 is the interface for a 3D vector.
type IVector3 interface {
	Add(vec Vector3)
	Sub(vec Vector3)
	Mul(vec Vector3)
	Div(vec Vector3)
	Scale(scale float64)
	Bounce()
	Normalize()
	AngleRadians() float64
	AngleDegrees() float64
	IsZero() bool
	Magnitude() float64
	MagnitudeSquared() float64
	Distance(vec Vector3) float64
	DistanceSquared(vec Vector3) float64
	Dot(vec Vector3) float64
	Lerp(vec Vector3, t float64)
	ClampMagnitude(maxValue float64)
	Clear()
	ToVector2() Vector2
}

// Vector3 represents a 3D vector with X, Y, and Z coordinates.
// It provides methods for vector operations.
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

// Add adds the values of another vector to this one.
func (v *Vector3) Add(vec Vector3) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
}

// Sub subtracts the values of another vector from this one.
func (v *Vector3) Sub(vec Vector3) {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
}

// Mul multiplies this vector by another vector.
func (v *Vector3) Mul(vec Vector3) {
	v.X *= vec.X
	v.Y *= vec.Y
	v.Z *= vec.Z
}

// Div divides this vector by another vector.
func (v *Vector3) Div(vec Vector3) {
	v.X /= vec.X
	v.Y /= vec.Y
	v.Z /= vec.Z
}

// Scale multiplies this vector by a scale.
func (v *Vector3) Scale(scale float64) {
	v.X *= scale
	v.Y *= scale
	v.Z *= scale
}

// Bounce inverts the direction of the vector.
func (v *Vector3) Bounce() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

// Normalize scales the vector to have a magnitude of 1.
func (v *Vector3) Normalize() {
	magnitudeSquared := v.X*v.X + v.Y*v.Y + v.Z*v.Z

	if magnitudeSquared != 0 {
		magnitude := math.Sqrt(magnitudeSquared)
		v.X /= magnitude
		v.Y /= magnitude
		v.Z /= magnitude
	}
}

// AngleRadians returns the angle in radians.
func (v Vector3) AngleRadians() float64 {
	return math.Atan2(v.Y, v.X)
}

// AngleDegrees returns the angle of the vector in degrees.
// This method ignores the Z axis and projects the vector onto the XY plane.
func (v Vector3) AngleDegrees() float64 {
	angle := math.Atan2(v.Y, v.X) * 180 / math.Pi

	if angle < 0 {
		angle += 360
	}

	return angle
}

// IsZero checks if all axes are zero.
func (v Vector3) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

// Magnitude returns the length of the vector.
func (v Vector3) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// MagnitudeSquared returns the squared magnitude of the vector.
func (v Vector3) MagnitudeSquared() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

// Distance returns the distance between this vector and another vector.
func (v Vector3) Distance(vec Vector3) float64 {
	dx := v.X - vec.X
	dy := v.Y - vec.Y
	dz := v.Z - vec.Z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// DistanceSquared returns the squared distance between this vector and another vector.
func (v Vector3) DistanceSquared(vec Vector3) float64 {
	dx := v.X - vec.X
	dy := v.Y - vec.Y
	dz := v.Z - vec.Z
	return dx*dx + dy*dy + dz*dz
}

// Dot returns the dot product.
// Positive = same direction, negative = opposite, zero = perpendicular.
func (v Vector3) Dot(vec Vector3) float64 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z
}

// Lerp interpolates between this vector and another vector.
func (v *Vector3) Lerp(vec Vector3, t float64) {
	v.X += (vec.X - v.X) * t
	v.Y += (vec.Y - v.Y) * t
	v.Z += (vec.Z - v.Z) * t
}

// ClampMagnitude limits the magnitude of the vector to a maximum value.
func (v *Vector3) ClampMagnitude(maxValue float64) {
	maxSquared := maxValue * maxValue
	magnitudeSquared := v.MagnitudeSquared()

	if magnitudeSquared == 0 || magnitudeSquared <= maxSquared {
		return
	}

	scale := maxValue / math.Sqrt(magnitudeSquared)
	v.X *= scale
	v.Y *= scale
	v.Z *= scale
}

// Clear sets the vector to zero.
func (v *Vector3) Clear() {
	v.X = 0
	v.Y = 0
	v.Z = 0
}

// ToVector2 converts the 3D vector to a 2D vector.
func (v Vector3) ToVector2() Vector2 {
	return Vector2{
		X: v.X,
		Y: v.Y,
	}
}
