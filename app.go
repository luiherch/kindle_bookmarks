package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type Library struct {
	Books      []Book
	Highlights []Highlight
}

func (l *Library) Sync() {
	// env := LoadEnv()
	// db := DbConfig{DatabaseId: env["db_id"], ApiSecret: env["notion_api_secret"]}
	var cfg DbConfig
	cfg.ReadConfig()
	var data DataComb
	for _, highlight := range l.Highlights {
		data = DataComb{Highlight: highlight, DbConfig: cfg}
		insertItem(data)
		time.Sleep(300 * time.Millisecond)
	}
}

func (cfg *DbConfig) UpdateConfig(field, value string) error {

	fmt.Println("field:" + field)
	fmt.Println("value:" + value)

	switch field {
	case "databaseId":
		cfg.DatabaseId = value
	case "apiSecret":
		cfg.ApiSecret = value
	}

	err := toJson(CONFIG_PATH, cfg)
	return err
}

func (cfg *DbConfig) ReadConfig() error {
	data, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		fmt.Println(err)
	}

	var db DbConfig
	if err := json.Unmarshal(data, &db); err != nil {
		fmt.Println(err)
	}
	cfg.DatabaseId = db.DatabaseId
	cfg.ApiSecret = db.ApiSecret
	return err
}
