package cerrors

import (
	"errors"
	"fmt"
)

func HttpCallError(err error) error {
	return errors.New(fmt.Sprintf("The HTTP request failed with error %s", err))
}

func UnmarshalError(err error) error {
	return errors.New(fmt.Sprintf("cannot unmarshal data: %s", err))
}

func ApiError(err string) error {
	return errors.New(fmt.Sprintf("FXCM REST error: %s", err))
}

func ExecuteError() error {
	return errors.New(fmt.Sprintf("FXCM cannot execute"))
}
