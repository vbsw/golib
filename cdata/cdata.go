/*
 *          Copyright 2025, Vitali Baumtrok.
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
	"runtime"
	"strconv"
	"unsafe"
)

// Collection is a collection of Processings.
type Collection struct {
	procs []Processing
	funcs []unsafe.Pointer
	data  []unsafe.Pointer
	Err   error
	dirty bool
}

// Processing is an interface to manage C data and C functions.
type Processing interface {
	CFunc() unsafe.Pointer
	CData() unsafe.Pointer
	SetCData(unsafe.Pointer)
	ToError(int64, int64, string) error
}

// Add adds Processing to procs.
func (collection *Collection) Add(proc Processing) {
	collection.procs = append(collection.procs, proc)
	collection.dirty = true
}

// Reset sets len(procs) = 0.
func (collection *Collection) Reset() {
	collection.procs = collection.procs[:0]
	collection.funcs = collection.funcs[:0]
	collection.data = collection.data[:0]
	collection.Err = nil
	collection.dirty = false
}

// Process calls CFunc in sequence from 0 to len(procs) - 1 and backwards from len(procs)-1 to 0.
// Processing forward is one pass. Processing backwards is another pass.
//
// If error occurs, CFunc are called from index len(procs)-1 to 0 (once) with negative pass number.
// It is expected, that CFunc performes cleanup.
func (collection *Collection) Process(passes int) {
	collection.Err = nil
	if len(collection.procs) > 0 && passes > 0 {
		var pinner runtime.Pinner
		var err1, err2 C.longlong
		var errInfo *C.char
		length := len(collection.procs)
		collection.ensureCap(length)
		for i, proc := range collection.procs {
			collection.funcs[i] = proc.CFunc()
			collection.data[i] = proc.CData()
		}
		// calling with &collection.data[0] may cause "pointer to unpinned Go pointer" error
		// https://github.com/PowerDNS/lmdb-go/issues/28
		for _, data := range collection.data {
			if data != nil {
				pinner.Pin(data)
			}
		}
		C.vbsw_cdata_proc(C.int(passes), C.int(len(collection.procs)), &collection.funcs[0], &collection.data[0], &err1, &err2, &errInfo)
		pinner.Unpin()
		if err1 != 0 {
			var errStr string
			err1Int64, err2Int64 := int64(err1), int64(err2)
			if errInfo != nil {
				errStr = C.GoString(errInfo)
				C.vbsw_cdata_free(unsafe.Pointer(errInfo))
			}
			for i := length - 1; i >= 0 && collection.Err == nil; i-- {
				collection.Err = collection.procs[i].ToError(err1Int64, err2Int64, errStr)
			}
			if collection.Err == nil {
				collection.Err = toError(err1Int64, err2Int64, errStr)
			}
		}
	}
	collection.dirty = false
}

// Assign calls SetCData on all Processings.
func (collection *Collection) Assign() {
	if !collection.dirty {
		for i, proc := range collection.procs {
			proc.SetCData(collection.data[i])
		}
	} else {
		panic("cdata outdated, call Process first")
	}
}

func (collection *Collection) ensureCap(length int) {
	if len(collection.funcs) <= length {
		if cap(collection.funcs) <= length {
			collection.funcs = make([]unsafe.Pointer, length)
			collection.data = make([]unsafe.Pointer, length)
		} else {
			collection.funcs = collection.funcs[:length]
			collection.data = collection.data[:length]
		}
	}
}

// toError returns error numbers/string as error.
func toError(err1, err2 int64, info string) error {
	var errStr string
	if err1 == 1 {
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
