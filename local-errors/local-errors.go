package local_errors

import (
	"errors"
	"fmt"
)

var NotFoundError = errors.New("object not found")

func NotFound(structName string) error {
	return fmt.Errorf(
		"error, %w: %s", NotFoundError, structName,
	)
}

func NoField(fieldName string) error {
	return fmt.Errorf(
		"no field: %s", fieldName,
	)
}
