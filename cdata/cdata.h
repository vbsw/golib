#ifndef VBSW_CDATA_H
#define VBSW_CDATA_H

#ifdef __cplusplus
extern "C" {
#endif

typedef struct { void **list; int err1, err2, ids_len, ids_cap, offs_len, offs_cap; char *err_str, *ids; int *props; void *set_func, *get_func; } cdata_t;
extern void vbsw_cdata_init(int passes, void **data, void **funcs, int length, long long *err1, long long *err2, char **err_str);
extern void vbsw_cdata_free(void *data);
extern void vbsw_cdata_testa(int pass, cdata_t *cdata);
extern void vbsw_cdata_testb(int pass, cdata_t *cdata);
extern void vbsw_cdata_testc(int pass, cdata_t *cdata);
extern void vbsw_cdata_testd(int pass, cdata_t *cdata);
extern void vbsw_cdata_teste(int pass, cdata_t *cdata);
extern void vbsw_cdata_testf(int pass, cdata_t *cdata);
extern void vbsw_cdata_testg(int pass, cdata_t *cdata);

#ifdef __cplusplus
}
#endif

#endif /* VBSW_CDATA_H */