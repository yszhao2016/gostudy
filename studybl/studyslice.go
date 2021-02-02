package studybl

import "fmt"

func init() {
	fmt.Println("studybl init test")
}

func Testslice() {
	var a []int
	var b = []int{1, 2, 3}
	var c = []string{"a", "b", "c"}

	// append 尾部追加
	a = append(a, 2)
	b = append(b, 4)

	// append 头部添加
	// c = append(c[1:2], "d")
	c = append([]string{"f", "e"}, c...)

	fmt.Println(a, b, c)
}

func Testslice2() {
	var a = []int{5, 6}
	a = append(a[:1], append([]int{8, 9}, a[1:]...)...) // 在第i个位置插入x
	fmt.Println(a)
	a = append(a[1:], append([]int{1, 2, 3}, a[1:]...)...) // 在第i个位置插入切片
	fmt.Println(a)
}
