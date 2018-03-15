package main

import (
	"fmt"
	// "time" import的东西必须要使用，不然编译错误
)

func main() {
	ch := make(chan string)

	// i := 0 声明变量并赋初值，自动根据赋值推导类型
	// goroutine可以承受的数量远远大于线程或者进程
	for i := 0; i < 5000; i++ {
		// go starts a goroutine
		go printHelloWorld(i, ch)
	}

	for {
		msg := <- ch
		fmt.Println(msg)
	}

	// 不等待的话，不会有任何输出，
	// 因为主函数的goroutine和开出的goroutine是并行执行的
	// 但是主函数的goroutine退出后，其他goroutine无论是否执行完都会被关闭
	// time.Sleep(100*time.Millisecond)
}

func printHelloWorld(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("Hello world from " +
			"goroutine %d!\n", i)
	}
}
