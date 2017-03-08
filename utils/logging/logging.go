// Package logging is reserved for all logging initializing functions.
//
package logging

import (
	"fmt"
	"io"
	"os"

	"../config"

	log "github.com/Sirupsen/logrus"
)

var logfile *os.File

// InitLog initialize the Logrus-Logger.
//
func InitLog() bool {
	if config.GetBoolValue("LoggingFileUse") {
		logFilename := config.GetStringValue("LoggingFile")
		if !fileExists(logFilename) {
			createFile(logFilename)
		}
		f, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		logfile = f
		if err != nil {
			log.Fatal(err)
			return false
		}

		log.SetOutput(io.Writer(logfile))
	}

	log.SetLevel(getLogLevel(config.GetStringValue("LoggingLevel")))

	return true
}

// CloseLog is for closing the used Log-File
//
func CloseLog() {
	fmt.Printf("Log-Datei schließen")
	logfile.Close()
	fmt.Printf("Log-Datei geschlossen")
}

// fileExists is a private function for checking if a Log-File still exists.
//
// PARAMETERS:
//    name - the filename that will be checked.
//
// The function returns a bool:
//    true - if a file exists
//    false - if it doesn't exists.
//
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// createFile is a private function for creating a new Log-File.
//
// PARAMETERS:
//    name - the filename of the new Log-File
//
func createFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}

// getLogLevel sets the Log-Level read from Config
//
func getLogLevel(logLevelString string) log.Level {
	if logLevelString == "error" {
		return log.ErrorLevel
	} else if logLevelString == "debug" {
		return log.DebugLevel
	} else if logLevelString == "info" {
		return log.InfoLevel
	} else if logLevelString == "fatal" {
		return log.FatalLevel
	} else if logLevelString == "panic" {
		return log.PanicLevel
	}

	return log.WarnLevel
}
