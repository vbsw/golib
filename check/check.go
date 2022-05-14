/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package check performs checks on files and their content.
package check

import (
	"bytes"
	"io"
	"os"
)

// FileExists returns true, if file exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

// FileContainsAll returns true, if file exists and contains all terms. If terms
// empty, false is returned.
func FileContainsAll(path string, buffer []byte, terms [][]byte) (bool, error) {
	termsCheck, lengthMax := termsCheckMax(terms)
	if lengthMax > 0 {
		file, err := os.Open(path)
		if err == nil {
			defer file.Close()
			var nRead int
			buffer = ensureBuffer(buffer, lengthMax)
			nRead, err = file.Read(buffer)
			for err == nil {
				if nRead < len(buffer) {
					return bufferContainsAllFinal(buffer, termsCheck), nil
				} else {
					if bufferContainsAll(buffer, termsCheck) {
						return true, nil
					} else {
						nProcessed := len(buffer) - lengthMax
						copy(buffer, buffer[nProcessed:])
						nRead, err = file.Read(buffer[lengthMax:])
						nRead += lengthMax
					}
				}
			}
			if err == io.EOF {
				err = nil
			}
		}
		return false, err
	}
	return false, nil
}

// FileContainsAny returns true, if file exists and contains any of terms. If terms
// empty, false is returned.
func FileContainsAny(path string, buffer []byte, terms [][]byte) (bool, error) {
	termsCheck, lengthMax := termsCheckMax(terms)
	if lengthMax > 0 {
		file, err := os.Open(path)
		if err == nil {
			defer file.Close()
			var nRead int
			buffer = ensureBuffer(buffer, lengthMax)
			nRead, err = file.Read(buffer)
			for err == nil {
				if bufferContainsAny(buffer, termsCheck) {
					return true, nil
				} else if nRead == len(buffer) {
					nProcessed := len(buffer) - lengthMax
					copy(buffer, buffer[nProcessed:])
					nRead, err = file.Read(buffer[lengthMax:])
					nRead += lengthMax
				} else {
					return false, nil
				}
			}
			if err == io.EOF {
				err = nil
			}
		}
		return false, err
	}
	return false, nil
}

// bufferContainsAllFinal returns true, if buffer contains all terms. Does not check all terms.
func bufferContainsAllFinal(buffer []byte, terms [][]byte) bool {
	for _, term := range terms {
		if len(term) > 0 && !bytes.Contains(buffer, term) {
			return false
		}
	}
	return true
}

// bufferContainsAll returns true, if buffer contains all terms.
func bufferContainsAll(buffer []byte, terms [][]byte) bool {
	hasAll := true
	for i, term := range terms {
		if len(term) > 0 {
			if bytes.Contains(buffer, term) {
				terms[i] = nil
			} else {
				hasAll = false
			}
		}
	}
	return hasAll
}

// bufferContainsAny returns true, if buffer contains any of terms.
func bufferContainsAny(buffer []byte, terms [][]byte) bool {
	for _, term := range terms {
		if len(term) > 0 && bytes.Contains(buffer, term) {
			return true
		}
	}
	return false
}

func parseOptions(options string) (bool, bool) {
	var beFile, beDirectory bool
	for _, r := range options {
		if r == 'f' {
			beFile = true
		} else if r == 'd' {
			beDirectory = true
		}
	}
	// if both false, then both allowed
	if beFile == beDirectory {
		return true, true
	}
	return beFile, beDirectory
}

func termsCheckMax(terms [][]byte) ([][]byte, int) {
	var max int
	termsCheck := make([][]byte, 0, len(terms))
	for _, term := range terms {
		length := len(term)
		if length > 0 {
			termsCheck = append(termsCheck, term)
			if length > max {
				max = length
			}
		}
	}
	return termsCheck, max
}

func ensureBuffer(bytes []byte, lengthMin int) []byte {
	if len(bytes) > lengthMin {
		return bytes
	}
	length := 1024 * 1024 * 4
	for length < lengthMin {
		length = length * 2
	}
	return make([]byte, length)
}
