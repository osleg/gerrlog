package errlog

import (
	"errors"
	"testing"
)

func TestCritical(t *testing.T) {
	err := errors.New("Critical")
	Err_panic(err)
}
