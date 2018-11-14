package ch4

import "fmt"

//Chapter information
func Chapter() {
L:
	println()
	println("第4章：引用类型")
	println("-----------------------")
	fmt.Println("1: 指针")
	fmt.Println("2: 函数")
	fmt.Println("3: 方法")
	// fmt.Println("4: 通道")/
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
