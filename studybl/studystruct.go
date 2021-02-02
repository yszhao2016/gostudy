package studybl

import (
	"encoding/json"
	"fmt"
)

type people struct {
	name  string
	sex   bool
	age   int
	hobby []string
	home
}

type home struct {
	p string
}

func (q *people) Song(name string) (res string) {
	res = "真的给我一记惊雷"
	fmt.Printf("%v唱了一首%v,观众觉得%v", q.name, name, res)
	return res

}

func (h *home) Open() {
	fmt.Println("我在home")
}

// tag可以为结构体的成员添加说明或者标签便于使用,这些说明可以通过反射获取到。
// 在前面提到了，结构体中的成员首字母小写对外不可见，
// 但是我们把成员定义为首字母大写这样与外界进行数据交互会带来极大的不便，此时tag带来了解决方法。

type Student1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Studystruct() {

	var test people
	test.name = "张三"
	test.sex = true
	test.age = 18
	test.hobby = []string{"篮球", "乒乓球", "看书"}

	test1 := people{
		name:  "zhangsan",
		sex:   true,
		age:   18,
		hobby: nil,
	}
	test.p = "json"
	test.Open()
	res := test1.Song("张三")
	println(res)
	test2 := new(people)

	var test3 *people

	test3 = &test1
	(*test3).name = "结构体指针"

	fmt.Println(test, test1, test2, test3)

	//结构体函数

	stu := Student1{Name: "wd", Age: 22}
	// fmt.Println(stu)
	data, err := json.Marshal(stu)
	fmt.Println()
	fmt.Println()
	if err == nil {
		fmt.Println("lalala")
		fmt.Println(string(data))
		fmt.Println("lalala")
	}

}
