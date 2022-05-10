package exchange

import "github.com/wspowell/errors"

var (
	errUncaughtDbError = errors.New("uncaught database error")
	errMissingLocation = errors.New("missing location")
)