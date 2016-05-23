package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	// Trace - log trace level output
	Trace *log.Logger
	// Info - log info level output
	Info *log.Logger
	// Warning - log warning level output
	Warning *log.Logger
	// Error - log error level output
	Error *log.Logger
)

func init() {
	// Initialize the custom logging options.
	// These are the defaults to use for testing
	// to avoid log output in the tests.
	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(ioutil.Discard,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(ioutil.Discard,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(ioutil.Discard,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// SetOutput - enables non-test logging output
func SetOutput() {
	Trace.SetOutput(ioutil.Discard)
	Info.SetOutput(os.Stdout)
	Warning.SetOutput(os.Stdout)
	Error.SetOutput(os.Stderr)
}
