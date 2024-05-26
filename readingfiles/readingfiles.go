package readingfiles

import (
	"bufio"
	"bytes"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tagSeperator         = "Tags: "
	bodySeperator        = "---\n"
)

type StubFailingFS struct {
}

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func (p Post) SanatisedTitle() string {
	return strings.ToLower(strings.Replace(strings.Replace(p.Title, ",", "", -1), " ", "-", -1))
}

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no! i always fail")
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	titleLine := readMetaLine(titleSeperator)
	descriptionLine := readMetaLine(descriptionSeperator)
	tagLine := readMetaLine(tagSeperator)

	bodyLines := readBody(scanner)

	post := Post{
		Title:       titleLine,
		Description: descriptionLine,
		Tags:        strings.Split(tagLine, ", "),
		Body:        bodyLines,
	}
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}
