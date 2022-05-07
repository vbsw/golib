/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package callback maps objects to ids. This allows to manage callbacks from C to Go.
package callback

// Callback holds objects identified by ids.
type Callback struct {
	used   []interface{}
	unused []int
}

// New returns an instance of Callback.
func New() *Callback {
	return new(Callback)
}

// Register returns a new id number for obj. obj will not be garbage collected until
// Unregister is called with this id.
func (cb *Callback) Register(obj interface{}) int {
	if len(cb.unused) == 0 {
		cb.used = append(cb.used, obj)
		return len(cb.used) - 1
	}
	indexLast := len(cb.unused) - 1
	indexObj := cb.unused[indexLast]
	cb.unused = cb.unused[:indexLast]
	cb.used[indexObj] = obj
	return indexObj
}

// Unregister makes the object no more identified by id.
// This object may be garbage collected, now.
func (cb *Callback) Unregister(id int) interface{} {
	obj := cb.used[id]
	cb.used[id] = nil
	cb.unused = append(cb.unused, id)
	return obj
}

// Obj returns object identified by id.
func (cb *Callback) Obj(id int) interface{} {
	return cb.used[id]
}
