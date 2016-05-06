package vpsc

import (
	"log"
	"sync"
	"syscall"
	"unsafe"
)

var (
	removeOverlaps *syscall.Proc
)

func init() {
	dll, err := syscall.LoadDLL("vpsc.dll")
	if err != nil {
		log.Println("could not load vpsc.dll")
		return
	}

	removeOverlaps, err = dll.FindProc("remove_overlaps")
	if err != nil {
		log.Println("remove_overlaps missing from vpsc.dll")
	}
}

// Ready reports if the library was successfully loaded.
// RemoveOverlaps panics if Ready() is false.
func Ready() bool {
	return removeOverlaps != nil
}

// must match windll/vpsc.h
type crect struct {
	x0, x1, y0, y1       float64
	allow_overlap, fixed uint8
}

// removeoverlaps of libvpsc is not fully thread safe because of the
// xBorder and yBorder static variables, we protect the calls with
// this mutex to avoid problems during its call.
var removeOverlapsMtx sync.Mutex

func RemoveOverlaps(rects []Rect) {
	if removeOverlaps == nil {
		panic("vpsc.dll is not loaded")
	}
	if len(rects) == 0 {
		return
	}

	rcv := make([]crect, len(rects))
	for i, n := range rects {
		rc := &rcv[i]
		x0, y0, x1, y1 := n.Position()
		rc.x0 = x0
		rc.x1 = x1
		rc.y0 = y0
		rc.y1 = y1
		rc.fixed = boolchar(n.Fixed())
		rc.allow_overlap = boolchar(n.Overlap())
	}

	removeOverlapsMtx.Lock()
	removeOverlaps.Call(uintptr(unsafe.Pointer(&rcv[0])), uintptr(len(rcv)))
	removeOverlapsMtx.Unlock()

	for i, n := range rects {
		rc := &rcv[i]
		n.SetPosition(rc.x0, rc.x1, rc.y0, rc.y1)
	}
}

func boolchar(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
