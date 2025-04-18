package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "test"}

	t.Run("know keys", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "test"
		assert(t, got, want)
	})
	t.Run("unknown key", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := "the key does not exist"

		if err == nil {
			t.Errorf("got no error, want %q", want)
		}

		assert(t, err.Error(), want)
	})
}

func assert(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
