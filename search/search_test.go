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
	index, found := Bool(list, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, false)
	index, found = Bool(list, false)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Bool(list, true)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	list = append(list, true)
	index, found = Bool(list, true)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	list = list[1:]
	index, found = Bool(list, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Bool(list, true)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestBoolDesc(t *testing.T) {
	var list []bool
	index, found := BoolDesc(list, false)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, true)
	index, found = BoolDesc(list, true)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = BoolDesc(list, false)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	list = append(list, false)
	index, found = BoolDesc(list, false)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	list = list[1:]
	index, found = BoolDesc(list, true)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = BoolDesc(list, false)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestByteA(t *testing.T) {
	var list []byte
	index, found := Byte(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 20, 30, 40)
	index, found = Byte(list, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteB(t *testing.T) {
	var list []byte
	index, found := Byte(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 20, 30, 40, 50)
	index, found = Byte(list, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Byte(list, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Byte(list, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteDescA(t *testing.T) {
	var list []byte
	index, found := ByteDesc(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 40, 30, 20, 10)
	index, found = ByteDesc(list, 45)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 40)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 30)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 15)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestByteDescB(t *testing.T) {
	var list []byte
	index, found := ByteDesc(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 40, 30, 20, 10, 0)
	index, found = ByteDesc(list, 45)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 40)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 30)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 15)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = ByteDesc(list, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestComplex128(t *testing.T) {
	var list []complex128
	index, found := Complex128(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 10+5i, 20, 20+5i)
	index, found = Complex128(list, 5+5i)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128(list, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128(list, 10+5i)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128(list, 10+10i)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128(list, 20)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128(list, 20+2i)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128(list, 20+5i)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestComplex128Desc(t *testing.T) {
	var list []complex128
	index, found := Complex128Desc(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 20+5i, 20, 10+5i, 10)
	index, found = Complex128Desc(list, 20+10i)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 20+5i)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 15+5i)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 10+5i)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 10+2i)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Complex128Desc(list, 10)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestFloat32(t *testing.T) {
	var list []float32
	index, found := Float32(list, 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, 10, 20, 30, 40)
	index, found = Float32(list, 5)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = Float32(list, 10)
	if index != 0 || found != true {
		t.Error(index, found)
	}
	index, found = Float32(list, 20)
	if index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = Float32(list, 25)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	index, found = Float32(list, 30)
	if index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = Float32(list, 35)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = Float32(list, 40)
	if index != 3 || found != true {
		t.Error(index, found)
	}
}

func TestBoolRngL(t *testing.T) {
	list := []bool{false}
	from := boolRngL(list, false, 0, 0)
	if from != 0 {
		t.Error(from)
	}
	list = append(list, false, false, false, false, true, true)
	from = boolRngL(list, false, 0, 3)
	if from != 0 {
		t.Error(from)
	}
	from = boolRngL(list, true, 0, 5)
	if from != 5 {
		t.Error(from)
	}
}

func TestBoolRngR(t *testing.T) {
	list := []bool{false}
	from := boolRngR(list, false, 0, 0)
	if from != 1 {
		t.Error(from)
	}
	list = append(list, false, false, false, false, true, true, true)
	from = boolRngR(list, false, 2, 6)
	if from != 5 {
		t.Error(from)
	}
	from = boolRngR(list, true, 6, 6)
	if from != 7 {
		t.Error(from)
	}
	list = append(list, true)
	from = boolRngR(list, true, 6, 7)
	if from != 8 {
		t.Error(from)
	}
}

func TestBoolRng(t *testing.T) {
	list := []bool{false, false, false, false, false, true, true, true}
	from, to, found := BoolRng(list, false)
	if from != 0 || to != 5 || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRng(list, true)
	if from != 5 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	list[5], list[6], list[7] = false, false, false
	from, to, found = BoolRng(list, false)
	if from != 0 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRng(list, true)
	if from != len(list) || to != len(list)+1 || found {
		t.Error(from, to, found)
	}
}

func TestBoolRngDesc(t *testing.T) {
	list := []bool{true, true, true, true, true, false, false, false}
	from, to, found := BoolRngDesc(list, true)
	if from != 0 || to != 5 || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRngDesc(list, false)
	if from != 5 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	list[5], list[6], list[7] = true, true, true
	from, to, found = BoolRngDesc(list, true)
	if from != 0 || to != len(list) || !found {
		t.Error(from, to, found)
	}
	from, to, found = BoolRngDesc(list, false)
	if from != len(list) || to != len(list)+1 || found {
		t.Error(from, to, found)
	}
}

func TestByteRngL(t *testing.T) {
	list := []byte{50}
	from := byteRngL(list, 50, 0, 0)
	if from != 0 {
		t.Error(from)
	}
	list = append(list, 50, 60, 60, 60, 70, 70, 70, 80)
	from = byteRngL(list, 60, 0, 3)
	if from != 2 {
		t.Error(from)
	}
	from = byteRngL(list, 70, 2, 6)
	if from != 5 {
		t.Error(from)
	}
}

func TestByteRngR(t *testing.T) {
	list := []byte{50}
	from := byteRngR(list, 50, 0, 0)
	if from != 1 {
		t.Error(from)
	}
	list = append(list, 50, 60, 60, 60, 70, 70, 70, 80)
	from = byteRngR(list, 60, 2, 8)
	if from != 5 {
		t.Error(from)
	}
	from = byteRngR(list, 70, 5, 8)
	if from != 8 {
		t.Error(from)
	}
}

func TestByteRng(t *testing.T) {
	list := []byte{50, 50, 54, 60, 60, 60, 60, 70, 70, 70, 80}
	from, to, found := ByteRng(list, 50)
	if from != 0 || to != 2 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 54)
	if from != 2 || to != 3 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 60)
	if from != 3 || to != 7 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 70)
	if from != 7 || to != 10 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 80)
	if from != 10 || to != 11 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 100)
	if from != 11 || to != 12 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 55)
	if from != 3 || to != 4 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRng(list, 65)
	if from != 7 || to != 8 || found {
		t.Error(from, to, found)
	}
}

