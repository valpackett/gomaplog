package gomaplog

type LogLevel uint8

const (
	Emergency LogLevel = 0
	Alert     LogLevel = 1
	Critical  LogLevel = 2
	Error     LogLevel = 3
	Warning   LogLevel = 4
	Notice    LogLevel = 5
	Info      LogLevel = 6
	Debug     LogLevel = 7
)

func (l LogLevel) String() string {
	switch l {
	case Emergency:
		return "EMERGENCY"
	case Alert:
		return "ALERT"
	case Critical:
		return "CRITICAL"
	case Error:
		return "ERROR"
	case Warning:
		return "WARNING"
	case Notice:
		return "NOTICE"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	default:
		return "INFO"
	}
}
