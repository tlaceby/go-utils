package arraylist

type ArrayList[T comparable] struct {
	data []T
}

// Creates a new arraylist. Optionally users can pass elements to initalize the array with
func New[T comparable](elements ...T) *ArrayList[T] {
	data := make([]T, len(elements))
	data = append(data, elements...)
	return &ArrayList[T]{data}
}

// Searches the array for the first occurance of a element. Returns the index of the element or -1 if not found.
func (arr *ArrayList[T]) IndexOf(element T) int {
	for index, el := range arr.data {
		if el == element {
			return index
		}
	}

	return -1
}

// Returns the element at a given index. Could panic if the index is out of bounds.
func (arr *ArrayList[T]) At(index uint64) T {
	return arr.data[index]
}

// Appends a new element to the array.
func (arr *ArrayList[T]) Push(element T) {
	arr.data = append(arr.data, element)
}

// Removes the last element from the array and returns it
func (arr *ArrayList[T]) Pop() T {
	last := arr.data[len(arr.data)-1]
	arr.data = arr.data[:len(arr.data)-1]
	return last
}

// Removes the first element in the array
func (arr *ArrayList[T]) Shift() T {
	first := arr.data[0]
	arr.data = arr.data[1:]
	return first
}

// Inserts an element into the beginning of the array
func (arr *ArrayList[T]) Unshift(element T) {
	if arr.Empty() {
		arr.data = append(arr.data, element)
		return
	}

	arr.data = append([]T{element}, arr.data...)
}

// Returns the length of the array
func (arr *ArrayList[T]) Length() int {
	return len(arr.data)
}

// Returns whether the length == 0
func (arr *ArrayList[T]) Empty() bool {
	return len(arr.data) == 0
}
