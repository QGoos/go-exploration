package webserver

import (
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get player score", func(t *testing.T) {
		database, closer := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer closer()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33
		AssertScoreEquals(t, got, want)
	})
	t.Run("store wins for an existing player", func(t *testing.T) {
		database, closer := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer closer()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		AssertScoreEquals(t, got, want)
	})
	t.Run("store wins for new players", func(t *testing.T) {
		database, closer := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer closer()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		store.RecordWin("Darius")

		got := store.GetPlayerScore("Darius")
		want := 1
		AssertScoreEquals(t, got, want)
	})
	t.Run("works with an empty file", func(t *testing.T) {
		database, closer := CreateTempFile(t, "")
		defer closer()

		_, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)
	})
	t.Run("league sorted", func(t *testing.T) {
		database, closer := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer closer()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
}
