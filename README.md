# go-leren-jiegouti

指针接收者（Pointer Receiver）
当一个方法使用指针接收者时，它能够修改接收者结构体的内容。在方法定义中，指针接收者写法如下：

go
复制代码
func (c *calc) display(newtext string) {
// 方法的实现代码
}
func (c *calc) 表示这是一个方法，方法的接收者为 calc 结构体类型。
(c *calc) 中的 *calc 表示接收者类型为 calc 结构体的指针。这意味着，调用该方法时，会操作结构体本身，而不是其副本，可以修改结构体中的字段内容。
为什么使用指针接收者？
使用指针接收者的主要原因有几点：

1.修改接收者的内容：指针接收者允许方法修改接收者结构体中的字段内容。在方法内部对接收者的任何更改都会直接反映在调用者持有的结构体实例上。

2.避免复制大对象：当结构体较大时，使用指针接收者避免了方法调用时的结构体复制，提高了性能效率。

3.一致性：如果某个方法需要修改接收者的内容或者对其执行特定的操作（例如更新 UI 元素），通常使用指针接收者是更自然的选择。

=============================
Explanation:
Pointer Receiver Methods:

func (p *Person) PrintInfo() 和 func (p *Person) Pin() 方法现在使用了指针接收者 \*Person。
这意味着方法可以修改 Person 结构体实例中的字段，并且方法调用时，会直接操作指向 Person 实例的指针。
Creating Pointer to Struct:

在 main 函数中，创建 person 变量时使用了结构体的指针 &Person{...}。
这允许方法调用能够修改 person 指向的实际结构体内容。
Calling Methods:

方法调用仍然保持不变，通过 person.PrintInfo() 和 person.Pin() 调用结构体的方法。
在方法内部，p.FirstName、p.LastName 和 p.Age 都可以直接访问和修改，因为 p 是 \*Person 类型，即指向 Person 结构体的指针。
Why Use Pointer Receivers?
Mutability: 指针接收者允许方法修改结构体的字段，而不仅仅是操作副本。
Performance: 避免在方法调用时对大结构体进行复制，因为指针传递的开销更小。
Consistency: 当方法需要修改接收者的状态时，使用指针接收者更为自然和一致。
