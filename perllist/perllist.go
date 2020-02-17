package perllist

// base type
type T1 = int

// T2 contains a slice of T1 values
type T2 struct {
	elements []T1
}

// NewL creates a new List of T1 values
// with optional initial values
func NewT2(values ...T1) *T2 {
	return &T2{values}
}

// Len returns the length of the underlying slice
func (l *T2) Len() int {
	return len(l.elements)
}

// Push appends an element at the end of the list
func (l *T2) Push(element T1) {
	l.elements = append(l.elements, element)
}

// Unshift adds an element at the beginning of the list
func (l *T2) Unshift(element T1) {
	newl := append([]T1{element}, l.elements...)
	l.elements = newl
}

// Shift returns the first element and removes it from the list
// returns zero value if list is empty
func (l *T2) Shift() T1 {
	if len(l.elements) == 0 {
		return *new(T1)
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// Pop returns the last element and removes it from the list
// returns zero value if list is empty
func (l *T2) Pop() T1 {
	if len(l.elements) == 0 {
		return *new(T1)
	}
	lastIndex := len(l.elements) - 1
	last := l.elements[lastIndex]
	l.elements = l.elements[:lastIndex]
	return last
}

// Foreach iterates over the slice by calling the given function
// providing the slice index and the value
func (l *T2) Foreach(f func(int, T1)) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// Grep iterates over the slice by calling the given function
// providing the slice index and the value
// and returning a new List that contains all element where
// the given function returned true
func (l *T2) Grep(f func(int, T1) bool) T2 {
	result := T2{}
	for index, value := range l.elements {
		if f(index, value) {
			result.elements = append(result.elements, value)
		}
	}
	return result
}

// Map iterates over the slice by calling the given function
// providing the slice index and the value
// and returning a new List that contains all element
// with the value returned by the function
func (l *T2) Map(f func(int, T1) T1) T2 {
	result := T2{}
	for index, value := range l.elements {
		result.elements = append(result.elements, f(index, value))
	}
	return result
}
