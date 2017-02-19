package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := filepath.Walk("src", compileIfMarkdownFile)
	if err != nil {
		panic(err)
	}
}

const markdownExtension = ".md"

func compileIfMarkdownFile(path string, info os.FileInfo, err error) (noError error) {
	if err != nil {
		fmt.Println(err)
		return
	}

	if info.IsDir() {
		return
	}

	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	kohdetiedosto := changeFirstDirectory(path, "html")

	if extension := filepath.Ext(path); extension == markdownExtension {
		kohdetiedosto = kohdetiedosto[:len(kohdetiedosto)-len(markdownExtension)] + ".html"
		fileContents = markdownHTMLläksi(fileContents)
	}

	os.MkdirAll(filepath.Dir(kohdetiedosto), 0777)
	err = ioutil.WriteFile(
		kohdetiedosto,
		fileContents,
		0666,
	)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func changeFirstDirectory(path, to string) string {
	tail := strings.SplitN(filepath.ToSlash(path), "/", 2)[1]
	return filepath.FromSlash(to + "/" + tail)
}
