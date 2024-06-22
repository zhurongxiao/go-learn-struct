package main

import "fmt"

// 定义一个函数打印人的信息
func PrintPersonInfo(firstName, lastName string, age int) {
	fmt.Printf("Name: %s %s, Age: %d\n", firstName, lastName, age)
}

func main() {
	// 直接调用函数打印信息
	PrintPersonInfo("John", "Doe", 30)
}
