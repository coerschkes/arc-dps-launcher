package logging

type Logger interface {
	/*
		Logs a message on the default log level.

		@param message - the message to log
	*/
	Log(message string)

	/*
		Logs a message to a specific log level.

		@param message - the message to log
		@param level - the log level
	*/
	LogWithLogLevel(message string, logLevel LOG_LEVEL)

	/*
		Sets the default log level.

		@param level - the log level to set
	*/
	SetDefaultLogLevel(level LOG_LEVEL)

	/*
		Sets the output file to the specified file.

		@param path - the path of the output file
	*/
	SetOutputFile(path string)
}
