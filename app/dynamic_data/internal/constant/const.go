package constant

import "time"

// collection
const (
	USER_COLL       = "users"
	DOCUMENT_COLL   = "documents"
	COLLECTION_COLL = "collections"
)

type DataType int32

const (
	STRING DataType = 1
	NUMBER DataType = 2
)

const (
	EXPORT_DIR    = "export"
	SQL_EXTENSION = ".sql"
)

const (
	BASE_DECIMAL = 10
)

const (
	BUCKET                      = "dynamic-data-bucket"
	DETELE_REMOTE_FILE_DURATION = 5 * time.Minute
)
