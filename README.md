准备部分较多
Application for go2 : 调用server的计算器返回给client结果

1.提前安装grpc(https://studygolang.com/articles/14920)，
protoc以及依赖项（go get -u github.com/golang/protobuf/{proto,protoc-gen-go}）

2.注意 GOPATH and GOROOT 配置（重要, 可在client go env命令查看）
GOPATH 为项目工程路径（go项目工程路径下标准为src、kg、 bin三同级目录）

3 protoc -I ${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/*/*.proto
生成proto对应的go代码，编译过程命令参数注意


