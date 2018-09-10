// The point of this package is to create a layer of abstraction
// on top of your favorite Go logging library, so that
// you can use your favorite logging library in Trinkets.
// A valid Logger implementation needs only 2 methods:
// * Info(string, ...interface{})
// * Error(string, ...interface{})
//
// To enable this you would use:
//
// logger.SetupLogBuilder(func(name string) logger.Logger {
// 	return your.Logger(name)
// })
//
package logger

// Logger types must support Info and Error functions
type Logger interface {
	Info(string, ...interface{})
	Error(string, ...interface{})
}

// Attrs are used for providing additional info in the log messages
type Attrs map[string]interface{}

// logBuilderType function builds new log instances
type logBuilderType func(name string) Logger

// NewLogger builder function
var NewLogger logBuilderType

// SetupLogBuilder helper is used for setting up the log builder function
func SetupLogBuilder(b logBuilderType) {
	NewLogger = b
}
