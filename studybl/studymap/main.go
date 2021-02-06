package main

import "fmt"

func main()  {
	var m map[string]interface{}
	var arr []map[string]interface{}
	m = map[string]interface{}{
		"abc" :1,
		"b" : 2,
		"c" :map[string]int{
			"a":1,
			"b":2,
			"c":3,
		},
	}
	arr = []map[string]interface{}{map[string]interface{}{"a":1, "b":2},map[string]interface{}{"a":1, "b":2}}
	fmt.Println(arr[0]["a"])
	fmt.Println(m["c"])
}
