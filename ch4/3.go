package ch4

import (
	"fmt"
	"math"
	"time"
)

//Point is a struct
type Point struct{ X, Y float64 }

//Distance is a method associated with Point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) scale(factor float64) {
	p.X *= factor
	p.Y *= factor
}

//F3 about Structs
func F3() {

	println()

	//4-7
	func() {
		println("4-7 方法")
		println("-----------------------")
		//Example
		println("Go语言中，方法是绑定了数据类型的函数：")
		println("const day = 24 * time.Hour")
		println("fmt.Println(day.Seconds())")
		const day = time.Hour * 24
		fmt.Println(">> ", day.Seconds())
		println("这里day是一个整型常量，但仍然可以拥有方法。")
		println()

		//Declaration
		println("定义一个方法：")
		println("func (p Point) Distance(q Point) float64 {")
		println("	return math.Hypot(q.X-p.X, q.Y-p.Y)")
		println("}")
		println("其中p被称为“接受者”，是方法调用的主体。")
		println("与命名函数一样，方法只能定义在包级别，不能定义在其他函数内部。")
		println("Go语言不允许定义匿名方法。")
		println()

		//Invoke
		println("调用方法：")
		println("p := Point{0, 0}")
		println("q := Point{3, 4}")
		println("fmt.Println(p.Distance(q))")
		println("fmt.Println(q.Distance(p))")
		p := Point{0, 0}
		q := Point{3, 4}
		fmt.Println(">> ", p.Distance(q))
		fmt.Println(">> ", q.Distance(p))
		println("p和q都是Point对象，因此都可以调用Distance方法。")
		println()

		//Pointer Receivers
		println("当需要操作调用者本身时，方法的接受者应该是一个指针：")
		println("func (p *Point) scale(factor float64) {")
		println("	p.X *= factor")
		println("	p.Y *= factor")
		println("}：")
		println("q.scale(2)")
		println("fmt.Println(q)")
		q.scale(2)
		fmt.Println(">> ", q)
		println()
	}()

	//4-8
	func() {
		println("4-8 ")
		println("----------------")

		println()
	}()

}
