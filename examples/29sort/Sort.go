package main

import (
	"sort"
	"fmt"
)

// 演示排序

func main() {

	// 基础数据类型的排序
	baseSort()

	// 使用函数自定义排序
	customSort()

}

func baseSort()  {

	// 直接排序 升序
	strs := []string{"x", "h", "k"}
	// 原地排序 直接修改对应的顺序
	sort.Strings(strs)
	// 打印排序后的结果
	fmt.Println("Strings => ", strs)

	ints := []int{8, 3, 7}
	sort.Ints(ints)
	fmt.Println("Ints => ", ints)

	// 判断是否已经排好序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	// ========逆序排序========
	list := []string{"x", "h", "k"}
	// 通过以下方式逆序排序
	sort.Sort(sort.Reverse(sort.StringSlice(list)))
	fmt.Println("reverse => ", list)

}

// 自定义排序函数
func customSort() {

	//在Java中，有一个Comparable接口，只要实现了它的compare方法即可实现对应的排序逻辑
	//在golang中，有个类似的机制，这里要实现的接口是sort.Interface
	//sort.Interface有三个方法需要实现：Len\Swap\Less
	//Less方法相当于上面提到的Java里的compare方法
	//这里实现一个将字符串切片按照字符串长度进行排序的例子
	elements := []string{"game", "reading", "watch"}
	sort.Sort(ByLength(elements))
	fmt.Println(elements) // [game watch reading]

}

// 定义一个自定义类型
type ByLength []string

// 实现sort.Interface接口
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
