package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널을 통해 받을 데이터의 타입을 명시
	c := make(chan string)
	people := [2]string{"boo", "young"}
	for _, person := range people {
		go Check(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// go Count("Boo")
	// Count("Young")
}

func Check (name string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- "Check" + name
}

func Count(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, i+1)
		time.Sleep(time.Second)
	}
}


