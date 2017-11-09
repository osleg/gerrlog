package errlog

import (
	"github.com/Sirupsen/logrus"
)

// Panic with given error
func Err_panic(err error) {
	Err_err(err, logrus.Panicf)
}

// Output critical with given error
func Err_critical(err error) {
	Err_err(err, logrus.Errorf)
}

// Output warning with given error
func Err_warning(err error) {
	Err_err(err, logrus.Warnf)
}

// Output info with given error
func Err_info(err error) {
	Err_err(err, logrus.Infof)
}

// Output debug with given error
func Err_debug(err error) {
	Err_err(err, logrus.Debugf)
}

// Output err with given error
func Err_err(err error, cb func(string, ...interface{})) error {
	if err != nil {
		if cb != nil {
			cb("%s", err.Error())
		}
		return err
	}
	return nil
}
