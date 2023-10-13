package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"
)

type DataComb struct {
	Highlight
	Db
}
type Db struct {
	DatabaseId string
	ApiSecret  string
}

func insertItem(data DataComb) {
	template_path := "./templates/insert_page.json"
	url := "https://api.notion.com/v1/pages"
	method := "POST"

	t, err := template.ParseFiles(template_path)

	fmt.Println(err)
	fmt.Println(t)
	var s strings.Builder
	err = t.Execute(&s, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	payload := strings.NewReader(s.String())

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	token := "Bearer " + data.ApiSecret
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}