func TestByteRngDesc(t *testing.T) {
	list := []byte{80, 80, 79, 70, 70, 70, 70, 60, 60, 60, 50}
	from, to, found := ByteRngDesc(list, 80)
	if from != 0 || to != 2 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 79)
	if from != 2 || to != 3 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 70)
	if from != 3 || to != 7 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 60)
	if from != 7 || to != 10 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 50)
	if from != 10 || to != 11 || !found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 40)
	if from != 11 || to != 12 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 75)
	if from != 3 || to != 4 || found {
		t.Error(from, to, found)
	}
	from, to, found = ByteRngDesc(list, 65)
	if from != 7 || to != 8 || found {
		t.Error(from, to, found)
	}
}

func TestStringOff(t *testing.T) {
	var list []string
	index, found := StringOff(list, "asdf", 0)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = StringOff(list, "asdf", 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, "", "aa", "ab", "dc", "dd")
	index, found = StringOff(list, "a", 0)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	index, found = StringOff(list, "b", 0)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "a", "aaa", "aab", "adc", "add"
	index, found = StringOff(list, "a", 1)
	if index != 1 || found != false {
		t.Error(index, found)
	}
	index, found = StringOff(list, "b", 1)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	index, found = StringOff(list, "", 1)
	if index != 0 || found != true {
		t.Error(index, found)
	}
}

