
package main

import (
	"fmt"
	"net/rpc"
	"../common"
)

func main() {
	var client, err = rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("连接不到服务器：", err)
	}
	var args = common.Args{40, 3}
	var result = common.Result{}
	fmt.Println("开始调用！")
	err = client.Call("MathService.Add", args, &result)
	if err != nil {
		fmt.Println("调用失败！", err)
	}
	fmt.Println("调用成功！结果：", result.Value)
}
