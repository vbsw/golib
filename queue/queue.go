/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package queue provides a simple First-In-First-Out queue.
package queue

const (
	PrioMin = int(^uint(0) >> 1)
	PrioMax = -PrioMin - 1
)

// TypedElement is an interface to control elements by type.
type TypedElement interface {
	OnQueuePut()
	OnQueueRemove()
}

// Queue is a First-In-First-Out buffer.
type Queue struct {
	prios      []int
	data       []interface{}
	indexRead  int
	indexWrite int
	empty      bool
}

// TypedQueue is a First-In-First-Out buffer for typed elements.
type TypedQueue struct {
	prios      []int
	data       []TypedElement
	indexRead  int
	indexWrite int
	empty      bool
}

// New returns a new queue. If capacity is negative,
// capacity is set to zero and queue is initialized
// with capacity of 8 at first Put.
func New(capacity int) *Queue {
	que := new(Queue)
	if capacity > 0 {
		que.prios = make([]int, capacity)
		que.data = make([]interface{}, capacity)
	}
	que.empty = true
	return que
}

// NewTyped returns a new typed queue. If capacity is negative,
// capacity is set to zero and queue is initialized
// with capacity of 8 at first Put.
func NewTyped(capacity int) *TypedQueue {
	que := new(TypedQueue)
	if capacity > 0 {
		que.prios = make([]int, capacity)
		que.data = make([]TypedElement, capacity)
	}
	que.empty = true
	return que
}

// Put appends element at the end of the queue.
func (que *Queue) Put(element interface{}) {
	index := que.prepareCapacity(que.indexWrite, 1, PrioMin)
	que.prios[index] = PrioMin
	que.data[index] = element
}

// PutPrio inserts element into queue. The position of inserted element
// is determined by parameter prio. A low prio number means hight priority.
// Element is put before other elements with lower priority, or
// after elements with same or higher priority.
func (que *Queue) PutPrio(prio int, element interface{}) {
	index := que.indexInsertion(prio)
	index = que.prepareCapacity(index, 1, prio)
	que.prios[index] = prio
	que.data[index] = element
}

// PutAll appends elements at the end of the queue.
func (que *Queue) PutAll(elements ...interface{}) {
	if len(elements) > 0 {
		if len(elements) == 1 {
			que.Put(elements[0])
		} else {
			index := que.prepareCapacity(que.indexWrite, len(elements), PrioMin)
			que.setPrio(PrioMin, index, index+len(elements))
			copy(que.data[index:], elements)
		}
	}
}

// PutAllPrio inserts elements into queue. The position of inserted elements
// is determined by parameter prio. A low prio number means hight priority.
// Elements are put before other elements with lower priority, or
// after elements with same or higher priority.
func (que *Queue) PutAllPrio(prio int, elements ...interface{}) {
	if len(elements) > 0 {
		if len(elements) == 1 {
			que.PutPrio(prio, elements[0])
		} else {
			index := que.indexInsertion(prio)
			index = que.prepareCapacity(index, len(elements), prio)
			que.setPrio(prio, index, index+len(elements))
			copy(que.data[index:], elements)
		}
	}
}

// First removes first element from queue and returns it.
func (que *Queue) First() interface{} {
	var data interface{}
	if !que.empty {
		data = que.data[que.indexRead]
		if que.indexRead+1 != len(que.data) {
			que.indexRead++
		} else {
			que.indexRead = 0
		}
		if que.indexRead == que.indexWrite {
			que.indexRead = 0
			que.indexWrite = 0
			que.empty = true
		}
	}
	return data
}

// All removes all elements from queue and returns them.
func (que *Queue) All() []interface{} {
	var data []interface{}
	if !que.empty {
		if que.indexRead < que.indexWrite {
			data = make([]interface{}, que.indexWrite-que.indexRead)
			copy(data, que.data[que.indexRead:que.indexWrite])
		} else {
			dataRight := que.data[que.indexRead:]
			dataLeft := que.data[:que.indexWrite]
			data = make([]interface{}, len(dataLeft)+len(dataRight))
			copy(data, dataRight)
			copy(data[len(dataRight):], dataLeft)
		}
		que.empty = true
	}
	que.indexRead = 0
	que.indexWrite = 0
	return data
}

// Size returns the number of elements in queue.
func (que *Queue) Size() int {
	var size int
	if !que.empty {
		if que.indexRead < que.indexWrite {
			size = que.indexWrite - que.indexRead
		} else {
			size = que.indexWrite + len(que.data) - que.indexRead
		}
	}
	return size
}

