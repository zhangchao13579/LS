package main

import (
	"fmt"
)

//编写一个将a[l]~a[r]从小到大排序的函数

func BubbleSort(a []int, l int, r int) {
	for i := l; i <= r; i++ {
		for j := l; j < r-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	var n int
	var l int
	var r int
	fmt.Print("输入符合要求的三个整数：")
	fmt.Scan(&n, &l, &r)
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Printf("请输入第%d个元素：", i+1)
		fmt.Scan(&a[i])
	}
	fmt.Println(a)
	b := a[:]
	BubbleSort(b, l, r)
	fmt.Println(b)
}
