// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package log4go

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	Global Logger
)

func init() {
	Global = NewDefaultLogger(DEBUG)
}

// Wrapper for (*Logger).LoadConfiguration
func LoadConfiguration(filename string, types ...string) {
	if len(types) > 0 && types[0] == "xml" {
		Global.LoadConfiguration(filename)
	} else {
		Global.LoadJsonConfiguration(filename)
	}
}

// Wrapper for (*Logger).AddFilter
func AddFilter(name string, lvl Level, writer LogWriter) {
	Global.AddFilter(name, lvl, writer)
}

// Wrapper for (*Logger).Close (closes and removes all logwriters)
func Close() {
	Global.Close()
}

// Wrapper for (*Logger).Flush (flush all logwriters)
func Flush() {
	Global.Flush()
}

func Crash(args ...interface{}) {
	if len(args) > 0 {
		msg := getMessage(strings.Repeat(" %v", len(args))[1:], args...)
		Global.Log(FATAL, getSource(), msg)
	}
	panic(args)
}

// Logs the given message and crashes the program
func Crashf(format string, args ...interface{}) {
	Global.Log(FATAL, getSource(), getMessage(format, args...))
	Global.Close() // so that hopefully the messages get logged
	panic(fmt.Sprintf(format, args...))
}

// Compatibility with `log`
func Exit(args ...interface{}) {
	if len(args) > 0 {
		msg := getMessage(strings.Repeat(" %v", len(args))[1:], args...)
		Global.Log(ERROR, getSource(), msg)
	}
	Global.Close() // so that hopefully the messages get logged
	os.Exit(0)
}

// Compatibility with `log`
func Exitf(format string, args ...interface{}) {
	Global.Log(ERROR, getSource(), getMessage(format, args...))
	Global.Close() // so that hopefully the messages get logged
	os.Exit(0)
}

// Compatibility with `log`
func Stderr(args ...interface{}) {
	if len(args) > 0 {
		msg := getMessage(strings.Repeat(" %v", len(args))[1:], args...)
		Global.Log(ERROR, getSource(), msg)
	}
}

// Compatibility with `log`
func Stderrf(format string, args ...interface{}) {
	Global.Log(ERROR, getSource(), getMessage(format, args...))
}

// Compatibility with `log`
func Stdout(args ...interface{}) {
	if len(args) > 0 {
		msg := getMessage(strings.Repeat(" %v", len(args))[1:], args...)
		Global.Log(INFO, getSource(), msg)
	}
}

// Compatibility with `log`
func Stdoutf(format string, args ...interface{}) {
	Global.Log(INFO, getSource(), getMessage(format, args...))
}

// Send a log message manually
// Wrapper for (*Logger).Log
func Log(lvl Level, source, message string) {
	Global.Log(lvl, source, message)
}

// Send a formatted log message easily
// Wrapper for (*Logger).Logf
func Logf(lvl Level, format string, args ...interface{}) {
	Global.Log(lvl, getSource(), getMessage(format, args...))
}

// Send a closure log message
// Wrapper for (*Logger).Logc
func Logc(lvl Level, closure func() string) {
	Global.Log(lvl, getSource(), closure())
}

// Utility for debug log messages
// When given a string as the first argument, this behaves like Logf but with the DEBUG log level (e.g. the first argument is interpreted as a format for the latter arguments)
// When given a closure of type func()string, this logs the string returned by the closure iff it will be logged.  The closure runs at most one time.
// When given anything else, the log message will be each of the arguments formatted with %v and separated by spaces (ala Sprint).
// Wrapper for (*Logger).Debug
func Debug(arg0 interface{}, args ...interface{}) {
	Global.Log(DEBUG, getSource(), getMessage(arg0, args...))
}

// Utility for info log messages (see Debug() for parameter explanation)
// Wrapper for (*Logger).Info
func Info(arg0 interface{}, args ...interface{}) {
	Global.Log(INFO, getSource(), getMessage(arg0, args...))
}

// Utility for warn log messages (returns an error for easy function returns) (see Debug() for parameter explanation)
// These functions will execute a closure exactly once, to build the error message for the return
// Wrapper for (*Logger).Warn
func Warn(arg0 interface{}, args ...interface{}) error {
	msg := getMessage(arg0, args...)
	Global.Log(WARNING, getSource(), msg)
	return errors.New(msg)
}

// Utility for error log messages (returns an error for easy function returns) (see Debug() for parameter explanation)
// These functions will execute a closure exactly once, to build the error message for the return
// Wrapper for (*Logger).Error
func Error(arg0 interface{}, args ...interface{}) error {
	msg := getMessage(arg0, args...)
	Global.Log(ERROR, getSource(), msg)
	return errors.New(msg)
}

// Utility for fatal log messages (returns an error for easy function returns) (see Debug() for parameter explanation)
// These functions will execute a closure exactly once, to build the error message for the return
// Wrapper for (*Logger).Fatal
func Fatal(arg0 interface{}, args ...interface{}) error {
	msg := getMessage(arg0, args...)
	Global.Log(FATAL, getSource(), msg)
	return errors.New(msg)
}
