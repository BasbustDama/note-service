package noteHandler

import (
	"net/http"
	"strconv"

	"github.com/BasbustDama/note-service/internal/handler/errors"
	"github.com/gin-gonic/gin"
)

type deleteUsecase interface {
	Delete(id int) error
}

func NewDeleteNote(usecase deleteUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		noteID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		if err := usecase.Delete(noteID); err != nil {
			errors.ErrorHandler(ctx, err)
			return
		}

		ctx.String(http.StatusOK, "Success")
	}
}
