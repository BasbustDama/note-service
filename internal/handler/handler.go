package handler

import (
	"net/http"

	noteHandler "github.com/BasbustDama/note-service/internal/handler/note"
	noteUsecase "github.com/BasbustDama/note-service/internal/usecase/note"
	"github.com/gin-gonic/gin"
)

func New(usecase noteUsecase.NoteUsecase) http.Handler {
	engine := gin.Default()

	customerGroup := engine.Group("/auth")
	{
		customerGroup.POST("/sign-up")
		customerGroup.GET("/sign-in")
	}

	noteGroup := engine.Group("/note")
	{
		noteGroup.POST("/", noteHandler.NewCreateNote(usecase))
		noteGroup.GET("/", noteHandler.NewGetListNote(usecase))
		noteGroup.GET("/:id", noteHandler.NewGetNote(usecase))
		noteGroup.PUT("/:id", noteHandler.NewUpdateNote(usecase))
		noteGroup.DELETE("/:id", noteHandler.NewDeleteNote(usecase))
	}

	return engine
}
