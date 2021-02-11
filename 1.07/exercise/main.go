package main

import (
	"fmt"

	"time"
)

var defaultOffset = 10

func main() {
	デバッグ := false
	日志级别 := "info"
	ይጀምሩ := time.Now()
	_A1_Μείγμα := " "
	fmt.Println(デバッグ, 日志级别, ይጀምሩ, _A1_Μείγμα)

	offset := 5

	fmt.Println(offset)

	offset = 10

	fmt.Println(offset)

	offset = offset + defaultOffset

	fmt.Println(offset)
}
