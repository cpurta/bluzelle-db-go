package bluzelledbgo

import "errors"

var (
	MISSING_MNEMONIC_ERROR      = errors.New("missing required mnemonic in provided config")
	MISSING_API_ENDPOINT_ERROR  = errors.New("missing required api endpoint in provided config")
	NIL_RESPONSE_RETURNED_ERROR = errors.New("query response was nil")
	NOT_IMPLEMENTED_ERROR       = errors.New("method not impelemented")
)
