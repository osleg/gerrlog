package errlog

import (
	"github.com/Sirupsen/logrus"
)

func Err_panic(err error) {
	Err_err(err, logrus.Panicf)
}

func Err_critical(err error) {
	Err_err(err, logrus.Errorf)
}

func Err_warning(err error) {
	Err_err(err, logrus.Warnf)
}

func Err_info(err error) {
	Err_err(err, logrus.Infof)
}

func Err_debug(err error) {
	Err_err(err, logrus.Debugf)
}

func Err_err(err error, cb func(string, ...interface{})) error {
	if err != nil {
		if cb != nil {
			cb("%s", err.Error())
		}
		return err
	}
	return nil
}
