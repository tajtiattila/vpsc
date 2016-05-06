// +build linux

#ifdef __cplusplus
extern "C" {
#endif

struct rect {
	double x0;
	double x1;
	double y0;
	double y1;

	char allow_overlap;
	char fixed;
};

void remove_overlaps(struct rect* r, unsigned n);

#ifdef __cplusplus
}
#endif
