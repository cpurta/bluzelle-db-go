package bluzelledbgo

import "errors"

var (
	MISSING_MNEMONIC_ERROR = errors.New("missing required mnemonic in provided config")
)
