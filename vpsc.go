package vpsc

type Node interface {
	// Rect reports the initial location
	Rect() (x0, x1, y0, y1 float64)

	// SetRect updates the node location
	SetRect(x0, x1, y0, y1 float64)

	// Overlap reports wether a node may be overlapped by others
	Overlap() bool

	// Fixed reports if the node should not be moved
	Fixed() bool
}
