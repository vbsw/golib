/*
 *        Copyright 2024, 2025 Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package cdata processes C data.
package cdata

// #include "cdata.h"
import "C"
import (
	"errors"
	"strconv"
	"unsafe"
	"fmt"
)

// CData is an interface to call C functions. CData are processed as list in sequence.
// First, all CProcFunc's are called. Then, the returned data are passed in sequence
// to SetCData.
//
// If an error occurs, CProcFunc's are called in reversed sequence with a negative
// pass number. (It is expected, that CProcFunc performes cleanup.) In this case
// SetCData is not called. The error is passed to all ToError in a reversed sequence
// until any returns not nil. (First parameter of ToError should be unique.)
type CData interface {
	CProcFunc() unsafe.Pointer
	SetCData(unsafe.Pointer)
	ToError(int64, int64, string) error
}

//export goDebug
func goDebug(a, b C.int, c, d C.ulong) {
	fmt.Println(a, b, c, d)
}

// Proc processes CData in sequence.
//
// The second parameter is the number of passes. The third parameter is the number of IDs in ID list.
// The fourth parameter is the minumum IDs word capacity in bytes.
//
// If second or third parameter is set to 0, then no memory is allocated to manage IDs and data.
// (Calling cdata_set or cdata_get will cause an error.)
//
// CProcFunc are called from index 0 to len(data)-1 and backwards from len(cdata)-1 to 0.
// Processing forward is one pass. Processing backwards is another pass. Only in first pass
// data (that will be passed to SetCData) can be set via cdata_set. If not set, it is initialized
// to NULL.
//
// If error occurs, SetCData isn't called and CProcFunc are called from index len(cdata)-1 to 0 (once).
// It is expected, that CProcFunc performes cleanup.
func Proc(data []CData, params ...int) error {
	var err error
	dataLen := len(data)
	passes, listCap, wordsCap := getParams(params, dataLen)
	if passes > 0 && dataLen > 0 {
		var err1, err2 C.longlong
		var errInfo *C.char
		dataC := make([]unsafe.Pointer, dataLen)
		funcs := make([]unsafe.Pointer, dataLen)
		for i, dat := range data {
			funcs[i] = dat.CProcFunc()
		}
		C.vbsw_cdata_proc(C.int(passes), &dataC[0], &funcs[0], C.int(dataLen), C.int(listCap), C.int(wordsCap), &err1, &err2, &errInfo)
		if err1 == 0 {
			for i, dat := range data {
				dat.SetCData(dataC[i])
			}
		} else {
			var errStr string
			err1Int64, err2Int64 := int64(err1), int64(err2)
			if errInfo != nil {
				errStr = C.GoString(errInfo)
				C.vbsw_cdata_free(unsafe.Pointer(errInfo))
			}
			for i := len(data) - 1; i >= 0 && err == nil; i-- {
				err = data[i].ToError(err1Int64, err2Int64, errStr)
			}
			if err == nil {
				err = toError(err1Int64, err2Int64, errStr)
			}
		}
	}
	return err
}

func getParams(params []int, dataLen int) (int, int, int) {
	var passes, listCap, wordsCap int
	if len(params) > 2 {
		if params[2] > 0 {
			wordsCap = params[2]
		}
	} else {
		wordsCap = dataLen * 56
	}
	if len(params) > 1 {
		if wordsCap > 0 && params[1] > 0 {
			listCap = params[1]
		}
	} else {
		listCap = dataLen
	}
	if len(params) > 0 && params[0] >= 0 {
		passes = params[0]
	} else {
		passes = 1
	}
	return passes, listCap, wordsCap
}

// toError returns error numbers/string as error.
func toError(err1, err2 int64, info string) error {
	var errStr string
	if err1 > 0 {
		/* 1 - 100, 1000001 - 1000100 */
		if err1 < 1000001 {
			errStr = "memory allocation failed"
		} else if err1 == 1000001 {
			errStr = "cdata_set has been called although no memory has been allocated to manage IDs"
		} else if err1 == 1000002 {
			errStr = "cdata_get has been called although no memory has been allocated to manage IDs"
		}
	}
	if len(errStr) == 0 {
		errStr = "unknown"
	}
	errStr = errStr + " (" + strconv.FormatInt(err1, 10)
	if err2 == 0 {
		errStr = errStr + ")"
	} else {
		errStr = errStr + ", " + strconv.FormatInt(err2, 10) + ")"
	}
	if len(info) > 0 {
		errStr = errStr + "; " + info
	}
	return errors.New(errStr)
}
