package errlog

import (
	"github.com/Sirupsen/logrus"
)

// ErrPanic Panic with given error
func ErrPanic(err error) {
	ErrErr(err, logrus.Panicf)
}

// ErrCritical Output critical with given error
func ErrCritical(err error) {
	ErrErr(err, logrus.Errorf)
}

// ErrWarning Output warning with given error
func ErrWarning(err error) {
	ErrErr(err, logrus.Warnf)
}

// ErrInfo Output info with given error
func ErrInfo(err error) {
	ErrErr(err, logrus.Infof)
}

// ErrDebug Output debug with given error
func ErrDebug(err error) {
	ErrErr(err, logrus.Debugf)
}

// ErrErr Output err with given error
func ErrErr(err error, cb func(string, ...interface{})) error {
	if err != nil {
		if cb != nil {
			cb("%s", err.Error())
		}
		return err
	}
	return nil
}
