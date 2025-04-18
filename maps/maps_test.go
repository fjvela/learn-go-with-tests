package maps

import "testing"

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word in dictionary")
	}
	assert(t, got, want)
}

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
