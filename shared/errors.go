package shared

import (
	"errors"
)

var (
	ErrNotInitialized     = errors.New("not initialized")
	ErrAlreadyInitialized = errors.New("already initialized")
	ErrProofNotExist      = errors.New("proof doesn't exist")
)
