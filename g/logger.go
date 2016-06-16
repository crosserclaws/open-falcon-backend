package g

import (
	"os"
	"sync"

	"github.com/Sirupsen/logrus"
)

var (
	logger *logrus.Entry
	Locker = new(sync.RWMutex)
)

func Logger() *logrus.Entry {
	Locker.RLock()
	defer Locker.RUnlock()
	return logger
}

func InitLogger(debugLevel bool) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stderr)
	if debugLevel {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
	logger = logrus.WithFields(logrus.Fields{})
}
