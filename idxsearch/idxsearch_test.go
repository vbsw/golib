/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package search

import (
	"testing"
)

func TestBool(t *testing.T) {
	var list []bool
	var indices []int
	index, found := Bool(list, indices, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, false)
	indices = append(indices, 0)
	index, found = Bool(list, indices, false)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Bool(list, indices, true)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	list = append(list, true)
	indices = append(indices, 1)
	index, found = Bool(list, indices, true)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	list = list[1:]
	indices = indices[:1]
	index, found = Bool(list, indices, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Bool(list, indices, true)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestBoolDesc(t *testing.T) {
	var list []bool
	var indices []int
	index, found := BoolDesc(list, indices, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, true)
	indices = append(indices, 0)
	index, found = BoolDesc(list, indices, true)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = BoolDesc(list, indices, false)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	list = append(list, false)
	indices = append(indices, 1)
	index, found = BoolDesc(list, indices, false)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	list = list[1:]
	indices = indices[:1]
	index, found = BoolDesc(list, indices, true)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = BoolDesc(list, indices, false)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestByteA(t *testing.T) {
	var list []byte
	var indices []int
	index, found := Byte(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 20, 30, 40)
	indices = append(indices, 0, 1, 2, 3)
	index, found = Byte(list, indices, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteB(t *testing.T) {
	var list []byte
	var indices []int
	index, found := Byte(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 20, 30, 40, 50)
	indices = append(indices, 0, 1, 2, 3, 4)
	index, found = Byte(list, indices, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, indices, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteDescA(t *testing.T) {
	var list []byte
	var indices []int
	index, found := ByteDesc(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 40, 30, 20, 10)
	indices = append(indices, 0, 1, 2, 3)
	index, found = ByteDesc(list, indices, 45)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 40)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 30)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 15)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteDescB(t *testing.T) {
	var list []byte
	var indices []int
	index, found := ByteDesc(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 40, 30, 20, 10, 0)
	indices = append(indices, 0, 1, 2, 3, 4)
	index, found = ByteDesc(list, indices, 45)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 40)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 30)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 15)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, indices, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestComplex128(t *testing.T) {
	var list []complex128
	var indices []int
	index, found := Complex128(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 10+5i, 20, 20+5i)
	indices = append(indices, 0, 1, 2, 3)
	index, found = Complex128(list, indices, 5+5i)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128(list, indices, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128(list, indices, 10+5i)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128(list, indices, 10+10i)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128(list, indices, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128(list, indices, 20+2i)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128(list, indices, 20+5i)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestComplex128Desc(t *testing.T) {
	var list []complex128
	var indices []int
	index, found := Complex128Desc(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 20+5i, 20, 10+5i, 10)
	indices = append(indices, 0, 1, 2, 3)
	index, found = Complex128Desc(list, indices, 20+10i)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, indices, 20+5i)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, indices, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, indices, 15+5i)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, indices, 10+5i)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, indices, 10+2i)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, indices, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestFloat32(t *testing.T) {
	var list []float32
	var indices []int
	index, found := Float32(list, indices, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 20, 30, 40)
	indices = append(indices, 0, 1, 2, 3)
	index, found = Float32(list, indices, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Float32(list, indices, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Float32(list, indices, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Float32(list, indices, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Float32(list, indices, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Float32(list, indices, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Float32(list, indices, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestBoolRngL(t *testing.T) {
	list := []bool{false}
	indices := []int{0}
	from := boolRngL(list, indices, false, 0, 0)
	if from != 0 {
		t.Error(from)
	}
	list = append(list, false, false, false, false, true, true)
	indices = append(indices, 1, 2, 3, 4, 5, 6)
	from = boolRngL(list, indices, false, 0, 3)
	if from != 0 {
		t.Error(from)
	}
	from = boolRngL(list, indices, true, 0, 5)
	if from != 5 {
		t.Error(from)
	}
}

func TestBoolRngR(t *testing.T) {
	list := []bool{false}
	indices := []int{0}
	from := boolRngR(list, indices, false, 0, 0)
	if from != 1 {
		t.Error(from)
	}
	list = append(list, false, false, false, false, true, true, true)
	indices = append(indices, 1, 2, 3, 4, 5, 6, 7)
	from = boolRngR(list, indices, false, 2, 6)
	if from != 5 {
		t.Error(from)
	}
	from = boolRngR(list, indices, true, 6, 6)
	if from != 7 {
		t.Error(from)
	}
	list = append(list, true)
	indices = append(indices, 8)
	from = boolRngR(list, indices, true, 6, 7)
	if from != 8 {
		t.Error(from)
	}
}

func TestBoolRng(t *testing.T) {
	list := []bool{false, false, false, false, false, true, true, true}
	indices := []int{0, 1, 2, 3, 4, 5, 6, 7}
	from, to, found := BoolRng(list, indices, false)
	if from != 0 || to != 5 || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRng(list, indices, true)
	if from != 5 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	list[5], list[6], list[7] = false, false, false
	from, to, found = BoolRng(list, indices, false)
	if from != 0 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRng(list, indices, true)
	if from != len(list) || to != len(list)+1 || found {
		t.Error(from, to, found)
	}
}

func TestBoolRngDesc(t *testing.T) {
	list := []bool{true, true, true, true, true, false, false, false}
	indices := []int{0, 1, 2, 3, 4, 5, 6, 7}
	from, to, found := BoolRngDesc(list, indices, true)
	if from != 0 || to != 5 || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRngDesc(list, indices, false)
	if from != 5 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	list[5], list[6], list[7] = true, true, true
	from, to, found = BoolRngDesc(list, indices, true)
	if from != 0 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRngDesc(list, indices, false)
	if from != len(list) || to != len(list)+1 || found {
		t.Error(from, to, found)
	}
}

func TestByteRngL(t *testing.T) {
	list := []byte{50}
	indices := []int{0}
	from := byteRngL(list, indices, 50, 0, 0)
	if from != 0 {
		t.Error(from)
	}
	list = append(list, 50, 60, 60, 60, 70, 70, 70, 80)
	indices = append(indices, 1, 2, 3, 4, 5, 6, 7, 8)
	from = byteRngL(list, indices, 60, 0, 3)
	if from != 2 {
		t.Error(from)
	}
	from = byteRngL(list, indices, 70, 2, 6)
	if from != 5 {
		t.Error(from)
	}
}

func TestByteRngR(t *testing.T) {
	list := []byte{50}
	indices := []int{0}
	from := byteRngR(list, indices, 50, 0, 0)
	if from != 1 {
		t.Error(from)
	}
	list = append(list, 50, 60, 60, 60, 70, 70, 70, 80)
	indices = append(indices, 1, 2, 3, 4, 5, 6, 7, 8)
	from = byteRngR(list, indices, 60, 2, 8)
	if from != 5 {
		t.Error(from)
	}
	from = byteRngR(list, indices, 70, 5, 8)
	if from != 8 {
		t.Error(from)
	}
}

func TestByteRng(t *testing.T) {
	list := []byte{50, 50, 54, 60, 60, 60, 60, 70, 70, 70, 80}
	indices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	from, to, found := ByteRng(list, indices, 50)
	if from != 0 || to != 2 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, indices, 54)
	if from != 2 || to != 3 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, indices, 60)
	if from != 3 || to != 7 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, indices, 70)
	if from != 7 || to != 10 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, indices, 80)
	if from != 10 || to != 11 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, indices, 100)
	if from != 11 || to != 12 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, indices, 55)
	if from != 3 || to != 4 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, indices, 65)
	if from != 7 || to != 8 || found {
		t.Error(from, to, found)
	}
}

func TestByteRngDesc(t *testing.T) {
	list := []byte{80, 80, 79, 70, 70, 70, 70, 60, 60, 60, 50}
	indices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	from, to, found := ByteRngDesc(list, indices, 80)
	if from != 0 || to != 2 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, indices, 79)
	if from != 2 || to != 3 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, indices, 70)
	if from != 3 || to != 7 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, indices, 60)
	if from != 7 || to != 10 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, indices, 50)
	if from != 10 || to != 11 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, indices, 40)
	if from != 11 || to != 12 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, indices, 75)
	if from != 3 || to != 4 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, indices, 65)
	if from != 7 || to != 8 || found {
		t.Error(from, to, found)
	}
}

