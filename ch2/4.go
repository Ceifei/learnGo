package ch2

import (
	"fmt"
	"strconv"
)

//F4 about Strings
func F4() {

	println()

	//2-10
	func() {
		println("2-10 字符串")
		println("-----------------------")
		println("字符串是最常用的数据类型之一。\nGo中字符串字面值用双引号包裹起来：")
		println("var str = \"Ceifei\"")
		var str = "Ceifei"
		fmt.Printf("%T\t%[1]s\t%d\n", str, len(str))
		println("双引号内允许转义字符，如\\\"")
		println("str = \"Cei\\\"fei\"")
		str = "Cei\"fei"
		fmt.Printf("%T\t%[1]s\t%d\n", str, len(str))
		println("Go语言中，单引号用于表示rune，因此在字符串中出现了单引号无需转义。")
		println("str = \"Cei'fei\"")
		str = "Cei'fei"
		fmt.Printf("%T\t%[1]s\t%d\n", str, len(str))
		println("Go语言中的字符串都是UTF-8编码，因此也支持Unicode字符转义：")
		println("str = \"\\u4e16\"")
		str = "\u4e16"
		fmt.Printf("%T\t%[1]s\t%d\n", str, len(str))
		println("注意，一个Unicode字符可能占据多个字节。通常一个汉字是三个字节。")

		println("如果想创建非转义的字符串，可以用反引号代替双引号：")
		println("str = `Ceifei\\nLardon`")
		str = `Ceifei\nLardon`
		fmt.Printf("%T\t%[1]s\t%d\n", str, len(str))
		println()
	}()

	//2-11
	func() {
		println("2-11 运算")
		println("----------------")
		println("Go内置的len()可以用来获取字符串的长度（字节数）:")
		println("str := \"Ceifei\"")
		println("println(len(str))")
		str := "Ceifei"
		println(len(str))

		println("如果要获取字符串中字符的个数，可以把字符串转化为字符数组，然后用len():")
		println("println(len([]rune(\"Ceifei徐磊\")))")
		println(len([]rune("Ceifei徐磊")))

		//rune
		println("Go语言中的字符串实际上一系列字节流，因此很容易获取每个字节的内容：")
		println("str := \"Ceifei\"")
		println("fmt.Printf(\"%d %c\\n\", str[0], str[0])")
		fmt.Printf("%T %[1]d %[1]c\n", str[0])
		println("注意，返回值是整型uint8，也就是byte。")
		println("")

		//Substrings
		println("还可以通过下标获取子字符串，如str[0:5]")
		println("str = \"Ceifei Lardon.\"")
		str = "Ceifei Lardon."
		fmt.Printf("%s\tstr[0:5]\n", str[0:5])
		fmt.Printf("%s\tstr[1:]\n", str[1:])
		fmt.Printf("%s\tstr[:10]\n", str[:10])
		fmt.Printf("%s\tstr[:]\n", str[:])
		println("注意，这里的下标是指字节位置，而不是字符的位置。")
		println()

		//range
		println("在for中搭配range，可以按字符遍历字符串：")
		str = "Ceifei 徐磊"
		for i, r := range str {
			fmt.Printf("%d\t%q\t%X\n", i, r, r)
		}

		println("字符串中的字符是不可修改的：")
		println("str[0]='c' //error")

		//Concatenation
		println("字符串可以通过通过 + 连接：")
		str = str + " Lardon"
		println(str)
		println("注意，字符不能和字符串连接：")
		println("str = str + 'C' //error")

		//Comparsion
		println("字符串可以直接进行比较：")
		println("fmt.Println(\"ceifei\" == \"Ceifei\")")
		fmt.Println("ceifei" == "Ceifei")

		//strings package
		println("strings包中包含大量操作字符串的函数。")
		println()
	}()

	//2-12
	func() {
		println("2-12 类型转换")
		println("----------------")
		println("Go中字符串与字节数组可以相互转化：")
		println("str := \"Ceifei\"")
		println("b := []byte(str)")
		println("fmt.Println(b, string(b))")
		str := "Ceifei"
		b := []byte(str)
		fmt.Println(">> ", b, string(b))
		println()

		//Integers
		println("整型转化为字符串：")
		println("i := 123")
		println("str = strconv.Itoa(i)")
		println("fmt.Printf(\"%T %[1]s\\n\", str)")
		i := 123
		str = strconv.Itoa(i)
		fmt.Printf(">> %T %[1]s\n", str)
		println()
		println("如果直接把一个整型强制转化，则返回一个字符：")
		println("fmt.Printf(\"%T %[1]s\\n\", string(123))")
		fmt.Printf(">> %T %[1]s\n", string(123))
		println("{的ASCII值正是123")
		println()

		//Float-point
		println("浮点型转化为字符串：")
		println("f := 12.3734")
		println("str = strconv.FormatFloat(f, 'f', 3, 64)")
		f := 12.3734
		str = strconv.FormatFloat(f, 'f', 3, 64)
		fmt.Printf(">> %T %[1]s\n", str)
		println("FormatFloat()第三个参数表示返回小数位数。")
		println()

		//Bool
		println("布尔型转化为字符串：")
		println("t := true")
		println("str = strconv.FormatBool(t)")
		t := true
		str = strconv.FormatBool(t)
		fmt.Printf(">> %T %[1]s\n", str)
		println()

		println("Go语言中复数不能直接转化为字符串。")
		println()
	}()

}