// prepareCapacity ensures enough capacity at given index. Returns new
// index, if data has been reallocated.
func (que *Queue) prepareCapacity(index, capAdd, prio int) int {
	indexNew := 0
	if que.empty {
		capNew := que.capacityNew(capAdd)
		if capNew > len(que.data) {
			que.prios = make([]int, capNew)
			que.data = make([]interface{}, capNew)
		}
		que.indexRead = 0
		que.indexWrite = capAdd
	} else {
		if index == que.indexWrite && (que.indexRead != que.indexWrite || que.prios[que.indexWrite] <= prio) {
			indexNew = que.prepareCapacityApd(capAdd)
		} else if index == que.indexRead {
			indexNew = que.prepareCapacityPrd(capAdd)
		} else {
			indexNew = que.prepareCapacityIns(index, capAdd)
		}
	}
	if que.indexWrite == len(que.data) {
		que.indexWrite = 0
	}
	que.empty = false
	return indexNew
}

// prepareCapacityApd ensures enough capacity for additional elements at the end of the queue.
func (que *Queue) prepareCapacityApd(capAdd int) int {
	var indexNew int
	if que.indexRead < que.indexWrite {
		if que.indexWrite+capAdd <= len(que.data) {
			indexNew = que.indexWrite
			que.indexWrite += capAdd
		} else {
			capNew := que.capacityNew(capAdd)
			if capNew == len(que.data) {
				copy(que.prios, que.prios[que.indexRead:que.indexWrite])
				copy(que.data, que.data[que.indexRead:que.indexWrite])
			} else {
				priosNew := make([]int, capNew)
				dataNew := make([]interface{}, capNew)
				copy(priosNew, que.prios[que.indexRead:que.indexWrite])
				copy(dataNew, que.data[que.indexRead:que.indexWrite])
				que.prios = priosNew
				que.data = dataNew
			}
			indexNew = que.indexWrite - que.indexRead
			que.indexRead = 0
			que.indexWrite = indexNew + capAdd
		}
	} else {
		if que.indexWrite+capAdd <= que.indexRead {
			indexNew = que.indexWrite
			que.indexWrite += capAdd
		} else {
			capNew := que.capacityNew(capAdd)
			priosNew := make([]int, capNew)
			dataNew := make([]interface{}, capNew)
			copy(priosNew, que.prios[que.indexRead:])
			copy(dataNew, que.data[que.indexRead:])
			if que.indexWrite > 0 {
				sizeRight := len(que.data) - que.indexRead
				copy(priosNew[sizeRight:], que.prios[:que.indexWrite])
				copy(dataNew[sizeRight:], que.data[:que.indexWrite])
			}
			indexNew = que.indexWrite + len(que.data) - que.indexRead
			que.indexRead = 0
			que.indexWrite = indexNew + capAdd
			que.prios = priosNew
			que.data = dataNew
		}
	}
	return indexNew
}

// prepareCapacityPrd ensures enough capacity for additional elements at the beginning of the queue.
func (que *Queue) prepareCapacityPrd(capAdd int) int {
	var indexNew int
	if que.indexRead < que.indexWrite {
		if capAdd <= que.indexRead {
			indexNew = que.indexRead - capAdd
			que.indexRead = indexNew
		} else {
			capNew := que.capacityNew(capAdd)
			if capNew == len(que.data) {
				if que.indexRead == 0 {
					indexNew = len(que.data) - capAdd
					que.indexRead = indexNew
				} else {
					copy(que.prios[capAdd:], que.prios[que.indexRead:que.indexWrite])
					copy(que.data[capAdd:], que.data[que.indexRead:que.indexWrite])
					que.indexWrite = capAdd + (que.indexWrite - que.indexRead)
					indexNew = 0
					que.indexRead = 0
				}
			} else {
				priosNew := make([]int, capNew)
				dataNew := make([]interface{}, capNew)
				copy(priosNew[capAdd:], que.prios[que.indexRead:que.indexWrite])
				copy(dataNew[capAdd:], que.data[que.indexRead:que.indexWrite])
				que.indexWrite = capAdd + (que.indexWrite - que.indexRead)
				indexNew = 0
				que.indexRead = 0
				que.prios = priosNew
				que.data = dataNew
			}
		}
	} else {
		if que.indexRead-capAdd >= que.indexWrite {
			indexNew = que.indexRead - capAdd
			que.indexRead = indexNew
		} else {
			capNew := que.capacityNew(capAdd)
			priosNew := make([]int, capNew)
			dataNew := make([]interface{}, capNew)
			copy(priosNew[capAdd:], que.prios[que.indexRead:])
			copy(dataNew[capAdd:], que.data[que.indexRead:])
			if que.indexWrite > 0 {
				sizeRight := len(que.data) - que.indexRead
				copy(priosNew[capAdd+sizeRight:], que.prios[:que.indexWrite])
				copy(dataNew[capAdd+sizeRight:], que.data[:que.indexWrite])
			}
			que.indexWrite = capAdd + (que.indexWrite + len(que.data) - que.indexRead)
			indexNew = 0
			que.indexRead = 0
			que.prios = priosNew
			que.data = dataNew
		}
	}
	return indexNew
}

