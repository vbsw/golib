/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package queue

import "testing"

func TestPutEmpty(t *testing.T) {
	q := New(0)
	q.Put(10)
	if len(q.data) == 0 {
		t.Error("queue not initialized")
	} else if q.indexRead == q.indexWrite {
		t.Error(len(q.data), q.indexRead, q.indexWrite)
	} else if q.Size() != 1 {
		t.Error(q.Size())
	} else {
		q.Put(20)
		if len(q.data) == 0 {
			t.Error("queue not initialized")
		} else if q.Size() != 2 {
			t.Error(q.Size())
		} else if q.indexRead == q.indexWrite {
			t.Error(len(q.data), q.indexRead, q.indexWrite)
		}
	}
}

func TestPutCap(t *testing.T) {
	q := New(1)
	q.Put(10)
	if len(q.data) != 1 {
		t.Error(len(q.data))
	} else if q.indexRead != q.indexWrite || q.indexRead != 0 {
		t.Error(len(q.data), q.indexRead, q.indexWrite)
	} else if !q.hasAny {
		t.Error("queue is marked empty")
	} else {
		q.Put(20)
		if len(q.data) == 1 {
			t.Error("queue not expanded")
		} else if q.indexRead != q.indexWrite || q.indexRead != 0 {
			t.Error(len(q.data), q.indexRead, q.indexWrite)
		}
	}
}

func TestResetA(t *testing.T) {
	q := New(0)
	q.Put(10)
	if len(q.data) == 0 {
		t.Error("queue not initialized")
	} else if q.indexRead == q.indexWrite {
		t.Error(len(q.data), q.indexRead, q.indexWrite)
	} else if q.Size() != 1 {
		t.Error(q.Size())
	} else {
		q.Put(20)
		q.Reset(0)
		if cap(q.data) != 0 {
			t.Error(cap(q.data))
		} else if q.Size() != 0 {
			t.Error(q.Size())
		} else if q.indexRead != q.indexWrite {
			t.Error(len(q.data), q.indexRead, q.indexWrite)
		} else if q.indexRead != 0 {
			t.Error(q.indexRead)
		}
	}
}

func TestResetB(t *testing.T) {
	q := New(0)
	q.Put(10)
	if len(q.data) == 0 {
		t.Error("queue not initialized")
	} else if q.indexRead == q.indexWrite {
		t.Error(len(q.data), q.indexRead, q.indexWrite)
	} else if q.Size() != 1 {
		t.Error(q.Size())
	} else {
		capPrev := cap(q.data)
		q.Put(20)
		q.Reset(capPrev)
		if cap(q.data) != capPrev {
			t.Error(cap(q.data), capPrev)
		} else if q.Size() != 0 {
			t.Error(q.Size())
		} else if q.indexRead != q.indexWrite {
			t.Error(len(q.data), q.indexRead, q.indexWrite)
		} else if q.indexRead != 0 {
			t.Error(q.indexRead)
		} else if q.data[0] != nil {
			t.Error("elements not cleared")
		}
	}
}

func TestSize(t *testing.T) {
	q := New(0)
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	q.Put(10)
	if q.Size() != 1 {
		t.Error(q.Size())
	}
	q.Put(20)
	if q.Size() != 2 {
		t.Error(q.Size())
	}
	q.First()
	if q.Size() != 1 {
		t.Error(q.Size())
	}
	q.First()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	q.First()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
}

func TestPutTwo(t *testing.T) {
	q := New(10)
	q.indexRead = 7
	q.indexWrite = 8
	q.hasAny = true
	q.Put(10)
	if q.indexRead != 7 {
		t.Error(q.indexRead)
	}
	if q.indexWrite != 9 {
		t.Error(q.indexWrite)
	}
	q.Put(20)
	if q.indexRead != 7 {
		t.Error(q.indexRead)
	}
	if q.indexWrite != 0 {
		t.Error(q.indexWrite)
	}
	q.Put(30)
	if q.indexRead != 7 {
		t.Error(q.indexRead)
	}
	if q.indexWrite != 1 {
		t.Error(q.indexWrite)
	}
}

func TestPutTwoSize(t *testing.T) {
	q := New(10)
	q.indexRead = 7
	q.indexWrite = 8
	q.hasAny = true
	if q.Size() != 1 {
		t.Error(q.Size())
	}
	q.Put(10)
	if q.Size() != 2 {
		t.Error(q.Size())
	}
	q.Put(20)
	if q.Size() != 3 {
		t.Error(q.Size())
	}
	q.First()
	if q.Size() != 2 {
		t.Error(q.Size())
	}
	q.First()
	if q.Size() != 1 {
		t.Error(q.Size())
	}
	q.First()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
}

