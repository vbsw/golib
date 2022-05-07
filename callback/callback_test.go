/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package callback

import (
	"testing"
)

func TestRegister(t *testing.T) {
	var cb Callback
	str := "something"
	objId := cb.Register(str)
	if len(cb.unused) != 0 {
		t.Error("wrong unused length:", len(cb.unused))
	}
	if objId != 0 {
		t.Error("wrong id:", objId)
	} else {
		if strBack, ok := cb.used[objId].(string); ok {
			if strBack != str {
				t.Error("string is wrong")
			}
		} else {
			t.Error("string cast failed")
		}
	}
}

func TestUnregister(t *testing.T) {
	var cb Callback
	str := "something"
	objId := cb.Register(str)
	if len(cb.unused) != 0 {
		t.Error("wrong unused length:", len(cb.unused))
	}
	if objId != 0 {
		t.Error("wrong id:", objId)
	} else {
		if strBack, ok := cb.Unregister(objId).(string); ok {
			if strBack != str {
				t.Error("returned string is wrong")
			}
		} else {
			t.Error("string cast failed")
		}
		if len(cb.unused) != 1 {
			t.Error("wrong unused length:", len(cb.unused))
		}
	}
}

func TestAll(t *testing.T) {
	var cb Callback
	strA := "something A"
	strB := "something B"
	strC := "something C"
	objIdA := cb.Register(strA)
	objIdB := cb.Register(strB)
	if len(cb.unused) != 0 {
		t.Error("wrong unused length:", len(cb.unused))
	}
	if objIdA != 0 {
		t.Error("wrong id A:", objIdA)
	} else if objIdB != 1 {
		t.Error("wrong id B:", objIdB)
	} else {
		if strBackA, ok := cb.Unregister(objIdA).(string); ok {
			if strBackA != strA {
				t.Error("returned string A is wrong")
			}
		} else {
			t.Error("string A cast failed")
		}
		if len(cb.unused) != 1 {
			t.Error("wrong unused length:", len(cb.unused))
		}
		objIdC := cb.Register(strC)
		if len(cb.unused) != 0 {
			t.Error("wrong unused length:", len(cb.unused))
		}
		if objIdC != 0 {
			t.Error("wrong id C:", objIdC)
		}
		objIdA = cb.Register(strA)
		if len(cb.unused) != 0 {
			t.Error("wrong unused length:", len(cb.unused))
		}
		if objIdA != 2 {
			t.Error("wrong id A:", objIdA)
		}
	}
}
