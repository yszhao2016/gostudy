package studybl

import "fmt"

//定义好返回值
func Studyfunc1(param1 int, param2 string) (res1 int, res2 string) {
	res1 = param1
	res2 = param2
	return
}

//任意个参数
func Studyfunc2(param1 int, param2 ...string) {

	for k, v := range param2 {
		fmt.Println(k, "=>", v)
	}

	fmt.Println()
	fmt.Println(param1)

}

//闭包函数

func Mo() func(int) {
	return func(num int) {
		fmt.Println("闭包函数", num)
	}
}
