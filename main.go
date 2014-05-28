package main

import (
    "net/http"
	"fmt"
	"io/ioutil"
)

type File struct {
	Name string
	Body []byte
}

func loadData(title string) (*File, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &File{Name: title, Body: body}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadData(title)
	if err!= nil {
		p = &File{Name: title}
	}
	fmt.Fprintf(w, string(p.Name))
	fmt.Fprintf(w, string(p.Body))
}

func main() {
	http.HandleFunc("/edit/", editHandler)
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Println("Error: ", err)
	}
}
