#ifndef VBSW_CDATA_H
#define VBSW_CDATA_H

#ifdef __cplusplus
extern "C" {
#endif

typedef struct { void *set_func, *get_func; char *err_str; void **all; long long err1, err2; int list_len, list_cap, words_len, words_cap; } cdata_t;
extern void vbsw_cdata_init(int passes, void **data, void **funcs, int length, int l_cap, int w_cap, long long *err1, long long *err2, char **err_str);
extern void vbsw_cdata_free(void *data);
extern void vbsw_cdata_testa(int pass, cdata_t *cdata);
extern void vbsw_cdata_testb(int pass, cdata_t *cdata);
extern void vbsw_cdata_testc(int pass, cdata_t *cdata);
extern void vbsw_cdata_testd(int pass, cdata_t *cdata);

#ifdef __cplusplus
}
#endif

#endif /* VBSW_CDATA_H */