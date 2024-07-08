## Hướng dẫn set up cho ubuntu

1. Install go language theo hướng dẫn https://go.dev/doc/install
2. chạy command `make set-up` để install bazel, buf, protobuf

## Hướng dẫn chạy từng service ở localhost
- Một ví dụ để chạy service user: `cd app/user && make run`
- Riêng đối với service gen-code: cần install prettier với command `npm install -g prettier`

## Buildify server

- User service: https://user-service.buildify.asia/ 
- Dynamic-data service: https://dynamic-data-service.buildify.asia/ 
- Gen-code service: https://gen-code-service.buildify.asia/ 
- File-mgt service: https://file-mgt-service.buildify.asia/

![architect](https://github.com/thesisK19/buildify/assets/91143821/7085027b-5f70-47c9-a5ab-2e71c102fb83)

