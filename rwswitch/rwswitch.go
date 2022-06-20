/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package rwswitch iterates over values 0, 1 and 2, differentiating between the read and write value. Read and write values are always distinct.
package rwswitch

type Switch struct {
	State  byte
	States [56]byte
}

func New() *Switch {
	sw := new(Switch)
	sw.Init()
	return sw
}

func (sw *Switch) Init() {
	sw.State = 0
	sw.States = [56]byte{0, 1, 2, 0, 10, 2, 2, 1, 3, 1, 2, 0, 3, 4, 1, 0, 6, 5, 1, 2, 0, 4, 1, 0, 6, 7, 0, 2, 13, 8, 0, 1, 9, 7, 0, 2, 9, 5, 1, 2, 10, 11, 0, 1, 9, 12, 0, 2, 13, 11, 0, 1, 13, 2, 2, 1}
}

// CurrRead returns current read index.
func (sw *Switch) CurrRead() int {
	indexCurr := sw.State * 4
	indexRead := sw.States[indexCurr+2]
	return int(indexRead)
}

// CurrWrite returns current write index.
func (sw *Switch) CurrWrite() int {
	indexCurr := sw.State * 4
	indexWrite := sw.States[indexCurr+3]
	return int(indexWrite)
}

// NextRead returns next read index.
func (sw *Switch) NextRead() int {
	indexCurr := sw.State * 4
	stateNext := sw.States[indexCurr+0]
	indexNext := stateNext * 4
	indexRead := sw.States[indexNext+2]
	sw.State = stateNext
	return int(indexRead)
}

// NextRead returns next write index.
func (sw *Switch) NextWrite() int {
	indexCurr := sw.State * 4
	stateNext := sw.States[indexCurr+1]
	indexNext := stateNext * 4
	indexWrite := sw.States[indexNext+3]
	sw.State = stateNext
	return int(indexWrite)
}
