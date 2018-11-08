package ch3

import (
	"fmt"
)

//F1 about Arrays and Slices
func F1() {

	println()

	//3-1
	func() {
		println("3-1 数组")
		println("----------------")
		println("数组（arrays）是多个同种类型数据的集合体。声明数组的方法如下：")
		println("var arr [3]int")
		var arr [3]int
		println("数组通过下标来访问其中的成员：")
		println("arr[0] = 12")
		arr[0] = 12
		println("如果下标越界，编译器将报错：")
		println("arr[3] = 30 //error")
		//arr[3] = 30
		println("内置函数len()可以获取数组的长度：")
		println("fmt.Printf(\"%T %[1]d\\n\", len(arr))")
		fmt.Printf(">> %T %[1]d\n", len(arr))
		println()

		//array literals
		println("可以使用数组字面值赋值：")
		println("arr = [3]int{22, 33}")
		arr = [3]int{22, 33}
		fmt.Println(">>", arr)
		println("注意，字面值的类型必需和数组完全一致，且字面值元素个数个数不能超过类型中声明的个数。")
		println()

		println("如果是声明数组时用字面量初始化，可以使用省略号：")
		println("arr2 := [...]int{1, 5, 10, 15, 25}")
		arr2 := [...]int{1, 5, 10, 15, 25}
		fmt.Println(">>", arr2)
		println("这个时候，数组的大小将根据字面值的实际值自动设置。")
		println()

		println("字面值还可以指定下标：")
		println("arr = [3]int{0: 77, 2: 88}")
		arr = [3]int{0: 77, 2: 88}
		fmt.Println(">>", arr)
		println("索引都必需是整型，从0开始。")
		println()

		println("数组的大小在声明的时候就确定了，之后也不能修改：")
		println("arr[]= 13 //error")
		//arr[]= 13 //error
		println()

		//Subarray
		println("与字符串一样，数组支持返回子数组：")
		println("arr3 := arr[1:2]")
		arr3 := arr[1:2]
		fmt.Println(">>", arr3)
		println()
	}()

	//3-2
	func() {
		println("3-2 切片")
		println("----------------")
		println("切片（Slices）的本质就是数组，但是大小可以动态调整：")
		println("var arr []int")
		var arr []int
		fmt.Printf(">> %T %[1]v %d\n", arr, cap(arr))
		arr = []int{12, 24, 36}
		fmt.Printf(">> %T %[1]v %d\n", arr, cap(arr))
		arr = []int{12, 24, 36, 48}
		fmt.Printf(">> %T %[1]v %d\n", arr, cap(arr))
		println("注意，切片虽然大小可以调整，但仍然下标不能越界：")
		println("arr[4] //error")
		//arr[4] //error
		println("Go语言内置函数cap()可以计算切片的容量大小：")
		fmt.Printf(">> %T %d\n", arr, cap(arr))
		println()

		//make()
		println("另一个内置函数make()可以快速创建切片：")
		println("arr2 := make([]int, 6, 10)")
		arr2 := make([]int, 6, 10)
		fmt.Printf(">> arr2 %T %[1]d %d\n", arr2, cap(arr2))
		println("make()也可以用来重置切片：")
		println("arr = make([]int, 10)")
		arr = make([]int, 4)
		fmt.Printf(">> arr %T %[1]d %d\n", arr, cap(arr))
		println()

		//append()
		println("内置函数append()可以快速添加元素到切片：")
		arr = append(arr, 99)
		fmt.Printf(">> arr %T %[1]d %d\n", arr, cap(arr))
		println("注意到，append()函数可以将切片自动扩容。另外，append()可以同时添加多个元素：")
		println("arr = append(arr, 100, 101)")
		arr = append(arr, 100, 101)
		fmt.Printf(">> arr %T %[1]d %d\n", arr, cap(arr))
		println("append()第二个参数也可以是切片：")
		arr = append(arr, arr2...)
		fmt.Printf(">> arr %T %[1]d %d\n", arr, cap(arr))
		println("注意，此时第二个参数后面要加省略号。")
		println()

		//copy()
		println("如果想复制一个切片，可以使用内置的copy()函数：")
		println("var arr3 = make([]int, len(arr))")
		println("copy(arr3, arr)")
		var arr3 = make([]int, len(arr))
		copy(arr3, arr)
		fmt.Printf(">> arr3 %T %[1]d %d\n", arr3, cap(arr3))
		println("这里copy()将arr的值复制给了arr3。注意，arr3的大小必需与arr一致，否则将复制失败。")
		println("还应注意，copy()只复制元素，不影响目标切片的容量。")
		println()

		println("Go语言没有为切片专门设计删除元素的方法，不过可以通过拼接子切片达到删除的目的：")
		println("i := 5")
		println("arr = append(arr[:i], arr[i+1:]...)")
		i := 5
		arr = append(arr[:i], arr[i+1:]...)
		fmt.Printf(">> arr %T %[1]d %d\n", arr, cap(arr))
		println("类似的操作还有插入。")
		println()
	}()
}
