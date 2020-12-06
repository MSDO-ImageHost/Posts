package auth

import "fmt"

var (
	InsufficentPermisions error = fmt.Errorf("Insufficent permissions")
)
