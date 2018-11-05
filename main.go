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
	fmt.Println("0: 退出")
	println("-----------------------")
	fmt.Print("请选择: ")

	//input
	var input string
	fmt.Scanln(&input)

	//Chapter 1
	if input == "1" {

		ch1.Chapter()
	}

	//Exit
	if input == "0" {
		os.Exit(1)
	}

	//Back to top
	goto L
}
