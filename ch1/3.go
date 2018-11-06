package ch1

//F3 knowledges about Program Structure in Go
func F3() {

	println()

	//1-6
	func() {
		println("1-6 变量声明")
		println("----------------")
		println("Go语言中，变量在使用前必需先声明，格式如下：")
		println("var i int = 12")
		println("其中i是变量的名字，int是变量的数据类型，12是初始化赋值。")
		println("Go语言是强类型语言，所有变量的数据类型一旦声明就不可改变。不过在声明变量时，可以由赋值来决定数据类型：")
		println("var i = 12")
		println("注意，此时变量i的类型省略了。")
		println("虽然Go语言中，变量声明时并不强制要求初始化，但未显式初始化的变量会被赋予空值：")
		println("var s string // s = \"\"")
		println("每种数据类型都有对应的空值，具体请查阅数据类型相关章节。")
		println("Go语言中声明变量的语法非常灵活，在函数内部可以使用短声明：")
		println("i := 100  // an int")
		println("t := 0.0  // a float-point number")
		println("还可以同时声明多个变量：")
		println("i, j := 0, 1")
		println("此时0赋给了变量i，1赋给了变量j")
		println()
	}()

	//1-7
	func() {
		println("1-7 变量生命周期")
		println("----------------")
		println("定义在函数之外的变量为全局变量，如：")
		println("var global int")
		println("func f() {")
		println("    var x int")
		println("    x = 1")
		println("    global = x")
		println("}")
		println("特别需要注意的是，定义在for和if语句里的变量只能在花括号之后可见：")
		println(`for t :=0.0; {
		}`)
		println()
	}()

}
