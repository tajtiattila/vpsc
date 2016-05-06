// Package vpsc is binding to libvpsc from Adaptagrams (http://www.adaptagrams.org).
// Libvpsc is a C++ library for solving for the Variable Placement with Separation Constraints problem.
package vpsc

// Rectangles represent the subject for overlap removal.
type Rectangles interface {
	// Len reports the number of rectangles.
	Len() int

	// Position reports the position of the ith Rect.
	Position(i int) (x0, y0, x1, y1 float64)

	// SetPosition updates the position of the ith Rect.
	SetPosition(i int, x0, y0, x1, y1 float64)

	// AllowOverlap reports if ith Rect may overlap with others.
	AllowOverlap(i int) bool

	// Fixed reports if the ith Rect should not be moved.
	Fixed(i int) bool
}
