package studybl

import "fmt"

type Animal interface {
	Eat()
	Run()
}

type Dog struct {
	Name string
	sex  int
}

type Cat struct {
	Name string
}

func (c Dog) Run() {
	fmt.Printf("%v在跑\n", c.Name)
}

func (c Dog) Eat() {
	fmt.Printf("%v在吃东西\n", c.Name)
}
func (c Cat) Run() {
	fmt.Printf("%v在跑\n", c.Name)
}

func (c Cat) Eat() {
	fmt.Printf("%v在吃东西\n", c.Name)
}

func Studyinterface() {

	// 方法一
	var a Animal
	var c Animal
	var d Animal
	var b Dog

	//b Dog
	b = Dog{
		Name: "大黄",
		sex:  3,
	}
	//c animal
	c = Dog{
		Name: "大黄2号",
		sex:  3,
	}
	d = Cat{
		Name: "Tom",
	}
	a = b
	a.Eat()
	a.Run()
	c.Run()
	c.Eat()
	Mytest(c)
	Mytest(d)

}

//  接口应用  范形
func Mytest(a Animal) {
	a.Eat()
	a.Run()
}
