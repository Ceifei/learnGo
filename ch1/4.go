package ch1

//F4 about Constants
func F4() {

	println()

	//1-8
	func() {
		println("1-8 常量定义")
		println("----------------")
		println("常量在定义的时候必需赋值，且所赋值必需是基本数据类型的表达式：")
		println("const pi = 3.14159")
		const pi = 3.14
		const PI = 3.1415
		println("Go语言中，常量也是区分大小写的。pi和PI不是一个常量")
		println("const pi = 3.14")
		println("const PI = 3.14159")
		println(pi, PI)
		println("一次定义多个常量，可以写作：")
		println("const (")
		println("   e = 2.71828")
		println("   pi = 3.14159")
		println(")")
		println()
	}()

	//1-9
	func() {
		println("1-9 常量表达式")
		println("----------------")
		println("常量的值可以是字面值的表达式，如：")
		println("const T = 1 + 2")
		const T = 1 + 2
		println(T)
		println("常量表达式中不能包含变量，也可能包含数组等复杂结构")
		println("var i = 12")
		println("const T2 = 1 + i // error")
		println("Go语言中预定义了一个常量值生成器：iota")
		println("const (")
		println("   Sunday = iota")
		println("   Monday")
		println("   Tuesday")
		println("   Wednesday")
		println("   Thursday")
		println("   Friday")
		println("   Saturday")
		println(")")
		println("这些常量将被从0到6赋值")
		println()
	}()

}
