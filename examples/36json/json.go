package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 演示json的处理
// Go 提供内置的 JSON 编解码支持，包括内置或者自定义类型与 JSON 数据之间的转化。
// 重要的就是以下2个函数：
// func Unmarshal(data []byte, v interface{}) error
// func Marshal(v interface{}) ([]byte, error)

func main() {

	// 基本数据类型的转换
	commonDemo()

	// 演示结构体转换
	structDemo()

	// map
	mapDemo()

	// Decode
	decodeDemo()

}

func decodeDemo() {
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple6": 5, "lettuce6": 7}
	enc.Encode(d)
}

func mapDemo() {

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

}

//
func structDemo() {

	u := user{Name:"张三", Age:18}
	structJ, _ := json.Marshal(u)
	fmt.Println(string(structJ))

	var u2 user
	json.Unmarshal(structJ, &u2)
	fmt.Println(u2)
	fmt.Println(u2.Name)

	var u3 user
	byt := []byte(`{"Name":"张三","Age":18}`)
	json.Unmarshal(byt, &u3)
	fmt.Println(u3)
	fmt.Println(u3.Name)

}

type user struct {
	Name string
	Age int
}

// 基本数据类型示例
func commonDemo() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("中国")
	fmt.Println(string(strB))

	// 逆向
	var b bool
	json.Unmarshal([]byte("true"), &b)
	fmt.Println(b)

	var i int
	json.Unmarshal([]byte("9"), &i)
	json.Unmarshal([]byte("str"), &i) // 这里转换失败 会返回0
	fmt.Println(i)

	var s string
	json.Unmarshal([]byte(`"中国"`), &s)
	fmt.Println("s => ", s)

	arr := []byte("中国")
	for _, c := range arr {
		fmt.Println("byte => ", c)
	}

	arr2 := []rune("中国")
	for _, c := range arr2 {
		fmt.Println("rune => ", c)
	}

}
