#########

go1

/*Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下： 函数必须是导出的(首字母大写) 必须有两个导出类型的参数，
 第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的 函数还要有一个返回值error 举个例子，正确的RPC函数格式如下： func (t *T) MethodName(argType T1, replyType *T2) error 
*/

rpc.HandleHTTP() //将Rpc绑定到HTTP协议上。 var ms = new(common.MathService) //实例化服务对像 rpc.Register(ms) //注册这个服务 fmt.Println("启动服务...") var address, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:1234") //定义TCP的服务承载地址 listener, err := net.ListenTCP("tcp", address) //监听TCP连接

go2
准备部分较多
Application for go2 : 调用server的计算器返回给client结果

1.提前安装grpc(https://studygolang.com/articles/14920)，
protoc以及依赖项（go get -u github.com/golang/protobuf/{proto,protoc-gen-go}）

2.注意 GOPATH and GOROOT 配置（重要, 可在client go env命令查看）
GOPATH 为项目工程路径（go项目工程路径下标准为src、kg、 bin三同级目录）

3 protoc -I ${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/*/*.proto
生成proto对应的go代码，编译过程命令参数注意


