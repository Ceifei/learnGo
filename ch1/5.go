package ch1

//F5 about Comments
func F5() {

	println()

	//1-10
	func() {
		println("1-10 注释")
		println("----------------")
		println("注释是保证程序可读性的必备技术，Go语言的注释方式与C语言完全一样。")
		println("// 单行注释")
		println(`/*
			多行注释
			多行注释
		*/ `)
		println("注意，多行注释是不能嵌套的。")
		println()
	}()
}
