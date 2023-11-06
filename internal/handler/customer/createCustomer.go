package customerHandler

import (
	"net/http"

	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/BasbustDama/note-service/internal/handler/errors"
	"github.com/gin-gonic/gin"
)

type createUsecase interface {
	Create(username string, password string) (entity.Customer, error)
}

func NewCreateCustomer(usecase createUsecase) gin.HandlerFunc {
	type requestBody struct {
		Username string `json:"username" binding:"required,max=255"`
		Password string `json:"password" binding:"required,maax=512"`
	}

	return func(ctx *gin.Context) {
		var request requestBody
		if err := ctx.ShouldBind(&request); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		customer, err := usecase.Create(request.Username, request.Password)
		if err != nil {
			errors.ErrorHandler(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, customer)
	}
}
