package arraylist

// Accepts a function to run for each element inside an array. Does not modify the array iteself
func (arr *ArrayList[T]) ForEach(fn func(el T, indx int)) {
	for index, value := range arr.data {
		fn(value, index)
	}
}

// Accepts a function to run on each element. The function should return the new updated value.
func (arr *ArrayList[T]) Map(fn func(el T) T) {
	for indx, value := range arr.data {
		arr.data[indx] = fn(value)
	}
}
