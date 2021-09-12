package main

import (
	"fmt"
	"strings"
	"time"
)

// GoComics Handle the gocomics comics
func GoComics(title string, slug string) {
	now := time.Now().Local()
	date := fmt.Sprintf("%02d/%02d/%02d", now.Year(), now.Month(), now.Day())
	url := fmt.Sprintf("https://www.gocomics.com/%s/%s", slug, date)

	data := FetchURL(url)
	if data == nil {
		return
	}

	text := string(data)
	pos := strings.Index(text, "item-comic-image")
	if pos == -1 {
		return
	}

	pos = IndexAt(text, "src=\"", pos)
	startPos := pos + 5
	endPos := IndexAt(text, "\"", startPos)
	link := text[startPos:endPos]
	HTMLHeader(title, url)
	HTMLAddComic(title, link, "")
	HTMLLineBreak("")
}