// prepareCapacityIns ensures enough capacity for additional elements at index in queue.
func (que *Queue) prepareCapacityIns(index, capAdd int) int {
	var indexNew int
	capNew := que.capacityNew(capAdd)
	if que.indexRead < que.indexWrite {
		if capNew == len(que.data) {
			sizeLeft := index - que.indexRead
			sizeRight := que.indexWrite - index
			if sizeLeft >= sizeRight {
				if capAdd <= len(que.data)-index {
					indexNew = que.moveRightSeq(index, capAdd, sizeRight)
				} else if capAdd <= index {
					indexNew = que.moveLeftSeq(index, capAdd, sizeLeft)
				} else {
					indexNew = que.moveAllSeq(index, capAdd, capNew)
				}
			} else {
				if capAdd <= index {
					indexNew = que.moveLeftSeq(index, capAdd, sizeLeft)
				} else if capAdd <= len(que.data)-index {
					indexNew = que.moveRightSeq(index, capAdd, sizeRight)
				} else {
					indexNew = que.moveAllSeq(index, capAdd, capNew)
				}
			}
		} else {
			indexNew = que.moveAllSeq(index, capAdd, capNew)
		}
	} else {
		if capNew == len(que.data) {
			if index < que.indexRead {
				indexNew = que.moveRightWrap(index, capAdd)
			} else {
				indexNew = que.moveLeftWrap(index, capAdd)
			}
		} else {
			indexNew = que.moveAllWrap(index, capAdd, capNew)
		}
	}
	return indexNew
}

func (que *Queue) moveRightSeq(index, capAdd, sizeRight int) int {
	indexRight := index + capAdd
	capRight := len(que.data) - indexRight
	capLeft := sizeRight - capRight
	indexWriteNew := que.indexWrite + capAdd
	if capRight > 0 {
		copy(que.prios[indexRight:], que.prios[index:que.indexWrite])
		copy(que.data[indexRight:], que.data[index:que.indexWrite])
	}
	if capLeft > 0 {
		copy(que.prios, que.prios[index+capRight:que.indexWrite])
		copy(que.data, que.data[index+capRight:que.indexWrite])
		indexWriteNew = capLeft
	}
	que.indexWrite = indexWriteNew
	return index
}

func (que *Queue) moveLeftSeq(index, capAdd, sizeLeft int) int {
	indexNew := index - capAdd
	indexLeft := indexNew - sizeLeft
	if indexLeft >= 0 {
		copy(que.prios[indexLeft:], que.prios[que.indexRead:index])
		copy(que.data[indexLeft:], que.data[que.indexRead:index])
		que.indexRead = indexLeft
	} else {
		capRight := sizeLeft - indexNew
		indexRight := len(que.data) - capRight
		if indexNew > 0 {
			copy(que.prios, que.prios[capAdd:index])
			copy(que.data, que.data[capAdd:index])
		}
		copy(que.prios[indexRight:], que.prios[que.indexRead:capAdd])
		copy(que.data[indexRight:], que.data[que.indexRead:capAdd])
		que.indexRead = indexRight
	}
	return indexNew
}

func (que *Queue) moveAllSeq(index, capAdd, capNew int) int {
	priosNew := make([]int, capNew)
	dataNew := make([]interface{}, capNew)
	indexNew := index - que.indexRead
	indexOff := indexNew + capAdd
	priosLeft := que.prios[que.indexRead:index]
	priosRight := que.prios[index:que.indexWrite]
	dataLeft := que.data[que.indexRead:index]
	dataRight := que.data[index:que.indexWrite]
	copy(priosNew, priosLeft)
	copy(priosNew[indexOff:], priosRight)
	copy(dataNew, dataLeft)
	copy(dataNew[indexOff:], dataRight)
	que.indexWrite = que.indexWrite - que.indexRead + capAdd
	que.indexRead = 0
	que.prios = priosNew
	que.data = dataNew
	return indexNew
}

