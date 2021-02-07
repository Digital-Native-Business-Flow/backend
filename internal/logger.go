package internal

import (
	"github.com/sirupsen/logrus"
)

const (
	infoColor  = 36
	errorColor = 31
)

func logDebugErrors(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func LogRequestResponse(code int, method, path, err string) {
	switch code {
	case 200:
		logrus.Infof("| \u001B[%d;1m%d\u001b[0m | %s %s", infoColor, code, method, path)
	default:
		logrus.Errorf("| \u001B[%d;1m%d\u001b[0m | %s %s | %s", errorColor, code, method, path, err)
	}
}
