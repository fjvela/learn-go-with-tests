package maps

import "testing"

func TestAdd(t *testing.T) {

	t.Run("add new key", func(t *testing.T) {

		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"

		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("add existing key", func(t *testing.T) {
		key := "thekeyexits"
		value := "this is just a test"

		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})

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

func TestUpdate(t *testing.T) {
	key := "test"
	value := "this is just a test"
	dictionary := Dictionary{key: value}
	newValue := "this is a new test"

	dictionary.Update(key, newValue)

	assertDefinition(t, dictionary, key, newValue)
}

func assert(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key string, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added definition:", err)
	}
	assert(t, got, value)
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
