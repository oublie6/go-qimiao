package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:8080/test", bytes.NewBuffer([]byte("{\"test\":\"我是客户端\"}")))
	req.Header["test"] = []string{"test", "test1"}
	res, _ := client.Do(req)
	body := res.Body
	b, _ := io.ReadAll(body)
	fmt.Println(string(b))
}
