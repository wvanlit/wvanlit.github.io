package parser

import (
	"bytes"
	"html/template"
	"site-generator/filesystem"
	"strings"

	"github.com/yuin/goldmark"
	goldmark_parser "github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"
	"go.abhg.dev/goldmark/wikilink"
)

type markdownPostParser struct {
	contentPath string
	parser      goldmark.Markdown
}

func (mpp markdownPostParser) Parse(file filesystem.File) (Post, error) {
	postPath := strings.TrimPrefix(file.FilePath, mpp.contentPath)
	postPath = strings.TrimSuffix(postPath, ".md") // TODO based on file type

	ctx := goldmark_parser.NewContext()

	var buf bytes.Buffer
	if err := mpp.parser.Convert(*file.Content, &buf, goldmark_parser.WithContext(ctx)); err != nil {
		return Post{}, err
	}

	d := frontmatter.Get(ctx)

	var meta struct {
		Title           string   `yaml:"title"`
		Tags            []string `yaml:"tags"`
		CreatedAt       string   `yaml:"date created"`
		TableOfContents bool     `yaml:"toc"`
		IsDraft         bool     `yaml:"draft"`
	}

	if err := d.Decode(&meta); err != nil {
		return Post{}, err
	}

	return Post{
		Title:               meta.Title,
		Path:                postPath,
		HtmlContent:         template.HTML(buf.String()),
		Tags:                meta.Tags,
		CreatedAt:           meta.CreatedAt,
		ShowTableOfContents: meta.TableOfContents,
		IsDraft:             meta.IsDraft,
	}, nil
}

func NewMarkdownPostParser(contentPath string) PostParser {
	parser := goldmark.New(
		goldmark.WithExtensions(
			&frontmatter.Extender{},
			&wikilink.Extender{},
		),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)

	return markdownPostParser{
		contentPath: contentPath,
		parser:      parser,
	}
}
