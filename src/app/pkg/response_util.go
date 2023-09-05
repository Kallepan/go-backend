package pkg

import (
	"github.com/kallepan/go-backend/app/constant"
	"github.com/kallepan/go-backend/app/domain/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.APIResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status int, message string, data T) dto.APIResponse[T] {
	return dto.APIResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
