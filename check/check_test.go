/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package check

import (
	"testing"
)

func TestBufferContainsAll(t *testing.T) {
	bytes := []byte("Hello, world!")
	terms := [][]byte{[]byte("llo"), []byte("He"), []byte(", w"), []byte("rld"), []byte("!")}
	termsCheck := make([][]byte, len(terms))
	copy(termsCheck, terms)
	hasAll := bufferContainsAll(bytes, termsCheck)
	if !hasAll {
		t.Error("not recognized")
	}
	terms[1] = []byte("he")
	copy(termsCheck, terms)
	hasAll = bufferContainsAll(bytes, termsCheck)
	if hasAll {
		t.Error("not recognized")
	}
	hasAll = bufferContainsAll(bytes[1:], termsCheck)
	copy(termsCheck, terms)
	if hasAll {
		t.Error("not recognized")
	}
	hasAll = bufferContainsAll(bytes[:0], termsCheck)
	if hasAll {
		t.Error("not recognized")
	}
}

func TestBufferContainsAny(t *testing.T) {
	bytes := []byte("Hello, world!")
	terms := [][]byte{[]byte("llo"), []byte("He"), []byte(", w"), []byte("rld"), []byte("!")}
	hasAny := bufferContainsAny(bytes, terms)
	if !hasAny {
		t.Error("not recognized")
	}
	terms[1] = []byte("he")
	hasAny = bufferContainsAny(bytes, terms)
	if !hasAny {
		t.Error("not recognized")
	}
	hasAny = bufferContainsAny(bytes[1:], terms)
	if !hasAny {
		t.Error("not recognized")
	}
	hasAny = bufferContainsAny(bytes[:0], terms)
	if hasAny {
		t.Error("not recognized")
	}
}
