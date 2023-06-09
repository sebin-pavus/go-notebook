package handler

import (
	"go-notebook/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type HandlerInterface interface {
	CreateUSer(c *gin.Context)
	CreateNewSession(c *gin.Context)
	GetNotes(c *gin.Context)
	CreateNote(c *gin.Context)
	DeleteNote(c *gin.Context)
}

type HandlerStruct struct {
	Users         map[string]model.User
	LoginRegistry map[string]bool
	Notes         map[string][]model.Note
}

func (h HandlerStruct) CreateUser(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	h.Users[user.Email] = user

	c.IndentedJSON(http.StatusOK, nil)
}

func (h HandlerStruct) CreateNewSession(c *gin.Context) {
	var userDetails model.CreateNewSessionRequest

	err := c.ShouldBindJSON(&userDetails)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	validate := validator.New()
	err = validate.Struct(userDetails)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	if userDetails.Password != h.Users[userDetails.Email].Password {
		errResponse := model.ErrorResponse{ErrorMessage: "Invalid username or password"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
		return
	}

	var uuidString string
	for true {
		uuidString = uuid.New().String()
		println(uuidString)
		_, exists := h.LoginRegistry[uuidString]
		if exists {
			continue
		} else {
			h.LoginRegistry[uuidString] = true
			println(uuidString)
			break
		}
	}
	println(uuidString)
	response := model.CreateNewSessionResponse{
		SerialId: uuidString,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (h HandlerStruct) GetNotes(c *gin.Context) {
	var req model.GetNotesRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	_, exists := h.LoginRegistry[req.SerialId]
	if !exists {
		errResponse := model.ErrorResponse{ErrorMessage: "Invalid sid"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
		return
	}

	response := model.GetNotesResponse{
		Notes: h.Notes[req.SerialId],
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (h HandlerStruct) CreateNote(c *gin.Context) {
	var req model.CreateNoteRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	_, exists := h.LoginRegistry[req.SerialId]
	if !exists {
		errResponse := model.ErrorResponse{ErrorMessage: "Invalid sid"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
		return
	}

	length := len(h.Notes[req.SerialId])
	id := uint32(1)
	if length > 0 {
		id = h.Notes[req.SerialId][length-1].Id + 1
	}

	var note model.Note

	note.Id = id
	note.NoteEntry = req.Note

	h.Notes[req.SerialId] = append(h.Notes[req.SerialId], note)

	response := model.CreateNoteResponse{
		Id: id,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (h HandlerStruct) DeleteNote(c *gin.Context) {
	var req model.DeleteNoteRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errResponse := model.ErrorResponse{ErrorMessage: err.Error()}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	_, exists := h.LoginRegistry[req.SerialId]
	if !exists {
		errResponse := model.ErrorResponse{ErrorMessage: "Invalid sid"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
		return
	}

	flag := 0
	// we can improve the time complexity by using binary search since the array is sorted on id
	for i, val := range h.Notes[req.SerialId] {
		if val.Id == req.Id {
			h.Notes[req.SerialId] = append(h.Notes[req.SerialId][:i], h.Notes[req.SerialId][i+1:]...)
			flag = 1
			break
		}
	}

	if flag == 0 {
		errResponse := model.ErrorResponse{ErrorMessage: "Invalid id"}
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
