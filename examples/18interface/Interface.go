package main

import "fmt"

// 本例演示接口的使用

// 实现一个接口 只需要实现接口中所有的方法

type animal interface {
	say(s string)
	run()
}

type human struct {
	name string
}

type dog struct {
	name string
}

func (self human) say(s string)  {
	fmt.Println(self.name + " 说: ", s)
}

func (self dog) say(s string)  {
	fmt.Println(self.name + " 说: ", s)
}

func (self human) run()  {
	fmt.Println("human is running...")
}

func (self dog) run()  {
	fmt.Println("dog is running...")
}

// 接受实现了animal接口的类型
func display(a animal)  {
	a.run()
	a.say("hello")
}

func main() {

	h := human{"larry"}
	h.say("不错")
	h.run()

	d := dog{"mao"}
	d.say("wang~")
	d.run()

	// 演示调用
	display(h)
	display(d)

	// TODO 接口方法实现接收器应该是值对象还是指针对象？ 待补充


}


