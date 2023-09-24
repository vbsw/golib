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
	"errors"
	"strconv"
	"unsafe"
)

type ErrorConvertor interface {
	ToError(int64, int64, string) error
}

type Initializer interface {
	FuncCInit() unsafe.Pointer
	SetCData(unsafe.Pointer)
}

type Parameters struct {
	ErrConv   ErrorConvertor
	Passes    int
	CListCap  int
	CWordsCap int
	Inits     []Initializer
}

type defaultErrConv struct {
}

func Init(params *Parameters) error {
	var err error
	initsLen := len(params.Inits)
	if params.Passes > 0 && initsLen > 0 {
		var err1, err2 C.longlong
		var errInfo *C.char
		data := make([]unsafe.Pointer, initsLen)
		funcs := make([]unsafe.Pointer, initsLen)
		listCap, wordsCap := params.capacities()
		for i, inz := range params.Inits {
			funcs[i] = inz.FuncCInit()
		}
		C.vbsw_cdata_init(C.int(params.Passes), &data[0], &funcs[0], C.int(initsLen), C.int(listCap), C.int(wordsCap), &err1, &err2, &errInfo)
		if err1 == 0 {
			for i, inz := range params.Inits {
				inz.SetCData(data[i])
			}
		} else {
			if errInfo == nil {
				err = params.errConv().ToError(int64(err1), int64(err2), "")
			} else {
				err = params.errConv().ToError(int64(err1), int64(err2), C.GoString(errInfo))
				C.vbsw_cdata_free(unsafe.Pointer(errInfo))
			}
		}
	}
	return err
}

func (params *Parameters) capacities() (int, int) {
	var listCap, wordsCap int
	initsLen := len(params.Inits)
	if params.CListCap >= initsLen {
		listCap = params.CListCap
	} else {
		listCap = initsLen
	}
	if params.CWordsCap >= initsLen {
		wordsCap = params.CWordsCap
	} else {
		wordsCap = initsLen * 60
	}
	return listCap, wordsCap
}

func (params *Parameters) errConv() ErrorConvertor {
	if params.ErrConv != nil {
		return params.ErrConv
	}
	return new(defaultErrConv)
}

func (errConv *defaultErrConv) ToError(err1, err2 int64, info string) error {
	var errStr string
	if err1 > 0 && err1 < 1000 {
		errStr = "memory allocation failed"
	} else {
		errStr = "unknown error"
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

type testA struct {
	state int
}

type testB struct {
}

type testC struct {
}

type testD struct {
}

func (ta *testA) FuncCInit() unsafe.Pointer {
	if ta.state == 0 {
		ta.state = 1
	} else {
		ta.state = 2
	}
	return unsafe.Pointer(C.vbsw_cdata_testa)
}

func (ta *testA) SetCData(unsafe.Pointer) {
	if ta.state == 1 {
		ta.state = 3
	} else {
		ta.state += 10
	}
}

func (tb *testB) FuncCInit() unsafe.Pointer {
	return unsafe.Pointer(C.vbsw_cdata_testb)
}

func (ta *testB) SetCData(unsafe.Pointer) {
}

func (tc *testC) FuncCInit() unsafe.Pointer {
	return unsafe.Pointer(C.vbsw_cdata_testc)
}

func (tc *testC) SetCData(unsafe.Pointer) {
}

func (td *testD) FuncCInit() unsafe.Pointer {
	return unsafe.Pointer(C.vbsw_cdata_testd)
}

func (td *testD) SetCData(unsafe.Pointer) {
}
