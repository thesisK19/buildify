SHELL := /bin/bash
 
GO_CMD_MAIN = cmd/*.go

run:
	export jwt_secret=eyJhbGciOiJIUzI1NiJ9 && go run $(GO_CMD_MAIN)

# TODO: pull raw component when run server