package constant

type Code uint32

const (
	Code_OK                  Code = 0
	Code_CANCELLED           Code = 1
	Code_UNKNOWN             Code = 2
	Code_INVALID_ARGUMENT    Code = 3
	Code_DEADLINE_EXCEEDED   Code = 4
	Code_NOT_FOUND           Code = 5
	Code_ALREADY_EXISTS      Code = 6
	Code_PERMISSION_DENIED   Code = 7
	Code_RESOURCE_EXHAUSTED  Code = 8
	Code_FAILED_PRECONDITION Code = 9
	Code_ABORTED             Code = 10
	Code_OUT_OF_RANGE        Code = 11
	Code_UNIMPLEMENTED       Code = 12
	Code_INTERNAL            Code = 13
	Code_UNAVAILABLE         Code = 14
	Code_DATA_LOSS           Code = 15
	Code_UNAUTHENTICATED     Code = 16
)

var codeToStr = map[Code]string{
	Code_OK:                  "OK",
	Code_CANCELLED:           "Canceled",
	Code_UNKNOWN:             "Unknown",
	Code_INVALID_ARGUMENT:    "InvalidArgument",
	Code_DEADLINE_EXCEEDED:   "DeadlineExceeded",
	Code_NOT_FOUND:           "NotFound",
	Code_ALREADY_EXISTS:      "AlreadyExists",
	Code_PERMISSION_DENIED:   "PermissionDenied",
	Code_RESOURCE_EXHAUSTED:  "ResourceExhausted",
	Code_FAILED_PRECONDITION: "FailedPrecondition",
	Code_ABORTED:             "Aborted",
	Code_OUT_OF_RANGE:        "OutOfRange",
	Code_UNIMPLEMENTED:       "Unimplemented",
	Code_INTERNAL:            "Internal",
	Code_UNAVAILABLE:         "Unavailable",
	Code_DATA_LOSS:           "DataLoss",
	Code_UNAUTHENTICATED:     "Unauthenticated",
}

func (x Code) String() string {
	return codeToStr[x]
}
