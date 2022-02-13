package logging

type Logger interface {

	/*
		Logs a message on the default og level.

		@aram message - the message to log
	*/
	Log(message string)

	/*
			Sts the output file to the specified file.

		@param path - the path of the output file
	*/
	SetOutputFile(path string)
}