func (que *Queue) moveRightWrap(index, capAdd int) int {
	indexRight := index + capAdd
	copy(que.prios[indexRight:], que.prios[index:que.indexWrite])
	copy(que.data[indexRight:], que.data[index:que.indexWrite])
	que.indexWrite += capAdd
	return index
}

func (que *Queue) moveLeftWrap(index, capAdd int) int {
	indexNew := index - capAdd
	sizeLeft := index - que.indexRead
	indexLeft := indexNew - sizeLeft
	copy(que.prios[indexLeft:], que.prios[que.indexRead:index])
	copy(que.data[indexLeft:], que.data[que.indexRead:index])
	que.indexRead = indexLeft
	return indexNew
}

func (que *Queue) moveAllWrap(index, capAdd, capNew int) int {
	var indexNew int
	priosNew := make([]int, capNew)
	dataNew := make([]interface{}, capNew)
	if index < que.indexRead {
		indexOffA := len(que.data) - que.indexRead
		indexNew = indexOffA + index
		indexOffB := indexNew + capAdd
		priosLeft := que.prios[que.indexRead:]
		priosRightA := que.prios[:index]
		priosRightB := que.prios[index:que.indexWrite]
		dataLeft := que.data[que.indexRead:]
		dataRightA := que.data[:index]
		dataRightB := que.data[index:que.indexWrite]
		copy(priosNew, priosLeft)
		copy(priosNew[indexOffA:], priosRightA)
		copy(priosNew[indexOffB:], priosRightB)
		copy(dataNew, dataLeft)
		copy(dataNew[indexOffA:], dataRightA)
		copy(dataNew[indexOffB:], dataRightB)
	} else {
		indexNew = index - que.indexRead
		indexOffA := indexNew + capAdd
		indexOffB := indexOffA + len(que.data) - index
		priosLeft := que.prios[que.indexRead:index]
		priosRightA := que.prios[index:]
		priosRightB := que.prios[:que.indexWrite]
		dataLeft := que.data[que.indexRead:index]
		dataRightA := que.data[index:]
		dataRightB := que.data[:que.indexWrite]
		copy(priosNew, priosLeft)
		copy(priosNew[indexOffA:], priosRightA)
		copy(priosNew[indexOffB:], priosRightB)
		copy(dataNew, dataLeft)
		copy(dataNew[indexOffA:], dataRightA)
		copy(dataNew[indexOffB:], dataRightB)
	}
	que.indexWrite = que.indexWrite + len(que.data) - que.indexRead + capAdd
	que.indexRead = 0
	que.prios = priosNew
	que.data = dataNew
	return indexNew
}

// indexInsertion returns index to insert new elements with prio.
func (que *Queue) indexInsertion(prio int) int {
	index := que.indexWrite
	if prio != PrioMin && !que.empty {
		if que.indexRead < que.indexWrite {
			index = que.indexInsertionFromTo(prio, que.indexRead, que.indexWrite)
		} else {
			index = que.indexInsertionFromTo(prio, que.indexRead, len(que.prios))
			if index == len(que.prios) {
				index = que.indexInsertionFromTo(prio, 0, que.indexWrite)
			}
		}
	}
	return index
}

// indexInsertionFromTo returns index to insert new elements with prio in specific range.
func (que *Queue) indexInsertionFromTo(prio, from, to int) int {
	// binary search
	left := from
	right := to - 1
	for left <= right {
		middle := (left + right) / 2
		value := que.prios[middle]
		if prio > value {
			left = middle + 1
		} else if prio < value {
			right = middle - 1
		} else {
			// search right side
			left = middle + 1
			for left <= right {
				middle = (left + right) / 2
				value = que.prios[middle]
				if prio == value {
					left = middle + 1
				} else {
					right = middle - 1
				}
			}
			return left
		}
	}
	return left
}

// capacityNew returns new capacity to hold additional elements.
func (que *Queue) capacityNew(capAdd int) int {
	capNew := len(que.data)
	capMin := que.Size() + capAdd
	if capNew < capMin {
		if capNew == 0 {
			capNew = 8
		}
		for capNew < capMin {
			capNew *= 2
		}
	}
	return capNew
}

