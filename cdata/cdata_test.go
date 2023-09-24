/*
 *          Copyright 2023, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package cdata

import (
	"errors"
	"testing"
)

type testErrConv struct {
	defaultErrConv
}

func TestCalls(t *testing.T) {
	var ta testA
	params := &Parameters{nil, 2, 1, 1, []Initializer{&ta}}
	err := Init(params)
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
	params := &Parameters{new(testErrConv), 2, 1, 1, []Initializer{&tb}}
	err := Init(params)
	if err == nil {
		t.Error("error missing")
	} else if err.Error() != "9100" {
		t.Error(err.Error())
	}
}

func TestNamingExpand(t *testing.T) {
	var tc testC
	params := &Parameters{nil, 1, 1, 60, []Initializer{&tc}}
	err := Init(params)
	if err == nil {
	} else {
		t.Error(err.Error())
	}
}

func TestInsert(t *testing.T) {
	var td testD
	params := &Parameters{nil, 2, 1, 1, []Initializer{&td}}
	err := Init(params)
	if err != nil {
		t.Error(err.Error())
	}
}

func (errConv *testErrConv) ToError(err1, err2 int64, info string) error {
	if err1 == 9100 {
		if info == "abc" {
			return errors.New("9100")
		}
		return errors.New("9103:" + info + ";")
	}
	return errConv.defaultErrConv.ToError(err1, err2, info)
}
