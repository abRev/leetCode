package heap

import (
	"leetcode/utils"
)

const (
	growthFactor = float32(2.0)
)

type Heap struct {
	elements   []interface{}
	Comparator utils.Comparator
	size       int
}

func NewWithIntComparator(values ...interface{}) *Heap {
	heap := &Heap{
		size:       1,
		Comparator: utils.IntComparator,
	}
	if len(values) > 0 {
		heap.Add(values...)
	}
	return heap
}

func (this *Heap) Add(values ...interface{}) {
	this.growBy(len(values))
	for _, value := range values {
		this.Push(value)
	}
}

func (this *Heap) Pop() interface{} {
	result := this.elements[1]
	this.size--
	this.elements[1] = this.elements[this.size]
	this.adjustmentPop(1)
	return result
}

func (this *Heap) Push(value interface{}) {
	this.elements[this.size] = value
	temp := this.size
	this.size++
	this.adjustment(temp)
}

func (this *Heap) adjustmentPop(size int) {
	if size*2 >= this.size {
		return
	}
	if size*2 == this.size-1 {
		if this.Comparator(this.elements[size], this.elements[2*size]) > 0 {
			temp := this.elements[size]
			this.elements[size] = this.elements[2*size]
			this.elements[2*size] = temp
		}
		return
	}
	if this.Comparator(this.elements[2*size+1], this.elements[2*size]) > 0 {
		if this.Comparator(this.elements[size], this.elements[2*size+1]) < 0 {
			temp := this.elements[size]
			this.elements[size] = this.elements[2*size+1]
			this.elements[2*size+1] = temp
			this.adjustmentPop(2*size + 1)
		} else if this.Comparator(this.elements[size], this.elements[2*size]) < 0 {
			temp := this.elements[size]
			this.elements[size] = this.elements[2*size]
			this.elements[2*size] = temp
			this.adjustmentPop(2 * size)
		}
	} else {
		if this.Comparator(this.elements[size], this.elements[2*size]) < 0 {
			temp := this.elements[size]
			this.elements[size] = this.elements[2*size]
			this.elements[2*size] = temp
			this.adjustmentPop(2 * size)
		} else if this.Comparator(this.elements[size], this.elements[2*size+1]) < 0 {
			temp := this.elements[size]
			this.elements[size] = this.elements[2*size+1]
			this.elements[2*size+1] = temp
			this.adjustmentPop(2*size + 1)
		}
	}
}

func (this *Heap) adjustment(sizeNow int) {
	if sizeNow <= 1 {
		return
	}
	if sizeNow%2 == 1 {
		if this.Comparator(this.elements[(sizeNow-1)/2], this.elements[sizeNow]) < 0 {
			temp := this.elements[(sizeNow-1)/2]
			this.elements[(sizeNow-1)/2] = this.elements[sizeNow]
			this.elements[sizeNow] = temp
			this.adjustment((sizeNow - 1) / 2)
		}
	} else {
		if this.Comparator(this.elements[sizeNow/2], this.elements[sizeNow]) < 0 {
			temp := this.elements[sizeNow/2]
			this.elements[sizeNow/2] = this.elements[sizeNow]
			this.elements[sizeNow] = temp
			this.adjustment(sizeNow / 2)
		}
	}
}

func (this *Heap) resize(cap int) {
	elements := make([]interface{}, cap, cap)
	copy(elements, this.elements)
	this.elements = elements
}

func (this *Heap) growBy(n int) {
	capNow := cap(this.elements)
	if capNow-n < this.size {
		newCap := int(growthFactor * (float32(n + capNow)))
		this.resize(newCap)
	}
}
