package ch2

import (
	"fmt"
	"strconv"
	"unsafe"
)

//F5 about Boolean
func F5() {

	println()

	//2-13
	func() {
		println("2-13 布尔型")
		println("-----------------------")
		println("布尔型的变量或常量只有两个值：true和false")
		println("var t = true")
		println("var f = false")
		var t = true
		var f = false
		fmt.Printf(">> %T\t%[1]v\t%d\n", t, unsafe.Sizeof(t))
		fmt.Printf(">> %T\t%[1]t\t%d\n", f, unsafe.Sizeof(f))
		println("Go语言严格区分大小写，因此True和False不是布尔型常量了。")
		println("未初始化的布尔变量，默认值是false。")
		var ub bool
		fmt.Printf(">> %T\t%[1]t\n", ub)
		println()

		//Comparsion statement
		println("除直接赋值外，表达式也可以返回布尔值：")
		println("fmt.Printf(\"%T\t%[1]t\\n\", 1 > 2)")
		println("fmt.Printf(\"%T\t%[1]t\\n\", 1 == 0)")
		println("除直接赋值外，表达式也可以返回布尔值：")
		fmt.Printf(">> %T\t%[1]t\n", 1 > 2)
		fmt.Printf(">> %T\t%[1]t\n", 1 == 0)
		println()

		//Convertion
		println("Go语言对数据类型的转换比较保守，大部分数据类型不能直接转换为布尔型：")
		println("if 1 {	} //error")
		println("bool(1) //error")
		println("字符串可以转换为布尔型：")
		println("b, _ = strconv.ParseBool(s)")
		var b bool
		var ss = []string{"1", "t", "T", "TRUE", "True", "true", "0", "f", "F", "FALSE", "False", "false", "2", "0.2", "TRue"}
		for _, s := range ss {
			b, _ = strconv.ParseBool(s)
			fmt.Printf(">> %7q\t%T(%t)\n", s, b, b)
		}
		println("注意，只有少数字符串可能转换为true，大部分都转换为false。")
		println()
	}()
}
