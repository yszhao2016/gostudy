package studybl

import "fmt"

func Studymap() {

	//map 三种赋值方法

	// 方法一
	var m map[string]string
	m = map[string]string{}

	m["name"] = "张三"
	m["sex"] = "女"

	// 方法二
	m1 := map[string]string{}

	// 方法三
	m2 := make(map[string]string)

	fmt.Println(m, m1, m2)

	m3 := map[int]interface{}{}
	m3[1] = 2
	m3[2] = [3]int{1, 2, 3}
	m3[3] = "abc"
	fmt.Println()
	fmt.Println(m3)
}
