package noteHandler

import (
	"net/http"

	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/gin-gonic/gin"
)

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
