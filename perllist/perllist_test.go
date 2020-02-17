package perllist

import (
	"testing"
)

func TestNewT1(t *testing.T) {
	cand := []struct {
		name   string
		values []T1
		length int
	}{
		{"empty", []T1{}, 0},
		{"eins", []T1{1}, 1},
		{"drei", []T1{1, 2, 3}, 3},
		{"fünf", []T1{1, 2, 3, 4, 5}, 5},
	}
	for _, c := range cand {
		t.Run(c.name, func(t *testing.T) {
			liste := NewT2(c.values...)
			if len(liste.elements) != c.length {
				t.Errorf("ERROR: wrong length, got %d, exp %d",
					len(liste.elements), c.length)
			}
		})
	}
}

func TestLenT1(t *testing.T) {
	cand := []struct {
		name   string
		values []T1
		length int
	}{
		{"empty", []T1{}, 0},
		{"eins", []T1{1}, 1},
		{"drei", []T1{1, 2, 3}, 3},
		{"fünf", []T1{1, 2, 3, 4, 5}, 5},
	}
	for _, c := range cand {
		t.Run(c.name, func(t *testing.T) {
			liste := NewT2(c.values...)
			if liste.Len() != c.length {
				t.Errorf("ERROR: wrong length, got %d, exp %d",
					liste.Len(), c.length)
			}
		})
	}
}

func TestPushT1(t *testing.T) {

	cand := []struct {
		name   string
		values []T1
	}{
		{"ein Element", []T1{1}},
		{"mehrere Element", []T1{1, 2, 3, 4}},
	}

	for _, c := range cand {
		liste := NewT2()
		t.Run(c.name, func(t *testing.T) {
			for i, val := range c.values {
				liste.Push(val)
				if len(liste.elements) != i+1 {
					t.Errorf("ERROR: wrong length: got %d, exp %d",
						len(liste.elements), i+1)
				}
			}
		})
	}
}

func TestCombinedSimpleT1(t *testing.T) {
	liste := NewT2()
	if liste.Len() != 0 {
		t.Errorf("ERROR: wrong length: got %d, exp %d", liste.Len(), 0)
	}
	liste.Push(1)
	if liste.Len() != 1 {
		t.Errorf("ERROR: wrong length: got %d, exp %d", liste.Len(), 1)
	}
	liste.Push(2)
	liste.Push(3)
	if liste.Len() != 3 {
		t.Errorf("ERROR: wrong length: got %d, exp %d", liste.Len(), 2)
	}
	val := liste.Pop()
	if val != 3 {
		t.Errorf("ERROR: wrong value: got %d, exp %d", val, 3)
	}
	if liste.Len() != 2 {
		t.Errorf("ERROR: wrong length: got %d, exp %d", liste.Len(), 2)
	}
	liste.Unshift(-1)
	if liste.Len() != 3 {
		t.Errorf("ERROR: wrong length: got %d, exp %d", liste.Len(), 3)
	}
	val = liste.Pop()
	if val != 2 {
		t.Errorf("ERROR: wrong value: got %d, exp %d", val, 2)
	}
	val = liste.Shift()
	if val != -1 {
		t.Errorf("ERROR: wrong value: got %d, exp %d", val, -1)
	}
	val = liste.Shift()
	if val != 1 {
		t.Errorf("ERROR: wrong value: got %d, exp %d", val, 1)
	}
	val = liste.Shift()
	if val != 0 {
		t.Errorf("ERROR: wrong value: got %d, exp %d", val, 0)
	}
	val = liste.Shift()
	if val != 0 {
		t.Errorf("ERROR: wrong value: got %d, exp %d", val, 0)
	}
	val = liste.Pop()
	if val != 0 {
		t.Errorf("ERROR: wrong value: got %d, exp %d", val, 0)
	}
}
