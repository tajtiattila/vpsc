// +build linux

package vpsc

/*
#cgo CXXFLAGS: -I ./adaptagrams/cola
#include "vpsc_unix.h"
*/
import "C"

import "sync"

// removeoverlaps of libvpsc is not fully thread safe because of the
// xBorder and yBorder static variables, we protect the calls with
// this mutex to avoid problems during its call.
var removeOverlapsMtx sync.Mutex

func RemoveOverlaps(nv []Node) error {
	if len(nv) == 0 {
		return nil
	}

	rcv := make([]C.struct_rect, len(nv))
	for i, n := range nv {
		rc := &rcv[i]
		x0, x1, y0, y1 := n.Rect()
		rc.x0 = C.double(x0)
		rc.x1 = C.double(x1)
		rc.y0 = C.double(y0)
		rc.y1 = C.double(y1)
		rc.fixed = boolchar(n.Fixed())
		rc.allow_overlap = boolchar(n.Overlap())
	}

	removeOverlapsMtx.Lock()
	C.remove_overlaps(&rcv[0], C.unsigned(len(rcv)))
	removeOverlapsMtx.Unlock()

	for i, n := range nv {
		rc := &rcv[i]
		x0 := float64(rc.x0)
		x1 := float64(rc.x1)
		y0 := float64(rc.y0)
		y1 := float64(rc.y1)
		n.SetRect(x0, x1, y0, y1)
	}

	return nil
}

func boolchar(b bool) C.char {
	if b {
		return 1
	}
	return 0
}
