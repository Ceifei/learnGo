package ch1

import "fmt"

//Chapter information
func Chapter() {
L:
	println()
	println("第1章：程序结构")
	println("-----------------------")
	fmt.Println("1: 标识符")
	fmt.Println("2: 包")
	fmt.Println("3: 变量")
	fmt.Println("4: 常量")
	fmt.Println("5: 注释")
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