// setPrio sets priority values in specific range.
func (que *Queue) setPrio(prio, from, to int) {
	prios := que.prios[from:to]
	for i := range prios {
		prios[i] = prio
	}
}

// Put appends element at the end of the queue. Function OnQueuePut is called
// on element after it has been added to the queue.
func (que *TypedQueue) Put(element TypedElement) {
	index := que.prepareCapacity(que.indexWrite, 1, PrioMin)
	que.prios[index] = PrioMin
	que.data[index] = element
	element.OnQueuePut()
}

// PutPrio inserts element into queue. Function OnQueuePut is called
// on element after it has been added to the queue. The position of inserted element
// is determined by parameter prio. A low prio number means hight priority.
// Element is put before other elements with lower priority, or
// after elements with same or higher priority.
func (que *TypedQueue) PutPrio(prio int, element TypedElement) {
	index := que.indexInsertion(prio)
	index = que.prepareCapacity(index, 1, prio)
	que.prios[index] = prio
	que.data[index] = element
	element.OnQueuePut()
}

// PutAll appends elements at the end of the queue. Function OnQueuePut is called
// on every element after they have been added to the queue.
func (que *TypedQueue) PutAll(elements ...TypedElement) {
	if len(elements) > 0 {
		if len(elements) == 1 {
			que.Put(elements[0])
		} else {
			index := que.prepareCapacity(que.indexWrite, len(elements), PrioMin)
			que.setPrio(PrioMin, index, index+len(elements))
			copy(que.data[index:], elements)
			for _, element := range elements {
				element.OnQueuePut()
			}
		}
	}
}

// PutAllPrio inserts elements into queue. Function OnQueuePut is called
// on every element after they have been added to the queue. The position of inserted elements
// is determined by parameter prio. A low prio number means hight priority.
// Elements are put before other elements with lower priority, or
// after elements with same or higher priority.
func (que *TypedQueue) PutAllPrio(prio int, elements ...TypedElement) {
	if len(elements) > 0 {
		if len(elements) == 1 {
			que.PutPrio(prio, elements[0])
		} else {
			index := que.indexInsertion(prio)
			index = que.prepareCapacity(index, len(elements), prio)
			que.setPrio(prio, index, index+len(elements))
			copy(que.data[index:], elements)
			for _, element := range elements {
				element.OnQueuePut()
			}
		}
	}
}

// First removes first element from queue and returns it. Function OnQueueRemove is called
// on element after it has been removed from the queue.
func (que *TypedQueue) First() TypedElement {
	var data TypedElement
	if !que.empty {
		data = que.data[que.indexRead]
		if que.indexRead+1 != len(que.data) {
			que.indexRead++
		} else {
			que.indexRead = 0
		}
		if que.indexRead == que.indexWrite {
			que.indexRead = 0
			que.indexWrite = 0
			que.empty = true
		}
		data.OnQueueRemove()
	}
	return data
}

// All removes all elements from queue and returns them. Function OnQueueRemove is called
// on every element after they have been removed from the queue.
func (que *TypedQueue) All() []TypedElement {
	var data []TypedElement
	if !que.empty {
		if que.indexRead < que.indexWrite {
			data = make([]TypedElement, que.indexWrite-que.indexRead)
			copy(data, que.data[que.indexRead:que.indexWrite])
		} else {
			dataRight := que.data[que.indexRead:]
			dataLeft := que.data[:que.indexWrite]
			data = make([]TypedElement, len(dataLeft)+len(dataRight))
			copy(data, dataRight)
			copy(data[len(dataRight):], dataLeft)
		}
		que.empty = true
	}
	que.indexRead = 0
	que.indexWrite = 0
	for _, element := range data {
		element.OnQueueRemove()
	}
	return data
}

// Size returns the number of elements in queue.
func (que *TypedQueue) Size() int {
	var size int
	if !que.empty {
		if que.indexRead < que.indexWrite {
			size = que.indexWrite - que.indexRead
		} else {
			size = que.indexWrite + len(que.data) - que.indexRead
		}
	}
	return size
}

