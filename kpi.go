package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".html"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {

	var filename string

	ext := strings.Split(title, ".")
	if len(ext) > 1 {
		filename = title
	} else {
		filename = title + ".html"
	}

	log.Printf("Loading page %v \n", filename)

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Page %v title not found => %v\n", title, err)
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/view/"):]
	log.Println("Handling title:", title)

	p, _ := loadPage(title)
	if p != nil {
		fmt.Fprintf(w, "%s", p.Body)
		return
	}

	fmt.Fprintf(w, "<h1>404</h1><div>Page not found</div>")

}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
