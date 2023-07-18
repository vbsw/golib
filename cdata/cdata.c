/*
 *          Copyright 2023, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

#include <stdlib.h>
#include <string.h>
#include "cdata.h"
#include "cdata_funcs.h"

static void cdata_set(cdata_t *const cdata, const char *const id, void *const data) {
	if (!cdata[0].err1) {
		const int length = cdata[0].length;
		const int curr = cdata[0].curr;
		const int last_idx = length - 1;
		int *const ids_offs = cdata[0].ids_props;
		const int ids_len = cdata[0].ids_len;
		const int id_new_len0 = id ? strlen(id) + 1 : 1;
		const int id_old_len0 = curr < last_idx ? ids_offs[curr+1] - ids_offs[curr] : ids_len - ids_offs[curr];
		const int id_diff = id_new_len0 - id_old_len0;
		const int id_len0_min = ids_len + id_diff;
		/* ensure capacity */
		if (id_len0_min > cdata[0].ids_cap) {
			int ids_cap_new = cdata[0].ids_cap;
			do {
				ids_cap_new *= 2;
			} while (id_len0_min < ids_cap_new);
			char *const ids_new = (char*)malloc((size_t)ids_cap_new);
			if (ids_new) {
				memcpy(ids_new, cdata[0].ids, (size_t)cdata[0].ids_len);
				free(cdata[0].ids);
				cdata[0].ids = ids_new;
				cdata[0].ids_cap = ids_cap_new;
			} else {
				cdata[0].err1 = 4;
			}
		}
		if (!cdata[0].err1) {
			char *const ids = cdata[0].ids;
			int *const ids_sorts = &cdata[0].ids_props[length];
			int ids_sort_len = cdata[0].sort_len;
			/* remove from sort */
			if (id_old_len0 > 1) {
				/* binary search */
				const char *const id_old = &ids[ids_offs[curr]]; int left = 0, right = ids_sort_len - 1;
				while (left <= right) {
					const int middle = (left + right) / 2;
					const int id_idx = ids_sorts[middle];
					const char *const id_curr = (const char*)&ids[ids_offs[id_idx]];
					const int result = strcmp(id_curr, id_old);
					if (result < 0) {
						left = middle + 1;
					} else if (result > 0) {
						right = middle - 1;
					} else {
						/* remove */
						const int rest_len = ids_sort_len - middle - 2;
						if (rest_len > 0) {
							int *const src = &ids_sorts[middle+1];
							int *const dst = &ids_sorts[middle];
							memmove(dst, src, sizeof(int) * (size_t)rest_len);
						}
						ids_sort_len--;
						break;
					}
				}
			}
			/* adjust space for new id */
			if (id_diff != 0) {
				int i;
				const int rest_len0 = curr < last_idx ? ids_len - (ids_offs[curr+1]) : 0;
				if (rest_len0 > 0) {
					char *const src = &ids[ids_offs[curr+1]];
					char *const dst = &src[id_diff];
					cdata[0].ids_len = ids_len + id_diff;
					memmove(dst, src, (size_t)rest_len0);
				}
				for (i = curr + 1; i < length; i++)
					ids_offs[i] += id_diff;
			}
			/* add new id */
			if (id_new_len0 > 1) {
				memcpy(&ids[ids_offs[curr]], id, (size_t)id_new_len0);
				/* add to sort */
				int left = 0, right = ids_sort_len - 1;
				while (left <= right) {
					const int middle = (left + right) / 2;
					const int id_idx = ids_sorts[middle];
					const char *const id_curr = (const char*)&ids[ids_offs[id_idx]];
					const int result = strcmp(id_curr, id);
					if (result < 0) {
						left = middle + 1;
					} else if (result > 0) {
						right = middle - 1;
					}
				}
				const int rest_len = ids_sort_len - left - 2;
				if (rest_len > 0) {
					int *const src = &ids_sorts[left];
					int *const dst = &ids_sorts[left+1];
					memmove(dst, src, sizeof(int) * (size_t)rest_len);
				}
				ids_sorts[left] = curr;
				ids_sort_len++;
			} else {
				ids[ids_offs[curr]] = 0;
			}
			/* others */
			cdata[0].list[curr] = data;
			cdata[0].sort_len = ids_sort_len;
		}
	}
}

static void* cdata_get(cdata_t *const cdata, const char *const id) {
	if (id) {
		const int length = cdata[0].length;
		const int last_idx = length - 1;
		const char *const ids = (const char*)cdata[0].ids;
		int *const ids_offs = cdata[0].ids_props;
		int *const ids_sorts = &cdata[0].ids_props[length];
		/* binary search */
		int left = 0, right = cdata[0].sort_len - 1;
		while (left <= right) {
			const int middle = (left + right) / 2;
			const int id_idx = ids_sorts[middle];
			const char *const id_curr = &ids[ids_offs[id_idx]];
			const int result = strcmp(id_curr, id);
			if (result < 0)
				left = middle + 1;
			else if (result > 0)
				right = middle - 1;
			else
				return cdata[0].list[id_idx];
		}
	}
	return NULL;
}

void vbsw_cdata_init(const int passes, void **const data, void **const funcs, const int length, long long *const err1, long long *const err2, char **const err_str) {
	if (passes > 0 && length > 0) {
		const size_t list_size = sizeof(void*) * (size_t)length;
		cdata_init_func_t *const init_funcs = (cdata_init_func_t*)funcs;
		cdata_t cdata; memset(&cdata, 0, sizeof(cdata_t));
		cdata.list = (void**)malloc(list_size);
		if (cdata.list) {
			const size_t ids_size = sizeof(char) * (size_t)length * 60;
			cdata.ids = (char*)malloc(ids_size);
			if (cdata.ids) {
				const size_t ids_props_size = sizeof(int) * (length * 2);
				cdata.ids_props = (int*)malloc(ids_props_size);
				if (cdata.ids_props) {
					int pass, i;
					int *const ids_offs = cdata.ids_props;
					int *const ids_sorts = &cdata.ids_props[length];
					cdata.length = length;
					cdata.ids_len = length;
					cdata.ids_cap = ids_size;
					cdata.set_func = cdata_set;
					cdata.get_func = cdata_get;
					memset(cdata.list, 0, list_size);
					memset(cdata.ids, 0, ids_size);
					for (i = 0; i < length; i++)
						ids_offs[i] = i;
					/* main */
					for (pass = 0; pass < passes && !cdata.err1; pass++) {
						/* forwards */
						for (i = 0; i < length && !cdata.err1 && !cdata.err2; i++) {
							cdata_init_func_t const init_func = init_funcs[i];
							if (init_func)
								init_func(pass, &cdata);
						}
						/* backwards */
						pass++;
						if (pass < passes) {
							for (i = length - 1; i >= 0 && !cdata.err1; i--) {
								cdata_init_func_t const init_func = init_funcs[i];
								if (init_func)
									init_func(pass, &cdata);
							}
						}
					}
					/* return */
					if (!cdata.err1) {
						for (i = 0; i < length; i++)
							data[i] = cdata.list[i];
					/* error handling */
					} else {
						for (i = length - 1; i >= 0; i--) {
							cdata_init_func_t const init_func = init_funcs[i];
							if (init_func)
								init_func(-1, &cdata);
						}
						err1[0] = cdata.err1;
						err2[0] = cdata.err2;
						err_str[0] = cdata.err_str;
					}
					/* clean up */
					free(cdata.list);
					free(cdata.ids);
					free(cdata.ids_props);
				} else {
					err1[0] = 3; free(cdata.list); free(cdata.ids);
				}
			} else {
				err1[0] = 2; free(cdata.list);
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
