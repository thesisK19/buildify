SHELL := /bin/bash

# buf mod init
# Always run buf mod update after adding a dependency to your buf.yaml.
buf-update:
	buf mod update

buf-gen:
	buf generate --path=app/


# bazel run //:gazelle -- update-repos 
update:
	go mod tidy
	bazel --max_idle_secs=30 run //:gazelle -- update-repos -from_file=go.mod -to_macro=go_repositories.bzl%go_repositories -build_file_proto_mode=disable -prune=true
	bazel --max_idle_secs=30 run //:gazelle -- fix

build:
	bazel build //...

set-up:
	sudo apt install bazel
	go install github.com/bufbuild/buf/cmd/buf@v1.17.0
	sudo apt install -y protobuf-compiler
	go mod tidy


	