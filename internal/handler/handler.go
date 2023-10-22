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
	GetList(offset int, limit int) ([]entity.Note, int, error)
	Update(id int, title *string, description *string) error
	Delete(id int) error
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
		if err := ctx.ShouldBind(&request); err != nil {
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

type getListUsecase interface {
	GetList(offset int, limit int) ([]entity.Note, int, error)
}

func getListNote(usecase getListUsecase) gin.HandlerFunc {
	type requestParams struct {
		Offset int `form:"offset" binding:"omitempty,gte=1"`
		Limit  int `form:"limit" binding:"required,oneof=10 20 50"`
	}

	return func(ctx *gin.Context) {
		var request requestParams
		if err := ctx.ShouldBindQuery(&request); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		noteList, count, err := usecase.GetList(request.Offset, request.Limit)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":  noteList,
			"count": count,
		})
	}
}

type updateUsecase interface {
	Update(id int, title *string, description *string) error
}

func updateNote(usecase updateUsecase) gin.HandlerFunc {
	type request struct {
		ID          int     `json:"-" binding:"required,gte=1"`
		Title       *string `json:"title" binding:"omitempty,min=1,max=255"`
		Description *string `json:"description" binding:"omitempty,max=4096"`
	}

	return func(ctx *gin.Context) {
		noteId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		request := request{ID: noteId}
		if err := ctx.ShouldBind(&request); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		if request.Title == nil && request.Description == nil {
			ctx.String(http.StatusBadRequest, "Title and description is null")
			return
		}

		err = usecase.Update(request.ID, request.Title, request.Description)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		ctx.String(http.StatusOK, "Success")
	}
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
