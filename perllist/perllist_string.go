package perllist

// base type

// T2 contains a slice of T1 values
type StringList struct {
	elements []string
}

// NewL creates a new List of T1 values
// with optional initial values
func NewStringList(values ...string) *StringList {
	return &StringList{values}
}

// Len returns the length of the underlying slice
func (l *StringList) Len() int {
	return len(l.elements)
}

// Push appends an element at the end of the list
func (l *StringList) Push(element string) {
	l.elements = append(l.elements, element)
}

// Unshift adds an element at the beginning of the list
func (l *StringList) Unshift(element string) {
	newl := append([]string{element}, l.elements...)
	l.elements = newl
}

// Shift returns the first element and removes it from the list
// returns zero value if list is empty
func (l *StringList) Shift() string {
	if len(l.elements) == 0 {
		return *new(string)
	}
	first := l.elements[0]
	l.elements = l.elements[1:]
	return first
}

// Pop returns the last element and removes it from the list
// returns zero value if list is empty
func (l *StringList) Pop() string {
	if len(l.elements) == 0 {
		return *new(string)
	}
	lastIndex := len(l.elements) - 1
	last := l.elements[lastIndex]
	l.elements = l.elements[:lastIndex]
	return last
}

// Foreach iterates over the slice by calling the given function
// providing the slice index and the value
func (l *StringList) Foreach(f func(int, string)) {
	for index, value := range l.elements {
		f(index, value)
	}
}

// Grep iterates over the slice by calling the given function
// providing the slice index and the value
// and returning a new List that contains all element where
// the given function returned true
func (l *StringList) Grep(f func(int, string) bool) StringList {
	result := StringList{}
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
func (l *StringList) Map(f func(int, string) string) StringList {
	result := StringList{}
	for index, value := range l.elements {
		result.elements = append(result.elements, f(index, value))
	}
	return result
}
