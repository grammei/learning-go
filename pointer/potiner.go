/*package main
import "fmt"

func swap(a, b *int) {
	b, a = a, b
	fmt.Println(*a, *b)
}
func swap1(a, b *int) {
	*b, *a = *a, *b
	fmt.Println(*a, *b)
}
func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
	swap1(&x, &y)
	fmt.Println(x, y)
}*/
package main
// 导入系统包
import (
	"flag"
	"fmt"
)
// 定义命令行参数
var mode = flag.String("mode", "", "process mode")
func main() {
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)
	str :=new(string)
	*str = "gaga"
	fmt.Println(*str)
}