func TestStringPfx(t *testing.T) {
	var list []string
	index, found := StringPfx(list, "asdf")
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, "", "aa", "ab", "dc", "dd")
	index, found = StringPfx(list, "a")
	if index != 1 && index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfx(list, "d")
	if index != 3 && index != 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfx(list, "b")
	if index != 3 || found != false {
		t.Error(index, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "a", "aaa", "aab", "adc", "add"
	index, found = StringPfx(list, "a")
	if index < 0 || index > 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfx(list, "aa")
	if index != 1 && index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfx(list, "ad")
	if index != 3 && index != 4 || found != true {
		t.Error(index, found)
	}
}

func TestStringPfxDesc(t *testing.T) {
	var list []string
	index, found := StringPfxDesc(list, "asdf")
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, "dd", "dc", "ab", "aa", "")
	index, found = StringPfxDesc(list, "a")
	if index != 2 && index != 3 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxDesc(list, "d")
	if index != 0 && index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxDesc(list, "b")
	if index != 2 || found != false {
		t.Error(index, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "add", "adc", "aab", "aaa", "a"
	index, found = StringPfxDesc(list, "a")
	if index < 0 || index > 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxDesc(list, "aa")
	if index != 2 && index != 3 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxDesc(list, "ad")
	if index != 0 && index != 1 || found != true {
		t.Error(index, found)
	}
}

func TestStringPfxOff(t *testing.T) {
	var list []string
	index, found := StringPfxOff(list, "asdf", 0)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = StringPfxOff(list, "asdf", 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, "_", ".aa", "_ab", ".dc", "_dd")
	index, found = StringPfxOff(list, "a", 1)
	if index != 1 && index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOff(list, "d", 1)
	if index != 3 && index != 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOff(list, "b", 1)
	if index != 3 || found != false {
		t.Error(index, found)
	}
	list[0], list[1], list[2], list[3], list[4] = ".a", "_aaa", ".aab", "_adc", ".add"
	index, found = StringPfxOff(list, "a", 1)
	if index < 0 || index > 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOff(list, "aa", 1)
	if index != 1 && index != 2 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOff(list, "ad", 1)
	if index != 3 && index != 4 || found != true {
		t.Error(index, found)
	}
}

func TestStringPfxOffDesc(t *testing.T) {
	var list []string
	index, found := StringPfxOffDesc(list, "asdf", 0)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	index, found = StringPfxOffDesc(list, "asdf", 10)
	if index != 0 || found != false {
		t.Error(index, found)
	}
	list = append(list, "_dd", ".dc", "_ab", ".aa", "_")
	index, found = StringPfxOffDesc(list, "a", 1)
	if index != 2 && index != 3 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOffDesc(list, "d", 1)
	if index != 0 && index != 1 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOffDesc(list, "b", 1)
	if index != 2 || found != false {
		t.Error(index, found)
	}
	list[0], list[1], list[2], list[3], list[4] = ".add", "_adc", ".aab", "_aaa", ".a"
	index, found = StringPfxOffDesc(list, "a", 1)
	if index < 0 || index > 4 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOffDesc(list, "aa", 1)
	if index != 2 && index != 3 || found != true {
		t.Error(index, found)
	}
	index, found = StringPfxOffDesc(list, "ad", 1)
	if index != 0 && index != 1 || found != true {
		t.Error(index, found)
	}
}

func TestStringRngOff(t *testing.T) {
	var list []string
	from, to, found := StringRngOff(list, "asdf", 0)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	from, to, found = StringRngOff(list, "asdf", 10)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	list = append(list, "", "aa", "ab", "dc", "dd")
	from, to, found = StringRngOff(list, "a", 0)
	if from != 1 || to != 2 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = StringRngOff(list, "a", 1)
	if from != 1 || to != 2 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = StringRngOff(list, "b", 0)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "a", "aaa", "aab", "adc", "add"
	from, to, found = StringRngOff(list, "a", 1)
	if from != 1 || to != 2 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = StringRngOff(list, "b", 1)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = StringRngOff(list, "", 1)
	if from != 0 || to != 1 || found != true {
		t.Error(from, to, found)
	}
}

func TestStringPfxRngOff(t *testing.T) {
	var list []string
	from, to, found := StringPfxRngOff(list, "asdf", 0)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	from, to, found = StringPfxRngOff(list, "asdf", 10)
	if from != 0 || to != 1 || found != false {
		t.Error(from, found)
	}
	list = append(list, "", "aa", "ab", "dc", "dd")
	from, to, found = StringPfxRngOff(list, "a", 0)
	if from != 1 || to != 3 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = StringPfxRngOff(list, "b", 0)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	list[0], list[1], list[2], list[3], list[4] = "a", "aaa", "aab", "adc", "add"
	from, to, found = StringPfxRngOff(list, "a", 1)
	if from != 1 || to != 3 || found != true {
		t.Error(from, to, found)
	}
	from, to, found = StringPfxRngOff(list, "b", 1)
	if from != 3 || to != 4 || found != false {
		t.Error(from, to, found)
	}
	from, to, found = StringPfxRngOff(list, "", 1)
	if from != 0 || to != len(list) || found != true {
		t.Error(from, to, found)
	}
}
