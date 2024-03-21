package main 

import (
	"fmt"
	"os"
	"net/http"
	"log"
)

type Page struct {
	Title string
	Body []byte
}

func (page *Page) save() error {
	filename := page.Title + ".txt"
	return os.WriteFile(filename, page.Body, 0600)
}

func loadPage(title string) (*Page, error) { 
	filename := title + ".txt"
	body, err := os.ReadFile(filename) 

	if err != nil {
		return nil, err
	}

	return &Page{Title :title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

