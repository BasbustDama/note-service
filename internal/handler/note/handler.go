package noteHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type noteUsecase interface {
	createUsecase
	getUsecase
	getListUsecase
	updateUsecase
	deleteUsecase
}

func New(usecase noteUsecase) http.Handler {
	engine := gin.Default()

	noteGroup := engine.Group("/note")
	{
		noteGroup.POST("/", createNote(usecase))
		noteGroup.GET("/", getListNote(usecase))
		noteGroup.GET("/:id", getNote(usecase))
		noteGroup.PUT("/:id", updateNote(usecase))
		noteGroup.DELETE("/:id", deleteNote(usecase))
	}

	return engine
}
