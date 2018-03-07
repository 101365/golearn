package main

import (
	"net/url"
	"fmt"
	"strings"
)

// 演示URL解析
func main() {

	//  URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 直接使用url包来解析，解析为URL对象
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// encoded query values, without '?'
	fmt.Println(u.RawQuery)
	// 参数解析
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
