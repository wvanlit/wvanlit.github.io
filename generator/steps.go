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

	writtenPosts := tmpl.GenerateSite(*outputArg, *draftsArg, posts, templates)

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
