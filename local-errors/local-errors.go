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
