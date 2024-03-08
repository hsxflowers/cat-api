package exceptions

import (
	"errors"
	"net/http"
)

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandleException(err error) ErrResponse {
	customErr, ok := err.(*Error)
	if !ok {
		return ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	switch err != nil {
	case
		errors.Is(customErr.CustomErr, ErrCatIdIsRequired),
		errors.Is(customErr.CustomErr, ErrTagIsRequired),
		errors.Is(customErr.CustomErr, ErrUrlIsNotValid),
		errors.Is(customErr.CustomErr, ErrTagIsNotValid),
		errors.Is(customErr.CustomErr, ErrBadRequest):
		return ErrResponse{
			Code:    http.StatusBadRequest,
			Message: customErr.CustomErr.Error(),
		}
	case
		errors.Is(customErr.CustomErr, ErrCreateCatInDB),
		errors.Is(customErr.CustomErr, ErrGetCatInDB),
		errors.Is(customErr.CustomErr, ErrListCatsInDB),
		errors.Is(customErr.CustomErr, ErrUpdateCatInDB),
		errors.Is(customErr.CustomErr, ErrDeleteCatInDB),
		errors.Is(customErr.CustomErr, ErrBindDataOnCreateCat),
		errors.Is(customErr.CustomErr, ErrBindDataOnUpdateCat):
		return ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: customErr.CustomErr.Error(),
		}
	case
		errors.Is(customErr.CustomErr, ErrBadData):
		return ErrResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: customErr.CustomErr.Error(),
		}
	case
		errors.Is(customErr.CustomErr, ErrCatNotFound):
		return ErrResponse{
			Code:    http.StatusNotFound,
			Message: customErr.CustomErr.Error(),
		}
	default:
		return ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		}
	}
}
