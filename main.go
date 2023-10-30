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

func main() {

	highlights := readHighlights()
	books := wrapHighlights(highlights)
	library := Library{Books: books, Highlights: highlights}
	var dbconfig DbConfig
	dbconfig.ReadConfig()
	fmt.Println("Hola Luis")
	fmt.Println(dbconfig)

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

	fmt.Println(largerTemplateStr)

	os.WriteFile("frontend/highlights.html", []byte(largerTemplateStr), os.ModePerm)

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
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
		println("Error:", err.Error())
	}
}
