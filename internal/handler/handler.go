package handler

import (
	"net/http"
	"strconv"

	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/gin-gonic/gin"
)

type noteUsecase interface {
	Create(title string, description string) (entity.Note, error)
	Get(id int) (entity.Note, error)
	Delete(id int) error
}

func New(usecase noteUsecase) http.Handler {
	engine := gin.Default()

	noteGroup := engine.Group("/note")
	{
		noteGroup.POST("/", createNote(usecase))
		noteGroup.GET("/", getListNote())
		noteGroup.GET("/:id", getNote(usecase))
		noteGroup.PUT("/:id", updateNote())
		noteGroup.DELETE("/:id", deleteNote(usecase))
	}

	return engine
}

type createUsecase interface {
	Create(title string, description string) (entity.Note, error)
}

func createNote(usecase createUsecase) gin.HandlerFunc {
	type requestBody struct {
		Title       string `json:"title" binding:"required,max=255"`
		Description string `json:"description" binding:"omitempty,max=4096"`
	}

	return func(ctx *gin.Context) {
		var request requestBody
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		note, err := usecase.Create(request.Title, request.Description)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, note)
	}
}

type getUsecase interface {
	Get(id int) (entity.Note, error)
}

func getNote(usecase getUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		noteId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		note, err := usecase.Get(noteId)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, note)
	}
}

func getListNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func updateNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

type deleteUsecase interface {
	Delete(id int) error
}

func deleteNote(usecase deleteUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		noteId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		if err := usecase.Delete(noteId); err != nil {
			errorHandler(ctx, err)
			return
		}

		ctx.String(http.StatusOK, "Success")
	}
}
