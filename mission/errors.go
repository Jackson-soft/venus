package mission

import "fmt"

var (
	ErrTaskNotFunc        = fmt.Errorf("handler must be a function")
	ErrNumberOfParameters = fmt.Errorf("number of provided parameters does not match expected")
	ErrTypeOfParameters   = fmt.Errorf("type of provided parameters does not match expected")
)
