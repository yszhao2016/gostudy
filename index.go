package main

import (
	"fmt"
	"studygo/studybl"
)

const SS = "234dfgdfg" //它是常量

func main() {
	var test string
	fmt.Println("----变量、常量学习-----")

	studybl.Studyconstant()

	//特殊只写变量 "_"，⽤于忽略值占位。
	_, test = studybl.Studyconstant1()
	fmt.Println(test)
	fmt.Println(SS)

	//枚举类型
	studybl.Studyiota()
	temp := studybl.Black
	studybl.Studyiota1(temp)
	studybl.Studyiota1(67)
	fmt.Println("----变量、常量学习-----")

	studybl.Aa = 123
	fmt.Println(studybl.Aa)
	fmt.Println("----Test Slice-----")
	studybl.Testslice()
	studybl.Testslice2()

	fmt.Println("----Test Slice-----")
	// studybl.myslice()
	// studybl.myslice()

	fmt.Println("----Study Map-----")

	studybl.Studymap()

	fmt.Println("----Study Map-----")

	fmt.Println("----Study function-----")

	fmt.Println(studybl.Studyfunc1(1, "test"))
	studybl.Studyfunc2(1, "a", "b", "c", "d")

	temp1 := []string{"aa", "bb", "cc", "dd"}

	//defer 延迟执行函数  最后执行
	defer studybl.Studyfunc2(2, temp1...)

	(func() {
		fmt.Println("我是自执行函数  应用场景  协程  应用")
	})()

	studybl.Mo()(4)
	fmt.Println("----Study function-----")

	fmt.Println("----指针学习-----")

	studybl.Studypointer()

	fmt.Println("----指针学习-----")

	fmt.Println("----结构体学习-----")
	studybl.Studystruct()
	fmt.Println("----结构体学习-----")

	fmt.Println()
	fmt.Println()
	fmt.Println("----接口学习-----")
	studybl.Studyinterface()
	fmt.Println("----接口学习-----")

	fmt.Println()
	fmt.Println()
	fmt.Println("----chan 学习-----")
	studybl.Studychannel()
	fmt.Println("----chan 学习-----")

	fmt.Println()
	fmt.Println()
	fmt.Println("----断言反射 学习-----")
	studybl.StudyAssertion()
	fmt.Println("----断言反射 学习-----")

	fmt.Println("----***-----")

	fmt.Println("----***-----")
}

//	空指针值 nil
