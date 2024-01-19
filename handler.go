package corey

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddContact(w http.ResponseWriter, r *http.Request) {
	contact := &Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	if contact.Email == "" || contact.Name == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, fmt.Errorf("email or name is required"))
		return
	}
	err = h.service.AddContact(contact)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, contact)
}

func (h *Handler) GetContact(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	contact, err := h.service.GetContact(uint(_id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, contact)
}

func (h *Handler) GetAllContact(w http.ResponseWriter, r *http.Request) {
	allContact, err := h.service.GetAllContact()
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, allContact)
}

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	task := &Task{}
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	if task.Title == "" || task.Description == "" || task.Reminder == nil || task.ContactID == 0 {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, fmt.Errorf("title, description,reminder or contact_id, is required"))
		return
	}
	err = h.service.AddTask(task)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, task)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	task, err := h.service.GetTask(uint(_id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, task)
}

func (h *Handler) GetAllTask(w http.ResponseWriter, r *http.Request) {
	allTask, err := h.service.GetAllTask()
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, allTask)
}
