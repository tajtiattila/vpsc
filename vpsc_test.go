package vpsc

import (
	"math"
	"math/rand"
	"testing"
)

func TestRemoveOverlaps(t *testing.T) {
	u, v := nodes(100)
	err := RemoveOverlaps(v)
	if err != nil {
		t.Error(err)
	}
	for i := range v {
		if d := dist(u[0], v[0]); d > 1e-6 {
			t.Log(i, d)
			t.Log(u[i].Rect())
			t.Log(v[i].Rect())
		}
	}
}

func nodes(n int) (u, v []Node) {
	u, v = make([]Node, n), make([]Node, n)
	for i := range v {
		var x, y, w, h float64
		if i < 10 {
			x, y = float64(i), float64(i)
			w, h = 5, 5
		} else {
			x, y = rand.Float64()*100, rand.Float64()*100
			w, h = 2+rand.Float64()*20, 2+rand.Float64()*20
		}
		v[i] = &rect{x, x + w, y, y + h}
		u[i] = &rect{x, x + w, y, y + h}
	}
	return
}

type rect struct {
	x0, x1, y0, y1 float64
}

func (r *rect) Rect() (x0, x1, y0, y1 float64) {
	return r.x0, r.x1, r.y0, r.y1
}

func (r *rect) SetRect(x0, x1, y0, y1 float64) {
	r.x0, r.x1, r.y0, r.y1 = x0, x1, y0, y1
}

func (r *rect) Overlap() bool { return false }
func (r *rect) Fixed() bool   { return false }

func dist(u, v Node) float64 {
	u0, u1, _, _ := u.Rect()
	v0, v1, _, _ := v.Rect()
	d0, d1 := u0-v0, u1-v1
	return math.Sqrt(d0*d0 + d1*d1)
}