func TestFirst(t *testing.T) {
	q := New(10)
	q.indexRead = 7
	q.indexWrite = 8
	q.hasAny = true
	if q.Size() != 1 {
		t.Error(q.Size())
	}
	q.Put(10)
	if q.Size() != 2 {
		t.Error(q.Size())
	}
	q.Put(20)
	if q.Size() != 3 {
		t.Error(q.Size())
	}
	q.First()
	if q.Size() != 2 {
		t.Error(q.Size())
	}
	q.First()
	if q.Size() != 1 {
		t.Error(q.Size())
	}
	q.First()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	q.First()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
}

func TestFirstTwo(t *testing.T) {
	q := New(10)
	q.indexRead = 8
	q.indexWrite = 8
	el := q.First()
	if el != nil {
		t.Error("element not nil")
	}
	arr := []int{10, 20, 30, 40, 50}
	q.indexRead = 7
	q.hasAny = true
	for _, num := range arr {
		q.Put(num)
	}
	q.First()
	for i, num := range arr {
		el = q.First()
		if el == nil {
			t.Error("in", i, "element is nil")
		} else if num2, ok := el.(int); ok && num2 != num {
			t.Error("in", i, "element is not", num, "but", el)
		}
	}
	el = q.First()
	if el != nil {
		t.Error("element is not nil")
	}
}

func TestPutAll(t *testing.T) {
	q := New(10)
	q.indexRead = 8
	q.indexWrite = 8
	q.PutAll(10, 20, 30)
	if q.indexRead != 0 {
		t.Error(q.indexRead)
	}
	if q.indexWrite != 3 {
		t.Error(q.indexWrite)
	}
	if q.Size() != 3 {
		t.Error(q.Size())
	}
}

func TestPutAllTwo(t *testing.T) {
	q := New(10)
	q.indexRead = 7
	q.indexWrite = 8
	q.hasAny = true
	q.Put(10)
	q.Put(20)
	q.Put(30)
	if q.indexRead != 7 {
		t.Error(q.indexRead)
	}
	if q.indexWrite != 1 {
		t.Error(q.indexWrite)
	}
	if q.Size() != 4 {
		t.Error(q.Size())
	}
	q.PutAll(40, 50, 60)
	if q.indexRead != 7 {
		t.Error(q.indexRead)
	}
	if q.indexWrite != 4 {
		t.Error(q.indexWrite)
	}
	if q.Size() != 7 {
		t.Error(q.Size())
	}
}

func TestAll(t *testing.T) {
	arr := []int{10, 20, 30}
	q := New(10)
	q.indexRead = 7
	q.indexWrite = 8
	q.hasAny = true
	q.PutAll(10, 20, 30)
	elements := q.All()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	if len(elements) != len(arr)+1 {
		t.Error(len(elements))
	} else {
		for i, num := range arr {
			el := elements[i+1]
			if el == nil {
				t.Error("in", i, "element is nil")
			} else if num2, ok := el.(int); ok && num2 != num {
				t.Error("in", i, "element is not", num, "but", el)
			}
		}
	}
}

func TestIndexInsertion(t *testing.T) {
	q := New(10)
	q.PutAll(10, 20, 30)
	q.prios[0] = 2
	q.prios[1] = 4
	q.prios[2] = 6
	index := q.indexInsertion(0)
	if index != 0 {
		t.Error(index)
	}
	index = q.indexInsertion(6)
	if index != 3 {
		t.Error(index)
	}
	index = q.indexInsertion(5)
	if index != 2 {
		t.Error(index)
	}
	index = q.indexInsertion(PrioMin)
	if index != 3 {
		t.Error(index)
	}
}

func TestPutPrioPrd(t *testing.T) {
	arr := []int{40, 10, 20, 30}
	q := New(len(arr))
	q.PutAll(10, 20, 30)
	q.PutPrio(2, 40)
	if q.Size() != len(arr) {
		t.Error(q.Size())
	}
	elements := q.All()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	if len(elements) != len(arr) {
		t.Error(len(elements))
	} else {
		for i, num := range arr {
			el := elements[i]
			if el == nil {
				t.Error("in", i, "element is nil")
			} else if num2, ok := el.(int); ok && num2 != num {
				t.Error("in", i, "element is not", num, "but", el)
			}
		}
	}
	arr = []int{50, 40, 10, 20, 30}
	q = New(len(arr) - 1)
	q.PutAll(10, 20, 30)
	q.PutPrio(2, 40)
	q.PutPrio(1, 50)
	if q.Size() != len(arr) {
		t.Error(q.Size())
	}
	elements = q.All()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	if len(elements) != len(arr) {
		t.Error(len(elements))
	} else {
		for i, num := range arr {
			el := elements[i]
			if el == nil {
				t.Error("in", i, "element is nil")
			} else if num2, ok := el.(int); ok && num2 != num {
				t.Error("in", i, "element is not", num, "but", el)
			}
		}
	}
}