func TestStringOff(t *testing.T) {
	var list []string
	var indices []int
	index, found := StringOff(list, indices, "asdf", 0)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = StringOff(list, indices, "asdf", 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, "", "aa", "ab", "dc", "dd")
	indices = append(indices, 0, 1, 2, 3, 4)
	index, found = StringOff(list, indices, "a", 0)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	index, found = StringOff(list, indices, "b", 0)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "a", "aaa", "aab", "adc", "add"
	index, found = StringOff(list, indices, "a", 1)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	index, found = StringOff(list, indices, "b", 1)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = StringOff(list, indices, "", 1)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestStringPfx(t *testing.T) {
	var list []string
	var indices []int
	index, found := StringPfx(list, indices, "asdf")
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, "", "aa", "ab", "dc", "dd")
	indices = append(indices, 0, 1, 2, 3, 4)
	index, found = StringPfx(list, indices, "a")
	if index != 1 && index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfx(list, indices, "d")
	if index != 3 && index != 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfx(list, indices, "b")
	if index != 3 || found != false {
		t.Error(index, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "a", "aaa", "aab", "adc", "add"
	index, found = StringPfx(list, indices, "a")
	if index < 0 || index > 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfx(list, indices, "aa")
	if index != 1 && index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfx(list, indices, "ad")
	if index != 3 && index != 4 || found != true {
		t.Error(index, found)
	}
}

func TestStringPfxDesc(t *testing.T) {
	var list []string
	var indices []int
	index, found := StringPfxDesc(list, indices, "asdf")
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, "dd", "dc", "ab", "aa", "")
	indices = append(indices, 0, 1, 2, 3, 4)
	index, found = StringPfxDesc(list, indices, "a")
	if index != 2 && index != 3 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxDesc(list, indices, "d")
	if index != 0 && index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxDesc(list, indices, "b")
	if index != 2 || found != false {
		t.Error(index, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "add", "adc", "aab", "aaa", "a"
	index, found = StringPfxDesc(list, indices, "a")
	if index < 0 || index > 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxDesc(list, indices, "aa")
	if index != 2 && index != 3 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxDesc(list, indices, "ad")
	if index != 0 && index != 1 || found != true {
		t.Error(index, found)
	}
}

func TestStringPfxOff(t *testing.T) {
	var list []string
	var indices []int
	index, found := StringPfxOff(list, indices, "asdf", 0)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = StringPfxOff(list, indices, "asdf", 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, "_", ".aa", "_ab", ".dc", "_dd")
	indices = append(indices, 0, 1, 2, 3, 4)
	index, found = StringPfxOff(list, indices, "a", 1)
	if index != 1 && index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOff(list, indices, "d", 1)
	if index != 3 && index != 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOff(list, indices, "b", 1)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	list[0], list[1], list[2], list[3], list[4] = ".a", "_aaa", ".aab", "_adc", ".add"
	index, found = StringPfxOff(list, indices, "a", 1)
	if index < 0 || index > 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOff(list, indices, "aa", 1)
	if index != 1 && index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOff(list, indices, "ad", 1)
	if index != 3 && index != 4 || found != true {
		t.Error(index, found)
	}
}

func TestStringPfxOffDesc(t *testing.T) {
	var list []string
	var indices []int
	index, found := StringPfxOffDesc(list, indices, "asdf", 0)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = StringPfxOffDesc(list, indices, "asdf", 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, "_dd", ".dc", "_ab", ".aa", "_")
	indices = append(indices, 0, 1, 2, 3, 4)
	index, found = StringPfxOffDesc(list, indices, "a", 1)
	if index != 2 && index != 3 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOffDesc(list, indices, "d", 1)
	if index != 0 && index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOffDesc(list, indices, "b", 1)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	list[0], list[1], list[2], list[3], list[4] = ".add", "_adc", ".aab", "_aaa", ".a"
	index, found = StringPfxOffDesc(list, indices, "a", 1)
	if index < 0 || index > 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOffDesc(list, indices, "aa", 1)
	if index != 2 && index != 3 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOffDesc(list, indices, "ad", 1)
	if index != 0 && index != 1 || found != true {
		t.Error(index, found)
	}
}

func TestStringRngOff(t *testing.T) {
	var list []string
	var indices []int
	from, to, found := StringRngOff(list, indices, "asdf", 0)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	from, to, found = StringRngOff(list, indices, "asdf", 10)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	list = append(list, "", "aa", "ab", "dc", "dd")
	indices = append(indices, 0, 1, 2, 3, 4)
	from, to, found = StringRngOff(list, indices, "a", 0)
	if from != 1 || to != 2 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = StringRngOff(list, indices, "a", 1)
	if from != 1 || to != 2 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = StringRngOff(list, indices, "b", 0)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "a", "aaa", "aab", "adc", "add"
	from, to, found = StringRngOff(list, indices, "a", 1)
	if from != 1 || to != 2 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = StringRngOff(list, indices, "b", 1)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = StringRngOff(list, indices, "", 1)
	if from != 0 || to != 1 || found != true {
		t.Error(from, to, found)
	}
}

func TestStringPfxRngOff(t *testing.T) {
	var list []string
	var indices []int
	from, to, found := StringPfxRngOff(list, indices, "asdf", 0)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	from, to, found = StringPfxRngOff(list, indices, "asdf", 10)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	list = append(list, "", "aa", "ab", "dc", "dd")
	indices = append(indices, 0, 1, 2, 3, 4)
	from, to, found = StringPfxRngOff(list, indices, "a", 0)
	if from != 1 || to != 3 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = StringPfxRngOff(list, indices, "b", 0)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "a", "aaa", "aab", "adc", "add"
	from, to, found = StringPfxRngOff(list, indices, "a", 1)
	if from != 1 || to != 3 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = StringPfxRngOff(list, indices, "b", 1)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = StringPfxRngOff(list, indices, "", 1)
	if from != 0 || to != len(list) || found != true {
		t.Error(from, to, found)
	}
}
