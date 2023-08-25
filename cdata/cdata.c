/*
 *          Copyright 2023, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

#include <stdlib.h>
#include <string.h>
#include "cdata.h"

#include "_cgo_export.h"

typedef void (*cdata_set_func_t)(cdata_t *cdata, const char *id, void *data);
typedef void (*cdata_get_func_t)(cdata_t *cdata, const char *id);
typedef void (*cdata_init_func_t)(int pass, cdata_t *cdata);

static void cdata_set(cdata_t *const cdata, const char *const id, void *const data) {
	const char *const ids = (const char*)cdata[0].ids;
	int *const offs = cdata[0].props;
	int *const sort = &cdata[0].props[cdata[0].list_cap];
	const int last_idx = cdata[0].list_len - 1;
	const char *const id0 = id ? id : "";
	/* binary search */
	int left = 0, right = last_idx;
	while (left <= right) {
		const int middle = (left + right) / 2;
		const int list_idx = sort[middle];
		const char *const id_curr = &ids[offs[list_idx]];
		const int result = strcmp(id_curr, id0);
		if (result < 0) {
			left = middle + 1;
		} else if (result > 0) {
			right = middle - 1;
		} else {
			cdata[0].list[list_idx] = data;
			break;
		}
	}
	/* add */
	if (left > right) {
		const int id0_len = (int)strlen(id0) + 1;
		/* ensure id capacity */
		if (cdata[0].ids_cap - cdata[0].ids_len < id0_len) {
			int ids_cap_new = cdata[0].ids_cap;
			do {
				ids_cap_new *= 2;
			} while(ids_cap_new < id0_len);
			char *const ids_new = (char*)malloc(sizeof(char) * (size_t)ids_cap_new);
			memcpy(ids_new, cdata[0].ids, sizeof(char) * (size_t)cdata[0].ids_len);
			free(cdata[0].ids);
			cdata[0].ids = ids_new;
			cdata[0].ids_cap = ids_cap_new;
		}
		/* ensure props capacity */
		if (cdata[0].list_len >= cdata[0].list_cap) {
			const int list_cap_new = cdata[0].list_cap * 2;
			int *const props_new = (int*)malloc(sizeof(int) * (size_t)(list_cap_new*2));
			void **const list_new = (void**)malloc(sizeof(void*) * (size_t)list_cap_new);
			memcpy(props_new, offs, sizeof(int) * (size_t)cdata[0].list_len);
			memcpy(list_new, cdata[0].list, sizeof(void*) * (size_t)cdata[0].list_len);
			if (left > 0)
				memcpy(&props_new[list_cap_new], sort, sizeof(int) * (size_t)left);
			if (left < cdata[0].list_len)
				memcpy(&props_new[list_cap_new+left+1], &sort[left], sizeof(int) * (size_t)(cdata[0].list_len-left));
			free(cdata[0].props);
			free(cdata[0].list);
			cdata[0].props = props_new;
			cdata[0].list = list_new;
			cdata[0].list_cap = list_cap_new;
		} else if (left < cdata[0].list_len) {
			memmove(&sort[left+1], &sort[left], sizeof(int) * (size_t)(cdata[0].list_len-left));
		}
		cdata[0].props[cdata[0].list_len] = cdata[0].ids_len;
		cdata[0].props[cdata[0].list_cap+left] = cdata[0].list_len;
		cdata[0].list[cdata[0].list_len] = data;
		memcpy(&cdata[0].ids[cdata[0].ids_len], id0, sizeof(char) * (size_t)id0_len);
		cdata[0].ids_len += id0_len;
		cdata[0].list_len++;
	}
}

static void* cdata_get(cdata_t *const cdata, const char *const id) {
/*
	if (id) {
		const int length = cdata[0].length;
		const int last_idx = length - 1;
		const char *const ids = (const char*)cdata[0].ids;
		int *const offs = cdata[0].props;
		int *const ids_sort = &cdata[0].props[length];
*/
		/* binary search */
/*
		int left = 0, right = cdata[0].sort_len - 1;
		while (left <= right) {
			const int middle = (left + right) / 2;
			const int id_idx = ids_sort[middle];
			const char *const id_curr = &ids[offs[id_idx]];
			const int result = strcmp(id_curr, id);
			if (result < 0)
				left = middle + 1;
			else if (result > 0)
				right = middle - 1;
			else
				return cdata[0].list[id_idx];
		}
	}
*/
	return NULL;
}

