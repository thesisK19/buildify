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

# run-user:
# 	bazel run //app/user/cmd:cmd

# run-gen-code:
# 	bazel run //app/gen_code/cmd:cmd

# bazel run //app/orders/fulfillment_router_service --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64

#  bazel build //app/user:base_binary_image
#  bazel run //app/user:base_binary_image

# bazel build //app/gen_code:base_binary_image
# bazel run //app/gen_code:base_binary_image
# docker run --rm -it -p3003:3003 bazel/app/gen_code:base_binary_image


# bazel build //app/gen_code/cmd:base_binary_image
# bazel run //app/gen_code/cmd:base_binary_image -- --norun
# docker run --rm -it -p3003:3003 bazel/app/gen_code:base_binary_image


# bazel build //app/gen_code:container_image
# bazel run //app/gen_code:container_image
# docker run --rm -it -p3003:3003 bazel/app/gen_code:container_image

# docker run --rm -it -p10080:10080 bazel/app/user:base_binary_image

# docker push buildify.azurecr.io/user:base_binary_image
# docker push buildify.azurecr.io/gen-code:base_binary_image

# docker rmi -f $(docker images -aq)


# docker tag bazel/app/user:base_binary_image buildify.azurecr.io/user:base_binary_image

# docker tag bazel/app/user:container_image gcr.io/thesis-378216/buildify-registry/user:base_binary_image

# docker tag bazel/app/gen_code:base_binary_image buildify.azurecr.io/gen-code:base_binary_image
# docker tag bazel/app/file_mgt:container_image buildify.azurecr.io/file_mgt:base_binary_image
# docker push buildify.azurecr.io/file_mgt:base_binary_image



# docker run -it bazel/app/gen_code:container_image




# 9,398,200.00
# 7,048,650.00 credit

#  bazel query @go_base_image//...   

#  docker exec -it ad019aef5836 /bin/sh

# bazel run //app/gen_code:push