// prepareCapacity ensures enough capacity at given index. Returns new
// index, if data has been reallocated.
func (que *TypedQueue) prepareCapacity(index, capAdd, prio int) int {
	indexNew := 0
	if que.empty {
		capNew := que.capacityNew(capAdd)
		if capNew > len(que.data) {
			que.prios = make([]int, capNew)
			que.data = make([]TypedElement, capNew)
		}
		que.indexRead = 0
		que.indexWrite = capAdd
	} else {
		if index == que.indexWrite && (que.indexRead != que.indexWrite || que.prios[que.indexWrite] <= prio) {
			indexNew = que.prepareCapacityApd(capAdd)
		} else if index == que.indexRead {
			indexNew = que.prepareCapacityPrd(capAdd)
		} else {
			indexNew = que.prepareCapacityIns(index, capAdd)
		}
	}
	if que.indexWrite == len(que.data) {
		que.indexWrite = 0
	}
	que.empty = false
	return indexNew
}

// prepareCapacityApd ensures enough capacity for additional elements at the end of the queue.
func (que *TypedQueue) prepareCapacityApd(capAdd int) int {
	var indexNew int
	if que.indexRead < que.indexWrite {
		if que.indexWrite+capAdd <= len(que.data) {
			indexNew = que.indexWrite
			que.indexWrite += capAdd
		} else {
			capNew := que.capacityNew(capAdd)
			if capNew == len(que.data) {
				copy(que.prios, que.prios[que.indexRead:que.indexWrite])
				copy(que.data, que.data[que.indexRead:que.indexWrite])
			} else {
				priosNew := make([]int, capNew)
				dataNew := make([]TypedElement, capNew)
				copy(priosNew, que.prios[que.indexRead:que.indexWrite])
				copy(dataNew, que.data[que.indexRead:que.indexWrite])
				que.prios = priosNew
				que.data = dataNew
			}
			indexNew = que.indexWrite - que.indexRead
			que.indexRead = 0
			que.indexWrite = indexNew + capAdd
		}
	} else {
		if que.indexWrite+capAdd <= que.indexRead {
			indexNew = que.indexWrite
			que.indexWrite += capAdd
		} else {
			capNew := que.capacityNew(capAdd)
			priosNew := make([]int, capNew)
			dataNew := make([]TypedElement, capNew)
			copy(priosNew, que.prios[que.indexRead:])
			copy(dataNew, que.data[que.indexRead:])
			if que.indexWrite > 0 {
				sizeRight := len(que.data) - que.indexRead
				copy(priosNew[sizeRight:], que.prios[:que.indexWrite])
				copy(dataNew[sizeRight:], que.data[:que.indexWrite])
			}
			indexNew = que.indexWrite + len(que.data) - que.indexRead
			que.indexRead = 0
			que.indexWrite = indexNew + capAdd
			que.prios = priosNew
			que.data = dataNew
		}
	}
	return indexNew
}

// prepareCapacityPrd ensures enough capacity for additional elements at the beginning of the queue.
func (que *TypedQueue) prepareCapacityPrd(capAdd int) int {
	var indexNew int
	if que.indexRead < que.indexWrite {
		if capAdd <= que.indexRead {
			indexNew = que.indexRead - capAdd
			que.indexRead = indexNew
		} else {
			capNew := que.capacityNew(capAdd)
			if capNew == len(que.data) {
				if que.indexRead == 0 {
					indexNew = len(que.data) - capAdd
					que.indexRead = indexNew
				} else {
					copy(que.prios[capAdd:], que.prios[que.indexRead:que.indexWrite])
					copy(que.data[capAdd:], que.data[que.indexRead:que.indexWrite])
					que.indexWrite = capAdd + (que.indexWrite - que.indexRead)
					indexNew = 0
					que.indexRead = 0
				}
			} else {
				priosNew := make([]int, capNew)
				dataNew := make([]TypedElement, capNew)
				copy(priosNew[capAdd:], que.prios[que.indexRead:que.indexWrite])
				copy(dataNew[capAdd:], que.data[que.indexRead:que.indexWrite])
				que.indexWrite = capAdd + (que.indexWrite - que.indexRead)
				indexNew = 0
				que.indexRead = 0
				que.prios = priosNew
				que.data = dataNew
			}
		}
	} else {
		if que.indexRead-capAdd >= que.indexWrite {
			indexNew = que.indexRead - capAdd
			que.indexRead = indexNew
		} else {
			capNew := que.capacityNew(capAdd)
			priosNew := make([]int, capNew)
			dataNew := make([]TypedElement, capNew)
			copy(priosNew[capAdd:], que.prios[que.indexRead:])
			copy(dataNew[capAdd:], que.data[que.indexRead:])
			if que.indexWrite > 0 {
				sizeRight := len(que.data) - que.indexRead
				copy(priosNew[capAdd+sizeRight:], que.prios[:que.indexWrite])
				copy(dataNew[capAdd+sizeRight:], que.data[:que.indexWrite])
			}
			que.indexWrite = capAdd + (que.indexWrite + len(que.data) - que.indexRead)
			indexNew = 0
			que.indexRead = 0
			que.prios = priosNew
			que.data = dataNew
		}
	}
	return indexNew
}

