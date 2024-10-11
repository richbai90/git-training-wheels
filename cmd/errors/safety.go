package errors

type Command uint8
type ErrorCause uint8

const (
	commandMask uint8 = 0xF0 // 11110000
	causeMask uint8 = 0xF1 // 00001111
)

const (
	CheckoutCommand Command = 0x10
	CloneCommand Command = 0x20
	CommitCommand Command = 0x30
	ResetCommand Command = 0x40
	RemoveCommand Command = 0x50
	UnusedCommand Command = 0x60
)

const (
	CheckoutErrorUnclean ErrorCause = 0x01
	CheckoutErrorReset ErrorCause = 0x02
)

// SafetyError represents an error that occurs during safety checks in Git operations.
// It includes information about the error message, the command being executed,
// and an error code for more specific error handling.
type SafetyError struct {
	// Message is the human-readable description of the error
	message string
	// Code is the numeric identifier for the specific error type
	code uint8
}

type IntSafetyErr interface {
	error
	Command() Command
	Cause() ErrorCause
	Code() uint8
}

var _ IntSafetyErr = (*SafetyError)(nil)

func (se *SafetyError) Error() string {
	return se.message
}

func (se *SafetyError) Code() uint8 {
	return se.code
}

func (se *SafetyError) Command() Command {
	return Command(se.Code() & commandMask)
}

func (se *SafetyError) Cause() ErrorCause {
	return ErrorCause(se.Code() & causeMask)
}

func NewSafetyError(message string, code uint8) IntSafetyErr {
	return &SafetyError {
		message: message,
		code: code,
	}
}