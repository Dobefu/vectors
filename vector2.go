package vectors

import (
	"math"
)

// IVector2 is the interface for a 2D vector.
type IVector2 interface {
	Add(vec Vector2)
	Sub(vec Vector2)
	Mul(vec Vector2)
	Div(vec Vector2)
	Bounce()
	Normalize()
	AngleRadians() float64
	AngleDegrees() float64
	IsZero() bool
	Magnitude() float64
	ClampMagnitude(maxValue float64)
	Clear()
	ToVector3() Vector2
}

// Vector2 provides a Vector2.
type Vector2 struct {
	IVector2

	X float64
	Y float64
}

// Add adds the values of another vector to this one.
func (v *Vector2) Add(vec Vector2) {
	v.X += vec.X
	v.Y += vec.Y
}

// Sub subtracts the values of another vector to this one.
func (v *Vector2) Sub(vec Vector2) {
	v.X -= vec.X
	v.Y -= vec.Y
}

// Mul multiplies the values of another vector to this one.
func (v *Vector2) Mul(vec Vector2) {
	v.X *= vec.X
	v.Y *= vec.Y
}

// Div divides the values of this vector by those of the other vector.
func (v *Vector2) Div(vec Vector2) {
	v.X /= vec.X
	v.Y /= vec.Y
}

// Bounce inverts the direction of the vector.
func (v *Vector2) Bounce() {
	v.X = -v.X
	v.Y = -v.Y
}

// Normalize scales the vector magnitude to 1.
func (v *Vector2) Normalize() {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)

	if magnitude != 0 {
		v.X /= magnitude
		v.Y /= magnitude
	}
}

// AngleRadians returns the angle of the vector in radians.
func (v Vector2) AngleRadians() float64 {
	return math.Atan2(v.Y, v.X)
}

// AngleDegrees returns the angle of the vector in degrees.
func (v Vector2) AngleDegrees() float64 {
	angle := math.Atan2(v.Y, v.X) * 180 / math.Pi

	return math.Mod(angle+360, 360)
}

// IsZero returns whether or not the vector is zero on all axes.
func (v *Vector2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// Magnitude returns the length of the vector.
func (v *Vector2) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

// ClampMagnitude limits the magnitude of the vector to a maximum value.
func (v *Vector2) ClampMagnitude(maxValue float64) {
	magnitude := v.Magnitude()

	if magnitude == 0 || magnitude <= maxValue {
		return
	}

	scale := maxValue / magnitude

	v.Mul(Vector2{
		X: scale,
		Y: scale,
	})
}

// Clear sets the vector to zero on all axes.
func (v *Vector2) Clear() {
	v.X = 0
	v.Y = 0
}

// ToVector3 converts the vector to a Vector3 with a Z value of 0.
func (v Vector2) ToVector3() Vector3 {
	return Vector3{
		X: v.X,
		Y: v.Y,
		Z: 0,
	}
}
