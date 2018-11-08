package ch3

import "fmt"

//Chapter information
func Chapter() {
L:
	println()
	println("第3章：复合数据类型")
	println("-----------------------")
	fmt.Println("1: 数组与切片")
	fmt.Println("2: 映射（map）")
	fmt.Println("3: 结构体")
	fmt.Println("0: 返回上一层")
	println("-----------------------")
	fmt.Print("请选择: ")

	var input string
	fmt.Scanln(&input)

	if input != "0" {

		switch input {
		case "1":
			F1()
		case "2":
			F2()
		case "3":
			F3()
		}

		goto L
	}
}
