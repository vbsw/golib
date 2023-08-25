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
	"fmt"
)

type ErrorConvertor interface {
	ToError(int64, int64, string) error
}

type CData interface {
	CInitFunc() unsafe.Pointer
	SetCData(unsafe.Pointer)
}

type Config struct {
	ErrorConv ErrorConvertor
	Passes int
	SortCap int
	WordCap int
}

type defaultErrConv struct {
}

func (cfg *Config) ensureValidity(dataLen int) {
	if cfg.SortCap < dataLen {
		cfg.SortCap = dataLen
	}
	if cfg.WordCap < dataLen {
		cfg.WordCap = dataLen
	}
	if cfg.ErrorConv == nil {
		cfg.ErrorConv = new(defaultErrConv)
	}
}

func CInit(cfg Config, data ...CData) error {
	var err error
	dataLen := len(data)
	if cfg.Passes > 0 && dataLen > 0 {
		var err1, err2 C.longlong
		var errInfo *C.char
		datas := make([]unsafe.Pointer, dataLen)
		funcs := make([]unsafe.Pointer, dataLen)
		cfg.ensureValidity(dataLen)
		for i, d := range data {
			funcs[i] = d.CInitFunc()
		}
		C.vbsw_cdata_init(C.int(cfg.Passes), &datas[0], &funcs[0], C.int(dataLen), C.int(cfg.SortCap), C.int(cfg.WordCap), &err1, &err2, &errInfo);
		if err1 == 0 {
			for i, d := range data {
				d.SetCData(datas[i])
			}
		} else {
			if errInfo == nil {
				err = cfg.ErrorConv.ToError(int64(err1), int64(err2), "")
			} else {
				err = cfg.ErrorConv.ToError(int64(err1), int64(err2), C.GoString(errInfo))
				C.vbsw_cdata_free(unsafe.Pointer(errInfo))
			}
		}
	}
	return err
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

func (ta *testA) CInitFunc() unsafe.Pointer {
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

func (tb *testB) CInitFunc() unsafe.Pointer {
	return unsafe.Pointer(C.vbsw_cdata_testb)
}

func (ta *testB) SetCData(unsafe.Pointer) {
}

func (tc *testC) CInitFunc() unsafe.Pointer {
	return unsafe.Pointer(C.vbsw_cdata_testc)
}

func (tc *testC) SetCData(unsafe.Pointer) {
}

//export goDebug
func goDebug(a, b, c, d C.int) {
	fmt.Println(a, b, c, d)
}
