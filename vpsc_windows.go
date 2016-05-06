package vpsc

import (
	"errors"
	"log"
	"sync"
	"syscall"
	"unsafe"
)

var (
	xerr           error
	removeOverlaps *syscall.Proc
)

func init() {
	dll, err := syscall.LoadDLL("vpsc.dll")
	if err != nil {
		xerr = errors.New("could not load vpsc.dll")
		log.Println(xerr)
		return
	}

	removeOverlaps, err = dll.FindProc("remove_overlaps")
	if err != nil {
		xerr = errors.New("remove_overlaps missing from vpsc.dll")
		log.Println(xerr)
	}
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

func RemoveOverlaps(nv []Node) error {
	if xerr != nil {
		return xerr
	}
	if len(nv) == 0 {
		return nil
	}

	rcv := make([]crect, len(nv))
	for i, n := range nv {
		rc := &rcv[i]
		x0, x1, y0, y1 := n.Rect()
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

	for i, n := range nv {
		rc := &rcv[i]
		n.SetRect(rc.x0, rc.x1, rc.y0, rc.y1)
	}

	return nil
}

func boolchar(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
