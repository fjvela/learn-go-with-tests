package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 20.0)
	want := 60.0
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 20.0)
	want := 200.0
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
