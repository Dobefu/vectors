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
	Bounce()
	Normalize()
	AngleRadians() float64
	AngleDegrees() float64
	IsZero() bool
	Magnitude() float64
	ClampMagnitude(maxValue float64)
	Clear()
	ToVector2() Vector2
}

// Vector3 provides a Vector3.
type Vector3 struct {
	IVector3

	X float64
	Y float64
	Z float64
}

// Add adds the values of another vector to this one.
func (v *Vector3) Add(vec Vector3) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
}

// Sub subtracts the values of another vector to this one.
func (v *Vector3) Sub(vec Vector3) {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
}

// Mul multiplies the values of another vector to this one.
func (v *Vector3) Mul(vec Vector3) {
	v.X *= vec.X
	v.Y *= vec.Y
	v.Z *= vec.Z
}

// Div divides the values of this vector by those of the other vector.
func (v *Vector3) Div(vec Vector3) {
	v.X /= vec.X
	v.Y /= vec.Y
	v.Z /= vec.Z
}

// Bounce inverts the direction of the vector.
func (v *Vector3) Bounce() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

// Normalize scales the vector magnitude to 1.
func (v *Vector3) Normalize() {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)

	if magnitude != 0 {
		v.X /= magnitude
		v.Y /= magnitude
		v.Z /= magnitude
	}
}

// AngleRadians returns the angle of the vector in radians.
func (v Vector3) AngleRadians() float64 {
	return math.Atan2(v.Y, v.X)
}

// AngleDegrees returns the angle of the vector in degrees.
func (v Vector3) AngleDegrees() float64 {
	angle := math.Atan2(v.Y, v.X) * 180 / math.Pi

	return math.Mod(angle+360, 360)
}

// IsZero returns whether or not the vector is zero on all axes.
func (v *Vector3) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

// Magnitude returns the length of the vector.
func (v *Vector3) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// ClampMagnitude limits the magnitude of the vector to a maximum value.
func (v *Vector3) ClampMagnitude(maxValue float64) {
	magnitude := v.Magnitude()

	if magnitude == 0 || magnitude <= maxValue {
		return
	}

	scale := maxValue / magnitude

	v.Mul(Vector3{
		X: scale,
		Y: scale,
		Z: scale,
	})
}

// Clear sets the vector to zero on all axes.
func (v *Vector3) Clear() {
	v.X = 0
	v.Y = 0
	v.Z = 0
}

// ToVector2 converts the vector to a Vector2, discarding the Z value.
func (v Vector3) ToVector2() Vector2 {
	return Vector2{
		X: v.X,
		Y: v.Y,
	}
}
