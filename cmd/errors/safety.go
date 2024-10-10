package errors

// SafetyError represents an error that occurs during safety checks in Git operations.
// It includes information about the error message, the command being executed,
// and an error code for more specific error handling.
type SafetyError struct {
	// Message is the human-readable description of the error
	message string
	// Command is the Git command that was being executed when the error occurred
	command string
	// Code is the numeric identifier for the specific error type
	code int
}

type IntSafetyErr interface {
	error
	Command() string
	Code() int
}

var _ IntSafetyErr = (*SafetyError)(nil)

func (se *SafetyError) Error() string {
	return se.message
}

func (se *SafetyError) Code() int {
	return se.code
}

func (se *SafetyError) Command() string {
	return se.command
}

func NewSafetyError(message string, code int, command string) IntSafetyErr {
	return &SafetyError {
		message: message,
		code: code,
		command: command,
	}
}