func TestPutPrioInsA(t *testing.T) {
	arr := []int{40, 10, 20, 30}
	q := New(len(arr))
	q.PutAll(20, 30)
	q.PutPrio(1, 40)
	q.PutPrio(2, 10)
	if q.Size() != len(arr) {
		t.Error(q.Size())
	}
	elements := q.All()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	if len(elements) != len(arr) {
		t.Error(len(elements))
	} else {
		for i, num := range arr {
			el := elements[i]
			if el == nil {
				t.Error("in", i, "element is nil")
			} else if num2, ok := el.(int); ok && num2 != num {
				t.Error("in", i, "element is not", num, "but", el)
			}
		}
	}
	arr = []int{50, 40, 10, 20, 30}
	q = New(len(arr) - 1)
	q.PutAll(10, 20, 30)
	q.PutPrio(2, 40)
	q.PutPrio(1, 50)
	if q.Size() != len(arr) {
		t.Error(q.Size())
	}
	elements = q.All()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	if len(elements) != len(arr) {
		t.Error(len(elements))
	} else {
		for i, num := range arr {
			el := elements[i]
			if el == nil {
				t.Error("in", i, "element is nil")
			} else if num2, ok := el.(int); ok && num2 != num {
				t.Error("in", i, "element is not", num, "but", el)
			}
		}
	}
}

func TestPutPrioInsB(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50, 60}
	q := New(len(arr))
	q.PutAll(50, 60)
	q.PutPrio(4, 40)
	q.PutPrio(3, 30)
	q.PutPrio(2, 20)
	q.PutPrio(1, 10)
	if q.Size() != len(arr) {
		t.Error(q.Size())
	}
	if len(q.data) != len(arr) {
		t.Error(len(q.data))
	}
	elements := q.All()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	if len(elements) != len(arr) {
		t.Error(len(elements))
	} else {
		for i, num := range arr {
			el := elements[i]
			if el == nil {
				t.Error("in", i, "element is nil")
			} else if num2, ok := el.(int); ok && num2 != num {
				t.Error("in", i, "element is not", num, "but", el)
			}
		}
	}
	arr = []int{10, 20, 30, 40, 50, 60}
	q = New(len(arr))
	q.PutAll(60)
	q.PutPrio(4, 40)
	q.PutPrio(3, 30)
	q.PutPrio(2, 20)
	q.PutPrio(1, 10)
	q.PutPrio(5, 50)
	if q.Size() != len(arr) {
		t.Error(q.Size())
	}
	if len(q.data) != len(arr) {
		t.Error(len(q.data))
	}
	elements = q.All()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	if len(elements) != len(arr) {
		t.Error(len(elements))
	} else {
		for i, num := range arr {
			el := elements[i]
			if el == nil {
				t.Error("in", i, "element is nil")
			} else if num2, ok := el.(int); ok && num2 != num {
				t.Error("in", i, "element is not", num, "but", el)
			}
		}
	}
}

func TestPutPrioInsC(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50, 60, 70}
	q := New(len(arr) - 1)
	q.PutAll(60, 70)
	q.PutPrio(4, 40)
	q.PutPrio(3, 30)
	q.PutPrio(2, 20)
	q.PutPrio(1, 10)
	if q.Size() != len(arr)-1 {
		t.Error(q.Size())
	}
	if len(q.data) != len(arr)-1 {
		t.Error(len(q.data))
	}
	q.PutPrio(5, 50)
	if q.Size() != len(arr) {
		t.Error(q.Size())
	}
	if len(q.data) <= len(arr)-1 {
		t.Error(len(q.data))
	}
	elements := q.All()
	if q.Size() != 0 {
		t.Error(q.Size())
	}
	if len(elements) != len(arr) {
		t.Error(len(elements))
	} else {
		for i, num := range arr {
			el := elements[i]
			if el == nil {
				t.Error("in", i, "element is nil")
			} else if num2, ok := el.(int); ok && num2 != num {
				t.Error("in", i, "element is not", num, "but", el)
			}
		}
	}
}
