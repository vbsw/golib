//          Copyright 2022, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)

// Package cb maps objects to ids. This allows to manage callbacks from C to Go.
package cb

var (
	objs []interface{}
)

// Register returns a new id number for obj. obj will not be garbage collected until
// Unregister is called with this id.
func Register(obj interface{}) int {
	for i, o := range objs {
		if o == nil {
			objs[i] = obj
			return i
		}
	}
	objs = append(objs, obj)
	return len(objs) - 1
}

// Unregister makes the object no more identified by id.
// This object may be garbage collected, now.
func Unregister(id int) interface{} {
	obj := objs[id]
	objs[id] = nil
	return obj
}

// Obj returns object identified by id.
func Obj(id int) interface{} {
	return objs[id]
}
