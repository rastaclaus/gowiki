package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is sample page")}
	p1.save()
	p2, err := loadPage("TestPage")
	fmt.Println(p2, err)
}
