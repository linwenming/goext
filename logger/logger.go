package logger

import (
	"github.com/sirupsen/logrus"
)

//Logger represents  the log interface
type Logger interface {
	Fatal(format ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
}

var (
	std Logger
	//slog *logrus.Logger
)

func InitLogger(path string) {
	plog := logrus.New()
	plog.Formatter = new(logrus.TextFormatter)
	plog.Level = logrus.DebugLevel

	output, err := New(path)
	if err == nil {
		plog.SetOutput(output)
	}

	log := plog.WithFields(logrus.Fields{
		"source": "app",
	})
	std = log
}

// SetLogger rewrites the default logger
func SetLogger(l Logger) {
	if l != nil {
		std = l
	}
}

func Fatal(args ...interface{}) {
	if std == nil {
		logrus.Fatal(args)
	} else {
		std.Fatal(args)
	}
}
func Fatalf(format string, args ...interface{}) {
	if std == nil {
		logrus.Fatalf(format, args)
	} else {
		std.Fatalf(format, args)
	}

}
func Fatalln(args ...interface{}) {
	if std == nil {
		logrus.Fatalln(args)
	} else {
		std.Fatalln(args)
	}
}

func Debug(args ...interface{}) {
	if std == nil {
		logrus.Debug(args)
	} else {
		std.Debug(args)
	}
}
func Debugf(format string, args ...interface{}) {
	if std == nil {
		logrus.Debugf(format, args)
	} else {
		std.Debugf(format, args)
	}
}
func Debugln(args ...interface{}) {
	if std == nil {
		logrus.Debugln(args)
	} else {
		std.Debugln(args)
	}
}

func Error(args ...interface{}) {
	if std == nil {
		logrus.Error(args)
	} else {
		std.Error(args)
	}
}
func Errorf(format string, args ...interface{}) {
	if std == nil {
		logrus.Errorf(format, args)
	} else {
		std.Errorf(format, args)
	}
}
func Errorln(args ...interface{}) {
	if std == nil {
		logrus.Errorln(args)
	} else {
		std.Errorln(args)
	}
}

func Info(args ...interface{}) {
	if std == nil {
		logrus.Info(args)
	} else {
		std.Info(args)
	}
}
func Infof(format string, args ...interface{}) {
	if std == nil {
		logrus.Infof(format, args)
	} else {
		std.Infof(format, args)
	}
}
func Infoln(args ...interface{}) {
	if std == nil {
		logrus.Infoln(args)
	} else {
		std.Infoln(args)
	}
}

func Warn(args ...interface{}) {
	if std == nil {
		logrus.Warn(args)
	} else {
		std.Warn(args)
	}
}
func Warnf(format string, args ...interface{}) {
	if std == nil {
		logrus.Warnf(format, args)
	} else {
		std.Warnf(format, args)
	}
}
func Warnln(args ...interface{}) {
	if std == nil {
		logrus.Warnln(args)
	} else {
		std.Warnln(args)
	}
}
