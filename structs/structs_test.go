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

	t.Run("Area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		got := Area(rectangle)
		want := 200.0
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("Area of circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := Area(circle)
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
