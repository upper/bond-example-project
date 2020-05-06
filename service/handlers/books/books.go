package books

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/upper/bond-example-project/app"
	"github.com/upper/bond-example-project/internal/model"
	"github.com/upper/bond-example-project/service/ws"
)

func getBookCtx(ctx context.Context) (model.Book, bool) {
	book, ok := ctx.Value("book").(model.Book)
	return book, ok
}

func setBookCtx(ctx context.Context, book *model.Book) context.Context {
	return context.WithValue(ctx, "book", *book)
}

func postBook(w http.ResponseWriter, r *http.Request) {
	var post model.Book
	err := ws.Bind(r, &post)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	newBook, err := app.Books(r.Context()).Create(&post)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, newBook)
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	booksPage, err := app.Books(r.Context()).Paginate()
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, booksPage)
}

func bookCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookID, _ := strconv.ParseUint(chi.URLParam(r, "bookID"), 10, 64)

		book, err := app.Books(r.Context()).Get(bookID)
		if err != nil {
			ws.Respond(w, http.StatusNotFound, nil)
			return
		}

		ctx := setBookCtx(r.Context(), book.Book)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getBook(w http.ResponseWriter, r *http.Request) {
	book, _ := getBookCtx(r.Context())
	ws.Respond(w, http.StatusOK, book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	book, _ := getBookCtx(r.Context())

	var patch model.Book
	err := ws.Bind(r, &patch)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	updatedBook, err := app.Books(r.Context()).Update(&book)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, updatedBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	book, _ := getBookCtx(r.Context())

	if err := app.Books(r.Context()).Delete(&book); err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, nil)
}

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", listBooks)
	r.Post("/", postBook)

	r.Route("/{bookID}", func(r chi.Router) {
		r.Use(bookCtx)

		r.Delete("/", deleteBook)
		r.Get("/", getBook)
		r.Post("/", updateBook)
	})

	return r
}
