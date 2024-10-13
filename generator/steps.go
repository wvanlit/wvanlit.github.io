package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	fs "site-generator/filesystem"
	parser "site-generator/parser"
	tmpl "site-generator/templates"
	"strings"
)

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

func parseTemplates(files []fs.File) []template.Template {
	templates, err := tmpl.ParseTemplates(files)

	if err != nil {
		panic(err)
	}

	return templates
}

func writeSite(posts []parser.Post, templates []template.Template) []string {
	const dirMode = 0040000 /* Dir */

	err := os.MkdirAll(*outputArg, dirMode)
	if err != nil {
		panic(err)
	}

	var postTemplate *template.Template = nil
	var indexTemplate *template.Template = nil

	for _, t := range templates {
		switch name := t.Name(); name {
		case "post.html":
			postTemplate = &t
		case "index.html":
			indexTemplate = &t
		}
	}

	if postTemplate == nil {
		panic("Post template cannot be found!")
	}

	if indexTemplate == nil {
		panic("Index template cannot be found!")
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

		if err = template.Execute(file, p); err != nil {
			panic(err)
		}

		writtenPosts = append(writtenPosts, postPath)
	}

	return writtenPosts
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
