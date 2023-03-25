SHELL := /bin/bash

start-server:
	go run app/user/cmd/main.go

# bazel run //:gazelle -- update-repos 
update:
	go mod tidy
	bazel --max_idle_secs=30 run //:gazelle -- update-repos -from_file=go.mod -to_macro=third_party/go_repositories.bzl%go_repositories -build_file_proto_mode=disable -prune=true
	bazel --max_idle_secs=30 run //:gazelle -- update

run-user:
	bazel run //app/user/cmd --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64

run-gen-code:
	bazel run //app/gen-code/cmd --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64