/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package insert inserts data in slices of basic type.
package insert

import (
	"unsafe"
)

// Bool inserts value in list.
func Bool(list []bool, index int, value bool) []bool {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]bool, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// BoolD2 inserts value in list.
func BoolD2(list [][]bool, index int, value []bool) [][]bool {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]bool, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// BoolD3 inserts value in list.
func BoolD3(list [][][]bool, index int, value [][]bool) [][][]bool {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]bool, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// BoolD4 inserts value in list.
func BoolD4(list [][][][]bool, index int, value [][][]bool) [][][][]bool {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]bool, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// BoolD5 inserts value in list.
func BoolD5(list [][][][][]bool, index int, value [][][][]bool) [][][][][]bool {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]bool, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// BoolN inserts values in list.
func BoolN(list []bool, index int, values []bool) []bool {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]bool, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// BoolND2 inserts values in list.
func BoolND2(list [][]bool, index int, values [][]bool) [][]bool {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]bool, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// BoolND3 inserts values in list.
func BoolND3(list [][][]bool, index int, values [][][]bool) [][][]bool {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]bool, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// BoolND4 inserts values in list.
func BoolND4(list [][][][]bool, index int, values [][][][]bool) [][][][]bool {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]bool, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// BoolND5 inserts values in list.
func BoolND5(list [][][][][]bool, index int, values [][][][][]bool) [][][][][]bool {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]bool, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Byte inserts value in list.
func Byte(list []byte, index int, value byte) []byte {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]byte, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// ByteD2 inserts value in list.
func ByteD2(list [][]byte, index int, value []byte) [][]byte {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]byte, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// ByteD3 inserts value in list.
func ByteD3(list [][][]byte, index int, value [][]byte) [][][]byte {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]byte, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// ByteD4 inserts value in list.
func ByteD4(list [][][][]byte, index int, value [][][]byte) [][][][]byte {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]byte, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// ByteD5 inserts value in list.
func ByteD5(list [][][][][]byte, index int, value [][][][]byte) [][][][][]byte {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]byte, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// ByteN inserts values in list.
func ByteN(list []byte, index int, values []byte) []byte {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]byte, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// ByteND2 inserts values in list.
func ByteND2(list [][]byte, index int, values [][]byte) [][]byte {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]byte, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// ByteND3 inserts values in list.
func ByteND3(list [][][]byte, index int, values [][][]byte) [][][]byte {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]byte, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// ByteND4 inserts values in list.
func ByteND4(list [][][][]byte, index int, values [][][][]byte) [][][][]byte {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]byte, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// ByteND5 inserts values in list.
func ByteND5(list [][][][][]byte, index int, values [][][][][]byte) [][][][][]byte {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]byte, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Complex64 inserts value in list.
func Complex64(list []complex64, index int, value complex64) []complex64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]complex64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Complex64D2 inserts value in list.
func Complex64D2(list [][]complex64, index int, value []complex64) [][]complex64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]complex64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Complex64D3 inserts value in list.
func Complex64D3(list [][][]complex64, index int, value [][]complex64) [][][]complex64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]complex64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Complex64D4 inserts value in list.
func Complex64D4(list [][][][]complex64, index int, value [][][]complex64) [][][][]complex64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]complex64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Complex64D5 inserts value in list.
func Complex64D5(list [][][][][]complex64, index int, value [][][][]complex64) [][][][][]complex64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]complex64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Complex64N inserts values in list.
func Complex64N(list []complex64, index int, values []complex64) []complex64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]complex64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Complex64ND2 inserts values in list.
func Complex64ND2(list [][]complex64, index int, values [][]complex64) [][]complex64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]complex64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Complex64ND3 inserts values in list.
func Complex64ND3(list [][][]complex64, index int, values [][][]complex64) [][][]complex64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]complex64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Complex64ND4 inserts values in list.
func Complex64ND4(list [][][][]complex64, index int, values [][][][]complex64) [][][][]complex64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]complex64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Complex64ND5 inserts values in list.
func Complex64ND5(list [][][][][]complex64, index int, values [][][][][]complex64) [][][][][]complex64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]complex64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Complex128 inserts value in list.
func Complex128(list []complex128, index int, value complex128) []complex128 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]complex128, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Complex128D2 inserts value in list.
func Complex128D2(list [][]complex128, index int, value []complex128) [][]complex128 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]complex128, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Complex128D3 inserts value in list.
func Complex128D3(list [][][]complex128, index int, value [][]complex128) [][][]complex128 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]complex128, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Complex128D4 inserts value in list.
func Complex128D4(list [][][][]complex128, index int, value [][][]complex128) [][][][]complex128 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]complex128, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Complex128D5 inserts value in list.
func Complex128D5(list [][][][][]complex128, index int, value [][][][]complex128) [][][][][]complex128 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]complex128, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Complex128N inserts values in list.
func Complex128N(list []complex128, index int, values []complex128) []complex128 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]complex128, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Complex128ND2 inserts values in list.
func Complex128ND2(list [][]complex128, index int, values [][]complex128) [][]complex128 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]complex128, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Complex128ND3 inserts values in list.
func Complex128ND3(list [][][]complex128, index int, values [][][]complex128) [][][]complex128 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]complex128, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Complex128ND4 inserts values in list.
func Complex128ND4(list [][][][]complex128, index int, values [][][][]complex128) [][][][]complex128 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]complex128, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Complex128ND5 inserts values in list.
func Complex128ND5(list [][][][][]complex128, index int, values [][][][][]complex128) [][][][][]complex128 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]complex128, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Error inserts value in list.
func Error(list []error, index int, value error) []error {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]error, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// ErrorD2 inserts value in list.
func ErrorD2(list [][]error, index int, value []error) [][]error {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]error, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// ErrorD3 inserts value in list.
func ErrorD3(list [][][]error, index int, value [][]error) [][][]error {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]error, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// ErrorD4 inserts value in list.
func ErrorD4(list [][][][]error, index int, value [][][]error) [][][][]error {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]error, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// ErrorD5 inserts value in list.
func ErrorD5(list [][][][][]error, index int, value [][][][]error) [][][][][]error {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]error, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// ErrorN inserts values in list.
func ErrorN(list []error, index int, values []error) []error {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]error, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// ErrorND2 inserts values in list.
func ErrorND2(list [][]error, index int, values [][]error) [][]error {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]error, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// ErrorND3 inserts values in list.
func ErrorND3(list [][][]error, index int, values [][][]error) [][][]error {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]error, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// ErrorND4 inserts values in list.
func ErrorND4(list [][][][]error, index int, values [][][][]error) [][][][]error {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]error, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// ErrorND5 inserts values in list.
func ErrorND5(list [][][][][]error, index int, values [][][][][]error) [][][][][]error {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]error, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Float32 inserts value in list.
func Float32(list []float32, index int, value float32) []float32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]float32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Float32D2 inserts value in list.
func Float32D2(list [][]float32, index int, value []float32) [][]float32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]float32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Float32D3 inserts value in list.
func Float32D3(list [][][]float32, index int, value [][]float32) [][][]float32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]float32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Float32D4 inserts value in list.
func Float32D4(list [][][][]float32, index int, value [][][]float32) [][][][]float32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]float32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Float32D5 inserts value in list.
func Float32D5(list [][][][][]float32, index int, value [][][][]float32) [][][][][]float32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]float32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Float32N inserts values in list.
func Float32N(list []float32, index int, values []float32) []float32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]float32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Float32ND2 inserts values in list.
func Float32ND2(list [][]float32, index int, values [][]float32) [][]float32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]float32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Float32ND3 inserts values in list.
func Float32ND3(list [][][]float32, index int, values [][][]float32) [][][]float32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]float32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Float32ND4 inserts values in list.
func Float32ND4(list [][][][]float32, index int, values [][][][]float32) [][][][]float32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]float32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Float32ND5 inserts values in list.
func Float32ND5(list [][][][][]float32, index int, values [][][][][]float32) [][][][][]float32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]float32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Float64 inserts value in list.
func Float64(list []float64, index int, value float64) []float64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]float64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Float64D2 inserts value in list.
func Float64D2(list [][]float64, index int, value []float64) [][]float64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]float64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Float64D3 inserts value in list.
func Float64D3(list [][][]float64, index int, value [][]float64) [][][]float64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]float64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Float64D4 inserts value in list.
func Float64D4(list [][][][]float64, index int, value [][][]float64) [][][][]float64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]float64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Float64D5 inserts value in list.
func Float64D5(list [][][][][]float64, index int, value [][][][]float64) [][][][][]float64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]float64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Float64N inserts values in list.
func Float64N(list []float64, index int, values []float64) []float64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]float64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Float64ND2 inserts values in list.
func Float64ND2(list [][]float64, index int, values [][]float64) [][]float64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]float64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Float64ND3 inserts values in list.
func Float64ND3(list [][][]float64, index int, values [][][]float64) [][][]float64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]float64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Float64ND4 inserts values in list.
func Float64ND4(list [][][][]float64, index int, values [][][][]float64) [][][][]float64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]float64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Float64ND5 inserts values in list.
func Float64ND5(list [][][][][]float64, index int, values [][][][][]float64) [][][][][]float64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]float64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int inserts value in list.
func Int(list []int, index int, value int) []int {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]int, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// IntD2 inserts value in list.
func IntD2(list [][]int, index int, value []int) [][]int {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]int, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// IntD3 inserts value in list.
func IntD3(list [][][]int, index int, value [][]int) [][][]int {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]int, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// IntD4 inserts value in list.
func IntD4(list [][][][]int, index int, value [][][]int) [][][][]int {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]int, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// IntD5 inserts value in list.
func IntD5(list [][][][][]int, index int, value [][][][]int) [][][][][]int {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]int, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// IntN inserts values in list.
func IntN(list []int, index int, values []int) []int {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]int, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// IntND2 inserts values in list.
func IntND2(list [][]int, index int, values [][]int) [][]int {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]int, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// IntND3 inserts values in list.
func IntND3(list [][][]int, index int, values [][][]int) [][][]int {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]int, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// IntND4 inserts values in list.
func IntND4(list [][][][]int, index int, values [][][][]int) [][][][]int {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]int, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// IntND5 inserts values in list.
func IntND5(list [][][][][]int, index int, values [][][][][]int) [][][][][]int {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]int, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int8 inserts value in list.
func Int8(list []int8, index int, value int8) []int8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]int8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int8D2 inserts value in list.
func Int8D2(list [][]int8, index int, value []int8) [][]int8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]int8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int8D3 inserts value in list.
func Int8D3(list [][][]int8, index int, value [][]int8) [][][]int8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]int8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int8D4 inserts value in list.
func Int8D4(list [][][][]int8, index int, value [][][]int8) [][][][]int8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]int8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int8D5 inserts value in list.
func Int8D5(list [][][][][]int8, index int, value [][][][]int8) [][][][][]int8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]int8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int8N inserts values in list.
func Int8N(list []int8, index int, values []int8) []int8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]int8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int8ND2 inserts values in list.
func Int8ND2(list [][]int8, index int, values [][]int8) [][]int8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]int8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int8ND3 inserts values in list.
func Int8ND3(list [][][]int8, index int, values [][][]int8) [][][]int8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]int8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int8ND4 inserts values in list.
func Int8ND4(list [][][][]int8, index int, values [][][][]int8) [][][][]int8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]int8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int8ND5 inserts values in list.
func Int8ND5(list [][][][][]int8, index int, values [][][][][]int8) [][][][][]int8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]int8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int16 inserts value in list.
func Int16(list []int16, index int, value int16) []int16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]int16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int16D2 inserts value in list.
func Int16D2(list [][]int16, index int, value []int16) [][]int16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]int16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int16D3 inserts value in list.
func Int16D3(list [][][]int16, index int, value [][]int16) [][][]int16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]int16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int16D4 inserts value in list.
func Int16D4(list [][][][]int16, index int, value [][][]int16) [][][][]int16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]int16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int16D5 inserts value in list.
func Int16D5(list [][][][][]int16, index int, value [][][][]int16) [][][][][]int16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]int16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int16N inserts values in list.
func Int16N(list []int16, index int, values []int16) []int16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]int16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int16ND2 inserts values in list.
func Int16ND2(list [][]int16, index int, values [][]int16) [][]int16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]int16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int16ND3 inserts values in list.
func Int16ND3(list [][][]int16, index int, values [][][]int16) [][][]int16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]int16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int16ND4 inserts values in list.
func Int16ND4(list [][][][]int16, index int, values [][][][]int16) [][][][]int16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]int16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int16ND5 inserts values in list.
func Int16ND5(list [][][][][]int16, index int, values [][][][][]int16) [][][][][]int16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]int16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int32 inserts value in list.
func Int32(list []int32, index int, value int32) []int32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]int32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int32D2 inserts value in list.
func Int32D2(list [][]int32, index int, value []int32) [][]int32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]int32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int32D3 inserts value in list.
func Int32D3(list [][][]int32, index int, value [][]int32) [][][]int32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]int32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int32D4 inserts value in list.
func Int32D4(list [][][][]int32, index int, value [][][]int32) [][][][]int32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]int32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int32D5 inserts value in list.
func Int32D5(list [][][][][]int32, index int, value [][][][]int32) [][][][][]int32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]int32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int32N inserts values in list.
func Int32N(list []int32, index int, values []int32) []int32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]int32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int32ND2 inserts values in list.
func Int32ND2(list [][]int32, index int, values [][]int32) [][]int32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]int32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int32ND3 inserts values in list.
func Int32ND3(list [][][]int32, index int, values [][][]int32) [][][]int32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]int32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int32ND4 inserts values in list.
func Int32ND4(list [][][][]int32, index int, values [][][][]int32) [][][][]int32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]int32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int32ND5 inserts values in list.
func Int32ND5(list [][][][][]int32, index int, values [][][][][]int32) [][][][][]int32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]int32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int64 inserts value in list.
func Int64(list []int64, index int, value int64) []int64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]int64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int64D2 inserts value in list.
func Int64D2(list [][]int64, index int, value []int64) [][]int64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]int64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int64D3 inserts value in list.
func Int64D3(list [][][]int64, index int, value [][]int64) [][][]int64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]int64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int64D4 inserts value in list.
func Int64D4(list [][][][]int64, index int, value [][][]int64) [][][][]int64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]int64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int64D5 inserts value in list.
func Int64D5(list [][][][][]int64, index int, value [][][][]int64) [][][][][]int64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]int64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// Int64N inserts values in list.
func Int64N(list []int64, index int, values []int64) []int64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]int64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int64ND2 inserts values in list.
func Int64ND2(list [][]int64, index int, values [][]int64) [][]int64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]int64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int64ND3 inserts values in list.
func Int64ND3(list [][][]int64, index int, values [][][]int64) [][][]int64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]int64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int64ND4 inserts values in list.
func Int64ND4(list [][][][]int64, index int, values [][][][]int64) [][][][]int64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]int64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Int64ND5 inserts values in list.
func Int64ND5(list [][][][][]int64, index int, values [][][][][]int64) [][][][][]int64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]int64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Interface inserts value in list.
func Interface(list []interface{}, index int, value interface{}) []interface{} {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]interface{}, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// InterfaceD2 inserts value in list.
func InterfaceD2(list [][]interface{}, index int, value []interface{}) [][]interface{} {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]interface{}, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// InterfaceD3 inserts value in list.
func InterfaceD3(list [][][]interface{}, index int, value [][]interface{}) [][][]interface{} {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]interface{}, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// InterfaceD4 inserts value in list.
func InterfaceD4(list [][][][]interface{}, index int, value [][][]interface{}) [][][][]interface{} {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]interface{}, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// InterfaceD5 inserts value in list.
func InterfaceD5(list [][][][][]interface{}, index int, value [][][][]interface{}) [][][][][]interface{} {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]interface{}, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// InterfaceN inserts values in list.
func InterfaceN(list []interface{}, index int, values []interface{}) []interface{} {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]interface{}, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// InterfaceND2 inserts values in list.
func InterfaceND2(list [][]interface{}, index int, values [][]interface{}) [][]interface{} {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]interface{}, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// InterfaceND3 inserts values in list.
func InterfaceND3(list [][][]interface{}, index int, values [][][]interface{}) [][][]interface{} {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]interface{}, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// InterfaceND4 inserts values in list.
func InterfaceND4(list [][][][]interface{}, index int, values [][][][]interface{}) [][][][]interface{} {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]interface{}, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// InterfaceND5 inserts values in list.
func InterfaceND5(list [][][][][]interface{}, index int, values [][][][][]interface{}) [][][][][]interface{} {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]interface{}, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Pointer inserts value in list.
func Pointer(list []unsafe.Pointer, index int, value unsafe.Pointer) []unsafe.Pointer {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]unsafe.Pointer, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// PointerD2 inserts value in list.
func PointerD2(list [][]unsafe.Pointer, index int, value []unsafe.Pointer) [][]unsafe.Pointer {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]unsafe.Pointer, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// PointerD3 inserts value in list.
func PointerD3(list [][][]unsafe.Pointer, index int, value [][]unsafe.Pointer) [][][]unsafe.Pointer {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]unsafe.Pointer, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// PointerD4 inserts value in list.
func PointerD4(list [][][][]unsafe.Pointer, index int, value [][][]unsafe.Pointer) [][][][]unsafe.Pointer {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]unsafe.Pointer, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// PointerD5 inserts value in list.
func PointerD5(list [][][][][]unsafe.Pointer, index int, value [][][][]unsafe.Pointer) [][][][][]unsafe.Pointer {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]unsafe.Pointer, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// PointerN inserts values in list.
func PointerN(list []unsafe.Pointer, index int, values []unsafe.Pointer) []unsafe.Pointer {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]unsafe.Pointer, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// PointerND2 inserts values in list.
func PointerND2(list [][]unsafe.Pointer, index int, values [][]unsafe.Pointer) [][]unsafe.Pointer {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]unsafe.Pointer, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// PointerND3 inserts values in list.
func PointerND3(list [][][]unsafe.Pointer, index int, values [][][]unsafe.Pointer) [][][]unsafe.Pointer {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]unsafe.Pointer, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// PointerND4 inserts values in list.
func PointerND4(list [][][][]unsafe.Pointer, index int, values [][][][]unsafe.Pointer) [][][][]unsafe.Pointer {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]unsafe.Pointer, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// PointerND5 inserts values in list.
func PointerND5(list [][][][][]unsafe.Pointer, index int, values [][][][][]unsafe.Pointer) [][][][][]unsafe.Pointer {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]unsafe.Pointer, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// Rune inserts value in list.
func Rune(list []rune, index int, value rune) []rune {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]rune, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// RuneD2 inserts value in list.
func RuneD2(list [][]rune, index int, value []rune) [][]rune {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]rune, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// RuneD3 inserts value in list.
func RuneD3(list [][][]rune, index int, value [][]rune) [][][]rune {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]rune, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// RuneD4 inserts value in list.
func RuneD4(list [][][][]rune, index int, value [][][]rune) [][][][]rune {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]rune, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// RuneD5 inserts value in list.
func RuneD5(list [][][][][]rune, index int, value [][][][]rune) [][][][][]rune {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]rune, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// RuneN inserts values in list.
func RuneN(list []rune, index int, values []rune) []rune {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]rune, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// RuneND2 inserts values in list.
func RuneND2(list [][]rune, index int, values [][]rune) [][]rune {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]rune, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// RuneND3 inserts values in list.
func RuneND3(list [][][]rune, index int, values [][][]rune) [][][]rune {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]rune, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// RuneND4 inserts values in list.
func RuneND4(list [][][][]rune, index int, values [][][][]rune) [][][][]rune {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]rune, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// RuneND5 inserts values in list.
func RuneND5(list [][][][][]rune, index int, values [][][][][]rune) [][][][][]rune {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]rune, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// String inserts value in list.
func String(list []string, index int, value string) []string {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]string, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// StringD2 inserts value in list.
func StringD2(list [][]string, index int, value []string) [][]string {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]string, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// StringD3 inserts value in list.
func StringD3(list [][][]string, index int, value [][]string) [][][]string {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]string, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// StringD4 inserts value in list.
func StringD4(list [][][][]string, index int, value [][][]string) [][][][]string {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]string, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// StringD5 inserts value in list.
func StringD5(list [][][][][]string, index int, value [][][][]string) [][][][][]string {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]string, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// StringN inserts values in list.
func StringN(list []string, index int, values []string) []string {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]string, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// StringND2 inserts values in list.
func StringND2(list [][]string, index int, values [][]string) [][]string {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]string, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// StringND3 inserts values in list.
func StringND3(list [][][]string, index int, values [][][]string) [][][]string {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]string, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// StringND4 inserts values in list.
func StringND4(list [][][][]string, index int, values [][][][]string) [][][][]string {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]string, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// StringND5 inserts values in list.
func StringND5(list [][][][][]string, index int, values [][][][][]string) [][][][][]string {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]string, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt inserts value in list.
func UInt(list []uint, index int, value uint) []uint {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]uint, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UIntD2 inserts value in list.
func UIntD2(list [][]uint, index int, value []uint) [][]uint {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]uint, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UIntD3 inserts value in list.
func UIntD3(list [][][]uint, index int, value [][]uint) [][][]uint {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]uint, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UIntD4 inserts value in list.
func UIntD4(list [][][][]uint, index int, value [][][]uint) [][][][]uint {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]uint, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UIntD5 inserts value in list.
func UIntD5(list [][][][][]uint, index int, value [][][][]uint) [][][][][]uint {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]uint, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UIntN inserts values in list.
func UIntN(list []uint, index int, values []uint) []uint {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]uint, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UIntND2 inserts values in list.
func UIntND2(list [][]uint, index int, values [][]uint) [][]uint {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]uint, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UIntND3 inserts values in list.
func UIntND3(list [][][]uint, index int, values [][][]uint) [][][]uint {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]uint, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UIntND4 inserts values in list.
func UIntND4(list [][][][]uint, index int, values [][][][]uint) [][][][]uint {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]uint, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UIntND5 inserts values in list.
func UIntND5(list [][][][][]uint, index int, values [][][][][]uint) [][][][][]uint {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]uint, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt8 inserts value in list.
func UInt8(list []uint8, index int, value uint8) []uint8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]uint8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt8D2 inserts value in list.
func UInt8D2(list [][]uint8, index int, value []uint8) [][]uint8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]uint8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt8D3 inserts value in list.
func UInt8D3(list [][][]uint8, index int, value [][]uint8) [][][]uint8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]uint8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt8D4 inserts value in list.
func UInt8D4(list [][][][]uint8, index int, value [][][]uint8) [][][][]uint8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]uint8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt8D5 inserts value in list.
func UInt8D5(list [][][][][]uint8, index int, value [][][][]uint8) [][][][][]uint8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]uint8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt8N inserts values in list.
func UInt8N(list []uint8, index int, values []uint8) []uint8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]uint8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt8ND2 inserts values in list.
func UInt8ND2(list [][]uint8, index int, values [][]uint8) [][]uint8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]uint8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt8ND3 inserts values in list.
func UInt8ND3(list [][][]uint8, index int, values [][][]uint8) [][][]uint8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]uint8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt8ND4 inserts values in list.
func UInt8ND4(list [][][][]uint8, index int, values [][][][]uint8) [][][][]uint8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]uint8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt8ND5 inserts values in list.
func UInt8ND5(list [][][][][]uint8, index int, values [][][][][]uint8) [][][][][]uint8 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]uint8, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt16 inserts value in list.
func UInt16(list []uint16, index int, value uint16) []uint16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]uint16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt16D2 inserts value in list.
func UInt16D2(list [][]uint16, index int, value []uint16) [][]uint16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]uint16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt16D3 inserts value in list.
func UInt16D3(list [][][]uint16, index int, value [][]uint16) [][][]uint16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]uint16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt16D4 inserts value in list.
func UInt16D4(list [][][][]uint16, index int, value [][][]uint16) [][][][]uint16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]uint16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt16D5 inserts value in list.
func UInt16D5(list [][][][][]uint16, index int, value [][][][]uint16) [][][][][]uint16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]uint16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt16N inserts values in list.
func UInt16N(list []uint16, index int, values []uint16) []uint16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]uint16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt16ND2 inserts values in list.
func UInt16ND2(list [][]uint16, index int, values [][]uint16) [][]uint16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]uint16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt16ND3 inserts values in list.
func UInt16ND3(list [][][]uint16, index int, values [][][]uint16) [][][]uint16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]uint16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt16ND4 inserts values in list.
func UInt16ND4(list [][][][]uint16, index int, values [][][][]uint16) [][][][]uint16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]uint16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt16ND5 inserts values in list.
func UInt16ND5(list [][][][][]uint16, index int, values [][][][][]uint16) [][][][][]uint16 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]uint16, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt32 inserts value in list.
func UInt32(list []uint32, index int, value uint32) []uint32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]uint32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt32D2 inserts value in list.
func UInt32D2(list [][]uint32, index int, value []uint32) [][]uint32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]uint32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt32D3 inserts value in list.
func UInt32D3(list [][][]uint32, index int, value [][]uint32) [][][]uint32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]uint32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt32D4 inserts value in list.
func UInt32D4(list [][][][]uint32, index int, value [][][]uint32) [][][][]uint32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]uint32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt32D5 inserts value in list.
func UInt32D5(list [][][][][]uint32, index int, value [][][][]uint32) [][][][][]uint32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]uint32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt32N inserts values in list.
func UInt32N(list []uint32, index int, values []uint32) []uint32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]uint32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt32ND2 inserts values in list.
func UInt32ND2(list [][]uint32, index int, values [][]uint32) [][]uint32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]uint32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt32ND3 inserts values in list.
func UInt32ND3(list [][][]uint32, index int, values [][][]uint32) [][][]uint32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]uint32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt32ND4 inserts values in list.
func UInt32ND4(list [][][][]uint32, index int, values [][][][]uint32) [][][][]uint32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]uint32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt32ND5 inserts values in list.
func UInt32ND5(list [][][][][]uint32, index int, values [][][][][]uint32) [][][][][]uint32 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]uint32, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt64 inserts value in list.
func UInt64(list []uint64, index int, value uint64) []uint64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([]uint64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt64D2 inserts value in list.
func UInt64D2(list [][]uint64, index int, value []uint64) [][]uint64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][]uint64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt64D3 inserts value in list.
func UInt64D3(list [][][]uint64, index int, value [][]uint64) [][][]uint64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][]uint64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt64D4 inserts value in list.
func UInt64D4(list [][][][]uint64, index int, value [][][]uint64) [][][][]uint64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][]uint64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt64D5 inserts value in list.
func UInt64D5(list [][][][][]uint64, index int, value [][][][]uint64) [][][][][]uint64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + 1
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		} else {
			listExt = make([][][][][]uint64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+1:], list[index:])
			listExt[index] = value
		}
	} else {
		listExt = append(list, value)
	}
	return listExt
}

// UInt64N inserts values in list.
func UInt64N(list []uint64, index int, values []uint64) []uint64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([]uint64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt64ND2 inserts values in list.
func UInt64ND2(list [][]uint64, index int, values [][]uint64) [][]uint64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][]uint64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt64ND3 inserts values in list.
func UInt64ND3(list [][][]uint64, index int, values [][][]uint64) [][][]uint64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][]uint64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt64ND4 inserts values in list.
func UInt64ND4(list [][][][]uint64, index int, values [][][][]uint64) [][][][]uint64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][]uint64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}

// UInt64ND5 inserts values in list.
func UInt64ND5(list [][][][][]uint64, index int, values [][][][][]uint64) [][][][][]uint64 {
	listExt := list
	if index < len(list) {
		lengthNew := len(list) + len(values)
		if lengthNew <= cap(list) {
			listExt = list[:lengthNew]
			copy(listExt[index+len(values):], list[index:])
		} else {
			listExt = make([][][][][]uint64, lengthNew, lengthNew+lengthNew>>1)
			copy(listExt, list[:index])
			copy(listExt[index+len(values):], list[index:])
		}
		copy(listExt[index:], values)
	} else {
		listExt = append(list, values...)
	}
	return listExt
}
