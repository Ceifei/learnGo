package ch4

import (
	"fmt"
	"unsafe"
)

func zero(ptr *int) {
	*ptr = 0
}

//F1 about Pointers
func F1() {

	println()

	//4-1
	func() {
		println("4-1 指针")
		println("----------------")
		println("指针（Pointers）存储着内存地址，指向另一个变量。指针变量通常定义就会初始化：")
		println("var i = 12")
		println("var ptr = &i")
		var i = 12
		var ptr = &i
		fmt.Printf(">> ptr %T %[1]v %d\n", ptr, unsafe.Sizeof(ptr))
		println("&是取地址操作符。注意，不是所有变量或对象都可以取地址。")
		println("通过*ptr可以访问指针执行的那个变量：")
		*ptr = 13
		fmt.Printf(">> *ptr %T %[1]v %d\n", *ptr, unsafe.Sizeof(*ptr))
		fmt.Printf(">> i %T %[1]v %d\n", i, unsafe.Sizeof(i))
		println("注意到，i和*ptr所代表的变量完全一致。")
		println()

		//new
		println("Go语言内置的new()函数可以快速创建各种类型的指针：")
		println("ptrInt := new([]int)")
		println("*ptrInt = []int{11, 22}")
		ptrInt := new([]int)
		*ptrInt = []int{11, 22}
		fmt.Printf(">> ptrInt %T %[1]v %d\n", ptrInt, unsafe.Sizeof(ptrInt))
		println("通过指针ptrInt，我们操作这个匿名数组。")
		println()

		//Type
		println("不同类型的指针不能赋值：")
		println("ptrstr := new(string)")
		println("ptr = ptrstr //error")
		//ptrstr := new(string)
		//ptr = ptrstr //error
		println()

		//Comparison
		println("类型相同的指针才可以比较，当指向相同地址时才相等：")
		println("ptr == &i")
		println("ptr == ptrInt //error")
		fmt.Println(">> ptr == &i ", ptr == &i)
		//fmt.Println(">> ptr == ptrInt ", ptr == ptrInt)
		println()

		//Function Parameters
		println("很多时候，我们希望函数修改外部变量值，这个时候需要用指针做参数：")
		println("func zero(ptr *int) { *ptr = 0}")
		println("x := 5")
		println("zero(&x)")
		x := 5
		zero(&x)
		fmt.Println(">> x =", x)
		println()

	}()

	//4-2
	func() {
		println("4-2 指针与结构体")
		println("----------------")
		println("结构体的域不能是自己，但可以是指向自己的指针：")
		println("type tree struct {")
		println("	value int")
		println("	left  *tree")
		println("}")
		type tree struct {
			value int
			left  *tree
			right *tree
		}
		println("这样就可以构建嵌套的数据结构。")
		println()
	}()
}
