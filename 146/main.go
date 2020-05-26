package main

import "fmt"

type DoubleList struct {
	Value int
	Head  *DoubleList
	Next  *DoubleList
}

type LRUCache struct {
	Size     int
	Capacity int
	Head     *DoubleList
	Last     *DoubleList
	Arr      map[int]int
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		Size:     0,
		Capacity: capacity,
		Arr:      make(map[int]int),
	}
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	result := -1
	if value, ok := this.Arr[key]; ok == true {
		result = value
		head := this.Head
		for head != this.Last {
			if head.Value == key {
				break
			}
			head = head.Next
		}
		// 把对应节点挪到双向链表的头部（分集中情况，本身节点就在头部，节点在尾部，节点在中间）
		if head != this.Head {
			if head == this.Last {
				this.Last = head.Head
				this.Last.Next = nil
			} else {
				head.Head.Next = head.Next
				head.Next.Head = head.Head
			}
			head.Head = nil
			head.Next = this.Head
			this.Head.Head = head
			this.Head = head
		}
	}
	return result
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) > -1 {
		this.Arr[key] = value
		return
	}
	if this.Size == this.Capacity {
		delete(this.Arr, this.Last.Value)
		this.Last = this.Last.Head
		if this.Last != nil && this.Last.Next != nil {
			this.Last.Next = nil
		}
	} else {
		this.Size++
	}
	this.Arr[key] = value
	head := DoubleList{
		Value: key,
		Next:  this.Head,
	}
	if this.Head != nil {
		this.Head.Head = &head
	}
	this.Head = &head
	if this.Last == nil {
		this.Last = this.Head
	}
}

func main() {
	lru := Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	fmt.Println(lru.Get(4))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(1))
	lru.Put(5, 5)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
	fmt.Println(lru.Get(5))
}
