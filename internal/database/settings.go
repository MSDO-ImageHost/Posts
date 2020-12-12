package database

import (
	"context"
	"time"
)

var (
	_LOG_TAG string = "Database:\t"

	timeOutCtx, cancel = context.WithTimeout(context.Background(), time.Duration(5*time.Second))
)
