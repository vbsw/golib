/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package ref converts bytes to string or string to bytes without reallocating data.
package ref

import (
	"reflect"
	"unsafe"
)

// Bytes returns string as byte slice, but without copying the bytes.
// This is useless, actually. The returned slice has no capacity and
// can not be sub sliced, then using the index on string to access the
// bytes works as well.
func Bytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// String returns byte slice as string, but without copying the bytes.
func String(bytes []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	strHeader := reflect.StringHeader{bytesHeader.Data, bytesHeader.Len}
	return *(*string)(unsafe.Pointer(&strHeader))
}
