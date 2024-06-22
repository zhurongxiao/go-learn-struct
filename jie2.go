package main

import "fmt"

// 定义一个结构体 Person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// 定义一个结构体的方法，使用指针接收者
func (p *Person) PrintInfo() {
	fmt.Printf("Name: %s %s, Age: %d\n", p.FirstName, p.LastName, p.Age)
}

// 修改结构体实例的 FirstName 字段为 "apple"
func (p *Person) Pin() {
	p.FirstName = "apple" // 修改 FirstName 为 "apple"
	fmt.Printf("Name: %s %s, Age: %d\n", p.FirstName, p.LastName, p.Age)
}

func main() {
	// 创建一个 Person 结构体的指针实例
	person := &Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// 调用结构体的方法打印信息
	person.PrintInfo()

	// 调用修改 FirstName 的方法
	person.Pin()

	// 打印修改后的信息
	person.PrintInfo()
}
