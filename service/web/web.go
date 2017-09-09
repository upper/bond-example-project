package main

import (
	"github.com/go-chi/chi"
	"github.com/upper/bond-example-project/service/web/routers"
	"log"
	"net/http"
)

const listenAddr = ":1999"

func main() {
	r := chi.NewRouter()

	r.Mount("/authors", routers.NewAuthorsRouter())
	r.Mount("/subjects", routers.NewSubjectsRouter())
	r.Mount("/books", routers.NewBooksRouter())

	log.Printf("Running server at %s", listenAddr)
	log.Fatal("http.ListenAndServe", http.ListenAndServe(listenAddr, r))
}
