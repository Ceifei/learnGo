package ch1

import "fmt"

//Chapter information
func Chapter() int {
L:
	println()
	println("第1章：程序结构")
	println("-----------------------")
	fmt.Println("1: 标识符")
	fmt.Println("2: 声明")
	fmt.Println("3: 变量")
	fmt.Println("3: 常量")
	fmt.Println("0: 返回上一层")
	println("-----------------------")
	fmt.Print("请选择: ")

	var input string
	fmt.Scanln(&input)

	if input == "0" {
		return 0
	} else {
		F1()
		goto L
	}
}
