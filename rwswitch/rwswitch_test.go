/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package rwswitch

import (
	"testing"
)

func TestRWRR(t *testing.T) {
	sw := New()
	readCurr := sw.CurrRead()
	writeCurr := sw.CurrWrite()
	if readCurr != 2 || writeCurr != 0 {
		t.Error(readCurr, writeCurr)
	}
	readNext := sw.NextRead()
	if readCurr != readNext {
		t.Error("changed read index", readCurr, readNext)
	}
	writeNext := sw.NextWrite()
	if writeCurr == writeNext {
		t.Error("unchanged write index", writeCurr)
	}
	readNext = sw.NextRead()
	if readCurr == readNext {
		t.Error("unchanged read index", readCurr)
	}
	readCurr = readNext
	readNext = sw.NextRead()
	if readCurr != readNext {
		t.Error("changed read index", readCurr, readNext)
	}
}

func TestAllWR(t *testing.T) {
	sw := New()
	writeCurr := sw.CurrWrite()
	readCurr := sw.CurrRead()
	for i := 0; i < 14*4; i++ {
		writeNext := sw.NextWrite()
		readNext := sw.NextRead()
		if writeCurr == writeNext {
			t.Error(i, "unchanged write index", writeCurr)
		}
		if readCurr == readNext {
			t.Error(i, "unchanged read index", writeCurr)
		}
		writeCurr = writeNext
		readCurr = readNext
	}
}

func TestAllWWR(t *testing.T) {
	sw := New()
	for i := 0; i < 14*4; i++ {
		writeCurr := sw.NextWrite()
		writeNext := sw.NextWrite()
		readNext := sw.NextRead()
		if writeCurr == writeNext {
			t.Error(i, "unchanged write index", writeCurr)
		}
		// previous write must be next read
		if writeCurr != readNext {
			t.Error(i, "read index not written", writeCurr, readNext)
		}
		if writeNext == readNext {
			t.Error(i, "read/write index equal", readNext)
		}
	}
}

func TestAllRWR(t *testing.T) {
	sw := New()
	for i := 0; i < 14*4; i++ {
		readCurr := sw.NextRead()
		writeNext := sw.NextWrite()
		readNext := sw.NextRead()
		if readCurr == readNext {
			t.Error(i, "unchanged read index", readCurr)
		}
		if readCurr == writeNext {
			t.Error(i, "read/write index equal", readCurr)
		}
		if writeNext == readNext {
			t.Error(i, "read/write index equal", readNext)
		}
	}
}
