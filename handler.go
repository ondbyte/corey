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
		render.JSON(w, r, err.Error())
		return
	}
	if contact.Email == "" || contact.Name == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, fmt.Errorf("email or name is required").Error())
		return
	}
	err = h.service.AddContact(contact)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, contact)
}

func (h *Handler) GetContact(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	contact, err := h.service.GetContact(uint(_id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, contact)
}

func (h *Handler) GetAllContact(w http.ResponseWriter, r *http.Request) {
	allContact, err := h.service.GetAllContact()
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, allContact)
}

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	task := &Task{}
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	if task.Title == "" || task.Description == "" || task.Reminder == nil || task.ContactID == 0 || task.Priority == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, fmt.Errorf("title, description,reminder, contact_id or priority is required").Error())
		return
	}
	if !IsValidPriority(task.Priority) {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, fmt.Errorf("invalid priority").Error())
		return
	}
	err = h.service.AddTask(task)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, fmt.Errorf("make sure contact exists %v", err).Error())
		return
	}
	render.JSON(w, r, task)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	task, err := h.service.GetTask(uint(_id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, task)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	task, err := h.service.GetTask(uint(_id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	err = h.service.DeleteTask(task)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, task)
}

func (h *Handler) GetAllTask(w http.ResponseWriter, r *http.Request) {
	allTask, err := h.service.GetAllTask()
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, allTask)
}
