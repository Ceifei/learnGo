package main

import (
	"fmt"
	"os"

	"github.com/Ceifei/learnGo/ch1"
)

func main() {

L:
	//Main menu
	println()
	println("Go语言学习笔记")
	println("-----------------------")
	fmt.Println("1: 程序结构")
	fmt.Println("2: 基本数据类型")
	fmt.Println("2: 复合数据类型")
	fmt.Println("3: 引用类型")
	fmt.Println("4: 函数")
	fmt.Println("0: exit")
	println("-----------------------")
	fmt.Print("请选择: ")

	//input
	var input string
	fmt.Scanln(&input)

	//Chapter 1
	if input == "1" {
	L1:
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
		//var inp string
		fmt.Scanln(&input)

		switch input {
		case "1":
			ch1.F1()
		case "2":
			//datatype.ShowFloat()
		case "3":
			//datatype.ShowBool()
		case "4":
			//datatype.ShowComplex()
		case "5":
			//datatype.ShowString()
			//datatype.ShowPrintf()
		default:
			goto L
		}

		goto L1
	}

	//Exit
	if input == "0" {
		os.Exit(1)
	}

	//Back to top
	goto L
}
