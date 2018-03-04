package main

import (
	"log"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/upper/bond-example-project/service/web/handlers/authors"
	"github.com/upper/bond-example-project/service/web/handlers/books"
	"github.com/upper/bond-example-project/service/web/handlers/subjects"
)

const listenAddr = ":1999"

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Mount("/authors", authors.NewRouter())
	r.Mount("/subjects", subjects.NewRouter())
	r.Mount("/books", books.NewRouter())

	log.Printf("Running server at %s", listenAddr)
	log.Fatal("http.ListenAndServe", http.ListenAndServe(listenAddr, r))
}
