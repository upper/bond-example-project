package subjects

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/upper/bond-example-project/app"
	"github.com/upper/bond-example-project/internal/model"
	"github.com/upper/bond-example-project/service/ws"
)

func getSubjectCtx(ctx context.Context) (model.Subject, bool) {
	subject, ok := ctx.Value("subject").(model.Subject)
	return subject, ok
}

func setSubjectCtx(ctx context.Context, subject *model.Subject) context.Context {
	return context.WithValue(ctx, "subject", *subject)
}

func postSubject(w http.ResponseWriter, r *http.Request) {
	var post model.Subject
	err := ws.Bind(r, &post)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	newSubject, err := app.Subjects(r.Context()).Create(&post)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, newSubject)
}

func listSubjects(w http.ResponseWriter, r *http.Request) {
	subjectsPage, err := app.Subjects(r.Context()).Paginate()
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, subjectsPage)
}

func subjectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		subjectID, _ := strconv.ParseUint(chi.URLParam(r, "subjectID"), 10, 64)

		subject, err := app.Subjects(r.Context()).Get(subjectID)
		if err != nil {
			ws.Respond(w, http.StatusNotFound, nil)
			return
		}

		ctx := setSubjectCtx(r.Context(), subject.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getSubject(w http.ResponseWriter, r *http.Request) {
	subject, _ := getSubjectCtx(r.Context())
	ws.Respond(w, http.StatusOK, subject)
}

func updateSubject(w http.ResponseWriter, r *http.Request) {
	subject, _ := getSubjectCtx(r.Context())

	var patch model.Subject
	err := ws.Bind(r, &patch)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	updatedSubject, err := app.Subjects(r.Context()).Update(&subject, &patch)
	if err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, updatedSubject)
}

func deleteSubject(w http.ResponseWriter, r *http.Request) {
	subject, _ := getSubjectCtx(r.Context())

	if err := app.Subjects(r.Context()).Delete(&subject); err != nil {
		ws.Respond(w, http.StatusInternalServerError, err)
		return
	}

	ws.Respond(w, http.StatusOK, nil)
}

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", listSubjects)
	r.Post("/", postSubject)

	r.Route("/{subjectID}", func(r chi.Router) {
		r.Use(subjectCtx)

		r.Delete("/", deleteSubject)
		r.Get("/", getSubject)
		r.Post("/", updateSubject)
	})

	return r
}
