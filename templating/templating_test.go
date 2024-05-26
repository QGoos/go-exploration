package templating

import (
	"bytes"
	"exploration/readingfiles"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRenderFiles(t *testing.T) {
	var (
		aPost = readingfiles.Post{
			Title:       "Sorry, Eh",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"sorry", "eh"},
		}
	)

	postRenderer, err := readingfiles.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("convert single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("render an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []readingfiles.Post{{Title: "Sorry, Eh"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = readingfiles.Post{
			Title:       "Sorry, Eh",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"sorry", "eh"},
		}
	)

	postRenderer, err := readingfiles.NewPostRenderer()

	if err != nil {
		b.Fatal()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}

func assertEqual(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
