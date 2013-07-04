package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
    Title string
    Body []byte
}

// to persist the Page
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

// to fetch the page from disk
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}


func main() {
    page, err := loadPage("pandas")
    if err != nil {
        fmt.Printf("%v\n", err)
    } else {
        fmt.Printf("<h1>%s</h1>\n<br><p>%s</p>\n", page.Title, page.Body)
    }
}
