package ch3

import (
	"fmt"
	"unsafe"
)

//F2 about Maps
func F2() {

	println()

	//3-3
	func() {
		println("3-3 映射")
		println("-----------------------")
		println("Go语言中，映射(maps)是哈希表的一种实现，定义方式如下：")
		var m map[int]string
		m2 := make(map[int]string)
		fmt.Printf(">> m  %T %d\n", m, unsafe.Sizeof(m))
		fmt.Printf(">> m2 %T %d\n", m2, unsafe.Sizeof(m2))
		println("其中，int是索引类型，string是值类型。")
		println()

		println("通过字面值初始化映射：")
		println("m = map[int]string{")
		println("	3: \"Kobe\",")
		println("	4: \"Byrant\",")
		println("}")
		m = map[int]string{
			3: "Kobe",
			4: "Byrant",
		}
		fmt.Printf(">> m  %v\n", m)
		println("注意到，字面值最后一项的逗号，这个在Go语言中是允许的。")
		println()
		println("初始化之后，可以通过下标访问映射的值：")
		println("m[0] = \"Ceifei\"")
		println("m[1] = \"Lardon\"")
		m[0] = "Ceifei"
		m[1] = "Lardon"
		fmt.Printf(">> m  %v\n", m)
		println()

		//delete()
		println("内置函数delete()用来删除映射的元素：")
		println("delete(m, 3)")
		delete(m, 3)
		fmt.Printf(">> m  %v\n", m)
		println("")

		//nil
		println("空映射与零值的映射是不同的：")
		println("var nilm map[string]int")
		println("fmt.Println(nilm == nil, len(nilm) == 0)")
		var nilm map[string]int
		fmt.Println(">>", nilm == nil, len(nilm) == 0)
		println("零值的映射未初始化，值还是nil，元素个数也0，此时也不能通过下标访问。")
		println("vm := map[string]int{}")
		println("fmt.Println(vm == nil, len(vm) == 0)")
		vm := map[string]int{}
		fmt.Println(">>", vm == nil, len(vm) == 0)
		println("空映射元素个数也是0，但值不是nil，元素个数也0，此时可以通过下标访问。")
		println("")
	}()

	//3-4
	func() {
		println("3-4 高级")
		println("----------------")
		println("可以使用range来遍历映射：")
		println("for key, val := range m")
		m := map[string]string{
			"age":    "20",
			"name":   "Ceifei",
			"height": "178",
		}
		for key, val := range m {
			fmt.Println(">> " + key + " : " + val)
		}
		println("")

		//Sets
		println("Go语言没有集合类型（Sets），可以通过映射来模拟：")
		println("set := make(map[string]bool)")
		set := make(map[string]bool)
		fmt.Printf(">> set  %v\n", set)
		println("")

		//Recursion
		println("映射的值可以是另一种映射，因此可以构建嵌套的映射：")
		//println("set := make(map[string]bool)")
		graph := make(map[string]map[string]bool)
		fmt.Printf(">> graph %T  %[1]v\n", graph)
		println("")
		println()
	}()
}
