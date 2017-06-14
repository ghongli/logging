package logging

const (
	ERROR_LOG_KEY = "ERROR"
	FATAL_LOG_KEY = "FATAL"
	PANIC_LOG_KEY = "PANIC"
	INFO_LOG_KEY  = "INFO"
	WARN_LOG_KEY  = "WARN"
)

type LoggerTag struct {
	name   string
	prefix string
}

func (l *LoggerTag) TagName() string {
	return l.name
}

func (l *LoggerTag) TagPrefix() string {
	return l.prefix
}

var logTagMap map[string]LoggerTag = map[string]LoggerTag{
	ERROR_LOG_KEY: {name: ERROR_LOG_KEY, prefix: "[" + ERROR_LOG_KEY + "]"},
	FATAL_LOG_KEY: {name: FATAL_LOG_KEY, prefix: "[" + FATAL_LOG_KEY + "]"},
	WARN_LOG_KEY:  {name: WARN_LOG_KEY, prefix: "[" + WARN_LOG_KEY + "]"},
	INFO_LOG_KEY:  {name: INFO_LOG_KEY, prefix: "[" + INFO_LOG_KEY + "]"},
	PANIC_LOG_KEY: {name: PANIC_LOG_KEY, prefix: "[" + PANIC_LOG_KEY + "]"},
}

func getErrorLogTag() LoggerTag {
	return logTagMap[ERROR_LOG_KEY]
}

func getFatalLogTag() LoggerTag {
	return logTagMap[FATAL_LOG_KEY]
}

func getWarnLogTag() LoggerTag {
	return logTagMap[WARN_LOG_KEY]
}

func getInfoLogTag() LoggerTag {
	return logTagMap[INFO_LOG_KEY]
}

func getPanicLogTag() LoggerTag {
	return logTagMap[PANIC_LOG_KEY]
}
