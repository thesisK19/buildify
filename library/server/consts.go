package server

const (
	/**
	|-------------------------------------------------------------------------
	| Due to issue "transport is closing", because the request get list orders of COV
	| have time response is larger time to live (age) of connection (many request are larger then 30s)
	| COV-182 Investigate bug Transport is Closing on COV
	| So, the value should be increase
	| See more detail: https://github.com/grpc/grpc-go#the-rpc-failed-with-error-code--unavailable-desc--transport-is-closing
	|-----------------------------------------------------------------------*/
	SERVER_OPTION_MAX_CONNECTION_AGE       = 120 // seconds
	SERVER_OPTION_MAX_CONNECTION_AGE_GRACE = 150 // seconds
	SERVER_OPTION_MAX_CONNECTION_IDLE      = 30  // seconds
	SERVER_OPTION_ENFORCEMENT_MIN_TIME     = 5   // seconds
)
