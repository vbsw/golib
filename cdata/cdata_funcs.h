/*
 *          Copyright 2023, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

typedef struct { void **list; int curr, length, err1, err2, ids_len, ids_cap, sort_len; char *err_str, *ids; int *ids_props; void *set_func, *get_func; } cdata_t;
typedef void (*cdata_set_func_t)(cdata_t *cdata, const char *id, void *data);
typedef void (*cdata_get_func_t)(cdata_t *cdata, const char *id);
typedef void (*cdata_init_func_t)(int pass, cdata_t *cdata);
