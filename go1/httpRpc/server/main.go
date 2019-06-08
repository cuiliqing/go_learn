package main

import (
	"../common"
	"fmt"
	"net/http"
	"net/rpc"
)

func main() {
	var ms = new (common.MathService)
	rpc.Register(ms)
	rpc.HandleHTTP()
	fmt.Println("start service ing...")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("service is done!")
}


