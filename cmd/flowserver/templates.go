package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func readFSFile(fs http.FileSystem, name string) (string, error) {
	file, err := fs.Open(name)
	if err != nil {
		return "", err
	}
	text, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(text), nil
}

func indexHandler(fs http.FileSystem, w http.ResponseWriter, r *http.Request) error {
	startFlow, err := readFSFile(fs, "/start.json")
	if err != nil {
		fmt.Printf("start doesn't exit\n")
		return err
	}
	indented := bytes.Buffer{}
	json.Indent(&indented, []byte(startFlow), "", "  ")

	indexTpl, err := readFSFile(fs, "/index.html")
	if err != nil {
		fmt.Printf("index doesn't exit\n")
		return err
	}

	t := template.New("index.html")
	t, err = t.Parse(indexTpl)
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, "index.html", indented.String())
}
