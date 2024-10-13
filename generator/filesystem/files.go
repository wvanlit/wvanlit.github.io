package filesystem

import (
	"log"
	"os"
	"path"
	"strings"
)

type fileType string

const (
	Unknown  fileType = "unknown"
	Markdown fileType = "md"
	Html     fileType = "html"
)

var knownFileTypes = []fileType{Markdown, Html}

var foldersToIgnore = map[string]bool{
	".obsidian": true,
}

type File struct {
	FilePath string
	FileType fileType
	Content  *[]byte
}

func readFile(filepath string) *[]byte {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return &content
}

func GetAllFiles(dir string) ([]File, error) {
	entries, err := os.ReadDir(dir)

	files := make([]File, 0)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		path := path.Join(dir, entry.Name())

		if entry.IsDir() {
			if foldersToIgnore[entry.Name()] {
				continue
			}

			nestedFiles, err := GetAllFiles(path)

			if err != nil {
				return nil, err
			}

			files = append(files, nestedFiles...)
		} else {
			fileparts := strings.Split(entry.Name(), ".")
			parsedFiletype := fileparts[len(fileparts)-1]

			filetype := Unknown

			for _, ft := range knownFileTypes {
				if ft == fileType(parsedFiletype) {
					filetype = ft
					break
				}
			}

			var content *[]byte = nil

			if filetype != Unknown {
				content = readFile(path)
			}

			files = append(files, File{
				FilePath: path,
				FileType: filetype,
				Content:  content,
			})
		}
	}

	return files, nil
}
