package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func toJson(file_path string, obj interface{}) {
	b, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	os.WriteFile(file_path, b, os.ModePerm)
}

func LoadEnv(filenames ...string) map[string]string {
	data, _ := os.ReadFile(".env")
	s := string(data)
	lines := strings.Split(s, "\n")
	m := make(map[string]string)
	for _, line := range lines {
		c := strings.Split(line, "=")
		m[c[0]] = strings.TrimSpace(c[1])
	}
	return m
}
