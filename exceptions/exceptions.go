package exceptions

import "errors"

// Cat errors
var ErrCatAlreadyExists = errors.New("cat: cat already exists")
var ErrCatIdIsRequired = errors.New("cat: cat_id is required and must be at least 3 characters")
var ErrTagIsRequired = errors.New("cat: tag is required and must be at least 3 characters")
var ErrUrlIsNotValid = errors.New("cat: url is not a valid URL")
var ErrTagIsNotValid = errors.New("cat: tag is not valid")
var ErrCatNotFound = errors.New("cat: cat not found")
var ErrBadData = errors.New("cat: unprocessable json")
var ErrBadRequest = errors.New("cat: can't update without valid field")
var ErrMissingField = errors.New("cat: can't create without valid field")
var ErrInternalServer = errors.New("cat: internal server error")

// Bind errors
var ErrBindDataOnCreateCat = errors.New("cat: error on bind cat request when creating cat")
var ErrBindDataOnUpdateCat = errors.New("cat: error on bind cat request when updating cat")

// DB errors
var ErrCreateCatInDB = errors.New("cat: error creating cat in the database")
var ErrUpdateCatInDB = errors.New("cat: error updating cat in the database")
var ErrDeleteCatInDB = errors.New("cat: error deleting cat in the database")
var ErrGetCatInDB = errors.New("cat: error getting cat in the database")
var ErrListCatsInDB = errors.New("cat: error listing cats in the database")