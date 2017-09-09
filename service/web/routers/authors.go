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

type authorBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getAuthorCtx(ctx context.Context) (*app.Author, bool) {
	a, ok := ctx.Value("author").(*app.Author)
	return a, ok
}

func setAuthorCtx(ctx context.Context, a *app.Author) context.Context {
	return context.WithValue(ctx, "author", a)
}

func postAuthor(w http.ResponseWriter, r *http.Request) {
	var data authorBody

	err := ws.Bind(r, &data)
	if err != nil {
		ws.Respond(w, 500, err)
		return
	}

	var author *app.Author
	operation := func(tx bond.Session) error {
		author = app.NewAuthor(&model.Author{
			FirstName: data.FirstName,
			LastName:  data.LastName,
		})

		return tx.Save(author)
	}

	if err := repo.Session.SessionTx(r.Context(), operation); err != nil {
		ws.Respond(w, 500, err)
		return
	}

	ws.Respond(w, 200, author)
}

func listAuthors(w http.ResponseWriter, r *http.Request) {
	var page *ws.Page

	operation := func(sess bond.Session) error {
		var authors []*app.Author
		var err error

		page, err = ws.Paginate(r, app.Authors(sess).Find(), &authors)
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

func authorCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorID := chi.URLParam(r, "authorID")

		var author app.Author
		err := app.Authors(repo.Session).Find(db.Cond{"id": authorID}).One(&author)
		if err != nil {
			ws.Respond(w, 404, nil)
			return
		}

		ctx := setAuthorCtx(r.Context(), &author)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAuthor(w http.ResponseWriter, r *http.Request) {
	author, ok := getAuthorCtx(r.Context())
	if !ok {
		ws.Respond(w, 500, nil)
		return
	}

	ws.Respond(w, 200, author)
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	author, ok := getAuthorCtx(r.Context())
	if !ok {
		ws.Respond(w, 500, nil)
		return
	}

	var data authorBody
	err := ws.Bind(r, &data)
	if err != nil {
		ws.Respond(w, 500, err)
		return
	}

	operation := func(sess bond.Session) error {
		if data.FirstName != "" {
			author.FirstName = data.FirstName
		}

		if data.LastName != "" {
			author.LastName = data.LastName
		}

		return sess.Save(author)
	}

	if err := repo.Session.SessionTx(r.Context(), operation); err != nil {
		ws.Respond(w, 500, err)
		return
	}

	ws.Respond(w, 200, author)
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	author, ok := getAuthorCtx(r.Context())
	if !ok {
		ws.Respond(w, 500, nil)
		return
	}

	operation := func(sess bond.Session) error {
		return sess.Delete(author)
	}

	if err := repo.Session.SessionTx(r.Context(), operation); err != nil {
		ws.Respond(w, 500, err)
		return
	}

	ws.Respond(w, 200, nil)
}

func NewAuthorsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", listAuthors)
	r.Post("/", postAuthor)

	r.Route("/{authorID}", func(r chi.Router) {
		r.Use(authorCtx)

		r.Delete("/", deleteAuthor)
		r.Get("/", getAuthor)
		r.Post("/", updateAuthor)
	})

	return r
}
