package main

//通过回调实现

import (
	"fmt"
	"time"
)

type FuncType func() string //用FuncType代表没有形参，有一个string型返回值的函数

func 欢迎来我家玩() string {
	time.Sleep(5 * time.Second)
	return "登dua郎"
}

func 打电动() string {
	return "输了啦，都是你害的"
}

func f(fType FuncType) (result string) {
	result = fType()
	return
}

func main() {
	s1 := f(打电动)
	fmt.Println(s1)
	s2 := f(欢迎来我家玩)
	fmt.Println(s2)

}
