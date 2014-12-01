package main

import "testing"

func TestLight(t *testing.T) {
	l := NewLight(NewVector(1, 2, 3))

	v := l.Position
	if v.x != 1 || v.y != 2 || v.z != 3 {
		t.Errorf("TestLight %v %v %v %v %v %v", 1, 2, 3, v.x, v.y, v.z)
	}
}
