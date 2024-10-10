package errors

// SafetyError represents an error that occurs during safety checks in Git operations.
// It includes information about the error message, the command being executed,
// and an error code for more specific error handling.
type SafteyError struct {
	// Message is the human-readable description of the error
	Message string
	// Command is the Git command that was being executed when the error occurred
	Command string
	// Code is the numeric identifier for the specific error type
	Code int
}
