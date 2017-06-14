package logging

type LogManager struct {
	loggers []Logger
}

func (logger *LogManager) GetPosition() Position {
	return POSITION_SINGLE
}

func (logger *LogManager) SetPosition(pos Position) {}

func (logger *LogManager) Error(v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Error(v...)
	}
	return content
}

func (logger *LogManager) Errorf(format string, v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Errorf(format, v...)
	}
	return content
}

func (logger *LogManager) Errorln(v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Errorln(v...)
	}
	return content
}

func (logger *LogManager) Fatal(v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Fatal(v...)
	}
	return content
}

func (logger *LogManager) Fatalf(format string, v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Fatalf(format, v...)
	}
	return content
}

func (logger *LogManager) Fatalln(v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Fatalln(v...)
	}
	return content
}

func (logger *LogManager) Panic(v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Panic(v...)
	}
	return content
}

func (logger *LogManager) Panicf(format string, v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Panicf(format, v...)
	}
	return content
}

func (logger *LogManager) Panicln(v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Panicln(v...)
	}
	return content
}

func (logger *LogManager) Info(v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Info(v...)
	}
	return content
}

func (logger *LogManager) Infof(format string, v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Infof(format, v...)
	}
	return content
}

func (logger *LogManager) Infoln(v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Infoln(v...)
	}
	return content
}

func (logger *LogManager) Warn(v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Warn(v...)
	}
	return content
}

func (logger *LogManager) Warnf(format string, v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Warnf(format, v...)
	}
	return content
}

func (logger *LogManager) Warnln(v ...interface{}) string {
	var content string
	for _, logger := range logger.loggers {
		content = logger.Warnln(v...)
	}
	return content
}
