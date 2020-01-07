# Go 语言学习笔记 -第7章


## 接口

```go
package main

import (
    "fmt"
)

//Usb 声明
//接口体现多态、高内聚低耦合思想
type Usb interface {
//声明两个没有实现的方法
//根据方法来判断
    Start()
    Stop()
}

//手机
type Phone struct {
}

//让Phone实现Usb接口的方法
func (p Phone) Start() {
    fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
    fmt.Println("手机停止工作...")
}

//照相机
type Camera struct {
}

//让Camera实现Usb接口的方法
func (c Camera) Start() {
    fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
    fmt.Println("相机停止工作...")
}

//计算机
type Computer struct {
}

//编写working方法,接收一个Usb接口类型变量
//只要是实现了Usb接口: 就是实现了Usb接口声明所有方法
func (c Computer) Working(usb Usb) {
    usb.Start()
    usb.Stop()

}
func main() {

    computer := Computer{}
    phone := Phone{}
    // camera = Camera{}

    computer.Working(phone)

}

```


##  Golang接口注意事项


1.接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）

```go
package main
import "fmt"

type Ainterface interface {
    Say()
}
type Stu struct {
    Name string
}
func (stu Stu) Say() {
    fmt.Printf("stu say()")
}
func main() {
    var stu Stu
    var a Ainterface = stu
    a.Say()
}
```


2.接口中所有的方法都没有方法体， 即都没有实现的方法

3.在Golang中，一个自定义类型需要将某个接口的所有方法都实现，我们所这个自定义类型实现了该接口

4.只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型

```go
package main
import "fmt"

type Ainterface interface {
    Say()
}
type Stu struct {
    Name string
}

func (stu Stu) Say() {
    fmt.Printf("stu say()")
}

type integer int
func (i integer) Say() {
    fmt.Println("integer  say i =", i)
}
func main() {
    var stu Stu
    var a Ainterface = stu
    a.Say()
    var i integer = 10
    var b Ainterface = i
    b.Say()
}
```


5.一个自定义类型可以实现多个接口

```go
package main
import "fmt"

type Ainterface interface {
    Say()
}
type Stu struct {
    Name string
}

func (stu Stu) Say() {
    fmt.Printf("stu say()")
}

type integer int
func (i integer) Say() {
    fmt.Println("integer  say i =", i)
}

type BInterface interface {
    Hello()
}
type Monster struct {
}

func (m Monster) Hello() {
    fmt.Println("Monster Hello()")
}
func (m Monster) Say() {
    fmt.Println("Monster Say()")
}

func main() {
    var stu Stu
    var a Ainterface = stu
    a.Say()
    var i integer = 10
    var b Ainterface = i
    b.Say()
    var monster Monster
    var a2 Ainterface = monster
    var b2 BInterface = monster
    a2.Say()
    b2.Hello()
}
```


6.Golang接口不能有任何变量

7.一个接口(比如A接口)可以继承多个别的接口(比如B,C接口），这时如果要实现A接口，也必须将B,C接口的方法也全部实现

```go
package main

import "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
type AInterface interface {
	BInterface
	CInterface
	test03()
}

//需要实现AInterface, 需要将BInterface,CInterface的方法都实现
type Stu struct {
}

func (stu Stu) test01() {
	fmt.Println("test01")
}
func (stu Stu) test02() {
	fmt.Println("test02")
}
func (stu Stu) test03() {
	fmt.Println("test03")
}
func main() {
	var stu Stu
	var a AInterface = stu
	a.test01()
}



```


8.interface类型默认是一个指针(引用类型),如果没有对interface初始化使用，那么会输出nil



9.空接口interface{} 没有任何方法， 因此所有类型都实现了空接口
可以把任何变量付给空接口

```go
package main

import "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
type AInterface interface {
	BInterface
	CInterface
	test03()
}

//需要实现AInterface, 需要将BInterface,CInterface的方法都实现
type Stu struct {
}

func (stu Stu) test01() {
	fmt.Println("test01")
}
func (stu Stu) test02() {
	fmt.Println("test02")
}
func (stu Stu) test03() {
	fmt.Println("test03")
}

type T interface {
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.test01()
	var t T = stu
	fmt.Println("t=", t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println("t2=", t2)
	fmt.Println("t=", t)
}

```