package cli

import "fmt"

var (
	ErrUnknownFlag  = fmt.Errorf("unknown flag")
	ErrMissingValue = fmt.Errorf("missing flag value")
	ErrInvalidValue = fmt.Errorf("invalid flag value")
)

type UnknownFlagError struct {
	Flag string
}

func (e UnknownFlagError) Error() string {
	return fmt.Sprintf("%s: %s", ErrUnknownFlag, e.Flag)
}

func (e UnknownFlagError) Is(targert error) bool {
	return targert == ErrUnknownFlag
}

// ---

type MissingValueError struct {
	Flag string
}

func (e MissingValueError) Error() string {
	return fmt.Sprintf("%s: %s", ErrMissingValue, e.Flag)
}

func (e MissingValueError) Is(targert error) bool {
	return targert == ErrMissingValue
}

// ---

type InvalidValueError struct {
	Flag string
}

func (e InvalidValueError) Error() string {
	return fmt.Sprintf("%s: %s", ErrInvalidValue, e.Flag)
}

func (e InvalidValueError) Is(targert error) bool {
	return targert == ErrInvalidValue
}
