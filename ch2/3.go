package ch2

import (
	"fmt"
	"unsafe"
)

//F3 about Complex numbers
func F3() {

	println()

	//2-7
	func() {
		println("2-7 复数")
		println("-----------------------")
		println("技术上来说，复数实际上是浮点数的扩展，复数的实部和虚部都是浮点数。Go支持2种复数：")
		fmt.Printf("类型\t\t字节\n")
		fmt.Printf("%T\t%d\n", complex64(1+2i), unsafe.Sizeof(complex64(1+2i)))
		fmt.Printf("%T\t%d\n", complex128(1+2i), unsafe.Sizeof(complex128(1+2i)))
		println("注意，没有complex型！")
		println("复数字面值的类型是complex128，复数型变量未赋值时为0.0 + 0.0i。")
		var cp complex128
		println(cp)
		println()
	}()

	//2-8
	func() {
		println("2-8 运算")
		println("----------------")
		println("Go内置了三个复数专用函数：complex()、real()和imag()。")
		println("complex()用来初始化复数：")
		println("var c128 = complex(3, 4) // 3+4i")
		var c128 = complex(3, 4)
		fmt.Println(c128)
		println("real()和imag()分别用来获取一个复数的实部和虚部：")
		println("fmt.Printf(\"%T\t%[1]f\\n\", real(c128))")
		println("fmt.Printf(\"%T\t%[1]f\\n\", imag(c128))")
		fmt.Printf("%T\t%[1]f\n", real(c128))
		fmt.Printf("%T\t%[1]f\n", imag(c128))
		var c64 complex64 = 7.2 + 4.5i
		println("fmt.Printf(\"%T\t%[1]f\\n\", real(c64))")
		println("fmt.Printf(\"%T\t%[1]f\\n\", imag(c64))")
		fmt.Printf("%T\t%[1]f\n", real(c64))
		fmt.Printf("%T\t%[1]f\n", imag(c64))
		println("注意返回的浮点数类型。")
		println()
	}()

	//2-6
	func() {
		println("2-9 类型转换")
		println("----------------")
		println("整型和浮点数很容易通过强制类型转换为复数：")
		fmt.Printf("%T\t%[1]f\n", complex128(1))
		fmt.Printf("%T\t%[1]f\n", complex128(1.2))

		println("字符串和布尔型都不能直接转换为复数，一般先转换为浮点数。")
		println()
	}()

}
