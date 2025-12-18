package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})
	t.Run("unkown word", func(t *testing.T) {
		_, got := dictionary.Search("unkown")
		if got == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := d.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, d, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{word: definition}
		err := d.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, word, definition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{}

		err := d.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	d := Dictionary{word: definition}
	newDefinition := "new definition"

	d.Update(word, newDefinition)

	assertDefinition(t, d, word, newDefinition)
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		d := Dictionary{word: "test definition"}

		err := d.Delete(word)
		assertError(t, err, nil)

		_, err = d.Search(word)
		assertError(t, err, ErrNotFound)
	})
	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		d := Dictionary{}

		err := d.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

// Assert Helpers

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error: %q, want: %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertString(t, got, definition)
}
