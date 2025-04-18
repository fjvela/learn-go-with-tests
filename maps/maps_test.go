package maps

import "testing"

func TestSearch(t *testing.T) {

	dict := map[string]string{"test": "test"}

	got := Search(dict, "test")
	want := "test"

	assert(t, got, want)
}

func assert(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
