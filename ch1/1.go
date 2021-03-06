package ch1

import "fmt"

//F1 knowledges about Program Structure in Go
func F1() {

	println()

	//1-1: Keywords
	func() {
		println("1-1 关键字")
		println("----------------")
		println("Go一共有25个关键字，按照用途可以分为以下三类：")
		println("程序声明(2)		import,package")
		println("实体声明(8)		chan,const,func,interface,map,struct,type,var")
		println("流程控制(15)		go,select,break,case,continue,default,defer,else,fallthrough,for,goto,if,range,return,switch")
		println("关键词在程序中具有特殊含义，不可用作变量名或函数名。")
		println()
	}()

	//1-2: Predeclared names
	func() {
		println("1-2 预定义标识符")
		println("----------------")
		println("除了关键字，Go还预定义若干标识符，在程序任何位置都可以直接使用。按照用途可以分为以下三类：")
		println("常量(4)		true,false,iota,nil")
		println("数据类型(20)	int,int8,int16,int32,int64")
		println("           	uint,uint8,uint16,uint32,uint64,uintptr")
		println("           	float32,float64,complex64,complex128")
		println("           	bool,byte,rune,string,error")
		println("函数(13)		make,len,cap,new,append,copy,close,delete")
		println("       		complex,real,imag")
		println("				panic,recover")
		println("预定义标识符虽然也具有特殊含义，但在实际编程时可以重新定义，覆盖Go语言内置的含义：")
		println("var true = 1")
		println("println(true) // 1")
		var true = 1
		println(true)
		println("再看一个例子：")
		println("var int = 12.3")
		println("fmt.Printf(\"%T %[1]f\\n\", int)")
		var int = 12.3
		fmt.Printf("%T %[1]f\n", int)
		println()
	}()

	//1-3: User-defined names
	func() {
		println("1-3 自定义标识符")
		println("----------------")
		println("Go语言中，自定义标识符必需以字母或下划线开头，后面跟若干个字母、数字或下划线，如：")
		println("i _i this32 我 text_chapter myName")
		println("特别说明这里说的字母，指Unicode字符集中标记为字母的字符，其中就包括汉字、日本和韩文。")
		println("注意，美元符号（$）不能成为标识符的一部分！")
		println()
	}()

}
