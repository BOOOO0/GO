package main

import (
	"fmt"
	"reflect"
	"strings"
)

type person struct{
	name string
	age int
	favFood []string
}

func plus (a int, b int) (int) {
	return a + b
}

func plusMinus (a int, b int) (int, int) {
	return a + b, a - b
}

func lenAndUppers (name string) (int, string) {
	return len(name), strings.ToUpper(name) 
}

func lenAndUppers2 (name string) (length int, uppers string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppers = strings.ToUpper(name)
	return
}


func spread (words ...string) {
	fmt.Println(words)
	fmt.Println(reflect.TypeOf(words))
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}
}

func superAdd (numbers ...int) int {
	for _, number := range numbers {
		fmt.Println( number)
	}
	return 1
}

func ageCheck (age int) bool {
	// if age > 18 {
	// 	return true
	// }
	// if문에서 변수 생성이 가능하다.
	if koreanAge := age + 2 ; koreanAge > 20 {
		return true
	}
	return false
}

func main() {
	const name string = "Boo"

	// 배열의 값을 타입 앞에 명시하고 = 없이 초기화
	names := [5]string{"정", "부", "영"}
	names[3] = "asdasd"
	names[4] = "xzcmjkzx"

	// 슬라이스
	names2 := []string{"정", "부", "영"}
	names2 = append(names2, "입니다.")

	// 맵
	boo := map[string]string{"name" : "boo", "age" : "19"}
	favFood := []string{"asd", "dfg"}
	boo0 := person{name: "boo", age: 18, favFood: favFood}


	fmt.Println("Hello")
	fmt.Println(name)
	fmt.Println(plus(1,2))
	fmt.Println(plusMinus(1,2))

	// 새로운 변수에 바로 값을 할당할 때 := 사용, 타입도 그때 값에 맞게 정해진다.
	length, upper := lenAndUppers(name)

	fmt.Println(length, upper)

	fmt.Println(lenAndUppers2(name))
	spread("boo", "young", "chung")
	superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(ageCheck(19))
	fmt.Println(names)
	fmt.Println(names2)
	fmt.Println(boo)

	// range로 순회시 map의 키와 밸류 각각 접근
	for key, value := range boo {
		fmt.Println(key, value)
	}
	fmt.Println(boo0, boo0.name)
}