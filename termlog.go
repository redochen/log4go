// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package log4go

import (
	"fmt"
	"io"
	"os"
	"time"
)

var stdout io.Writer = os.Stdout

//ConsoleLogWriter This is the standard writer that prints to standard output.
type ConsoleLogWriter struct {
	format string
	w      chan *LogRecord
}

//NewConsoleLogWriter This creates a new ConsoleLogWriter
func NewConsoleLogWriter() *ConsoleLogWriter {
	consoleWriter := &ConsoleLogWriter{
		format: "[%T %D] [%C] [%L] (%S) %M",
		w:      make(chan *LogRecord, LogBufferLength),
	}
	go consoleWriter.run(stdout)
	return consoleWriter
}

//SetFormat set format for NewConsoleLogWriter
func (c *ConsoleLogWriter) SetFormat(format string) {
	c.format = format
}

func (c *ConsoleLogWriter) run(out io.Writer) {
	for rec := range c.w {
		fmt.Fprint(out, FormatLogRecord(c.format, rec))
	}
}

//LogWrite This is the ConsoleLogWriter's output method.  This will block if the output
// buffer is full.
func (c *ConsoleLogWriter) LogWrite(rec *LogRecord) {
	c.w <- rec
}

// Close stops the logger from sending messages to standard output.  Attempts to
// send log messages to this logger after a Close have undefined behavior.
func (c *ConsoleLogWriter) Close() {
	close(c.w)
	time.Sleep(50 * time.Millisecond) // Try to give console I/O time to complete
}

//Flush flush content of ConsoleLogWriter to console
func (c *ConsoleLogWriter) Flush() {

}
