package main

import "fmt"

// 定义一个结构体 Person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// 定义一个结构体的方法
func (p Person) PrintInfo() {
	fmt.Printf("Name: %s %s, Age: %d\n", p.FirstName, p.LastName, p.Age)
}

func (p Person) Pin() {
	fmt.Printf("Name: %s %s, Age: %d\n", p.LastName, p.FirstName, p.Age)
}

func main() {
	// 创建一个 Person 结构体的实例
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// 调用结构体的方法打印信息
	person.PrintInfo()

	person.Pin()
}
