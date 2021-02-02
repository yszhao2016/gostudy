package studybl

import "fmt"

/*
缓存区
go
去过一次就没了
close for range
select  //随机执行
*/
func Studychannel() {

	// 定义缓存区
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch)

	//无缓冲区chan
	ch1 := make(chan int)
	go func() {
		ch1 <- 22
	}()

	fmt.Println(<-ch1)

	// 打印没有，就去存
	ch2 := make(chan int, 5)

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i
		}
	}()

	for j := 0; j < 10; j++ {
		fmt.Println(<-ch2)
	}

	ch3 := make(chan int, 3)
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	//不关闭(close)  for range 会报错
	close(ch3)
	for value := range ch3 {
		fmt.Println(value)
	}
	ch <- 23
	ch2 <- 3
	// chan  selelct 用法

	select {
	case <-ch:
		fmt.Println("ch")
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	default:
		fmt.Println("啥也没。。。")
	}

	//
	ch4 := make(chan int)
	var readc <-chan int = ch4
	var writec chan<- int = ch4
	go SetChan(writec)
	GetChan(readc)
}

func SetChan(writec chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("我是在set函数里面")
		writec <- i

	}

}

func GetChan(readc <-chan int) {
	for j := 0; j < 10; j++ {
		fmt.Println("我是在get函数里面,取出setchan 返回给我信息是", <-readc)
	}

}
