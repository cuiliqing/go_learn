package main

import (
	"../common"
	"fmt"
	"net/rpc"
)

func main() {
	var args = common.Args{17,23}
	var result = common.Result{}

	var client, err = rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("connect failed.", err)
	}

	err = client.Call("MathService.Divide", args, &result)
	if err != nil {
		fmt.Println("call failed:", err)
	}
	fmt.Println("call result:", result.Value)
}
