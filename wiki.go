package main

import (
	//"fmt"
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


func main() {
    page := &Page{"cats", []byte("an amusing organism\n")}
    page.save()    
}
