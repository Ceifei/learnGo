package ch4

import (
	"fmt"
	"unsafe"
)

func getName(name string) string {
	return name
}
func sum(i ...int) int {
	total := 0
	for _, val := range i {
		total += val
	}
	return total
}

func swap(x, y int) (int, int) {
	return y, x
}
func add(x, y int) (z int) {
	z = x + y
	return
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func r(i int) {
	fmt.Println(i)
}

//F2 Is about Functions
func F2() {

	println()

	//4-3
	func() {
		println("4-3 函数")
		println("-----------------------")
		println("函数（Functions）是极其重要的一种程序实体和数据类型。定义方法如下：")
		println("func getName(name string) string {")
		println("	return name")
		println("}")
		fmt.Printf(">> getName %T %d\n", getName, unsafe.Sizeof(getName))
		println("这里定义了一个名为getName的函数，接受一个类型为string的参数name，返回类型是string。")
		println("Go语言中的函数名，实际上是一个引用类型的变量，指向了栈中函数体执行的入口，类似JavaScirpt中的函数。")
		println()
		println("因此，Go语言允许将函数名赋值给另外一个变量，从而执行同一个函数：")
		println("var fn = getName")
		var fn = getName
		fmt.Printf(">> fn %T\n", fn)
		fmt.Printf(">> fn(\"Ceifei\") %T %[1]s\n", fn("Ceifei"))
		println("正因如此，Go语言函数可以作为参数传递，也作为返回值返回。")
		println()

		//Parameters
		println("Go语言不支持参数默认值，所有参数必需在调用时初始化：")
		println("getName()")
		//getName()  //error
		println("多个连续参数，如果类型相同，可以采用简写形式：")
		println("func f(i, j int, s, t string) {}")
		println("i和j都是整型，s和t都是字符串。")
		println()
		println("Go语言支持可变参数，即不固定个数的参数：")
		println("func sum(i ...int) {}")
		println("sum()\t", sum())
		println("sum(1)\t", sum(1))
		println("sum(1, 2)\t", sum(1, 2))
		println("注意，可变参数包括0个参数的情形。一个函数最多声明一个可变参数。")
		println()
		println("可变参数还可以接受数组，效果等同全部元素一个一个传入：")
		println("sum(values...)")
		println("注意，数组名后加是省略号，否则传递的就是一个数组变量了。")
		println()

		//Return Values
		println("函数通常都有返回值，但也可以不返回：")
		println("func onreturn() {}")
		println("Go语言支持函数返回多个值：")
		println("func swap(x, y int) (int, int) {")
		println("	return y, x")
		println("}")
		println("x, y := swap(11, 22)")
		x, y := swap(11, 22)
		fmt.Println(">> x y", x, y)
		println("注意，返回类型的写法。如果有多个值，就必需用括号。")
		println()
		println("Go语言还可以返回命名的参数：")
		println("func add(x, y int) (z int) {")
		println("	z = x + y")
		println("	return")
		println("}")
		println("z := add(11, 22)")
		z := add(11, 22)
		fmt.Println(">> z = ", z)
		println("这个时候返回变量也必需用括号。同时返回语句可以留空。")
		println()

		//Function Type
		println("Go语言中，函数可以作为一种变量数据类型：")
		println("var f func(int, int) int")
		println("此时，f还是nil值，因此不能调用。")
		println("f = add")
		println("f(2, 28)")
		var f func(int, int) int
		f = add
		fmt.Println(f(2, 28))
		println("被实体函数赋值后，f就可以调用了。")
		println()
	}()

	//4-4
	func() {
		println("4-4 匿名函数")
		println("----------------")
		println("Go语言中，命名函数只能是包环境下定义，不能在其他函数里定义。")
		println("但是函数字面值却不限制，可以用在任何表达式中：")
		println("var f = func(x, y int) int { return x * y }")
		var f = func(x, y int) int {
			return x * y
		}
		fmt.Println(">> f(5,6) ", f(5, 6))
		println("函数字面值返回的值就叫做匿名函数，这里f可以调用这个匿名函数。")
		println()

		//Local variables
		println("使用匿名函数，技术上又叫做闭包。使用闭包最大的好处在于，闭包内部可以访问外部变量：")
		println("func squares() func() int {")
		println("	var x int")
		println("	return func() int {")
		println("		x++")
		println("		return x * x")
		println("	}")
		println("}")
		println("fn := squares()")
		fn := squares()
		fmt.Println(">> fn() ", fn())
		fmt.Println(">> fn() ", fn())
		fmt.Println(">> fn() ", fn())
		fmt.Println(">> fn() ", fn())
		fmt.Println(">> fn() ", fn())
		println("fn指向的是squares()内部的那个匿名函数，当fn()执行时，访问了squares内部的变量x。")
		println("由于fn的存在，导致当squares内部的匿名函数执行完后，变量x没有被回收，并继续保持状态。")
		println()

		//Immediately execution
		println("匿名函数支持立即执行，从无需赋值给变量：")
		println("func() { println(\"I am an anonymous function!\") }()")
		func() { println(">> I am an anonymous function!") }()
		println("注意，最后加了一个括号。匿名函数内部的变量不会与外部冲突，因此匿名函数很常用。")
		println()
		println()
	}()

	//4-5
	func() {
		println("4-5 随机与延迟执行")
		println("----------------")
		println("Go语言提供了两种非常有特色的函数执行方式：defer和go。")
		println("defer用于函数延迟执行执行：")
		func() {
			defer r(1)
			defer r(2)
			defer r(3)
		}()
		println("defer修饰的语句，会在宿主函数返回后或函数体结束之后执行。多个defer时，先声明的后执行。")
		println()
	}()

	//4-6
	func() {
		println("4-6 异常处理")
		println("----------------")
		println("Go将运行时错误称为“疼痛”（panic），相应的错误处理机制叫做“恢复”（recover）。")
		println("总体而言，Go语言的异常处理机制还不成熟，有待以后的发展。")
		println()
	}()
}
