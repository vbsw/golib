#ifndef VBSW_CDATA_H
#define VBSW_CDATA_H

#ifdef __cplusplus
extern "C" {
#endif

extern void vbsw_cdata_proc(int passes, int length, void **funcs, void **data, long long *err1, long long *err2, char **err_str);
extern void vbsw_cdata_free(void *data);

#ifdef __cplusplus
}
#endif

#endif /* VBSW_CDATA_H */