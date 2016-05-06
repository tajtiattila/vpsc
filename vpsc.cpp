// +build linux

#include "vpsc_unix.h"

#include <set>
#include "libvpsc/rectangle.h"
#include "libvpsc/assertions.h"

//#include <cstdio>

using namespace vpsc;

void remove_overlaps(struct rect* rc, unsigned n) {
	Rectangles rs(n);
    std::set<unsigned> fixed;
	struct rect* p = rc;
	for (unsigned i = 0; i < n; i++, p++) {
        //printf("%d %f %f %f %f %d\n", i, p->x0, p->x1, p->y0, p->y1, p->fixed);
		rs[i] = new Rectangle(p->x0, p->x1, p->y0, p->y1, p->allow_overlap);
        if (p->fixed) {
            fixed.insert(i);
        }
	}
    removeoverlaps(rs, fixed);
	p = rc;
	for (unsigned i = 0; i < n; i++, p++) {
        Rectangle const* r = rs[i];
        p->x0 = r->getMinX();
        p->x1 = r->getMaxX();
        p->y0 = r->getMinY();
        p->y1 = r->getMaxY();
    }
}

