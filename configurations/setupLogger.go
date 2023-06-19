package configurations

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}
