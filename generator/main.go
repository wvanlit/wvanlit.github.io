package main

import (
	"flag"
	"fmt"
	"html/template"
	fs "site-generator/filesystem"
	"site-generator/parser"
)

var contentArg = flag.String("content", "", "Path to content directory")
var templatesArg = flag.String("templates", "", "Path to template directory")
var outputArg = flag.String("output", "", "Path to output directory")
var verboseArg = flag.Bool("verbose", false, "Enable verbose logging")
var serveArg = flag.Bool("serve", false, "Serve the output after generating")
var serverPortArg = flag.Int("port", 8080, "Port for serving content")

func init() {
	flag.Parse()
}

func main() {
	contentFiles := readFiles(*contentArg)

	reportSectionDone(
		"🗃️  found content...",
		contentFiles,
		func(f fs.File) string { return "\t" + f.FilePath },
	)

	posts := parsePosts(contentFiles)

	reportSectionDone(
		"📖 parsed posts...",
		posts,
		func(p parser.Post) string { return fmt.Sprintf("\t%s @ %s", p.Title, p.Path) },
	)

	templateFiles := readFiles(*templatesArg)

	reportSectionDone(
		"🗃️  found templates...",
		templateFiles,
		func(f fs.File) string { return "\t" + f.FilePath },
	)

	templates := parseTemplates(templateFiles)

	reportSectionDone(
		"🩻  parsed templates...",
		templates,
		func(t template.Template) string { return "\t" + t.Name() + t.DefinedTemplates() },
	)

	writtenPosts := writeSite(posts, templates)

	reportSectionDone(
		"🖋️  wrote posts",
		writtenPosts,
		func(s string) string { return s },
	)

	serveSiteIfEnabled()
}
