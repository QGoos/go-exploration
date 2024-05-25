package readingfiles

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestReadingFiles(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Sorry
eh`
		secondBody = `Title: Post 2
Description: Description 2
Tags: python, pyspark
---
Hey
yo`
	)
	t.Run("Read Files", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		assertEqualLen(t, posts, fs)
	})

	t.Run("New Blog Posts", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		want := Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body:        "Sorry\neh",
		}

		assertPost(t, got, want)

	})
}

func assertEqualLen(t testing.TB, got []Post, want fstest.MapFS) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("Got %d but wanted %d", len(got), len(want))
	}
}

func assertPost(t *testing.T, got, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
