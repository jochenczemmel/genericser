package perllist

import (
	"testing"
)

func TestNewstring(t *testing.T) {
	cand := []struct {
		name   string
		values []string
		length int
	}{
		{"empty", []string{}, 0},
		{"eins", []string{"1"}, 1},
		{"drei", []string{"1", "2", "3"}, 3},
		{"fünf", []string{"1", "2", "3", "4", "5"}, 5},
	}
	for _, c := range cand {
		t.Run(c.name, func(t *testing.T) {
			liste := NewStringList(c.values...)
			if len(liste.elements) != c.length {
				t.Errorf("ERROR: wrong length, got %d, exp %d",
					len(liste.elements), c.length)
			}
		})
	}
}

func TestLenstring(t *testing.T) {
	cand := []struct {
		name   string
		values []string
		length int
	}{
		{"empty", []string{}, 0},
		{"eins", []string{"1"}, 1},
		{"drei", []string{"1", "2", "3"}, 3},
		{"fünf", []string{"1", "2", "3", "4", "5"}, 5},
	}
	for _, c := range cand {
		t.Run(c.name, func(t *testing.T) {
			liste := NewStringList(c.values...)
			if liste.Len() != c.length {
				t.Errorf("ERROR: wrong length, got %d, exp %d",
					liste.Len(), c.length)
			}
		})
	}
}

func TestPushstring(t *testing.T) {

	cand := []struct {
		name   string
		values []string
	}{
		{"ein Element", []string{"1"}},
		{"mehrere Element", []string{"1", "2", "3", "4"}},
	}

	for _, c := range cand {
		liste := NewStringList()
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

func TestCombinedSimplestring(t *testing.T) {
	liste := NewStringList()
	if liste.Len() != 0 {
		t.Errorf("ERROR: wrong length: got %d, exp %d", liste.Len(), 0)
	}
	liste.Push("1")
	if liste.Len() != 1 {
		t.Errorf("ERROR: wrong length: got %d, exp %d", liste.Len(), 1)
	}
	liste.Push("2")
	liste.Push("3")
	if liste.Len() != 3 {
		t.Errorf("ERROR: wrong length: got %d, exp %d", liste.Len(), 2)
	}
	val := liste.Pop()
	if val != "3" {
		t.Errorf("ERROR: wrong value: got %s, exp %s", val, "3")
	}
	if liste.Len() != 2 {
		t.Errorf("ERROR: wrong length: got %d, exp %d", liste.Len(), 2)
	}
	liste.Unshift("-1")
	if liste.Len() != 3 {
		t.Errorf("ERROR: wrong length: got %d, exp %d", liste.Len(), 3)
	}
	val = liste.Pop()
	if val != "2" {
		t.Errorf("ERROR: wrong value: got %s, exp %s", val, "2")
	}
	val = liste.Shift()
	if val != "-1" {
		t.Errorf("ERROR: wrong value: got %s, exp %s", val, "-1")
	}
	val = liste.Shift()
	if val != "1" {
		t.Errorf("ERROR: wrong value: got %s, exp %s", val, "1")
	}
	val = liste.Shift()
	if val != "" {
		t.Errorf("ERROR: wrong value: got %s, exp %s", val, "")
	}
	val = liste.Shift()
	if val != "" {
		t.Errorf("ERROR: wrong value: got %s, exp %s", val, "")
	}
	val = liste.Pop()
	if val != "" {
		t.Errorf("ERROR: wrong value: got %s, exp %s", val, "")
	}
}
