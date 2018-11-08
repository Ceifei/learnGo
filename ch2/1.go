package ch2

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

//F1 about Integers
func F1() {

	println()

	//2-1
	func() {
		println("2-1 整型")
		println("----------------")
		println("数据类型是一切编程语言的基础。Go语言中，数据类型分为四类：")
		println("基本数据、复合数据、引用数据和接口数据。")
		println("其中基本数据又包括数字、字符串和布尔值。")
		println("Go语言中，数字类型包括整型、浮点型和复数，本节学习整型。")
		println("Go语言一共支持11种类型的整型，基本情况如下表：")
		fmt.Printf("类型\t字节\t最小值\t\t\t最大值\n")
		fmt.Printf("%T\t%d\t%d\t\t\t%d\n", int8(1), unsafe.Sizeof(int8(1)), -1<<7, 1<<7-1)
		fmt.Printf("%T\t%d\t%d\t\t\t%d\n", int16(1), unsafe.Sizeof(int16(1)), -1<<15, 1<<15-1)
		fmt.Printf("%T\t%d\t%d\t\t%d\n", int32(1), unsafe.Sizeof(int32(1)), -1<<31, 1<<31-1)
		fmt.Printf("%T\t%d\t%d\t%d\n", int64(1), unsafe.Sizeof(int64(1)), -1<<63, 1<<63-1)
		fmt.Printf("%T\t%d\t%d\t\t\t%d\n", uint8(1), unsafe.Sizeof(uint8(1)), 0, 1<<8-1)
		fmt.Printf("%T\t%d\t%d\t\t\t%d\n", uint16(1), unsafe.Sizeof(uint16(1)), 0, 1<<16-1)
		fmt.Printf("%T\t%d\t%d\t\t\t%d\n", uint32(1), unsafe.Sizeof(uint32(1)), 0, 1<<32-1)
		fmt.Printf("%T\t%d\t%d\t\t\t%s\n", uint64(1), unsafe.Sizeof(uint64(1)), 0, "18446744073709551615")
		fmt.Printf("%T\t%d\t%d\n", uintptr(1), unsafe.Sizeof(uintptr(1)), 0)
		fmt.Printf("%T\t%d\n", int(1), unsafe.Sizeof(int(1)))
		fmt.Printf("%T\t%d\n", uint(1), unsafe.Sizeof(uint(1)))
		println("int和uint的取值范围与平台有关，在64位CPU上，分别等同int64和uint64。")
		println("此外，Go语言还定义两种常用的别名。byte是uint8的别名，rune是int32的别名。")
		fmt.Printf("%T\t%d\n", byte(1), unsafe.Sizeof(byte(1)))
		fmt.Printf("%T\t%d\n", rune(1), unsafe.Sizeof(rune(1)))
		println("各种整型在初始化时没有赋值，则会被自动赋值为0。")
		println()
	}()

	//2-2
	func() {
		println("2-2 基本运算")
		println("----------------")

		println("Go语言是强类型语言，因此不能类型的变量不能直接运算。")
		println("var i8 int8 = 3")
		println("var u8 uint8 = 7")
		println("var s = i8 + u8 //error")
		println("fmt.Println(int(i8) + int(u8))")
		var i8 int8 = 3
		var u8 uint8 = 7
		fmt.Println(int(i8) + int(u8))

		println("整型的字面值类型是int")
		fmt.Printf("%T %[1]d\n", 144)

		println("范围较小的整型运算要注意越界的结果可能出乎意料:")
		println("var u8 uint8 = 255")
		println("fmt.Println(u8, u8, u8*u8)")
		u8 = 255
		fmt.Println(u8, u8+1, u8*u8)
		println("var i8 int8 = 127")
		println("fmt.Println(i8, i8, i8*i8)")
		i8 = 127
		fmt.Println(i8, i8+1, i8*i8)
		println("注意，初始化时，如果字面值超出范围，则编译器报错：")
		println("var i8 int8 = 128 //error")
		println()

		println("整型与整型之间的除法结果也为整数型：")
		println("fmt.Println(7/3, 8/-5, -10/-2)")
		fmt.Println(7/3, 8/-5, -10/-2)
		println("所有结果都被取整了，小数部分完全丢弃。这个与数学常识不符。")
		println()

		println("整型的取余操作较为特殊：")
		println("fmt.Println(-5%3, -5%-3)")
		fmt.Println(-5%3, -5%-3)
		println("结果的正负性只与被除数有关，与除数无关。这个与数学常识也不符。")
		println()

		println("整型是精确的，因此经常被用来做比较操作，结果为布尔值：")
		println("fmt.Println(1 == 1, 12 > 7, 0 > 0)")
		fmt.Println(1 == 1, 12 > 7, 0 > 0)

		println("类似C++，Go语言还支持变量自增运算：")
		println("i8++")
		i8++
		fmt.Println(i8)
		println()
	}()

	//2-3
	func() {
		println("2-3 类型转换")
		println("----------------")
		println("本节介绍各种数据类型转换为整型的方法。")
		println("整型字面值转换为整型，直接用类型名就可以了：")
		println("var i = int8(111)")
		var i = int8(111)
		fmt.Printf("%T %[1]d\n", i)

		println("浮点数字面值不能直接转换为整型：")
		println("var i32 = int(17.3) //error")
		println("浮点型的变量是可以转换为整型的：")
		println("var f float32 = 17.5")
		println("fmt.Printf(\"%T %[1]d\\n\", int8(f))")
		var f float32 = 17.5
		fmt.Printf("%T %[1]d\n", int8(f))
		println("注意到，浮点数的小数部分被完全丢弃。")
		println()

		println("需要四舍五入时，可以使用math.Round()")
		u8 := uint8(math.Round(12.5))
		println("u8 = uint8(math.Round(12.5))")
		fmt.Printf("u8\t%T\t%d\n", u8, u8)
		u8 = uint8(math.Round(12.4))
		println("u8 = uint8(math.Round(12.4))")
		fmt.Printf("u8\t%T\t%d\n", u8, u8)
		println()

		println("需要向上取整时，可以使用math.Ceil()")
		u8 = uint8(math.Ceil(12.1))
		fmt.Printf("u8\t%T\t%d\n", u8, u8)
		println()

		//String
		println("字符串转换为整型需要使用strconv.Atoi()")
		println("如果字符串不是合法的形式，则返回整型0")
		println("var str = \"-42\"")
		println("j, _ := strconv.Atoi(str)")
		var str = "-42"
		j, _ := strconv.Atoi(str)
		fmt.Printf("j\t%T\t%d\n", j, j)
		println("注意返回整型是int")
		println("还可以使用strconv.ParseInt()和strconv.ParseUint()")
		println()

		println("Go语言中，布尔型不能转换为整型。")
		println()

	}()

}
