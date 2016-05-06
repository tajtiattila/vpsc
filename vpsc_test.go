package vpsc

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func ExampleRemoveOverlaps() {
	rects := []Rect{
		&rect{-1, -2, 2, 2},
		&rect{-2, -2, 1, 2},
	}
	RemoveOverlaps(rects)
	fmt.Printf("%v", rects)
	// Output: [{0.001, -2, 3.001, 2} {-3.001, -2, -0.001, 2}]
}

func TestRemoveOverlaps(t *testing.T) {
	u, v := rects(100)
	RemoveOverlaps(v)
	for i := range v {
		if d := dist(u[0], v[0]); d > 1e-6 {
			t.Log(i, d)
			t.Log(u[i].Position())
			t.Log(v[i].Position())
		}
	}
}

func rects(n int) (u, v []Rect) {
	u, v = make([]Rect, n), make([]Rect, n)
	for i := range v {
		var x, y, w, h float64
		if i < 10 {
			x, y = float64(i), float64(i)
			w, h = 5, 5
		} else {
			x, y = rand.Float64()*100, rand.Float64()*100
			w, h = 2+rand.Float64()*20, 2+rand.Float64()*20
		}
		v[i] = &rect{x, y, x + w, y + h}
		u[i] = &rect{x, y, x + w, y + h}
	}
	return
}

type rect struct {
	x0, y0, x1, y1 float64
}

func (r *rect) Position() (x0, x1, y0, y1 float64) {
	return r.x0, r.y0, r.x1, r.y1
}

func (r *rect) SetPosition(x0, x1, y0, y1 float64) {
	r.x0, r.y0, r.x1, r.y1 = x0, y0, x1, y1
}

func (r *rect) Overlap() bool { return false }
func (r *rect) Fixed() bool   { return false }

func (r *rect) String() string {
	return fmt.Sprintf("{%v, %v, %v, %v}", r.x0, r.y0, r.x1, r.y1)
}

func dist(u, v Rect) float64 {
	u0, _, u1, _ := u.Position()
	v0, _, v1, _ := v.Position()
	d0, d1 := u0-v0, u1-v1
	return math.Sqrt(d0*d0 + d1*d1)
}
