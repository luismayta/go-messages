package main

import (
	"github.com/russross/blackfriday"
	"net/http"
)

func main() {
	http.HandleFunc("/markdown", GeneraDesdeMarkdown)
	http.Handle("/", http.FileServer(http.Dir("publico")))
	http.ListenAndServe(":8080", nil)
}

func GeneraDesdeMarkdown(rw http.ResponseWriter, r *http.Request) {
	html := blackfriday.MarkdownCommon([]byte(r.FormValue("cuerpo")))
	rw.Write(html)
}
