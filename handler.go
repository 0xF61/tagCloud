package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate := template.Must(template.ParseFiles("tpl/index.tpl"))
	w.WriteHeader(http.StatusOK)

	data := IndexPageData{PageTitle: "Welcome", Context: "Hello World"}
	indexTemplate.Execute(w, data)
}

func tagHandler(w http.ResponseWriter, r *http.Request) {
	tagTemplate := template.Must(template.ParseFiles("tpl/tag.tpl"))

	var url string
	w.WriteHeader(http.StatusOK)

	url = r.FormValue("url")
	if url == "" {
		url = "https://tr.lipsum.com/"
	}

	fmt.Println(r.FormValue("url"))
	minFrequency := r.FormValue("minFrequency")

	freq, err := strconv.Atoi(minFrequency)
	if err != nil {
		freq = 3
	}

	tags := getPage(url, w, r, freq)

	data := TagPageData{
		PageTitle: "Tags",
		Tags:      tags,
	}

	tagTemplate.Execute(w, data)
}
