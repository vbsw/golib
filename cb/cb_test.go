/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package cb

import (
	"testing"
)

func TestAll(t *testing.T) {
	strA := "something A"
	strB := "something B"
	objA := interface{}(strA)
	objB := interface{}(strB)
	objIdA := Register(objA)
	objIdB := Register(objB)
	if objIdA != 0 {
		t.Error("id A is wrong")
	} else if objIdB != 1 {
		t.Error("id B is wrong")
	} else {
		objBackA := Obj(objIdA)
		objBackB := Obj(objIdB)
		if objBackA == objBackB {
			t.Error("object A and B are equal")
		} else {
			if strBackA, ok := objBackA.(string); ok {
				if strBackA != strA {
					t.Error("returned string A is wrong")
				}
			} else {
				t.Error("string A cast failed")
			}
			if strBackB, ok := objBackB.(string); ok {
				if strBackB != strB {
					t.Error("returned string B is wrong")
				}
			} else {
				t.Error("string B cast failed")
			}
			objBackA2 := Unregister(objIdA)
			if objBackA2 != objBackA {
				t.Error("released object A returned is wrong")
			} else {
				objIdA2 := Register(objA)
				if objIdA2 != objIdA {
					t.Error("id A2 is wrong")
				}
			}
		}
	}
}
