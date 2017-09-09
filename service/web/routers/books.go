package routers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/upper/bond-example-project/app"
	"github.com/upper/bond-example-project/model"
	"github.com/upper/bond-example-project/repo"
	"github.com/upper/bond-example-project/service/web/ws"
	"net/http"

	"upper.io/bond"
	"upper.io/db.v3"
)

type bookBody struct {
	Title     string `json:"title"`
	AuthorID  uint64 `json:"author_id"`
	SubjectID uint64 `json:"subject_id"`
}

func getBookCtx(ctx context.Context) (*app.Book, bool) {
	a, ok := ctx.Value("book").(*app.Book)
	return a, ok
}

func setBookCtx(ctx context.Context, a *app.Book) context.Context {
	return context.WithValue(ctx, "book", a)
}

func postBook(w http.ResponseWriter, r *http.Request) {
	var data bookBody

	err := ws.Bind(r, &data)
	if err != nil {
		ws.Respond(w, 500, err)
		return
	}

	var book *app.Book
	operation := func(tx bond.Session) error {
		book = app.NewBook(&model.Book{
			Title:     data.Title,
			AuthorID:  data.AuthorID,
			SubjectID: data.SubjectID,
		})

		return tx.Save(book)
	}

	if err := repo.Session.SessionTx(r.Context(), operation); err != nil {
		ws.Respond(w, 500, err)
		return
	}

	ws.Respond(w, 200, book)
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	var page *ws.Page

	operation := func(sess bond.Session) error {
		var books []*app.Book
		var err error

		page, err = ws.Paginate(r, app.Books(sess).Find(), &books)
		if err != nil {
			return err
		}

		return nil
	}

	if err := repo.Session.SessionTx(r.Context(), operation); err != nil {
		ws.Respond(w, 500, err)
		return
	}

	ws.Respond(w, 200, page)
}

func bookCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookID := chi.URLParam(r, "bookID")

		var book app.Book
		err := app.Books(repo.Session).Find(db.Cond{"id": bookID}).One(&book)
		if err != nil {
			ws.Respond(w, 404, nil)
			return
		}

		ctx := setBookCtx(r.Context(), &book)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getBook(w http.ResponseWriter, r *http.Request) {
	book, _ := getBookCtx(r.Context())
	ws.Respond(w, 200, book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	book, _ := getBookCtx(r.Context())

	var data bookBody
	err := ws.Bind(r, &data)
	if err != nil {
		ws.Respond(w, 500, err)
		return
	}

	operation := func(sess bond.Session) error {
		if data.Title != "" {
			book.Title = data.Title
		}

		if data.AuthorID > 0 {
			book.AuthorID = data.AuthorID
		}

		if data.SubjectID > 0 {
			book.SubjectID = data.SubjectID
		}

		return sess.Save(book)
	}

	if err := repo.Session.SessionTx(r.Context(), operation); err != nil {
		ws.Respond(w, 500, err)
		return
	}

	ws.Respond(w, 200, book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	book, _ := getBookCtx(r.Context())

	operation := func(sess bond.Session) error {
		return sess.Delete(book)
	}

	if err := repo.Session.SessionTx(r.Context(), operation); err != nil {
		ws.Respond(w, 500, err)
		return
	}

	ws.Respond(w, 200, nil)
}

func NewBooksRouter() http.Handler {
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
