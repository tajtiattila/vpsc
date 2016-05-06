// Package vpsc is binding to libvpsc from Adaptagrams (http://www.adaptagrams.org).
// Libvpsc is a C++ library for solving for the Variable Placement with Separation Constraints problem.
package vpsc

// Rect is a subject for overlap removal.
type Rect interface {
	// Position reports the position of this Rect.
	Position() (x0, y0, x1, y1 float64)

	// SetPosition updates the position of this Rect.
	SetPosition(x0, y0, x1, y1 float64)

	// Overlap reports if this Rect may overlap with others.
	Overlap() bool

	// Fixed reports if the this Rect should not be moved.
	Fixed() bool
}
