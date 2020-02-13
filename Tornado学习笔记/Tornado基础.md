# Tornado基础

<!-- TOC -->

- [Tornado基础](#tornado%e5%9f%ba%e7%a1%80)
  - [是什么](#%e6%98%af%e4%bb%80%e4%b9%88)
  - [安装](#%e5%ae%89%e8%a3%85)
  - [测试案例](#%e6%b5%8b%e8%af%95%e6%a1%88%e4%be%8b)
    - [带日志的例子](#%e5%b8%a6%e6%97%a5%e5%bf%97%e7%9a%84%e4%be%8b%e5%ad%90)
  - [优缺点](#%e4%bc%98%e7%bc%ba%e7%82%b9)
  - [实现原理](#%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86)

<!-- /TOC -->


## 是什么

Facebook发布了开源网络服务器框架Tornado

**Tornado由Python编写**，是一款轻量级的Web服务器，同时又是一个开发框架。

采用非阻塞I/O模型(epoll)，主要是为了应对高并发 访问量而被开发出来



## 安装

使用pip 命令进行安装：

```
pip install torando
```



## 测试案例

```python
import tornado.ioloop
import tornado.web

class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.write("Hello, world")

application = tornado.web.Application([
    (r"/", MainHandler),
])

if __name__ == "__main__":
    application.listen(8888)
    tornado.ioloop.IOLoop.instance().start()
```



### 带日志的例子



```python
import tornado.ioloop
import tornado.web
import logging

def cfgLogging():
    # create logger
    fmt = "%(asctime)s - %(name)s - %(levelname)s - %(message)s"
    logging.basicConfig(format=fmt, level=logging.DEBUG)

class MainHandler(tornado.web.RequestHandler):

    def get(self):
        self.write("Hello, world")

application = tornado.web.Application([
    (r"/", MainHandler),
])

if __name__ == "__main__":
    cfgLogging()
    logging.debug('* TCP Service started on port: 8888... ')
    application.listen(8888)
    tornado.ioloop.IOLoop.instance().start()

```



## 优缺点

Tornado 相对于其他框架的优缺点是什么?

**异步非阻塞，单进程并发, 高性能**

**没有ORM, 没有Session支持，没有Django 自动化后台**



## 实现原理

核心技术： **IOLoop, IOStream, HTTPConnection**

> 暂时我们把ioloop理解为一个事件容器. 
>
> 用户把socket和回调函数注册到容器中, 容器内部会轮询socket, 一旦某个socket
>
> 可以读写, 就调用回调函数来处理socket的读写事件.
>
> IOStream对socket的读写做了一层封装, 通过使用两个缓冲区, 实现对socket的异步读写.
>
> HttpConnection类专门用来处理http请求



后面会详细说明