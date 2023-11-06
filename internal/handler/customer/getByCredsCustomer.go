package customerHandler

import (
	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/gin-gonic/gin"
)

type getByCredsUsecase interface {
	GetByCreds(username string, passwordHash string) (entity.Customer, error)
}

func NewGetByCreds(usecase getByCredsUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
