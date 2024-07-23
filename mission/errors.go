package mission

import (
	"errors"
)

var (
	ErrTaskNotFunc        = errors.New("handler must be a function")
	ErrNumberOfParameters = errors.New("number of provided parameters does not match expected")
	ErrTypeOfParameters   = errors.New("type of provided parameters does not match expected")
)
