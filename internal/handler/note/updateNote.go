package noteHandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
