package conf

// default log path
const (
	DefaultAppLogPath   = "logs/app.log"
	DefaultErrorLogPath = "logs/error.log"
)

// LogConfig holds log-related configuration.
type LogConfig struct {
	Level        string `yaml:"level"`          // debug/info/warn/error
	Outputs      string `yaml:"outputs"`        // comma separated: stdout, file
	AppLogPath   string `yaml:"app_log_path"`   // log file path
	ErrorLogPath string `yaml:"error_log_path"` // error log file path
}
