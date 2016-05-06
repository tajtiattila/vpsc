package vpsc_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/tajtiattila/vpsc"
)

func ExampleRemoveOverlaps() {
	v := []rect{
		{-1, -2, 2, 2},
		{-2, -2, 1, 2},
	}
	vpsc.RemoveOverlaps(rects(v))
	fmt.Printf("%v", v)
	// Output: [{0.001 -2 3.001 2} {-3.001 -2 -0.001 2}]
}

func TestRemoveOverlaps(t *testing.T) {
	u, v := makeRects(100)
	vpsc.RemoveOverlaps(rects(v))
	for i := range v {
		if d := dist(u[0], v[0]); d > 1e-6 {
			t.Log(i, d)
			t.Log(rects(u).Position(i))
			t.Log(rects(v).Position(i))
		}
	}
}

func makeRects(n int) (u, v []rect) {
	u, v = make([]rect, n), make([]rect, n)
	for i := range v {
		var x, y, w, h float64
		if i < 10 {
			x, y = float64(i), float64(i)
			w, h = 5, 5
		} else {
			x, y = rand.Float64()*100, rand.Float64()*100
			w, h = 2+rand.Float64()*20, 2+rand.Float64()*20
		}
		v[i] = rect{x, y, x + w, y + h}
		u[i] = rect{x, y, x + w, y + h}
	}
	return
}

type rect struct {
	x0, y0, x1, y1 float64
}

type rects []rect

func (v rects) Len() int { return len(v) }

func (v rects) Position(i int) (x0, y0, x1, y1 float64) {
	r := v[i]
	return r.x0, r.y0, r.x1, r.y1
}

func (v rects) SetPosition(i int, x0, y0, x1, y1 float64) {
	r := &v[i]
	r.x0, r.y0, r.x1, r.y1 = x0, y0, x1, y1
}

func (r rects) AllowOverlap(i int) bool { return false }
func (r rects) Fixed(i int) bool        { return false }

func dist(u, v rect) float64 {
	u0, u1 := u.x0, u.x1
	v0, v1 := v.x0, v.x1
	d0, d1 := u0-v0, u1-v1
	return math.Sqrt(d0*d0 + d1*d1)
}
