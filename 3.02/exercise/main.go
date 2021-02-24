package main

import (
	"fmt"
	//"runtime"
)

//func showMemoryUsedByInts() {
//	var list []int
//
//	for i := 0; i < 10000000; i++ {
//		list = append(list, 100)
//
//		var m runtime.MemStats
//
//		runtime.ReadMemStats(&m)
//
//		fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
//
//	}
//}

func main() {

	// showMemoryUsedByInts()

	var (
		a int     = 100
		b float32 = 100
		c float64 = 100
	)

	fmt.Println(a/3, "\n", b/3, "\n", c/3)
	fmt.Println((a/3)*3, "\n", (b/3)*3, "\n", (c/3)*3)

}
