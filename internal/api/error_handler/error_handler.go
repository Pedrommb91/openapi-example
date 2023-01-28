package error_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Pedrommb91/openapi-example/internal/api/openapi"
	"github.com/Pedrommb91/openapi-example/pkg/errors"
)

func newError(errType, msg, path string, status int) openapi.Error {
	return openapi.Error{
		Error:     errType,
		Message:   msg,
		Path:      path,
		Status:    int32(status),
		Timestamp: time.Now(),
	}
}

func HandleError(ctx *gin.Context, err error) {
	er, ok := err.(*errors.Error)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.AbortWithStatusJSON(er.Kind.Int(), newError(er.Kind.String(), er.Error(), ctx.FullPath(), er.Kind.Int()))
}