// prepareCapacityIns ensures enough capacity for additional elements at index in queue.
func (que *TypedQueue) prepareCapacityIns(index, capAdd int) int {
	var indexNew int
	capNew := que.capacityNew(capAdd)
	if que.indexRead < que.indexWrite {
		if capNew == len(que.data) {
			sizeLeft := index - que.indexRead
			sizeRight := que.indexWrite - index
			if sizeLeft >= sizeRight {
				if capAdd <= len(que.data)-index {
					indexNew = que.moveRightSeq(index, capAdd, sizeRight)
				} else if capAdd <= index {
					indexNew = que.moveLeftSeq(index, capAdd, sizeLeft)
				} else {
					indexNew = que.moveAllSeq(index, capAdd, capNew)
				}
			} else {
				if capAdd <= index {
					indexNew = que.moveLeftSeq(index, capAdd, sizeLeft)
				} else if capAdd <= len(que.data)-index {
					indexNew = que.moveRightSeq(index, capAdd, sizeRight)
				} else {
					indexNew = que.moveAllSeq(index, capAdd, capNew)
				}
			}
		} else {
			indexNew = que.moveAllSeq(index, capAdd, capNew)
		}
	} else {
		if capNew == len(que.data) {
			if index < que.indexRead {
				indexNew = que.moveRightWrap(index, capAdd)
			} else {
				indexNew = que.moveLeftWrap(index, capAdd)
			}
		} else {
			indexNew = que.moveAllWrap(index, capAdd, capNew)
		}
	}
	return indexNew
}

func (que *TypedQueue) moveRightSeq(index, capAdd, sizeRight int) int {
	indexRight := index + capAdd
	capRight := len(que.data) - indexRight
	capLeft := sizeRight - capRight
	indexWriteNew := que.indexWrite + capAdd
	if capRight > 0 {
		copy(que.prios[indexRight:], que.prios[index:que.indexWrite])
		copy(que.data[indexRight:], que.data[index:que.indexWrite])
	}
	if capLeft > 0 {
		copy(que.prios, que.prios[index+capRight:que.indexWrite])
		copy(que.data, que.data[index+capRight:que.indexWrite])
		indexWriteNew = capLeft
	}
	que.indexWrite = indexWriteNew
	return index
}

func (que *TypedQueue) moveLeftSeq(index, capAdd, sizeLeft int) int {
	indexNew := index - capAdd
	indexLeft := indexNew - sizeLeft
	if indexLeft >= 0 {
		copy(que.prios[indexLeft:], que.prios[que.indexRead:index])
		copy(que.data[indexLeft:], que.data[que.indexRead:index])
		que.indexRead = indexLeft
	} else {
		capRight := sizeLeft - indexNew
		indexRight := len(que.data) - capRight
		if indexNew > 0 {
			copy(que.prios, que.prios[capAdd:index])
			copy(que.data, que.data[capAdd:index])
		}
		copy(que.prios[indexRight:], que.prios[que.indexRead:capAdd])
		copy(que.data[indexRight:], que.data[que.indexRead:capAdd])
		que.indexRead = indexRight
	}
	return indexNew
}

func (que *TypedQueue) moveAllSeq(index, capAdd, capNew int) int {
	priosNew := make([]int, capNew)
	dataNew := make([]TypedElement, capNew)
	indexNew := index - que.indexRead
	indexOff := indexNew + capAdd
	priosLeft := que.prios[que.indexRead:index]
	priosRight := que.prios[index:que.indexWrite]
	dataLeft := que.data[que.indexRead:index]
	dataRight := que.data[index:que.indexWrite]
	copy(priosNew, priosLeft)
	copy(priosNew[indexOff:], priosRight)
	copy(dataNew, dataLeft)
	copy(dataNew[indexOff:], dataRight)
	que.indexWrite = que.indexWrite - que.indexRead + capAdd
	que.indexRead = 0
	que.prios = priosNew
	que.data = dataNew
	return indexNew
}

