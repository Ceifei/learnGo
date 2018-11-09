package main

import (
	"fmt"
	"os"

	"github.com/Ceifei/learnGo/ch1"
	"github.com/Ceifei/learnGo/ch2"
	"github.com/Ceifei/learnGo/ch3"
	"github.com/Ceifei/learnGo/ch4"
)

func main() {

L:
	//Main menu
	println()
	println("Go语言学习笔记")
	println("-----------------------")
	fmt.Println("1: 程序结构")
	fmt.Println("2: 基本数据类型")
	fmt.Println("3: 复合数据类型")
	fmt.Println("4: 引用类型")
	// fmt.Println("5: 接口")
	// fmt.Println("6: Go工具")
	// fmt.Println("7: 反射")
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

	//Chapter 2
	if input == "2" {
		ch2.Chapter()
	}

	//Chapter 3
	if input == "3" {
		ch3.Chapter()
	}

	//Chapter 4
	if input == "4" {
		ch4.Chapter()
	}

	//Exit
	if input == "0" {
		os.Exit(1)
	}

	//Back to top
	goto L
}
