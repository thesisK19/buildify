package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Level         string `json:"level" yaml:"level" mapstructure:"level"`
	ReportCaller  bool   `json:"report_caller" yaml:"report_caller" mapstructure:"report_caller"`
	DisableColors bool   `json:"disable_colors" yaml:"disable_colors" mapstructure:"disable_colors"`
}

func DefaultConfig() Config {
	return Config{
		Level:         logrus.DebugLevel.String(),
		ReportCaller:  false,
		DisableColors: false,
	}
}

func (c Config) Build() (*logrus.Logger, error) {
	log := logrus.New()

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the debug severity or above
	// ALL < TRACE < DEBUG < INFO < WARN < ERROR < FATAL < OFF
	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}
	log.SetLevel(level)

	// logrus show line number
	log.SetReportCaller(c.ReportCaller)

	// format color
	log.SetFormatter(&logrus.TextFormatter{
		// ForceColors:   true,
		DisableColors: c.DisableColors,
	})

	return log, nil
}