func (que *TypedQueue) moveRightWrap(index, capAdd int) int {
	indexRight := index + capAdd
	copy(que.prios[indexRight:], que.prios[index:que.indexWrite])
	copy(que.data[indexRight:], que.data[index:que.indexWrite])
	que.indexWrite += capAdd
	return index
}

func (que *TypedQueue) moveLeftWrap(index, capAdd int) int {
	indexNew := index - capAdd
	sizeLeft := index - que.indexRead
	indexLeft := indexNew - sizeLeft
	copy(que.prios[indexLeft:], que.prios[que.indexRead:index])
	copy(que.data[indexLeft:], que.data[que.indexRead:index])
	que.indexRead = indexLeft
	return indexNew
}

func (que *TypedQueue) moveAllWrap(index, capAdd, capNew int) int {
	var indexNew int
	priosNew := make([]int, capNew)
	dataNew := make([]TypedElement, capNew)
	if index < que.indexRead {
		indexOffA := len(que.data) - que.indexRead
		indexNew = indexOffA + index
		indexOffB := indexNew + capAdd
		priosLeft := que.prios[que.indexRead:]
		priosRightA := que.prios[:index]
		priosRightB := que.prios[index:que.indexWrite]
		dataLeft := que.data[que.indexRead:]
		dataRightA := que.data[:index]
		dataRightB := que.data[index:que.indexWrite]
		copy(priosNew, priosLeft)
		copy(priosNew[indexOffA:], priosRightA)
		copy(priosNew[indexOffB:], priosRightB)
		copy(dataNew, dataLeft)
		copy(dataNew[indexOffA:], dataRightA)
		copy(dataNew[indexOffB:], dataRightB)
	} else {
		indexNew = index - que.indexRead
		indexOffA := indexNew + capAdd
		indexOffB := indexOffA + len(que.data) - index
		priosLeft := que.prios[que.indexRead:index]
		priosRightA := que.prios[index:]
		priosRightB := que.prios[:que.indexWrite]
		dataLeft := que.data[que.indexRead:index]
		dataRightA := que.data[index:]
		dataRightB := que.data[:que.indexWrite]
		copy(priosNew, priosLeft)
		copy(priosNew[indexOffA:], priosRightA)
		copy(priosNew[indexOffB:], priosRightB)
		copy(dataNew, dataLeft)
		copy(dataNew[indexOffA:], dataRightA)
		copy(dataNew[indexOffB:], dataRightB)
	}
	que.indexWrite = que.indexWrite + len(que.data) - que.indexRead + capAdd
	que.indexRead = 0
	que.prios = priosNew
	que.data = dataNew
	return indexNew
}

// indexInsertion returns index to insert new elements with prio.
func (que *TypedQueue) indexInsertion(prio int) int {
	index := que.indexWrite
	if prio != PrioMin && !que.empty {
		if que.indexRead < que.indexWrite {
			index = que.indexInsertionFromTo(prio, que.indexRead, que.indexWrite)
		} else {
			index = que.indexInsertionFromTo(prio, que.indexRead, len(que.prios))
			if index == len(que.prios) {
				index = que.indexInsertionFromTo(prio, 0, que.indexWrite)
			}
		}
	}
	return index
}

// indexInsertionFromTo returns index to insert new elements with prio in specific range.
func (que *TypedQueue) indexInsertionFromTo(prio, from, to int) int {
	// binary search
	left := from
	right := to - 1
	for left <= right {
		middle := (left + right) / 2
		value := que.prios[middle]
		if prio > value {
			left = middle + 1
		} else if prio < value {
			right = middle - 1
		} else {
			// search right side
			left = middle + 1
			for left <= right {
				middle = (left + right) / 2
				value = que.prios[middle]
				if prio == value {
					left = middle + 1
				} else {
					right = middle - 1
				}
			}
			return left
		}
	}
	return left
}

// capacityNew returns new capacity to hold additional elements.
func (que *TypedQueue) capacityNew(capAdd int) int {
	capNew := len(que.data)
	capMin := que.Size() + capAdd
	if capNew < capMin {
		if capNew == 0 {
			capNew = 8
		}
		for capNew < capMin {
			capNew *= 2
		}
	}
	return capNew
}

// setPrio sets priority values in specific range.
func (que *TypedQueue) setPrio(prio, from, to int) {
	prios := que.prios[from:to]
	for i := range prios {
		prios[i] = prio
	}
}
