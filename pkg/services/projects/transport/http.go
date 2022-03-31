package transport

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"lastimplementation.com/internal/validate"
	"lastimplementation.com/pkg/services/projects"
	"lastimplementation.com/pkg/services/projects/logger"
	"lastimplementation.com/pkg/services/projects/models"
	"lastimplementation.com/pkg/services/projects/store"
)

type handler struct {
	l               logger.Logger
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
	ph := handler{l.WithPrefix("transport"), ps}
	s := r.PathPrefix("/projects").Subrouter()
	s.HandleFunc("", ph.Add).Methods("POST", "OPTIONS")
	s.HandleFunc("", ph.GetAll).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", ph.Get).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", ph.Update).Methods("PATCH", "OPTIONS")
	s.HandleFunc("/{id:[0-9]+}", ph.Delete).Methods("DELETE")
	s.HandleFunc("/{id:[0-9]+}/files", ph.GetFiles).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}/files", ph.UpdateFiles).Methods("PUT", "OPTIONS")
	s.Use(mux.CORSMethodMiddleware(s))
	s.Use(corsAccessHeader)
	s.Use(jsonContentHeader)
}

// Get gets a single project.
func (ph *handler) Add(rw http.ResponseWriter, h *http.Request) {
	ph.l.Trace("add project request started")
	var p models.Project
	if err := p.FromJSON(h.Body); err != nil {
		ph.handleError(err, rw)
		return
	}
	if err := validate.Get().Struct(p); err != nil {
		ph.l.Error("add project", "reading input values", err)
		ph.writeError(rw, http.StatusBadRequest, err)
		return
	}
	pId, err := ph.ProjectsService.Add(context.Background(), p)
	if err != nil {
		ph.handleError(err, rw)
		return
	}
	p.Id = pId
	if err := p.ToJSON(rw); err != nil {
		ph.handleError(err, rw)
	}
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
		ph.writeError(rw, http.StatusBadRequest, err)
		return
	}
	pl, err := ph.ProjectsService.GetAll(context.Background(), qp)
	if err != nil {
		ph.handleError(err, rw)
		return
	}
	if err := pl.ToJSON(rw); err != nil {
		ph.handleError(err, rw)
	}
}

// Get gets a single project.
func (ph *handler) Get(rw http.ResponseWriter, h *http.Request) {
	ph.l.Trace("get project request started")
	id, err := idVar(mux.Vars(h))
	if err != nil {
		ph.writeError(rw, http.StatusBadRequest, err)
		return
	}
	p, err := ph.ProjectsService.Get(context.Background(), id)
	if err != nil {
		ph.handleError(err, rw)
		return
	}
	if err := p.ToJSON(rw); err != nil {
		ph.handleError(err, rw)
	}
}

