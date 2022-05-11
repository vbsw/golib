/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package idxsearch provides indexed binary search for slices of basic type.
package idxsearch

import (
	"strings"
	"unsafe"
)

// Bool searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and sorted from false to true.
func Bool(list []bool, indices []int, element bool) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value != element {
			if element {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			return middle, true
		}
	}
	return left, false
}

// BoolDesc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and sorted from true to false.
func BoolDesc(list []bool, indices []int, element bool) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value != element {
			if element {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			return middle, true
		}
	}
	return left, false
}

// BoolRng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be sorted from false to true.
func BoolRng(list []bool, indices []int, element bool) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if element != value {
			if element {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			from := boolRngL(list, indices, element, left, middle-1)
			to := boolRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// BoolRngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be sorted from true to false.
func BoolRngDesc(list []bool, indices []int, element bool) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if element != value {
			if !element {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			from := boolRngDescL(list, indices, element, left, middle-1)
			to := boolRngDescR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Byte searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Byte(list []byte, indices []int, element byte) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// ByteDesc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func ByteDesc(list []byte, indices []int, element byte) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// ByteRng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func ByteRng(list []byte, indices []int, element byte) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := byteRngL(list, indices, element, left, middle-1)
			to := byteRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// ByteRngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func ByteRngDesc(list []byte, indices []int, element byte) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := byteRngL(list, indices, element, left, middle-1)
			to := byteRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex128 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Complex128(list []complex128, indices []int, element complex128) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex128Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Complex128Desc(list []complex128, indices []int, element complex128) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex128Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Complex128Rng(list []complex128, indices []int, element complex128) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			from := complex128RngL(list, indices, element, left, middle-1)
			to := complex128RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex128RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Complex128RngDesc(list []complex128, indices []int, element complex128) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			from := complex128RngL(list, indices, element, left, middle-1)
			to := complex128RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex64 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Complex64(list []complex64, indices []int, element complex64) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex64Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Complex64Desc(list []complex64, indices []int, element complex64) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Complex64Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Complex64Rng(list []complex64, indices []int, element complex64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal > valueReal {
			left = middle + 1
		} else if elementReal < valueReal {
			right = middle - 1
		} else if elementImag > valueImag {
			left = middle + 1
		} else if elementImag < valueImag {
			right = middle - 1
		} else {
			from := complex64RngL(list, indices, element, left, middle-1)
			to := complex64RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Complex64RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Complex64RngDesc(list []complex64, indices []int, element complex64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementReal := real(element)
	elementImag := imag(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		valueReal := real(value)
		valueImag := imag(value)
		if elementReal < valueReal {
			left = middle + 1
		} else if elementReal > valueReal {
			right = middle - 1
		} else if elementImag < valueImag {
			left = middle + 1
		} else if elementImag > valueImag {
			right = middle - 1
		} else {
			from := complex64RngL(list, indices, element, left, middle-1)
			to := complex64RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float32 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Float32(list []float32, indices []int, element float32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Float32Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Float32Desc(list []float32, indices []int, element float32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Float32Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Float32Rng(list []float32, indices []int, element float32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := float32RngL(list, indices, element, left, middle-1)
			to := float32RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float32RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Float32RngDesc(list []float32, indices []int, element float32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := float32RngL(list, indices, element, left, middle-1)
			to := float32RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float64 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Float64(list []float64, indices []int, element float64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Float64Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Float64Desc(list []float64, indices []int, element float64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Float64Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Float64Rng(list []float64, indices []int, element float64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := float64RngL(list, indices, element, left, middle-1)
			to := float64RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Float64RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Float64RngDesc(list []float64, indices []int, element float64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := float64RngL(list, indices, element, left, middle-1)
			to := float64RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Int(list []int, indices []int, element int) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// IntDesc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func IntDesc(list []int, indices []int, element int) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// IntRng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func IntRng(list []int, indices []int, element int) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := intRngL(list, indices, element, left, middle-1)
			to := intRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// IntRngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func IntRngDesc(list []int, indices []int, element int) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := intRngL(list, indices, element, left, middle-1)
			to := intRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int16 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Int16(list []int16, indices []int, element int16) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int16Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Int16Desc(list []int16, indices []int, element int16) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int16Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Int16Rng(list []int16, indices []int, element int16) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := int16RngL(list, indices, element, left, middle-1)
			to := int16RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int16RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Int16RngDesc(list []int16, indices []int, element int16) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := int16RngL(list, indices, element, left, middle-1)
			to := int16RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int32 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Int32(list []int32, indices []int, element int32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int32Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Int32Desc(list []int32, indices []int, element int32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int32Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Int32Rng(list []int32, indices []int, element int32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := int32RngL(list, indices, element, left, middle-1)
			to := int32RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int32RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Int32RngDesc(list []int32, indices []int, element int32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := int32RngL(list, indices, element, left, middle-1)
			to := int32RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int64 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Int64(list []int64, indices []int, element int64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int64Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Int64Desc(list []int64, indices []int, element int64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int64Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Int64Rng(list []int64, indices []int, element int64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := int64RngL(list, indices, element, left, middle-1)
			to := int64RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int64RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Int64RngDesc(list []int64, indices []int, element int64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := int64RngL(list, indices, element, left, middle-1)
			to := int64RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int8 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Int8(list []int8, indices []int, element int8) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int8Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func Int8Desc(list []int8, indices []int, element int8) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// Int8Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func Int8Rng(list []int8, indices []int, element int8) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := int8RngL(list, indices, element, left, middle-1)
			to := int8RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Int8RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func Int8RngDesc(list []int8, indices []int, element int8) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := int8RngL(list, indices, element, left, middle-1)
			to := int8RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Pointer searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Pointer(list []unsafe.Pointer, indices []int, element unsafe.Pointer) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementUIntPtr := uintptr(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		valueUIntPtr := uintptr(list[index])
		if elementUIntPtr > valueUIntPtr {
			left = middle + 1
		} else if elementUIntPtr < valueUIntPtr {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// PointerDesc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func PointerDesc(list []unsafe.Pointer, indices []int, element unsafe.Pointer) (int, bool) {
	left := 0
	right := len(indices) - 1
	elementUIntPtr := uintptr(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		valueUIntPtr := uintptr(list[index])
		if elementUIntPtr < valueUIntPtr {
			left = middle + 1
		} else if elementUIntPtr > valueUIntPtr {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// PointerRng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func PointerRng(list []unsafe.Pointer, indices []int, element unsafe.Pointer) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementUIntPtr := uintptr(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		valueUIntPtr := uintptr(list[index])
		if elementUIntPtr > valueUIntPtr {
			left = middle + 1
		} else if elementUIntPtr < valueUIntPtr {
			right = middle - 1
		} else {
			from := pointerRngL(list, indices, element, left, middle-1)
			to := pointerRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// PointerRngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func PointerRngDesc(list []unsafe.Pointer, indices []int, element unsafe.Pointer) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	elementUIntPtr := uintptr(element)
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		valueUIntPtr := uintptr(list[index])
		if elementUIntPtr < valueUIntPtr {
			left = middle + 1
		} else if elementUIntPtr > valueUIntPtr {
			right = middle - 1
		} else {
			from := pointerRngL(list, indices, element, left, middle-1)
			to := pointerRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// Rune searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func Rune(list []rune, indices []int, element rune) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// RuneDesc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func RuneDesc(list []rune, indices []int, element rune) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// RuneRng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func RuneRng(list []rune, indices []int, element rune) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := runeRngL(list, indices, element, left, middle-1)
			to := runeRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// RuneRngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func RuneRngDesc(list []rune, indices []int, element rune) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := runeRngL(list, indices, element, left, middle-1)
			to := runeRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// String searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func String(list []string, indices []int, element string) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// StringDesc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func StringDesc(list []string, indices []int, element string) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// StringOff searches elements in list starting at offset matching element, i.e.
// list[i][offset:offset+len(element)] == element. Returns start and end index,
// i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be unique and in ascending order
// regarding offset.
func StringOff(list []string, indices []int, element string, offset int) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset {
			value = value[offset:]
			if value < element {
				left = middle + 1
			} else if value > element {
				right = middle - 1
			} else {
				return middle, true
			}
		} else {
			left = middle + 1
		}
	}
	return left, false
}

// StringOffDesc searches elements in list starting at offset matching element, i.e.
// list[i][offset:offset+len(element)] == element. Returns start and end index,
// i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be unique and in descending order
// regarding offset.
func StringOffDesc(list []string, indices []int, element string, offset int) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset {
			value = value[offset:]
			if value > element {
				left = middle + 1
			} else if value < element {
				right = middle - 1
			} else {
				return middle, true
			}
		} else {
			right = middle - 1
		}
	}
	return left, false
}

// StringPfx searches element in list by prefix and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique by prefix and in ascending order.
func StringPfx(list []string, indices []int, element string) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else {
			if strings.HasPrefix(value, element) {
				return middle, true
			} else {
				right = middle - 1
			}
		}
	}
	return left, false
}

// StringPfxDesc searches element in list by prefix and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique by prefix and in descending order.
func StringPfxDesc(list []string, indices []int, element string) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			right = middle - 1
		} else {
			if strings.HasPrefix(value, element) {
				return middle, true
			} else {
				left = middle + 1
			}
		}
	}
	return left, false
}

// StringPfxOff searches elements in list by prefix starting at offset matching element, i.e.
// HasPrefix(list[i][offset:offset+len(element)], element). Returns start and end index,
// i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be unique by prefix and in ascending order,
// both regarding offset.
func StringPfxOff(list []string, indices []int, element string, offset int) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset {
			value = value[offset:]
			if value < element {
				left = middle + 1
			} else {
				if strings.HasPrefix(value, element) {
					return middle, true
				} else {
					right = middle - 1
				}
			}
		} else {
			left = middle + 1
		}
	}
	return left, false
}

// StringPfxOffDesc searches elements in list by prefix starting at offset matching element, i.e.
// HasPrefix(list[i][offset:offset+len(element)], element). Returns start and end index,
// i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be unique by prefix and in descending order,
// both regarding offset.
func StringPfxOffDesc(list []string, indices []int, element string, offset int) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset {
			value = value[offset:]
			if value < element {
				right = middle - 1
			} else {
				if strings.HasPrefix(value, element) {
					return middle, true
				} else {
					left = middle + 1
				}
			}
		} else {
			right = middle - 1
		}
	}
	return left, false
}

// StringPfxRngOff searches elements in list starting at offset matching element, i.e.
// HasPrefix(list[i][offset:offset+len(element)], element). Returns start and end index,
// i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be in ascending order
// regarding offset.
func StringPfxRngOff(list []string, indices []int, element string, offset int) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset {
			value = value[offset:]
			if value < element {
				left = middle + 1
			} else if strings.HasPrefix(value, element) {
				from := stringPfxOffL(list, indices, element, left, middle-1, offset)
				to := stringPfxOffR(list, indices, element, middle+1, right, offset)
				return from, to, true
			} else {
				right = middle - 1
			}
		} else {
			left = middle + 1
		}
	}
	return left, left + 1, false
}

// StringPfxRngOffDesc searches elements in list starting at offset matching element, i.e.
// HasPrefix(list[i][offset:offset+len(element)], element). Returns start and end index,
// i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be in descending order
// regarding offset.
func StringPfxRngOffDesc(list []string, indices []int, element string, offset int) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset {
			value = value[offset:]
			if value < element {
				right = middle - 1
			} else if strings.HasPrefix(value, element) {
				from := stringPfxOffL(list, indices, element, left, middle-1, offset)
				to := stringPfxOffR(list, indices, element, middle+1, right, offset)
				return from, to, true
			} else {
				left = middle + 1
			}
		} else {
			right = middle - 1
		}
	}
	return left, left + 1, false
}

// StringRng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func StringRng(list []string, indices []int, element string) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := stringRngL(list, indices, element, left, middle-1)
			to := stringRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// StringRngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func StringRngDesc(list []string, indices []int, element string) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := stringRngL(list, indices, element, left, middle-1)
			to := stringRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// StringRngOff searches elements in list starting at offset matching element, i.e.
// list[i][offset:offset+len(element)] == element. Returns start and end index,
// i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be in ascending order
// regarding offset.
func StringRngOff(list []string, indices []int, element string, offset int) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset {
			value = value[offset:]
			if value < element {
				left = middle + 1
			} else if value > element {
				right = middle - 1
			} else {
				from := stringOffL(list, indices, element, left, middle-1, offset)
				to := stringOffR(list, indices, element, middle+1, right, offset)
				return from, to, true
			}
		} else {
			left = middle + 1
		}
	}
	return left, left + 1, false
}

// StringRngOffDesc searches elements in list starting at offset matching element, i.e.
// list[i][offset:offset+len(element)] == element. Returns start and end index,
// i.e. start <= i < end. Start index is inclusive, end index is exclusive.
// Third return value is true, if element is in list, otherwise false and the index
// returned is the insert index. Elements in list must be in descending order
// regarding offset.
func StringRngOffDesc(list []string, indices []int, element string, offset int) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset {
			value = value[offset:]
			if value > element {
				left = middle + 1
			} else if value < element {
				right = middle - 1
			} else {
				from := stringOffL(list, indices, element, left, middle-1, offset)
				to := stringOffR(list, indices, element, middle+1, right, offset)
				return from, to, true
			}
		} else {
			right = middle - 1
		}
	}
	return left, left + 1, false
}

// UInt searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func UInt(list []uint, indices []int, element uint) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UIntDesc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UIntDesc(list []uint, indices []int, element uint) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UIntRng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UIntRng(list []uint, indices []int, element uint) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := uintRngL(list, indices, element, left, middle-1)
			to := uintRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UIntRngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UIntRngDesc(list []uint, indices []int, element uint) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := uintRngL(list, indices, element, left, middle-1)
			to := uintRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt16 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func UInt16(list []uint16, indices []int, element uint16) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt16Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UInt16Desc(list []uint16, indices []int, element uint16) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt16Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UInt16Rng(list []uint16, indices []int, element uint16) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := uint16RngL(list, indices, element, left, middle-1)
			to := uint16RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt16RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UInt16RngDesc(list []uint16, indices []int, element uint16) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := uint16RngL(list, indices, element, left, middle-1)
			to := uint16RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt32 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func UInt32(list []uint32, indices []int, element uint32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt32Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UInt32Desc(list []uint32, indices []int, element uint32) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt32Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UInt32Rng(list []uint32, indices []int, element uint32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := uint32RngL(list, indices, element, left, middle-1)
			to := uint32RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt32RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UInt32RngDesc(list []uint32, indices []int, element uint32) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := uint32RngL(list, indices, element, left, middle-1)
			to := uint32RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt64 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func UInt64(list []uint64, indices []int, element uint64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt64Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UInt64Desc(list []uint64, indices []int, element uint64) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt64Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UInt64Rng(list []uint64, indices []int, element uint64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := uint64RngL(list, indices, element, left, middle-1)
			to := uint64RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt64RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UInt64RngDesc(list []uint64, indices []int, element uint64) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := uint64RngL(list, indices, element, left, middle-1)
			to := uint64RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt8 searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func UInt8(list []uint8, indices []int, element uint8) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt8Desc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UInt8Desc(list []uint8, indices []int, element uint8) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UInt8Rng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UInt8Rng(list []uint8, indices []int, element uint8) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := uint8RngL(list, indices, element, left, middle-1)
			to := uint8RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UInt8RngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UInt8RngDesc(list []uint8, indices []int, element uint8) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := uint8RngL(list, indices, element, left, middle-1)
			to := uint8RngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UIntPtr searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func UIntPtr(list []uintptr, indices []int, element uintptr) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UIntPtrDesc searches element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in descending order.
func UIntPtrDesc(list []uintptr, indices []int, element uintptr) (int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// UIntPtrRng searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in ascending order.
func UIntPtrRng(list []uintptr, indices []int, element uintptr) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			from := uintptrRngL(list, indices, element, left, middle-1)
			to := uintptrRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

// UIntPtrRngDesc searches element in list and returns matching start and end index. Start index
// is inclusive, end index is exclusive. Third return value is true, if element is in
// list, otherwise false and the index returned is the insert index. Elements in list
// must be in descending order.
func UIntPtrRngDesc(list []uintptr, indices []int, element uintptr) (int, int, bool) {
	left := 0
	right := len(indices) - 1
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value > element {
			left = middle + 1
		} else if value < element {
			right = middle - 1
		} else {
			from := uintptrRngL(list, indices, element, left, middle-1)
			to := uintptrRngR(list, indices, element, middle+1, right)
			return from, to, true
		}
	}
	return left, left + 1, false
}

func boolRngL(list []bool, indices []int, element bool, left, right int) int {
	if element {
		for left <= right {
			middle := (left + right) / 2
			index := indices[middle]
			value := list[index]
			if value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return left
}

func boolRngR(list []bool, indices []int, element bool, left, right int) int {
	if !element {
		for left <= right {
			middle := (left + right) / 2
			index := indices[middle]
			value := list[index]
			if value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		return left
	}
	return right + 1
}

func boolRngDescL(list []bool, indices []int, element bool, left, right int) int {
	if !element {
		for left <= right {
			middle := (left + right) / 2
			index := indices[middle]
			value := list[index]
			if !value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return left
}

func boolRngDescR(list []bool, indices []int, element bool, left, right int) int {
	if element {
		for left <= right {
			middle := (left + right) / 2
			index := indices[middle]
			value := list[index]
			if !value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		return left
	}
	return right + 1
}

func byteRngL(list []byte, indices []int, element byte, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func byteRngR(list []byte, indices []int, element byte, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func complex128RngL(list []complex128, indices []int, element complex128, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func complex128RngR(list []complex128, indices []int, element complex128, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func complex64RngL(list []complex64, indices []int, element complex64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func complex64RngR(list []complex64, indices []int, element complex64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func float32RngL(list []float32, indices []int, element float32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func float32RngR(list []float32, indices []int, element float32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func float64RngL(list []float64, indices []int, element float64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func float64RngR(list []float64, indices []int, element float64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func intRngL(list []int, indices []int, element int, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func intRngR(list []int, indices []int, element int, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int16RngL(list []int16, indices []int, element int16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int16RngR(list []int16, indices []int, element int16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int32RngL(list []int32, indices []int, element int32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int32RngR(list []int32, indices []int, element int32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int64RngL(list []int64, indices []int, element int64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int64RngR(list []int64, indices []int, element int64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func int8RngL(list []int8, indices []int, element int8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func int8RngR(list []int8, indices []int, element int8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func pointerRngL(list []unsafe.Pointer, indices []int, element unsafe.Pointer, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func pointerRngR(list []unsafe.Pointer, indices []int, element unsafe.Pointer, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func runeRngL(list []rune, indices []int, element rune, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func runeRngR(list []rune, indices []int, element rune, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func stringRngL(list []string, indices []int, element string, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func stringRngR(list []string, indices []int, element string, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func stringPfxOffL(list []string, indices []int, element string, left, right, offset int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset && strings.HasPrefix(value[offset:], element) {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func stringPfxOffR(list []string, indices []int, element string, left, right, offset int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset && strings.HasPrefix(value[offset:], element) {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func stringOffL(list []string, indices []int, element string, left, right, offset int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset && value[offset:] == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func stringOffR(list []string, indices []int, element string, left, right, offset int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if len(value) >= offset && value[offset:] == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uintRngL(list []uint, indices []int, element uint, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uintRngR(list []uint, indices []int, element uint, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint16RngL(list []uint16, indices []int, element uint16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint16RngR(list []uint16, indices []int, element uint16, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint32RngL(list []uint32, indices []int, element uint32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint32RngR(list []uint32, indices []int, element uint32, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint64RngL(list []uint64, indices []int, element uint64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint64RngR(list []uint64, indices []int, element uint64, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uint8RngL(list []uint8, indices []int, element uint8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uint8RngR(list []uint8, indices []int, element uint8, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func uintptrRngL(list []uintptr, indices []int, element uintptr, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func uintptrRngR(list []uintptr, indices []int, element uintptr, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		index := indices[middle]
		value := list[index]
		if value == element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}
