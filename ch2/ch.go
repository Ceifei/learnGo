package ch2

import "fmt"

//Chapter information
func Chapter() {
L:
	println()
	println("第2章：基本数据类型")
	println("-----------------------")
	fmt.Println("1: 整型")
	fmt.Println("2: 浮点型")
	fmt.Println("3: 复数")
	fmt.Println("4: 字符串")
	fmt.Println("5: 布尔型")
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
		}

		goto L
	}
}
