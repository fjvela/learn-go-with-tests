package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}

	got := Perimeter(rectangle)
	want := 60.0
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}
	got := Area(rectangle)
	want := 200.0
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
