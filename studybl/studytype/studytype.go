package main

import "fmt"

/**
	func 的基本构成元素
		func  (p myType )  funcName ( a, b int , c string ) ( r , s int ) {
			return
		}

	其中：
		关键字———func // 这个是定义函数的关键字
		函数拥有者—(p myType) // 这个是此函数的拥有者，下面解释（此项可省略）
		方法名———funcName // 这个是定义函数的名字
		入参———— a,b int,b string // 这个是定义函数的入参
		返回值——— r,s int // 这个是定义函数的返回值，golang可以返回多个值
		函数体——— { }

		重点说说这个函数拥有者(p myType)，这个是相较于C/C++比较特殊的地方。
		为特定类型定义函数，即为类型对象定义方法
		在Go中通过给函数标明所属类型，来给该类型定义方法，上面的 （p myType） 即表示给myType声明了一个方法， p myType 不是必须的。
		如果没有，则纯粹是一个函数。
 */
type myint int

type myfunc func(interface{}) interface{}

func (test myfunc) abc(a int) string  {
	return "string"
}

func (a myint) testmyint(b int) int{
	fmt.Println(a)
	return b
}

func (a *myint) testmyint1() {
	fmt.Println(*a)
}
func main()  {
	var a myint
	var b myfunc
	a = 100
	fmt.Println(a.testmyint(100))

	fmt.Println(b.abc(1))

	(&a).testmyint1()
	a.testmyint1()
}



