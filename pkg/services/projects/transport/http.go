package transport

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"lastimplementation.com/pkg/services/projects"
	"lastimplementation.com/pkg/services/projects/logger"
	"lastimplementation.com/pkg/services/projects/models"
	"lastimplementation.com/pkg/services/projects/store"
)

type handler struct {
	l               logger.Logger
	ctx             context.Context
	ProjectsService projects.Service
}

// Activate ...
func Activate(ctx context.Context, r *mux.Router, db *sql.DB, reset bool) {
	// Setup service.
	l := logger.New("projects", true)
	pdb := store.New(l, db)
	ps := projects.New(l, pdb)
	if reset {
		if err := ps.ResetRepo(ctx); err != nil {
			l.Error("resetting the projects service: %v", err)
		}
	}

	// Setup handlers.
	ph := handler{l, ctx, ps}
	s := r.PathPrefix("/projects").Subrouter()
	s.HandleFunc("", ph.GetAll).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", ph.Get).Methods("GET")
	s.Use(mux.CORSMethodMiddleware(s))
	s.Use(corsAccessHeader)
}

// GetAll gets all the projects.
func (ph *handler) GetAll(rw http.ResponseWriter, h *http.Request) {
	ph.l.Trace("get all projects request started")
	qp, err := models.NewSearchQP(
		h.FormValue("q"),
		h.FormValue("page"),
		h.FormValue("limit"),
	)
	if err != nil {
		ph.l.Error("get all projects", "reading form values", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	ps, err := ph.ProjectsService.GetAll(ph.ctx, qp)
	if err != nil {
		rw.WriteHeader(handleError(err))
		return
	}
	json.NewEncoder(rw).Encode(ps)
}

// Get gets a single project.
func (ph *handler) Get(rw http.ResponseWriter, h *http.Request) {
	ph.l.Trace("get project request started")
	vars := mux.Vars(h)
	idv, ok := vars["id"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idv)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := ph.ProjectsService.Get(ph.ctx, id)
	if err != nil {
		rw.WriteHeader(handleError(err))
		return
	}
	json.NewEncoder(rw).Encode(p)
}

// handleError allows us to map errors defined internally to appropriate HTTP error codes and JSON responses
func handleError(e error) int {
	switch e {
	case projects.ErrProjectNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
