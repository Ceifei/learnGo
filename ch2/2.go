package ch2

import (
	"fmt"
	"strconv"
	"unsafe"
)

//F2 about Float-point numbers
func F2() {

	println()

	//2-4
	func() {
		println("2-4 浮点型")
		println("-----------------------")
		println("Go支持2种浮点型数据：float32和float64。")
		fmt.Printf("类型\t字节\t最小值\t\t最大值\n")
		fmt.Printf("%T\t%d\t%e\t%e\n", float32(1),
			unsafe.Sizeof(float32(1)),
			1.401298464324817070923729583289916131280e-45,
			3.40282346638528859811704183484516925440e+38)
		fmt.Printf("%T\t%d\t%e\t%e\n", float64(1),
			unsafe.Sizeof(float64(1)),
			1.797693134862315708145274237317043567981e+308,
			4.940656458412465441765687928682213723651e-324)
		println("注意，没有float型！")
		println("浮点数字面值的类型是float64，浮点型变量未赋值时为0.0。")
		println()
	}()

	//2-5
	func() {
		println("2-5 基本运算")
		println("----------------")

		println("Go语言中，包含浮点数的常量表达式都是float64类型")
		fmt.Printf("1 + 1.0\t%T\t%f\n", 1+1.0, 1+1.0)
		fmt.Printf("1.0 / 6\t%T\t%f\n", 1.0/6, 1.0/6)
		println()

		println("如果表达式中只有32位的浮点数变量，则表达式是float32类型")
		println("var f32 float32 = 12.42")
		println("fmt.Printf(\"f32 + 1.2\t%T\t%[1]f\\n\", f32+1.2)")
		var f32 float32 = 12.42
		fmt.Printf("f32 + 1.2\t%T\t%[1]f\n", f32+1.2)
		fmt.Printf("f32 + 1.2\t%T\t%[1]f\n", f32/5)
		println()

		println("32位浮点数与64位浮点数不能运算")
		println("f := float64(1) + float32(1) //error")
		println()

		println("浮点数不能参与取余操作：")
		println("fmt.Println(-5.0%3) //error")
		println()

		println("浮点数是不精确的，因此切勿用来做比较操作，结果可能出乎意料：")
		println("fmt.Println(1.1+2.2 == 3.3)")
		fmt.Println(1.1+2.2 == 3.3)

		println("浮点数也支持变量自增运算：")
		f32++
		println("f32++")
		fmt.Println(f32)
		println()
	}()

	//2-6
	func() {
		println("2-6 类型转换")
		println("----------------")
		println("本节介绍各种数据类型转换为浮点型的方法。")
		println()

		println("整型转浮点数比较容易")
		println("var u8 uint8 = 12")
		println("f = float32(u8)")
		var u8 uint8 = 12
		f := float32(u8)
		fmt.Printf("f\t%T\t%f\n", f, f)
		println()

		println("布尔型不能转化为浮点型")
		println("f = float32(ture) //error")
		//var t = true
		//f = float32(t) //error
		println()

		println("字符串不能直接转化为浮点型，需要借助strconv.ParseFloat()")
		println("var s = \"0.1\"")
		println("f64, _ := strconv.ParseFloat(s, 32)")
		var s = "0.1"
		f64, _ := strconv.ParseFloat(s, 32)
		fmt.Printf("f64\t%T\t%[1]f\n", f64)
		println()

	}()

}
