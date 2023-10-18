package handler

import (
	"net/http"

	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/gin-gonic/gin"
)

func errorHandler(ctx *gin.Context, err error) {
	if appErr, ok := err.(entity.AppError); ok {
		switch appErr {
		case entity.ErrorBadRequest:
			ctx.String(http.StatusBadRequest, "Bad Request Error")
			return
		case entity.ErrorNotFound:
			ctx.String(http.StatusNotFound, "Not Found Error")
			return
		case entity.ErrorInternal:
			fallthrough
		default:
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}

	ctx.String(http.StatusInternalServerError, "Internal Server Error")
}
