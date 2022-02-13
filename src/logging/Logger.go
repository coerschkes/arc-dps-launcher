package logging

type Logger interface {

	/*
		Logs a message.

		@param message - the message to log
	*/
	Log(message string)

	/*
		Logs an error.

		@param err - the error to log
	*/
	LogError(err error)

	/*
		Sets the output file to te specified fil.

		@param path - the path of the output fie
	*/
	SetOutputFile(path string)
}
