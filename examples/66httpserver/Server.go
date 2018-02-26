package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world!!!")
}

// 启动运行  在浏览器访问 http://localhost:9090/?name=张三 即可显示：Hello world!!! 后台会打印张三。
func main() {

	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}