package main

import (
	"fmt"
	"time"
)

// 装弹函数
func f1(n int, 装弹情况 chan bool, 瞄准情况 chan bool) {
	for i := 0; i < n; i++ {
		if <-装弹情况 {
			time.Sleep(1 * time.Second)
			fmt.Print("装弹->")
			瞄准情况 <- true
		}
	}
}

// 瞄准函数
func f2(n int, 瞄准情况 chan bool, 发射情况 chan bool) {
	for i := 0; i < n; i++ {
		if <-瞄准情况 {
			time.Sleep(1 * time.Second)
			fmt.Print("瞄准->")
			发射情况 <- true
		}
	}
}

// 发射函数
func f3(n int, 继续开炮 chan bool, 发射情况 chan bool, 装弹情况 chan bool) {
	defer close(继续开炮)
	for i := 0; i < n; i++ {
		if <-发射情况 {
			time.Sleep(1 * time.Second)
			fmt.Println("发射!")
			装弹情况 <- true
		}
	}
}

func main() {
	//确定开炮数
	var n int
	fmt.Println("输入开炮数：")
	fmt.Scan(&n)

	装弹情况 := make(chan bool, 1)
	瞄准情况 := make(chan bool)
	发射情况 := make(chan bool)
	继续开炮 := make(chan bool)

	go f1(n, 装弹情况, 瞄准情况)
	go f2(n, 瞄准情况, 发射情况)
	go f3(n, 继续开炮, 发射情况, 装弹情况)

	装弹情况 <- true
	<-继续开炮
}
