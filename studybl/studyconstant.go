package studybl

import "fmt"

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func Studyconstant() {
	var a int = 3
	var b, c = "123", 123
	// var c, d string

	//自动定义类型  并且复制
	e := "123"
	fmt.Println(a, b, c, e)

	s := `ab

	s:=123
	cedf`
	println(s)
}

func Studyconstant1() (int, string) {
	return 1, "abc"
}

func Studyiota() {
	type ByteSize float64

	const (
		_           = iota             // 忽略了定义，初始化还是从0开始
		KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
		MB                             // 1 << (10*2)
		GB                             // 1 << (10*3)
		TB                             // 1 << (10*4)
		PB                             // 1 << (10*5)
		EB                             // 1 << (10*6)
		ZB                             // 1 << (10*7)
		YB                             // 1 << (10*8)
	)

	//const iota
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          //2
		Wednesday        //3
		Thursday         //4
		Friday           //5
		Saturday         //6
	)

	fmt.Println("TB ==", TB)
	fmt.Println("Monday = ", Monday)
}

func Studyiota1(c Color) {
	println("color 类型枚举常量", c)
}
