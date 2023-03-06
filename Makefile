SHELL := /bin/bash

start-server:
	go run app/user/cmd/main.go

# bazel run //:gazelle -- update-repos 
