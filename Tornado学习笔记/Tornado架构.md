# Tornado 架构
<!-- TOC -->

- [Tornado 架构](#tornado-%e6%9e%b6%e6%9e%84)
  - [概述](#%e6%a6%82%e8%bf%b0)
  - [传统Web服务器架构](#%e4%bc%a0%e7%bb%9fweb%e6%9c%8d%e5%8a%a1%e5%99%a8%e6%9e%b6%e6%9e%84)
    - [例子](#%e4%be%8b%e5%ad%90)
    - [socket 介绍](#socket-%e4%bb%8b%e7%bb%8d)
    - [socket原理分析](#socket%e5%8e%9f%e7%90%86%e5%88%86%e6%9e%90)
  - [Tornado 架构](#tornado-%e6%9e%b6%e6%9e%84-1)
    - [架构分析](#%e6%9e%b6%e6%9e%84%e5%88%86%e6%9e%90)

<!-- /TOC -->
## 概述

tornado的http服务器采用的是: 

**多进程 + 非阻塞 + epoll + pre-fork 模型**



## 传统Web服务器架构



![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/web_normal.jpg)



简单描述：

1. 创建listen socket, 在指定的监听端口, 等待客户端请求的到来
2. listen socket接受客户端的请求, 得到client socket, 接下来通过client socket与客户端通信
3.  处理客户端的请求, 首先从client socket读取http请求的协议头, 如果是post协议, 还可能要读取客户端上传的数据, 然后处理请求, 准备好客户端需要的数据, 通过client socket写给客户端



### 例子

```python
import socket


def handle_request(client):
    buf = client.recv(1024)
    print(buf)

    client.send("HTTP/1.1 200 OK\r\n\r\n")
    client.send("Hello, World")


def main():
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.bind(('localhost', 8080))
    sock.listen(5)

    while True:
        connection, address = sock.accept()
        handle_request(connection)
        connection.close()


if __name__ == '__main__':
    main()

```



### socket 介绍

当客户端和服务器使用TCP协议进行通信时，客户端封装一个请求对象req，将请求对象req序列化成字节数组，然

后通过套接字socket将字节数组发送到服务器，服务器通过套接字socket读取到字节数组，再反序列化成请求对象

req，进行处理，处理完毕后，生成一个响应对应res，将响应对象res序列化成字节数组，然后通过套接字将自己

数组发送给客户端，客户端通过套接字socket读取到自己数组，再反序列化成响应对象。



通信框架往往可以将序列化的过程隐藏起来





### socket原理分析

这里使用动画来说明：

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/socket.gif)

**细节**

我们平时用到的套接字其实只是一个引用(一个对象ID)，这个套接字对象实际上是放在操作系统内核中。

这个套接字对象内部有两个重要的缓冲结构，一个是读缓冲(read buffer)，一个是写缓冲(write buffer)，它们都是有限大小的数组结构。

当我们对客户端的socket写入字节数组时(序列化后的请求消息对象req)，是将字节数组拷贝到内核区套接字对象

的write buffer中，内核网络模块会有单独的线程负责不停地将write buffer的数据拷贝到网卡硬件，网卡硬件再将

数据送到网线，经过一些列路由器交换机，最终送达服务器的网卡硬件中。

同样，服务器内核的网络模块也会有单独的线程不停地将收到的数据拷贝到套接字的read buffer中等待用户层来

读取。最终服务器的用户进程通过socket引用的read方法将read buffer中的数据拷贝到用户程序内存中进行反序

列化成请求对象进行处理。然后服务器将处理后的响应对象走一个相反的流程发送给客户端

这里就不再具体描述



**阻塞**

我们注意到write buffer空间都是有限的，所以如果应用程序往套接字里写的太快，这个空间是会满的。

一旦满了，写操作就会阻塞，直到这个空间有足够的位置腾出来。

不过有了NIO(非阻塞IO)，写操作也可以不阻塞，能写多少是多少，通过返回值来确定到底写进去多少，那些没有写进去的内容用户程序会缓存起来，后续会继续重试写入。

同样我们也注意到read buffer的内容可能会是空的。

这样套接字的读操作(一般是读一个定长的字节数组)也会阻塞，直到read buffer中有了足够的内容(填充满字节数组)才会返回。

有了NIO，就可以有多少读多少，无须阻塞了。

读不够的，后续会继续尝试读取。



**ack**

那上面这张图就展现了套接字的全部过程么？显然不是，数据的确认过程(ack)就完全没有展现。

比如当写缓冲的内容拷贝到网卡后，是不会立即从写缓冲中将这些拷贝的内容移除的，而要等待对方的ack过来之后才会移除。

如果网络状况不好，ack迟迟不过来，写缓冲很快就会满的。



**包头**

细心的同学可能注意到图中的消息req被拷贝到网卡的时候变成了大写的REQ，这是为什么呢？

因为这两个东西已经不是完全一样的了。

内核的网络模块会将缓冲区的消息进行分块传输，如果缓冲区的内容太大，是会被拆分成多个独立的小消息包的。

并且还要在每个消息包上附加上一些额外的头信息，比如源网卡地址和目标网卡地址、消息的序号等信息，到了接收端需要对这些消息包进行重新排序组装去头后才会扔进读缓冲中。



**速率**

还有个问题那就是如果读缓冲满了怎么办，网卡收到了对方的消息要怎么处理？一般的做法就是丢弃掉不给对方ack，对方如果发现ack迟迟没有来，就会重发消息。

那缓冲为什么会满？是因为消息接收方处理的慢而发送方生产的消息太快了，这时候tcp协议就会有个动态窗口调整算法来限制发送方的发送速率，使得收发效率趋于匹配。

如果是udp协议的话，消息一丢那就彻底丢了。





## Tornado 架构

纯socket的服务性能是非常弱的，这时候就需要高效的框架Tornado， 它 是专门为处理异步进程而构建的

先看一个例子：



```python
import tornado.httpserver
import tornado.ioloop


def handle_request(request):

    message = "Hello World from Tornado Http Server"
    request.write("HTTP/1.1 200 OK\r\nContent-Length: %d\r\n\r\n%s" %
                  (len(message), message))

    request.finish()


http_server = tornado.httpserver.HTTPServer(handle_request)
http_server.listen(8080)
tornado.ioloop.IOLoop.instance().start()

```



主要的组件：

> httpserver - 服务于 web 模块的一个非常简单的 HTTP 服务器的实现

> iostream - 对非阻塞式的 socket 的简单封装，以方便常用读写操作

> ioloop - 核心的 I/O 循环



### 架构分析



![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/tornado.png)



Tornado服务器有3大核心模块:

**1.IOLoop**

与我们上面那个简陋的http服务器不同, Tornado为了实现高并发和高性能, 使用了一个**IOLoop来处理socket的读写事件, IOLoop基于epoll, 可以高效的响应网络事件. 这是Tornado高效的保证.** 

暂时我们把ioloop理解为一个事件容器. 

用户把socket和回调函数注册到容器中, 容器内部会轮询socket, 一旦某个socket可以读写, 就调用回调函数来处理socket的读写事件.





2.**IOStream**

为了在处理请求的时候, 实现对socket的异步读写, Tornado实现了IOStream类, 用来处理socket的异步读写



3.**HTTPConnection**

这个类用来处理http的请求, 包括读取http请求头, 读取post过来的数据, 调用用户自定义的处理方法,以及把响应数据写给客户端socket



这三大组件的细节部分会在后面的文章中讲解
