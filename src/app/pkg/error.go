package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/kallepan/go-backend/app/constant"

	"github.com/gin-gonic/gin"
)

func PanicException_(key int, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%d: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func PanicHandler(ctx *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		keyStr := strArr[0]
		key, _ := strconv.Atoi(keyStr)
		msg := strings.Trim(strArr[1], " ")

		switch key {
		case
			constant.DataNotFound.GetResponseStatus():
			ctx.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			ctx.Abort()
		case
			constant.Unauthorized.GetResponseStatus():
			ctx.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
			ctx.Abort()
		default:
			ctx.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
			ctx.Abort()
		}
	}
}
