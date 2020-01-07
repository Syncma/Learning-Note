# Go 语言学习笔记 -第6章


## 方法

```go

package main

import "fmt"

//结构体
type Person struct {
    Name string
}

//给A类型绑定一个方法
//小写p 表示哪个Pesrson变量调用, 就是它的副本
//小写p(最好写person)
func (p Person) test() {
    p.Name = "jack"
    fmt.Println("Person() name=", p.Name)
}

func main() {

    var p Person
    p.Name = "tom"
    p.test() //调用方法
    fmt.Println("main() name=", p.Name)
}

```