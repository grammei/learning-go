package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++
		// 返回一个累加值
		return value
	}
}
// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为150
	//rand.Seed(time.Now().Unix())

	hp := rand.Intn(151) - -50
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}
func main() {

	// 准备一个字符串
	str := "hello world"
	// 创建一个匿名函数
	foo := func() {

		// 匿名函数中访问str
		str = "hello dude"
	}
	// 调用匿名函数
	foo()
	fmt.Println(str)

	// 创建一个累加器, 初始值为0
	accumulator := Accumulate(0)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 创建一个累加器, 初始值为10
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)
	fmt.Printf("%p\n", accumulator2)

	rand.Seed(time.Now().Unix())
	// 创建一个玩家生成器
	generator1 := playerGen("high")
	generator2 := playerGen("low")
	// 返回玩家的名字和血量
	name1, hp1 := generator1()
	name2, hp2 := generator2()
	// 打印值
	fmt.Println(name1, name2,hp1,hp2)
}