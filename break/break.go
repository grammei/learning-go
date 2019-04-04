package main
import "fmt"
func main() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				//break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}

OuterLoop1:
	for i1 := 0; i1 < 2; i1++ {
		for j1 := 0; j1 < 5; j1++ {
			switch j1 {
			case 2:
				fmt.Println(i1, j1)
				continue OuterLoop1
			}
		}
	}
}