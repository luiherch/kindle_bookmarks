package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Highlight struct {
	BookTitle   string `json:"title"`
	BookAuthor  string `json:"author"`
	Page        int    `json:"page"`
	Position    string `json:"position"`
	AddedOn     string `json:"added_on"`
	Highlighted string `json:"highlighted"`
}

type Book struct {
	BookTitle  string      `json:"title"`
	BookAuthor string      `json:"author"`
	Highlights []Highlight `json:"highlights"`
}

type CompositeKey struct {
	Title  string
	Author string
}

func parseMeta(meta string, h Highlight) Highlight {
	meta_parts := strings.Split(meta, "|")
	length := len(meta_parts)

	if length == 2 {
		re := regexp.MustCompile("[0-9]+")
		page, err := strconv.Atoi(re.FindAllString(meta_parts[0], -1)[0])
		if err != nil {
			fmt.Println("Error during conversion")
			panic("A")
		}
		h.Page = page
		h.AddedOn = strings.TrimSpace(meta_parts[1])
	}
	if length == 3 {
		re := regexp.MustCompile("[0-9]+")
		page, err := strconv.Atoi(re.FindAllString(meta_parts[0], -1)[0])
		if err != nil {
			fmt.Println("Error during conversion")
			panic("A")
		}
		h.Page = page
		h.Position = strings.TrimSpace(meta_parts[1])
		h.AddedOn = strings.TrimSpace(meta_parts[2])
	}
	return h
}

func splitAuthor(t string) (string, string) {
	pattern := `^(.*?)\s*\((.*?)\)$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(t)

	if len(matches) >= 3 {
		title := matches[1]
		author := matches[2]

		return title, author
	} else {
		fmt.Println("No match found")
	}
	return t, ""
}

func wrapHighlights(highlights []Highlight) []Book {
	h_map := make(map[CompositeKey][]Highlight)
	for _, h := range highlights {
		key := CompositeKey{Title: h.BookTitle, Author: h.BookAuthor}
		h_map[key] = append(h_map[key], h)
	}
	var books []Book
	for k, v := range h_map {
		b := Book{BookTitle: k.Title, BookAuthor: k.Author, Highlights: v}
		books = append(books, b)
	}
	return books
}

func readHighlights() []Highlight {
	//clippings_path := "E:\\documents"
	clippings_path := "C:\\Users\\luish\\Documents\\Projects\\kindle_bookmarks"
	files, err := os.ReadDir(clippings_path)
	if err != nil {
		log.Fatal(err)
	}

	var line_delim string = "\r\n"
	var text string

	for _, file := range files {
		if file.Name() == "My Clippings.txt" {
			fmt.Println("Clippings found!")
			dat, _ := os.ReadFile(clippings_path + "\\My Clippings.txt")
			text = string(dat)
		}
	}

	// Remove BOM
	BOM := "\ufeff"
	if strings.HasPrefix(text, BOM) {
		fmt.Println("Removing BOM")
		text = strings.TrimPrefix(text, BOM)
	}

	clips := strings.Split(text, "==========\r\n")

	highlights := make([]Highlight, 0)
	for _, clip := range clips {
		lines := strings.Split(clip, line_delim)
		if len(lines) < 5 {
			continue
		}
		h := Highlight{}
		title, author := splitAuthor(lines[0])
		h.BookTitle = strings.TrimSpace(title)
		h.BookAuthor = strings.TrimSpace(author)
		h.Highlighted = lines[3]
		h = parseMeta(lines[1], h)

		highlights = append(highlights, h)
	}

	return highlights
}
