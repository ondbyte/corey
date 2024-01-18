package corey

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddContact(c *gin.Context) {
	contact := &Contact{}
	err := c.Bind(contact)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, c.Errors)
		return
	}
	err = h.service.AddContact(contact)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, c.Errors)
		return
	}
	c.JSON(http.StatusOK, contact)
}
func (h *Handler) GetContact(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		c.Error(fmt.Errorf("id should be a valid uint"))
		return
	}
	contact, err := h.service.GetContact(uint(_id))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, c.Errors)
		return
	}
	c.JSON(http.StatusOK, contact)
}

func (h *Handler) GetAllContact(c *gin.Context) {
	allContact, err := h.service.GetAllContact()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, c.Errors)
		return
	}
	c.JSON(http.StatusOK, allContact)
}

func (h *Handler) AddTask(c *gin.Context) {
	task := &Task{}
	err := c.Bind(task)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, c.Errors)
		return
	}
	err = h.service.AddTask(task)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, c.Errors)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *Handler) GetTask(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		c.Error(fmt.Errorf("id should be a valid uint"))
		return
	}
	task, err := h.service.GetTask(uint(_id))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, c.Errors)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *Handler) GetAllTask(c *gin.Context) {
	allTasks, err := h.service.GetAllTask()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, c.Errors)
		return
	}
	c.JSON(http.StatusOK, allTasks)
}
