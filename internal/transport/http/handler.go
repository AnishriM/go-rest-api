package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/AnishriM/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Handler - stores pointer to our comment service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - an object to store responses from our APIs
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func BasicAuth(orginal func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("Basic auth endpoint hit")
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			orginal(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}
	}
}

// SetupRoutes - sets up all the routes for out application
func (h *Handler) SetupRoutes() {
	logrus.Info("Setting up the routes.")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", BasicAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.UpdateComment)).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.DeleteComment)).Methods("DELETE")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(w, Response{Message: "I'm alive!"}); err != nil {
			panic(err)
		}
	})
}

// GetComment - retrives the comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UNIT from ID", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))

	if err != nil {
		sendErrorResponse(w, "Error retriving comment by ID", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// GetAllComments - retrives comment from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Error retriving all comments", err)
		return
	}
	if err := sendOkResponse(w, comments); err != nil {
		panic(err)
	}
}

// PostComment -  add new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var newcomment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&newcomment); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	comment, err := h.Service.PostComment(newcomment)
	if err != nil {
		sendErrorResponse(w, "Failed to POST the comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "error occurred while paring the id", err)
		return
	}
	var newcomment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&newcomment); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	comment, err := h.Service.UpdateComment(uint(commentId), newcomment)

	if err != nil {
		sendErrorResponse(w, "Failed to update the comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// DeleteComment - Delete comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UNIT from ID", err)
		return
	}

	if err := h.Service.DeleteComment(uint(commentID)); err != nil {
		sendErrorResponse(w, "Failed to delete the comment", err)
		return
	}

	if err := sendOkResponse(w, Response{Message: "Successfully deleted the comment"}); err != nil {
		panic(err)
	}
}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
