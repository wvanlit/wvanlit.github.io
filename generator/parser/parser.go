package parser

import (
	"html/template"
	"site-generator/filesystem"
)

type Post struct {
	Title               string
	Path                string
	HtmlContent         template.HTML // TODO replace with generic ast?
	Tags                []string
	CreatedAt           string
	ShowTableOfContents bool
}

type PostParser interface {
	Parse(file filesystem.File) (Post, error)
}
