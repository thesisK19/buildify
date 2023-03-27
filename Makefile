SHELL := /bin/bash

start-server:
	go run app/user/cmd/main.go

# bazel run //:gazelle -- update-repos 
update:
	go mod tidy
	bazel --max_idle_secs=30 run //:gazelle -- update-repos -from_file=go.mod -to_macro=third_party/go_repositories.bzl%go_repositories -build_file_proto_mode=disable -prune=true
	bazel --max_idle_secs=30 run //:gazelle -- update

run-user:
	bazel run //app/user/cmd:cmd -- server

run-gen-code:
	bazel run //app/gen-code/cmd:cmd -- server

# bazel run //app/orders/fulfillment_router_service --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64

#  bazel build //app/user:base_binary_image
#  bazel run //app/user:base_binary_image

# bazel run //app/gen-code:base_binary_image

# docker run --rm -it -p10080:10080 bazel/app/user:base_binary_image
# docker run --rm -it -p10080:10080 bazel/app/user:base_binary_image

# docker push buildify.azurecr.io/user:base_binary_image
# docker push buildify.azurecr.io/gen-code:base_binary_image

# docker rmi -f $(docker images -aq)


# docker tag user:base_binary_image buildify.azurecr.io/user:base_binary_image

# docker tag gen-code:base_binary_image buildify.azurecr.io/gen-code:base_binary_image