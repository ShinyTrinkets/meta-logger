// DefaultLogger is so simple that it's laughable
package logger

import "fmt"

// DefaultLogger is a basic "fmt.print" logger
type DefaultLogger struct {
	// Name by which the logger is identified
	Name string
}

// Info prints information to the screen
func (l *DefaultLogger) Info(msg string, v ...interface{}) {
	fmt.Println("INFO:", msg, v)
}

// Error prints an error message
func (l *DefaultLogger) Error(msg string, v ...interface{}) {
	fmt.Println("ERROR:", msg, v)
}
