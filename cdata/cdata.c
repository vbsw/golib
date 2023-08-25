/*
 *          Copyright 2023, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

#include <stdlib.h>
#include <string.h>
#include "cdata.h"

typedef struct { void **all; char *err_str; void *set_func, *get_func; long long err1, err2; int list_len, list_cap, words_len, words_cap; } cdata_t;
typedef void (*cdata_set_func_t)(cdata_t *cdata, void *data, const char *id);
typedef void* (*cdata_get_func_t)(cdata_t *cdata, const char *id);
typedef void (*cdata_init_func_t)(int pass, cdata_t *cdata);

static void cdata_set(cdata_t *const cdata, const char *const id, void *const data) {
	const int list_cap = cdata[0].list_cap;
	void **const all = cdata[0].all;
	int *const offs = (int*)(&all[list_cap]);
	int *const sort = &offs[list_cap];
	char *const words = (char*)&sort[list_cap];
	const char *const id0 = id ? id : "";
	/* binary search */
	int left = 0, right = cdata[0].list_len - 1;
	while (left <= right) {
		const int middle = (left + right) / 2;
		const int list_idx = sort[middle];
		char *const id_curr = &words[offs[list_idx]];
		const int result = strcmp(id_curr, id0);
		if (result < 0) {
			left = middle + 1;
		} else if (result > 0) {
			right = middle - 1;
		} else {
			/* reaplace */
			all[list_idx] = data;
			return;
		}
	}
	/* insert */
	const int list_len = cdata[0].list_len;
	const int words_len = cdata[0].words_len;
	const int words_cap = cdata[0].words_cap;
	const int id0_len = (int)strlen(id0) + 1;
	if (list_len >= list_cap || words_len + id0_len > words_cap) {
		const int list_cap_new = (list_len < list_cap) ? list_cap : list_cap * 2;
		int words_cap_new = words_cap;
		while(words_len + id0_len > words_cap_new)
			words_cap_new *= 2;
		const size_t size_new = sizeof(void*) * (size_t)list_cap_new + sizeof(int) * (size_t)(list_cap_new*2) + sizeof(char) * (size_t)words_cap_new;
		void **const all_new = (void**)malloc(size_new);
		if (all_new) {
			int *const offs_new = (int*)(&all_new[list_cap_new]);
			int *const sort_new = &offs_new[list_cap_new];
			char *const words_new = (char*)(&offs_new[list_cap_new*2]);
			memcpy(all_new, all, sizeof(void*) * (size_t)list_len);
			memcpy(offs_new, offs, sizeof(int) * (size_t)list_len);
			if (left > 0)
				memcpy(sort_new, sort, sizeof(int) * (size_t)left);
			if (left < list_len)
				memcpy(&sort_new[left+1], &sort[left], sizeof(int) * (size_t)(list_len-left));
			memcpy(words_new, words, sizeof(char) * (size_t)words_len);
			free(cdata[0].all);
			cdata[0].all = all_new;
			cdata[0].list_cap = list_cap_new;
			cdata[0].words_cap = words_cap_new;
			offs_new[list_len] = words_len;
			sort_new[left] = list_len;
			memcpy(&words_new[words_len], id0, sizeof(char) * (size_t)id0_len);
		} else {
			cdata[0].err1 = 2;
			return;
		}
	} else {
		offs[list_len] = words_len;
		sort[left] = list_len;
		memcpy(&words[words_len], id0, sizeof(char) * (size_t)id0_len);
	}
	cdata[0].all[cdata[0].list_len] = data;
	cdata[0].list_len++;
	cdata[0].words_len += id0_len;
}

static void* cdata_get(cdata_t *const cdata, const char *const id) {
	const int list_cap = cdata[0].list_cap;
	void **const all = cdata[0].all;
	int *const offs = (int*)(&all[list_cap]);
	int *const sort = &offs[list_cap];
	char *const words = (char*)&sort[list_cap];
	const char *const id0 = id ? id : "";
	/* binary search */
	int left = 0, right = cdata[0].list_len - 1;
	while (left <= right) {
		const int middle = (left + right) / 2;
		const int list_idx = sort[middle];
		char *const id_curr = &words[offs[list_idx]];
		const int result = strcmp(id_curr, id0);
		if (result < 0) {
			left = middle + 1;
		} else if (result > 0) {
			right = middle - 1;
		} else {
			return all[list_idx];
		}
	}
	return NULL;
}

void vbsw_cdata_init(const int passes, void **const data, void **const funcs, const int length, const int l_cap, const int w_cap, long long *const err1, long long *const err2, char **const err_str) {
	if (passes > 0 && length > 0) {
		const size_t list_size = sizeof(void*) * (size_t)l_cap;
		cdata_init_func_t *const init_funcs = (cdata_init_func_t*)funcs;
		cdata_t cdata; memset(&cdata, 0, sizeof(cdata_t));
		cdata.all = (void**)malloc(list_size + sizeof(int) * (size_t)(l_cap*2) + sizeof(char) * (size_t)w_cap);
		if (cdata.all) {
			int pass, i;
			cdata.list_cap = l_cap;
			cdata.words_cap = w_cap;
			cdata.set_func = cdata_set;
			cdata.get_func = cdata_get;
			memset(cdata.all, 0, list_size);
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
					data[i] = cdata.all[i];
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
			free(cdata.all);
		} else {
			err1[0] = 1;
		}
	}
}

void vbsw_cdata_free(void *const data) {
	if (data)
		free(data);
}
