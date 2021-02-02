package studybl

import "fmt"

func Studypointer() {

	var a int
	var b *int
	a = 123
	b = &a
	*b = 456
	a = 789
	fmt.Println(b)
	fmt.Println(a, *b)

	//数组指针和指针数组
	var c [5]int
	var d *[5]int
	c = [5]int{1, 2, 3, 4, 5}
	d = &c
	fmt.Println(c[1], d[1])

	var dpp [5]*int
	in1 := 1
	in2 := 2
	in3 := 1
	in4 := 2
	in5 := 1

	dpp = [5]*int{&in1, &in2, &in3, &in4, &in5}

	fmt.Println(dpp[1])
	fmt.Println(*dpp[1])
}
