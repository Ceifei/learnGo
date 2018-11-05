package ch1

//F2 knowledges about Program Structure in Go
func F2() {

	println()

	//1-3: Packages
	func() {
		println("1-3 包声明")
		println("----------------")
		println("Go语言中，每一个源代码文件都必需在开头声明包。")
		println("一般来说，包的名字就是源代码文件所在文件夹的名字：")
		println("package ch1")
		println("注意，这里ch1是不带引号的。一个文件只能声明一个包。")
		println("一个文件夹下的源代码文件可以属于不同的包，不过习惯上只声明一个包。")
		println("main()函数所在的文件必需声明为main包，一个文件夹下只能有一个main包。")
		println("也就是说，main包是不可能被引用到其他包里的。")
		println()
	}()

	//1-4: Import Declarations
	func() {
		println("1-4 引用包")
		println("----------------")
		println("Go语言的源代码文件可以引用若干其他包：")
		println("import \"fmt\"")
		println("import \"ch1\"")
		println("注意，这里包名必需用双引号括起来。习惯上，引用多个包可以写作：")
		println("import(")
		println("        \"fmt\"")
		println("        \"os\"")
		println("        ")
		println("        \"github.com/Ceifei/learnGo/ch1\"")
		println("       )")
		println("这里括号中的一行表示一个包，注意其中的空行用来区分标准包与非标准包。")
		println("当多个包名称相同时，可以使用别名来区分：")
		println("import(")
		println("       mrand \"math/rand\"")
		println("       crand \"crypto/rand\"")
		println("       )")
		println("别名本身不含引号，该别名只当前源代码文件中有效")
		println("Go语言不允许引用包之后又不在代码中使用。如果无需在代码中使用包，可以使用_作为包的别名：")
		println("import _ \"image/png\"")
	}()

}
