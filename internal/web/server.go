package web

import (
	"go-notebook/internal/model"
	"go-notebook/internal/web/handler"

	"github.com/gin-gonic/gin"
)

func NewServer(router *gin.Engine) {
	newHandler := handler.HandlerStruct{Users: make(map[string]model.User), LoginRegistry: make(map[string]bool), Notes: make(map[string][]model.Note)}
	router.POST("/signup", newHandler.CreateUser)
	router.POST("/login", newHandler.CreateNewSession)
	router.GET("/notes", newHandler.GetNotes)
	router.POST("/notes", newHandler.CreateNote)
	router.DELETE("/notes", newHandler.DeleteNote)
}
