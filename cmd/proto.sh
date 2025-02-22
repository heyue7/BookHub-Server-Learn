#!/bin/bash

if [[ "$(dirname $0)" == "." ]]; then
    cd ..
fi

[ ! -d proto-pb ] && mkdir -p proto-pb

[ -d proto-pb ] && find proto-pb -name "*.pb.go" | xargs rm -f

# 1. 生成 .pb.go 文件
protoc ./proto/**/**/*.proto \
    --proto_path=./proto \
    --go_out=paths=source_relative:./proto-pb  # 生成普通的 Go 文件

# 2. 生成 gRPC 代码
protoc ./proto/**/**/*.proto \
    --proto_path=./proto \
    --go-grpc_out=paths=source_relative:./proto-pb  # 生成 gRPC 代码

cd proto-pb

# 3. 删除所有生成代码中的 ",omitempty"
find . -name "*.pb.go" | xargs sed -i '' 's/,omitempty//g'