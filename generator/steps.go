package main

import (
	"fmt"
	htmlTemplate "html/template"
	"net/http"
	"os"
	"os/exec"
	"path"
	fs "site-generator/filesystem"
	parser "site-generator/parser"
	tmpl "site-generator/templates"
	"strings"
)

func reportDone(text string) {
	fmt.Println(text)
}

func reportSectionDone[K any](section string, items []K, sprint func(K) string) {
	fmt.Println(section)

	if *verboseArg {
		fmt.Println()

		for _, i := range items {
			fmt.Println(sprint(i))
		}

		fmt.Println()
	}
}

func readFiles(path string) []fs.File {
	files, err := fs.GetAllFiles(path)

	if err != nil {
		panic(err)
	}

	return files
}

func parsePosts(files []fs.File) []parser.Post {
	var markdownParser = parser.NewMarkdownPostParser(*contentArg)

	posts := make([]parser.Post, 0)

	for _, file := range files {
		if file.FileType == fs.Markdown {
			post, err := markdownParser.Parse(file)

			if err != nil {
				panic(err)
			}

			posts = append(posts, post)
		} else {
			panic("Cannot parse non-markdown files yet!")
		}
	}

	return posts
}

func parseTemplates(files []fs.File) []htmlTemplate.Template {
	templates, err := tmpl.ParseTemplates(files)

	if err != nil {
		panic(err)
	}

	return templates
}

func writeSite(posts []parser.Post, templates []htmlTemplate.Template) []string {
	const dirMode = 0040000 /* Dir */

	err := os.MkdirAll(*outputArg, dirMode)
	if err != nil {
		panic(err)
	}

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

	navSet := make(map[string]bool, 0)

	for _, p := range posts {
		navSet[path.Dir(p.Path)] = true
	}

	nav := make([]string, 0)

	for n := range navSet {
		if strings.Contains(n, string(os.PathSeparator)) || n == "." {
			continue
		}
		nav = append(nav, n)
	}

	writtenPosts := make([]string, 0)

	for _, p := range posts {

		postPath := path.Join(*outputArg, p.Path+".html")

		dir := path.Dir(postPath)

		os.MkdirAll(dir, dirMode)

		file, err := os.Create(postPath)

		if err != nil {
			panic(err)
		}

		template := postTemplate

		if strings.HasSuffix(p.Path, "index") {
			template = indexTemplate
		}

		builder := new(strings.Builder)

		if err = template.Execute(builder, p); err != nil {
			panic(err)
		}

		inner := builder.String()
		builder = new(strings.Builder)

		type layout struct {
			Title       string
			Nav         []string
			HtmlContent htmlTemplate.HTML
		}

		if err = layoutTemplate.Execute(builder, layout{Title: p.Title, HtmlContent: htmlTemplate.HTML(inner), Nav: nav}); err != nil {
			panic(err)
		}

		if _, err := file.WriteString(builder.String()); err != nil {
			panic(err)
		}

		writtenPosts = append(writtenPosts, postPath)
	}

	return writtenPosts
}

func copyStaticFiles() {
	// TODO do not use cp (won't work on windows)
	cmd := exec.Command("cp", "-r", path.Join(*staticArg, ".")+"/.", *outputArg)

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func serveSiteIfEnabled() {
	if *serveArg {
		fmt.Printf("\n💻  serving at http://localhost:%d\n", *serverPortArg)
		http.Handle("/", http.FileServer(http.Dir(*outputArg)))

		err := http.ListenAndServe(fmt.Sprintf(":%d", *serverPortArg), nil)
		if err != nil {
			panic(err)
		}
	}
}
