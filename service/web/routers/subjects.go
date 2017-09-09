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

type subjectBody struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func getSubjectCtx(ctx context.Context) (*app.Subject, bool) {
	a, ok := ctx.Value("subject").(*app.Subject)
	return a, ok
}

func setSubjectCtx(ctx context.Context, a *app.Subject) context.Context {
	return context.WithValue(ctx, "subject", a)
}

func postSubject(w http.ResponseWriter, r *http.Request) {
	var data subjectBody

	err := ws.Bind(r, &data)
	if err != nil {
		ws.Respond(w, 500, err)
		return
	}

	var subject *app.Subject
	operation := func(tx bond.Session) error {
		subject = app.NewSubject(&model.Subject{
			Name:     data.Name,
			Location: data.Location,
		})

		return tx.Save(subject)
	}

	if err := repo.Session.SessionTx(r.Context(), operation); err != nil {
		ws.Respond(w, 500, err)
		return
	}

	ws.Respond(w, 200, subject)
}

func listSubjects(w http.ResponseWriter, r *http.Request) {
	var page *ws.Page

	operation := func(sess bond.Session) error {
		var subjects []*app.Subject
		var err error

		page, err = ws.Paginate(r, app.Subjects(sess).Find(), &subjects)
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

func subjectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		subjectID := chi.URLParam(r, "subjectID")

		var subject app.Subject
		err := app.Subjects(repo.Session).Find(db.Cond{"id": subjectID}).One(&subject)
		if err != nil {
			ws.Respond(w, 404, nil)
			return
		}

		ctx := setSubjectCtx(r.Context(), &subject)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getSubject(w http.ResponseWriter, r *http.Request) {
	subject, _ := getSubjectCtx(r.Context())
	ws.Respond(w, 200, subject)
}

func updateSubject(w http.ResponseWriter, r *http.Request) {
	subject, _ := getSubjectCtx(r.Context())

	var data subjectBody
	err := ws.Bind(r, &data)
	if err != nil {
		ws.Respond(w, 500, err)
		return
	}

	operation := func(sess bond.Session) error {
		if data.Name != "" {
			subject.Name = data.Name
		}

		if data.Location != "" {
			subject.Location = data.Location
		}

		return sess.Save(subject)
	}

	if err := repo.Session.SessionTx(r.Context(), operation); err != nil {
		ws.Respond(w, 500, err)
		return
	}

	ws.Respond(w, 200, subject)
}

func deleteSubject(w http.ResponseWriter, r *http.Request) {
	subject, _ := getSubjectCtx(r.Context())

	operation := func(sess bond.Session) error {
		return sess.Delete(subject)
	}

	if err := repo.Session.SessionTx(r.Context(), operation); err != nil {
		ws.Respond(w, 500, err)
		return
	}

	ws.Respond(w, 200, nil)
}

func NewSubjectsRouter() http.Handler {
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
