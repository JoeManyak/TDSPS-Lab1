package local_errors

import (
	"errors"
	"fmt"
)

func NotFound(structName string) error {
	return errors.New(
		fmt.Sprintf("error, object not found: %s", structName),
	)
}

func NoField(fieldName string) error {
	return errors.New(fmt.Sprintf("no field: %s", fieldName))
}

func Invalid(fieldName string) error {
	return errors.New(fmt.Sprintf("invalid field: %s", fieldName))
}
