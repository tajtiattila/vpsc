#ifdef __cplusplus
extern "C" {
#endif

#ifdef BUILDING_VPSC_DLL
#define VPSC_DLL __declspec(dllexport)
#else
#define VPSC_DLL __declspec(dllimport)
#endif

struct rect {
	double x0;
	double x1;
	double y0;
	double y1;

	char allow_overlap;
	char fixed;
};

__stdcall VPSC_DLL
void remove_overlaps(struct rect* r, unsigned n);

#ifdef __cplusplus
}
#endif
