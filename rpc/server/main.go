package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Req struct {
	NumOne int
	NumTwo int
}

type Res struct {
	Num int
}

func main() {
	req := Req{NumOne: 1, NumTwo: 2}
	var res Res
	client, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//client.Call("Server.Add", req, &res)
	ca := client.Go("Server.Add", req, &res, nil)
tag:
	for {
		select {
		case <-ca.Done:
			break tag
		default:
			time.Sleep(time.Second)
			fmt.Println("我等着")
		}
	}
	fmt.Println(res)
}
