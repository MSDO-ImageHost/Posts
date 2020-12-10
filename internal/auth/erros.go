package auth

import "fmt"

var (
	InsufficentPermisions error = fmt.Errorf("Insufficent permissions")
	InvalidAuthToken      error = fmt.Errorf("Invalid authentication")
)
