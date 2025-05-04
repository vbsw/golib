/*
 *          Copyright 2025, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

#include <stdlib.h>
#include <string.h>
#include "cdata.h"

typedef void (*cdata_proc_func_t)(int pass, int index, int length, void** data, const char** ids, long long* err1, long long* err2, char** err_str);

void vbsw_cdata_proc(const int passes, const int length, void** const funcs, void** const data, long long* const err1, long long* const err2, char** const err_str) {
	if (passes > 0 && length > 0) {
		const size_t ids_size = sizeof(const char*) * (size_t)length;
		const char** const ids = (const char**)malloc(ids_size);
		if (ids) {
			int pass, i;
			cdata_proc_func_t* const proc_funcs = (cdata_proc_func_t*)funcs;
			memset((void*)ids, 0, ids_size);
			/* main */
			for (pass = 0; pass < passes;) {
				/* forward */
				for (i = 0; i < length && err1[0] == 0; i++) {
					cdata_proc_func_t const proc_func = proc_funcs[i];
					if (proc_func)
						proc_func(pass, i, length, data, ids, err1, err2, err_str);
				}
				/* backwards */
				if (err1[0] == 0) {
					pass++;
					if (pass < passes) {
						for (i = length - 1; i >= 0 && err1[0] == 0; i--) {
							cdata_proc_func_t const proc_func = proc_funcs[i];
							if (proc_func)
								proc_func(pass, i, length, data, ids, err1, err2, err_str);
						}
					}
				}
				if (err1[0] == 0)
					pass++;
				else
					break;
			}
			/* error handling */
			if (err1[0]) {
				pass = -(pass + 1);
				for (i = length - 1; i >= 0; i--) {
					cdata_proc_func_t const proc_func = proc_funcs[i];
					if (proc_func)
						proc_func(pass, i, length, data, ids, err1, err2, err_str);
				}
			}
			free(ids);
		} else {
			err1[0] = 1;
		}
	}
}

void vbsw_cdata_free(void *const data) {
	if (data)
		free(data);
}
