package studybl

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	Class string
	Test  int
	User
}

func (q User) SayName(name string) {
	fmt.Println("我的名字叫", name)

}

func StudyAssertion() {
	u := User{
		Name: "lisi",
		Age:  16,
		Sex:  true,
	}
	s := Student{Class: "4年级3班", User: u}
	// check(u)
	// check(s)
	// check1(u)
	// check2(s)
	check3(&s)
}

func check(tp interface{}) {
	//断言
	switch tp.(type) {
	case User:
		tp.(User).SayName("用户张三")
		break
	case Student:
		tp.(Student).SayName("学生张武")
		break
	default:
		fmt.Println("断言失败")
	}
}

func check1(tp interface{}) {
	a := reflect.TypeOf(tp)
	b := reflect.ValueOf(tp)
	fmt.Println(a)
	for i := 0; i < b.NumField(); i++ {
		fmt.Println(b.Field(i))
	}

}

func check2(tp interface{}) {
	a := reflect.TypeOf(tp)
	b := reflect.ValueOf(tp)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b.FieldByIndex([]int{0}), b.FieldByIndex([]int{1, 0}))
	// fmt.Println(b.FieldByName("class"))
	ty := a.Kind()
	if ty == reflect.Struct {
		fmt.Println("Struct")
	} else if ty == reflect.Array {
		fmt.Println("Array")
	} else if ty == reflect.Bool {
		fmt.Println("Bool")
	}
}
func check3(tp interface{}) {
	a := reflect.ValueOf(tp)
	e := a.Elem()
	e.FieldByName("Test").SetInt(12)
	e.FieldByName("Class").SetString("受到法律的是否")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(e)

	v := reflect.ValueOf(tp)
	m := v.Method(0)
	m.Call([]reflect.Value{reflect.ValueOf("大大大")})
	fmt.Println(m)
}
