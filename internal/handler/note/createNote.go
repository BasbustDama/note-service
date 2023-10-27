package noteHandler

import (
	"net/http"

	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/gin-gonic/gin"
)

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
