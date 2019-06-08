/*Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下：
函数必须是导出的(首字母大写)
必须有两个导出类型的参数，
第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
函数还要有一个返回值error
举个例子，正确的RPC函数格式如下：
func (t *T) MethodName(argType T1, replyType *T2) error
*/

rpc.HandleHTTP() //将Rpc绑定到HTTP协议上。
var ms = new(common.MathService) //实例化服务对像
   rpc.Register(ms)                 //注册这个服务
   fmt.Println("启动服务...")
   var address, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:1234") //定义TCP的服务承载地址
   listener, err := net.ListenTCP("tcp", address)               //监听TCP连接
   