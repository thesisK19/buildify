package constant

type Code string

const (
	Code_OK                  Code = "OK"
	Code_CANCELLED           Code = "CANCELLED"
	Code_UNKNOWN             Code = "UNKNOWN"
	Code_INVALID_ARGUMENT    Code = "INVALID_ARGUMENT"
	Code_DEADLINE_EXCEEDED   Code = "DEADLINE_EXCEEDED"
	Code_NOT_FOUND           Code = "NOT_FOUND"
	Code_ALREADY_EXISTS      Code = "ALREADY_EXISTS"
	Code_PERMISSION_DENIED   Code = "PERMISSION_DENIED"
	Code_UNAUTHENTICATED     Code = "UNAUTHENTICATED"
	Code_RESOURCE_EXHAUSTED  Code = "RESOURCE_EXHAUSTED"
	Code_FAILED_PRECONDITION Code = "FAILED_PRECONDITION"
	Code_ABORTED             Code = "ABORTED"
	Code_OUT_OF_RANGE        Code = "OUT_OF_RANGE"
	Code_UNIMPLEMENTED       Code = "UNIMPLEMENTED"
	Code_INTERNAL            Code = "INTERNAL"
	Code_UNAVAILABLE         Code = "UNAVAILABLE"
	Code_DATA_LOSS           Code = "DATA_LOSS"
)

func (x Code) String() string {
	return string(x)
}
