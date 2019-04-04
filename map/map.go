package main
import (
	"fmt"
	"sync"
)
func main() {
	sc1 := make(map[string]int)
	v, ok := sc1["route"]
	fmt.Println(v,ok)
	sc1["route"] = 66
	v1, ok1 := sc1["route"]
	fmt.Println(v1,ok1)
	delete(sc1, "route")
	v2, ok2 := sc1["route"]
	fmt.Println(v2,ok2)

	// fatal error: concurrent map read and map write
	// 一种效率较高的并发安全的 sync.Map
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	v3,ok3:=scene.Load("london")
	fmt.Println(v3,ok3)
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}