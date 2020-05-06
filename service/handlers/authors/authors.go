package authors

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/upper/bond-example-project/app"
	"github.com/upper/bond-example-project/internal/model"
	"github.com/upper/bond-example-project/service/ws"
)

func getAuthorCtx(ctx context.Context) (model.Author, bool) {
	author, ok := ctx.Value("author").(model.Author)
	return author, ok
}

func setAuthorCtx(ctx context.Context, author *model.Author) context.Context {
	return context.WithValue(ctx, "author", *author)
}

func postAuthor(w http.ResponseWriter, r *http.Request) {
	var post model.Author
	err := ws.Bind(r, &post)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	newAuthor, err := app.Authors(r.Context()).Create(&post)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, newAuthor)
}

func listAuthors(w http.ResponseWriter, r *http.Request) {
	authorsPage, err := app.Authors(r.Context()).Paginate()
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, authorsPage)
}

func authorCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorID, _ := strconv.ParseUint(chi.URLParam(r, "authorID"), 10, 64)

		author, err := app.Authors(r.Context()).Get(authorID)
		if err != nil {
			ws.Respond(w, http.StatusNotFound, nil)
			return
		}

		ctx := setAuthorCtx(r.Context(), author.Author)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAuthor(w http.ResponseWriter, r *http.Request) {
	author, _ := getAuthorCtx(r.Context())
	ws.Respond(w, http.StatusOK, author)
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	author, _ := getAuthorCtx(r.Context())

	var patch model.Author
	err := ws.Bind(r, &patch)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	updatedAuthor, err := app.Authors(r.Context()).Update(&author)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, updatedAuthor)
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	author, _ := getAuthorCtx(r.Context())

	if err := app.Authors(r.Context()).Delete(&author); err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, nil)
}

func NewRouter() http.Handler {
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
