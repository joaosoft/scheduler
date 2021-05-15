package models

import (
	"github.com/joaosoft/errors"
	"github.com/joaosoft/web"
)

var (
	ErrorNotFound    = errors.New(errors.LevelError, web.StatusNotFound, "user not found")
	ErrorInvalidType = errors.New(errors.LevelError, web.StatusNotFound, "invalid type")
)
