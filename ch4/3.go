package ch4

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

//F3 about Structs
func F3() {

	println()

	//3-5
	func() {
		println("3-5 结构体")
		println("-----------------------")
		//Declaration
		println("结构体与数组类似，但其属性的数据类型可以不同：")
		println("type Person struct {")
		println("	Name string")
		println("	Age  int")
		println("	Sex  string")
		println("}")
		type Person struct {
			Name string
			Age  int
			Sex  string
		}
		println("这里定义了一个名为Person的结构体，包含三个属性。注意，每个属性结尾没有逗号。")
		println("Go语言中，结构体是对象的模板，根据结构体可以定义对象：")
		println("var p Person")
		var p Person
		fmt.Printf(">> p %T %[1]v %d\n", p, unsafe.Sizeof(p))
		println("注意到，一个没有赋值的结构体对象占据了40个字节。")
		println()

		//Initialization
		println("结构体对象通过点号访问其属性：")
		println("p.Name = \"Ceifei\"")
		println("p.Age = 22")
		p.Name = "Ceifei"
		p.Age = 22
		fmt.Printf(">> p %T %[1]v %d\n", p, unsafe.Sizeof(p))
		println()

		println("可以通过结构体字面量来初始化对象：")
		p = Person{
			Name: "Ceifei Lardon",
			Age:  21,
			Sex:  "M"}
		fmt.Printf(">> p %T %[1]v %d\n", p, unsafe.Sizeof(p))
		println("还可以把字面值简化：")
		println("p = Person{\"Kobe Bryant\", 21, \"M\"}")
		p = Person{"Kobe Bryant", 21, "M"}
		fmt.Printf(">> p %T %[1]v %d\n", p, unsafe.Sizeof(p))
		println("花括号内的值与结构体属性的个数和顺序一致。")
		println()

		//Unexported fields
		println("Go语言中，结构体属性的首字母大小写决定了可访问性：")
		println("type T struct{a, b int}")
		type T struct{ a, b int }
		println("T中的属性都是小写字母开头，因此在包外均不能被访问。")
		println()

		//Comparsion
		println("结构体对象可以进行比较：")
		println("p1 := Point{1, 2}")
		println("p2 := Point{1, 2}")
		type Point struct{ X, Y int }
		p1 := Point{1, 2}
		p2 := Point{1, 2}
		fmt.Println(">> p1 == p2", p1 == p2)
		println("只有两个对象的所有属性值相等时，对象才相等：")
		println("p2.X = 2")
		p2.X = 2
		fmt.Println(">> p1 == p2", p1 == p2)
		println()
	}()

	//3-6
	func() {
		println("3-6 高级")
		println("----------------")
		//Embedding
		println("一个结构体可以嵌入另外一个结构体：")
		type Field struct {
			Name    string
			Type    string
			Default string
			IsNull  int
		}
		type Table struct {
			Name string
			Field
		}
		println("Table结构体嵌入了Field结构体，注意定义的顺序。")
		var tb Table
		tb.Field.Name = "ID"
		tb.Field.Type = "INT"
		tb.Field.IsNull = 0
		fmt.Printf(">> tb %T %#[1]v %d\n", tb, unsafe.Sizeof(tb))
		println("运用嵌入，可以构建复杂的树状结构体")
		println()

		//JSON
		println("结构体与JSON类似，Go语言支持在结构体中添加JSON标记：")
		type Movie struct {
			Title string
			Year  int  `json:"released"`
			Color bool `json:"color,omitempty"`
		}
		println("反括号内就是JSON标记，用于转换json字符串。")
		var movies = []Movie{
			{"Casablanca", 1942, false},
			{"Cool Hand Luke", 1967, true},
			{"Bullitt", 1968, true},
		}
		json, _ := json.Marshal(movies)
		fmt.Printf("%s\n", json)
		println()
	}()

}
