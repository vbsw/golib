/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package cdata

import (
	"errors"
	"testing"
	"time"
)

type testErrConv struct {
	defaultErrConv
}

func TestNamingExpansion(t *testing.T) {
	var tc testC
	var cfg Config = Config{nil, 2, 1, 60}
	err := CInit(cfg, &tc)
	time.Sleep(time.Second)
	if err == nil {
	} else {
		t.Error(err.Error())
	}
}

func TestCalls(t *testing.T) {
	var ta testA
	var cfg Config = Config{nil, 2, 1, 1}
	err := CInit(cfg, &ta)
	time.Sleep(time.Second)
	if err == nil {
		if ta.state != 3 {
			t.Error("order failed:", ta.state)
		}
	} else {
		t.Error(err.Error())
	}
}

func TestErrors(t *testing.T) {
	var tb testB
	var cfg Config = Config{new(testErrConv), 2, 1, 1}
	err := CInit(cfg, &tb)
	time.Sleep(time.Second)
	if err == nil {
		t.Error("error missing")
	} else if err.Error() != "9000" {
		t.Error(err.Error())
	}
}

func (errConv *testErrConv) ToError(err1, err2 int64, info string) error {
	if err1 == 9000 {
		if info == "abc" {
			return errors.New("9000")
		}
		return errors.New("9003:" + info + ";")
	}
	return errConv.defaultErrConv.ToError(err1, err2, info)
}
