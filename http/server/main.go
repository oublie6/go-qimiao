package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Write([]byte("我收到了给你返回GET"))
	case "POST":
		b, _ := io.ReadAll(req.Body)
		//header := res.Header()
		//header["test"] = []string{"test1", "test2"}
		//res.WriteHeader(http.StatusBadGateway)
		fmt.Println(req.Header["Test"])
		res.Write(b)
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler)
	http.ListenAndServe(":8080", mux)
}
