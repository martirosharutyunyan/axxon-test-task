package response

import (
	"github.com/gin-gonic/gin"
	httpErrors "github.com/martirosharutyunyan/axxon-test-task/pkg/common/http-errors"
	"net/http"
)

type Context struct {
	*gin.Context
}

type handlerFuncType func(ctx *Context)

func Handler(handler handlerFuncType) func(*gin.Context) {
	return func(context *gin.Context) {
		ctx := &Context{context}
		handler(ctx)
	}
}

func (ctx *Context) GetUser() (any, bool) {
	user, exists := ctx.Get("user")

	if !exists {
		return nil, false
	}

	return user, true
}

func (ctx *Context) StatusOk() {
	ctx.Status(http.StatusOK)
}

func (ctx *Context) Error(err error) {
	if httpError, ok := err.(*httpErrors.HTTPError); err != nil && ok {
		ctx.AbortWithStatusJSON(httpError.StatusCode, httpError)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusInternalServerError,
		httpErrors.NewHTTPError(http.StatusInternalServerError, "Error is not httpError type please fix"),
	)
}

func (ctx *Context) BadRequest(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusBadRequest, message))
}

func (ctx *Context) Conflict(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusConflict, message))
}

func (ctx *Context) Internal(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusInternalServerError, message))
}

func (ctx *Context) NotFound(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusNotFound, message))
}

func (ctx *Context) UnprocessableEntity(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusUnprocessableEntity, message))
}

func (ctx *Context) Forbidden(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusForbidden, message))
}

func (ctx *Context) Unauthorized(message string) {
	ctx.Error(httpErrors.NewHTTPError(http.StatusUnauthorized, message))
}