void vbsw_cdata_init(const int passes, void **const data, void **const funcs, const int length, const int l_cap, const int w_cap, long long *const err1, long long *const err2, char **const err_str) {
	if (passes > 0 && length > 0) {
		cdata_init_func_t *const init_funcs = (cdata_init_func_t*)funcs;
		cdata_t cdata; memset(&cdata, 0, sizeof(cdata_t));
		const int list_cap = l_cap;
		cdata.props = (int*)malloc(sizeof(int) * (size_t)(list_cap*2));
		if (cdata.props) {
			const size_t list_size = sizeof(void*) * (size_t)list_cap;
			cdata.list = (void**)malloc(list_size);
			if (cdata.list) {
				const int ids_cap = w_cap;
				cdata.ids = (char*)malloc(sizeof(char) * (size_t)ids_cap);
				if (cdata.ids) {
					int pass, i;
					cdata.ids_cap = ids_cap;
					cdata.list_cap = list_cap;
					cdata.set_func = cdata_set;
					cdata.get_func = cdata_get;
					memset(cdata.list, 0, list_size);
					/* main */
					for (pass = 0; pass < passes;) {
						/* forwards */
						for (i = 0; i < length && !cdata.err1; i++) {
							cdata_init_func_t const init_func = init_funcs[i];
							if (init_func)
								init_func(pass, &cdata);
						}
						/* backwards */
						if (!cdata.err1) {
							pass++;
							if (pass < passes) {
								for (i = length - 1; i >= 0 && !cdata.err1; i--) {
									cdata_init_func_t const init_func = init_funcs[i];
									if (init_func)
										init_func(pass, &cdata);
								}
							}
						}
						if (!cdata.err1)
							pass++;
						else
							break;
					}
					/* return */
					if (!cdata.err1) {
						for (i = 0; i < length; i++)
							data[i] = cdata.list[i];
					/* error handling */
					} else {
						pass = -(pass + 1);
						for (i = length - 1; i >= 0; i--) {
							cdata_init_func_t const init_func = init_funcs[i];
							if (init_func)
								init_func(pass, &cdata);
						}
						err1[0] = cdata.err1;
						err2[0] = cdata.err2;
						err_str[0] = cdata.err_str;
					}
					/* clean up */
					free(cdata.list);
					free(cdata.ids);
					free(cdata.props);
				} else {
					err1[0] = 3; free(cdata.props); free(cdata.list);
				}
			} else {
				err1[0] = 2; free(cdata.props);
			}
		} else {
			err1[0] = 1;
		}
	}
}

void vbsw_cdata_free(void *const data) {
	if (data)
		free(data);
}

void vbsw_cdata_testa(const int pass, cdata_t *const cdata) {
}

void vbsw_cdata_testb(const int pass, cdata_t *const cdata) {
	if (pass == 1) {
		cdata[0].err_str = (char*)malloc(sizeof(char) * 4);
		if (cdata[0].err_str) {
			cdata[0].err1 = 9000;
			cdata[0].err_str[0] = 'a';
			cdata[0].err_str[1] = 'b';
			cdata[0].err_str[2] = 'c';
			cdata[0].err_str[3] = 0;
		} else {
			cdata[0].err1 = 9001;
		}
	} else if (pass < 0 && pass != -2) {
		cdata[0].err1 = 9002;
	}
}

void vbsw_cdata_testc(const int pass, cdata_t *const cdata) {
	cdata_set_func_t const set = (cdata_set_func_t)cdata[0].set_func;
	set(cdata, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", NULL); /* 50x */
	set(cdata, "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", NULL);
	set(cdata, "cccccccccccccccccccccccccccccccccccccccccccccccccc", NULL);
	set(cdata, "dddddddddddddddddddddddddddddddddddddddddddddddddd", NULL);
	set(cdata, "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", NULL);
	set(cdata, "ffffffffffffffffffffffffffffffffffffffffffffffffff", NULL);
	set(cdata, "gggggggggggggggggggggggggggggggggggggggggggggggggg", NULL);
}

void vbsw_cdata_testd(const int pass, cdata_t *const cdata) {
}

void vbsw_cdata_teste(const int pass, cdata_t *const cdata) {
}

void vbsw_cdata_testf(const int pass, cdata_t *const cdata) {
}

void vbsw_cdata_testg(const int pass, cdata_t *const cdata) {
}
