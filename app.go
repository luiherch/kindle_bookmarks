package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
		cfg.DatabaseId = EncryptDecrypt(XOR_KEY, value)
	case "apiSecret":
		cfg.ApiSecret = EncryptDecrypt(XOR_KEY, value)
	}

	err := toJson(CONFIG_PATH, cfg)
	return err
}

func (cfg *DbConfig) ReadConfig() error {
	err := CheckFilePath(CONFIG_PATH)
	data, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		fmt.Println(err)
	}

	var db DbConfig
	if err := json.Unmarshal(data, &db); err != nil {
		fmt.Println(err)
	}
	cfg.DatabaseId = EncryptDecrypt(XOR_KEY, db.DatabaseId)
	cfg.ApiSecret = EncryptDecrypt(XOR_KEY, db.ApiSecret)
	return err
}

func CheckFilePath(filePath string) error {
	dir := filepath.Dir(filePath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
		fmt.Println("Directory created successfully:", dir)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating file: %v", err)
		}
		defer file.Close()
		fmt.Println("File created successfully:", filePath)
		var cfg DbConfig
		cfg.UpdateConfig("", "")
	}

	return nil
}
