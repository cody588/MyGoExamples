package main

import "fmt"

func main() {
	//make时，不指定cap，默认和len长度一样，且为len长度个0值
	a := make([]int, 5)
	fmt.Printf("a:%v, cap(a)=%d,len(a)=%d\n", a, cap(a), len(a))
	//make声明时，指定cap，len=0，则没有0值，且最大容量为5
	b := make([]int, 0, 5)
	fmt.Printf("b:%v, cap(b)=%d,len(b)=%d\n", b, cap(b), len(b))
	//对slice再取slice，则len和cap都变为了新的slice的长度，且有len个值，不足时补0
	b = b[:cap(b)]
	fmt.Printf("b:%v, cap(b)=%d,len(b)=%d\n", b, cap(b), len(b))
	b = b[1:]
	fmt.Printf("b:%v, cap(b)=%d,len(b)=%d\n", b, cap(b), len(b))
	b = b[1:]
	fmt.Printf("b:%v, cap(b)=%d,len(b)=%d\n", b, cap(b), len(b))
}

// package main

// import "fmt"

// func main() {
// 	a := make([]int, 5)
// 	printSlice("a", a)

// 	b := make([]int, 0, 5)
// 	printSlice("b", b)

// 	c := b[:2]
// 	printSlice("c", c)

// 	d := c[2:5]
// 	printSlice("d", d)
// }

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }
