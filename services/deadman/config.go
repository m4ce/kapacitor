package deadman

import (
	"time"

	"github.com/influxdata/config"
)

const (
	// Default deadman's switch interval
	DefaultInterval = config.Duration(time.Second * 10)
	// Default deadman's switch threshold
	DefaultThreshold = float64(0)
	// Default deadman's switch id
	DefaultId = "node 'NODE_NAME' in task '{{ .TaskName }}'"
	// Default deadman's switch message
	DefaultMessage = "{{ .ID }} is {{ if eq .Level \"OK\" }}alive{{ else }}dead{{ end }}: {{ index .Fields \"collected\" | printf \"%0.3f\" }} points/INTERVAL."
)

type Config struct {
	Interval  config.Duration `toml:"interval"`
	Threshold float64         `toml:"threshold"`
	Id        string          `toml:"id"`
	Message   string          `toml:"message"`
	Global    bool            `toml:"global"`
}

func NewConfig() Config {
	return Config{
		Interval:  DefaultInterval,
		Threshold: DefaultThreshold,
		Id:        DefaultId,
		Message:   DefaultMessage,
	}
}
