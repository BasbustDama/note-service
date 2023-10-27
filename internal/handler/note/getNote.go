package noteHandler

import (
	"net/http"
	"strconv"

	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/gin-gonic/gin"
)

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
