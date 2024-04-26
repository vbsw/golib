/*
 *          Copyright 2024, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package cdata initializes data in C.
//
// Inspect the implementation for more information on how to use this package.
package cdata

// #include "cdata.h"
import "C"
import (
	"errors"
	"strconv"
	"unsafe"
)

// ErrorConv converts error numbers/strings to error.
type ErrorConv interface {
	ToError(int64, int64, string) error
}

// CData is an interface that provides a function to initialize C data.
// After initialization the C data is passed to the function SetCData.
// SetCData is called regardless of whether C data is nil or not (i.e. has been initialized or not).
type CData interface {
	CInitFunc() unsafe.Pointer
	SetCData(unsafe.Pointer)
}

// Params holds the initialization parameters.
type Params struct {
	ErrConvs []ErrorConv
	Passes   int
	ListCap  int
	WordsCap int
	Data     []CData
}

// Init processes params.Data forwards from index 0 to len(params.Data)-1 and backwards from len(params.Data)-1 to 0.
// Processing forwards is one pass. Processing backwards is another pass. How many passes is set via params.Passes.
// params.ListCap and params.WordsCap is used to initialize internal data. Each will be at least >= len(params.Data).
// If error occurs, SetCData isn't called. (It is expected, that the C initialization function already performed the cleanup.)
func Init(params *Params) error {
	var err error
	dataLen := len(params.Data)
	if params.Passes > 0 && dataLen > 0 {
		var err1, err2 C.longlong
		var errInfo *C.char
		data := make([]unsafe.Pointer, dataLen)
		funcs := make([]unsafe.Pointer, dataLen)
		listCap, wordsCap := params.capacities()
		for i, inz := range params.Data {
			funcs[i] = inz.CInitFunc()
		}
		C.vbsw_cdata_init(C.int(params.Passes), &data[0], &funcs[0], C.int(dataLen), C.int(listCap), C.int(wordsCap), &err1, &err2, &errInfo)
		if err1 == 0 {
			for i, inz := range params.Data {
				inz.SetCData(data[i])
			}
		} else {
			var errStr string
			err1Int64, err2Int64 := int64(err1), int64(err2)
			if errInfo != nil {
				errStr = C.GoString(errInfo)
				C.vbsw_cdata_free(unsafe.Pointer(errInfo))
			}
			for i := 0; i < len(params.ErrConvs) && err == nil; i++ {
				err = params.ErrConvs[i].ToError(err1Int64, err2Int64, errStr)
			}
			if err == nil {
				err = toError(err1Int64, err2Int64, errStr)
			}
		}
	}
	return err
}

// capacities returns the minimum values for capacities.
func (params *Params) capacities() (int, int) {
	var listCap, wordsCap int
	dataLen := len(params.Data)
	if params.ListCap >= dataLen {
		listCap = params.ListCap
	} else {
		listCap = dataLen
	}
	if params.WordsCap >= dataLen {
		wordsCap = params.WordsCap
	} else {
		wordsCap = dataLen * 60
	}
	return listCap, wordsCap
}

// toError returns error numbers/string as error.
func toError(err1, err2 int64, info string) error {
	var errStr string
	if err1 > 0 && err1 < 1000000 {
		errStr = "memory allocation failed"
	} else {
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
