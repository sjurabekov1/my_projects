// Package logger represents a generic logging interface
package logger

import "hyssa/hyssa_go_billing_service/pkg/logger"

// Log is a package level variable, every program should access logging function through "Log"
var (
	Log logger.Logger
)

// SetLogger is the setter for log variable, it should be the only way to assign value to log
func SetLogger(newLogger logger.Logger) {
	Log = newLogger
}