// Update updates a project.
func (ph *handler) Update(rw http.ResponseWriter, h *http.Request) {
	ph.l.Trace("update project request started")
	id, err := idVar(mux.Vars(h))
	if err != nil {
		ph.writeError(rw, http.StatusBadRequest, err)
		return
	}
	var projectDetails models.ProjectDetails
	if err := projectDetails.FromJSON(h.Body); err != nil {
		ph.handleError(err, rw)
		return
	}
	filled := projectDetails.FilledProps()
	if len(filled) == 0 {
		ph.l.Error("update project", "empty input value", err)
		ph.writeError(rw, http.StatusBadRequest, fmt.Errorf("empty project"))
		return
	}
	if err := validate.Get().StructPartial(projectDetails, filled...); err != nil {
		ph.l.Error("update project", "reading input values", err)
		ph.writeError(rw, http.StatusBadRequest, err)
		return
	}
	if err := ph.ProjectsService.Update(context.Background(), id, projectDetails); err != nil {
		ph.handleError(err, rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

// Delete deletes a project.
func (ph *handler) Delete(rw http.ResponseWriter, h *http.Request) {
	ph.l.Trace("delete project request started")
	id, err := idVar(mux.Vars(h))
	if err != nil {
		ph.writeError(rw, http.StatusBadRequest, err)
		return
	}
	if err := ph.ProjectsService.Delete(context.Background(), id); err != nil {
		ph.handleError(err, rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

// GetFiles writes a list of existing files for a project.
func (ph *handler) GetFiles(rw http.ResponseWriter, h *http.Request) {
	log := ph.l.WithPrefix("get project files")
	log.Trace("request started")
	id, err := idVar(mux.Vars(h))
	if err != nil {
		log.Error("project id", err)
		ph.writeError(rw, http.StatusBadRequest, err)
		return
	}
	files, err := ph.ProjectsService.GetFiles(context.Background(), id)
	if err != nil {
		ph.handleError(err, rw)
		return
	}
	if err := files.ToJSON(rw); err != nil {
		ph.handleError(err, rw)
	}
}

// UpdateFiles updates the current code files.
func (ph *handler) UpdateFiles(rw http.ResponseWriter, h *http.Request) {
	log := ph.l.WithPrefix("update project files")
	log.Trace("request started")
	id, err := idVar(mux.Vars(h))
	if err != nil {
		log.Error("project id", err)
		ph.writeError(rw, http.StatusBadRequest, err)
		return
	}
	var files models.CodeFiles
	if err := files.FromJSON(h.Body); err != nil {
		log.Error("failed to decode body", err)
		ph.writeError(rw, http.StatusBadRequest, projects.ErrDecodeBody)
		return
	}
	if len(files) > models.MaximumCodeFiles {
		err := projects.NewError(fmt.Sprintf("total of code files exceeded the maximum limit (%d)", models.MaximumCodeFiles))
		log.Error(err)
		ph.writeError(rw, http.StatusBadRequest, err)
		return
	}
	for _, file := range files {
		if err := validate.Get().Struct(file); err != nil {
			err := fmt.Errorf("invalid input file %q: %w", file.Name, err)
			log.Error(err)
			ph.writeError(rw, http.StatusBadRequest, err)
			return
		}
	}
	if err := ph.ProjectsService.UpdateFiles(context.Background(), id, files); err != nil {
		ph.handleError(err, rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func idVar(vars map[string]string) (int, error) {
	idv, ok := vars["id"]
	if !ok {
		return -1, nil
	}
	id, err := strconv.Atoi(idv)
	if err != nil {
		return -1, fmt.Errorf("invalid id %d: invalid integer", id)
	}
	if id < 1 {
		return -1, fmt.Errorf("invalid id %d: minimum 1", id)
	}
	return id, nil
}

// handleError allows us to map errors defined internally to appropriate HTTP error codes and JSON responses
func (ph *handler) handleError(err error, rw http.ResponseWriter) {
	outboundErr, ok := err.(projects.OutboundError)
	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	switch outboundErr {
	case projects.ErrProjectTimeout:
		ph.writeResponse(rw, http.StatusRequestTimeout, outboundErr)
	case projects.ErrProjectNotFound:
		ph.writeResponse(rw, http.StatusNotFound, outboundErr)
	case projects.ErrAddProjectDuplicatedName:
		ph.writeResponse(rw, http.StatusBadRequest, outboundErr)
	default:
		ph.writeResponse(rw, http.StatusInternalServerError, outboundErr)
	}
}

func (ph *handler) writeResponse(rw http.ResponseWriter, status int, body models.EncodeJSON) {
	rw.WriteHeader(status)
	if body == nil {
		return
	}
	if err := body.ToJSON(rw); err != nil {
		ph.l.Error("emitting response", "encoding response body", err, fmt.Sprintf("response: [Status %d]: %v", status, body))
	}
}

func (ph *handler) writeError(rw http.ResponseWriter, status int, err error) {
	rw.WriteHeader(status)
	ph.writeResponse(rw, status, projects.NewError(err.Error()))
}
