package noteHandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
