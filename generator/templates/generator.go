package templates

import (
	htmlTemplate "html/template"
	"os"
	"path"
	"site-generator/parser"
	"slices"
	"strings"
)

func GenerateSite(
	outputPath string,
	posts []parser.Post,
	templates []htmlTemplate.Template) []string {
	const dirMode = 0755 /* Dir */

	namedTemplates := getNamedTemplates(templates)

	nav := getNavigationHeader(posts)

	writtenPosts := make([]string, 0)

	for _, p := range posts {

		postPath := path.Join(outputPath, p.Path+".html")

		dir := path.Dir(postPath)

		os.MkdirAll(dir, dirMode)

		file, err := os.Create(postPath)

		if err != nil {
			panic(err)
		}

		inner := namedTemplates.generatePost(p, posts)

		page := namedTemplates.wrapInLayout(layoutContext{Title: p.Title, HtmlContent: htmlTemplate.HTML(inner), Nav: nav})

		if _, err := file.WriteString(page); err != nil {
			panic(err)
		}

		writtenPosts = append(writtenPosts, postPath)
	}

	return writtenPosts
}

func getNavigationHeader(posts []parser.Post) []string {
	navSet := make(map[string]bool, 0)

	for _, p := range posts {
		navSet[path.Dir(p.Path)] = true
	}

	nav := make([]string, 0)

	for n := range navSet {
		if strings.Contains(n, string(os.PathSeparator)) || n == "." {
			continue
		}
		nav = append(nav, strings.ToLower(n))
	}

	slices.SortFunc(nav, strings.Compare)

	return nav
}

type postContext struct {
	parser.Post

	Children []parser.Post
}

type layoutContext struct {
	Title       string
	Nav         []string
	HtmlContent htmlTemplate.HTML
}

func (nt namedTemplates) generatePost(post parser.Post, posts []parser.Post) string {
	template := nt.post

	if strings.HasSuffix(post.Path, "index") {
		template = nt.index
	}

	builder := new(strings.Builder)

	ctx := postContext{Post: post}

	if ctx.ShowTableOfContents {
		ctx.Children = make([]parser.Post, 0)
		dir := path.Dir(post.Path)

		for _, child := range posts {
			if strings.Contains(child.Path, dir) && child.Path != post.Path {
				ctx.Children = append(ctx.Children, child)
			}
		}
	}

	if err := template.Execute(builder, ctx); err != nil {
		panic(err)
	}

	return builder.String()
}

func (nt namedTemplates) wrapInLayout(context layoutContext) string {
	builder := new(strings.Builder)

	if err := nt.layout.Execute(builder, context); err != nil {
		panic(err)
	}

	return builder.String()
}

type namedTemplates struct {
	layout htmlTemplate.Template
	post   htmlTemplate.Template
	index  htmlTemplate.Template
}

func getNamedTemplates(templates []htmlTemplate.Template) namedTemplates {
	var layoutTemplate *htmlTemplate.Template = nil
	var postTemplate *htmlTemplate.Template = nil
	var indexTemplate *htmlTemplate.Template = nil

	for _, t := range templates {
		switch name := t.Name(); name {
		case "layout.html":
			layoutTemplate = &t
		case "post.html":
			postTemplate = &t
		case "index.html":
			indexTemplate = &t
		}
	}

	if layoutTemplate == nil {
		panic("Layout template cannot be found!")
	}

	if postTemplate == nil {
		panic("Post template cannot be found!")
	}

	if indexTemplate == nil {
		panic("Index template cannot be found!")
	}
	return namedTemplates{
		layout: *layoutTemplate,
		post:   *postTemplate,
		index:  *indexTemplate,
	}
}
