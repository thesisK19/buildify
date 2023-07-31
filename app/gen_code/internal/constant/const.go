package constant

import "time"

const (
	BASE_DECIMAL = 10
)

const (
	INVALID_PAGE_PATH   = ""
	INVALID_COMP_PATH   = ""
	ROOT_ID_PREFIX      = "ROOT"
	ROOT_COMP_ID_PREFIX = "ROOT_C"
	PROPS               = "props"
	REF_PROPS           = "refProps"
	KEY_CHILDREN        = "***"
	KEY_CHILDREN_PROPS  = "***"
	START_CUR_INDEX     = 0
)

const (
	EXPORT_DIR         = "export"
	THEME_DIR          = `src/theme`
	DATABASE_DIR       = `src/database`
	ROUTES_DIR         = `src/routes`
	PAGES_DIR          = `src/pages`
	USER_COMPONENT_DIR = `src/user-components`
	INDEX_TSX          = `index.tsx`
	INDEX_TS           = `index.ts`
	PROPS_TSX          = `props.tsx`
	CONFIG_TSX         = `config.tsx`
	COMPONENT_DIR      = `src/components`
	REACT_JS_BASE_DIR  = "base/reactJS"
	SOURCE_CODE        = "source_code"
	ZIP_EXTENSION      = ".zip"
)

const (
	BUCKET                      = "gen-code-bucket"
	DETELE_REMOTE_FILE_DURATION = 5 * time.Minute
)
