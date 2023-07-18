/*
 *          Copyright 2023, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package cdata initializes data in C.
package cdata

// #include "cdata.h"
import "C"
import (
	"unsafe"
)

type ErrorConvertor interface {
	ToError(int64, int64, string) error
}

type CData interface {
	CInitFunc() unsafe.Pointer
	SetCData(unsafe.Pointer)
}

func CInit(errConv ErrorConvertor, passes int, data ...CData) error {
	var err error
	dataLen := len(data)
	if passes > 0 && dataLen > 0 {
		var err1, err2 C.longlong
		var errStr *C.char
		datas := make([]unsafe.Pointer, dataLen)
		funcs := make([]unsafe.Pointer, dataLen)
		for i, d := range data {
			funcs[i] = d.CInitFunc()
		}
		C.vbsw_cdata_init(C.int(passes), &datas[0], &funcs[0], C.int(dataLen), &err1, &err2, &errStr);
		if err1 == 0 {
			for i, d := range data {
				d.SetCData(datas[i])
			}
		} else if errStr == nil {
			err = errConv.ToError(int64(err1), int64(err2), "")
		} else {
			err = errConv.ToError(int64(err1), int64(err2), C.GoString(errStr))
			C.vbsw_cdata_free(unsafe.Pointer(errStr))
		}
	}
	return err
}
