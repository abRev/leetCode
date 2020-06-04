package stack

const (
	growthFactor = float32(2.0)
)

type Stack struct {
	size     int
	elements []interface{}
}

func New(values ...interface{}) *Stack {
	stack := &Stack{}
	if len(values) > 0 {
		stack.Add(values...)
	}
	return stack
}

func (this *Stack) Pop() interface{} {
	this.size--
	value := this.elements[this.size]
	return value
}

func (this *Stack) Head() interface{} {
	return this.elements[this.size-1]
}

func (this *Stack) Push(value interface{}) {
	this.growBy(1)
	this.elements[this.size] = value
	this.size++
}

func (this *Stack) Add(values ...interface{}) {
	this.growBy(len(values))
	for _, value := range values {
		this.elements[this.size] = value
		this.size++
	}
}

func (this *Stack) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, this.elements)
	this.elements = newElements
}

func (this *Stack) growBy(n int) {
	currentCapacity := cap(this.elements)
	if this.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		this.resize(newCapacity)
	}
}
