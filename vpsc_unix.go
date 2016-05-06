// +build linux

package vpsc

/*
#cgo CXXFLAGS: -I ./adaptagrams/cola
#include "vpsc_unix.h"
*/
import "C"

import "sync"

// Ready reports if the library was successfully loaded.
// RemoveOverlaps panics if Ready() is false.
func Ready() bool {
	return true
}

// removeoverlaps of libvpsc is not fully thread safe because of the
// xBorder and yBorder static variables, we protect the calls with
// this mutex to avoid problems during its call.
var removeOverlapsMtx sync.Mutex

func RemoveOverlaps(rects Rectangles) {
	if rects.Len() == 0 {
		return
	}

	rcv := make([]C.struct_rect, rects.Len())
	for i := 0; i < rects.Len(); i++ {
		rc := &rcv[i]
		x0, y0, x1, y1 := rects.Position(i)
		rc.x0 = C.double(x0)
		rc.x1 = C.double(x1)
		rc.y0 = C.double(y0)
		rc.y1 = C.double(y1)
		rc.fixed = boolchar(rects.Fixed(i))
		rc.allow_overlap = boolchar(rects.AllowOverlap(i))
	}

	removeOverlapsMtx.Lock()
	C.remove_overlaps(&rcv[0], C.unsigned(len(rcv)))
	removeOverlapsMtx.Unlock()

	for i := 0; i < rects.Len(); i++ {
		rc := &rcv[i]
		x0 := float64(rc.x0)
		x1 := float64(rc.x1)
		y0 := float64(rc.y0)
		y1 := float64(rc.y1)
		rects.SetPosition(i, x0, y0, x1, y1)
	}
}

func boolchar(b bool) C.char {
	if b {
		return 1
	}
	return 0
}
