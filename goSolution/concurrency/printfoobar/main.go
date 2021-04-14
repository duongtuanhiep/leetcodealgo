package main

import (
	"fmt"
	"sync"
)

/*
Question 1115: https://leetcode.com/problems/print-foobar-alternately/
**/

func main() {
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number received : %v \nResult: %v\nMutex res: %v\nNo Mutex res: %v", num, PrintFooBar(num), PrintFooBarMutex(num), PrintFooBarNoMutex(num))
	// fmt.Printf("Number received : %v \nResult: %v\nMutex res: %v\n", num, PrintFooBar(num), PrintFooBarMutex(num))
}

func PrintFooBar(n int) string {
	var res string
	foo := PrintFoo(n)
	bar := PrintBar(n)
	for i := 0; i < n; i++ {
		res += <-foo
		res += <-bar
	}
	return res
}

func PrintFoo(n int) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < n; i++ {
			c <- "foo"
		}
	}()
	return c
}

func PrintBar(n int) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < n; i++ {
			c <- "bar"
		}
	}()
	return c
}

func PrintFooBarNoMutex(n int) string {
	var res string
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			res = PrintFooStr(res)
			res = PrintBarStr(res)
			wg.Done()
		}()
	}
	wg.Wait()
	return res
}

func PrintFooBarMutex(n int) string {
	lock := &sync.Mutex{}
	wg := sync.WaitGroup{}
	var res string
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			lock.Lock()
			res = PrintFooStr(res)
			res = PrintBarStr(res)
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return res
}

func PrintFooStr(str string) string {
	return str + "foo"
}

func PrintBarStr(str string) string {
	return str + "bar"
}
