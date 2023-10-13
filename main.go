package main

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {
	ebiten_main()
	// env := LoadEnv()
	// highlights := readHighlights()
	// books := wrapHighlights(highlights)
	// toJson("highlights.json", highlights)
	// toJson("books.json", books)
	// db := Db{DatabaseId: env["db_id"], ApiSecret: env["notion_api_secret"]}
	// var data DataComb
	// for _, highlight := range highlights {
	// 	data = DataComb{Highlight: highlight, Db: db}
	// 	insertItem(data)
	// 	time.Sleep(100 * time.Millisecond)
	// }
}
