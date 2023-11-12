package main

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS
var CONFIG_PATH = "configs/data.json"
var XOR_KEY = "ShDRq+Xd1qZrUElVi6T36zpw/PXMHfdO"

func main() {

	var highlights []Highlight
	var books []Book
	var dbconfig DbConfig

	kindle_path, err := getKindlePath()
	if err != nil {
		fmt.Println("No Kindle detected")
	} else {
		highlights = readHighlights(kindle_path)
		books = wrapHighlights(highlights)
	}
	library := Library{Books: books, Highlights: highlights}

	dbconfig.ReadConfig()

	templateStr := `<details>
		<summary>{{.BookTitle}}</summary>
		{{range .Highlights}}
		<article>
			<cite>{{.Highlighted}}</cite>
			<footer>{{.BookAuthor}}</footer>
		</article>
		{{end}}
	</details>`

	t, _ := template.New("book").Parse(templateStr)
	var generatedHTML bytes.Buffer

	for _, item := range books {
		err := t.Execute(&generatedHTML, item)
		if err != nil {
			log.Fatal(err)
		}
	}

	parent_html, _ := os.ReadFile("templates/highlights_template.html")
	largerTemplateStr := strings.Replace(string(parent_html), "<!-- Insert html here -->", generatedHTML.String(), 1)

	os.WriteFile("frontend/highlights.html", []byte(largerTemplateStr), os.ModePerm)

	app := NewApp()

	err = wails.Run(&options.App{
		Title:  "Kindle Bookmarks",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			&library,
			&dbconfig,
		},
	})

	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}
