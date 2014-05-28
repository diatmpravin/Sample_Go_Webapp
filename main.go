package main

import (
    "net/http"
	"fmt"
	"io/ioutil"
	"html/template"
)

type File struct {
	Name string
	Body []byte
}

func (f File) save() error {
	name := f.Name + ".txt"
	return ioutil.WriteFile(name, f.Body, 0666)
}

func loadData(name string) (*File, error) {
	filename := name + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &File{Name: name, Body: body}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Path[len("/edit/"):]
    p, err := loadData(name)
	if err!= nil {
		p = &File{Name: name}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &File{Name: name, Body: []byte(body)}
	p.save()
    http.Redirect(w, r, "/view/"+name , http.StatusFound)
}

func viewHanlder(w http.ResponseWriter, r *http.Request){
	name := r.URL.Path[len("/view/"):]
	p, err := loadData(name)
	if err!=nil {
		fmt.Println("Error: ", err)
	}
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/view/", viewHanlder)
	http.HandleFunc("/save/", saveHandler)
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Println("Error: ", err)
	}
}
