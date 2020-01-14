# WSGI学习笔记

<!-- TOC -->

- [WSGI学习笔记](#wsgi%e5%ad%a6%e4%b9%a0%e7%ac%94%e8%ae%b0)
  - [WSGI 是什么](#wsgi-%e6%98%af%e4%bb%80%e4%b9%88)
  - [WSGI Server是什么](#wsgi-server%e6%98%af%e4%bb%80%e4%b9%88)
  - [Web服务器](#web%e6%9c%8d%e5%8a%a1%e5%99%a8)
  - [WSGI工作机制](#wsgi%e5%b7%a5%e4%bd%9c%e6%9c%ba%e5%88%b6)
    - [服务器端](#%e6%9c%8d%e5%8a%a1%e5%99%a8%e7%ab%af)
    - [应用程序](#%e5%ba%94%e7%94%a8%e7%a8%8b%e5%ba%8f)
  - [WSGI优缺点](#wsgi%e4%bc%98%e7%bc%ba%e7%82%b9)
    - [优点](#%e4%bc%98%e7%82%b9)
    - [缺点](#%e7%bc%ba%e7%82%b9)
  - [HTTP基础知识](#http%e5%9f%ba%e7%a1%80%e7%9f%a5%e8%af%86)
  - [引出WSGI](#%e5%bc%95%e5%87%bawsgi)
  - [WSGI 组件](#wsgi-%e7%bb%84%e4%bb%b6)
    - [Web server side](#web-server-side)
    - [app side](#app-side)
    - [Middleware](#middleware)
  - [WSGI简单实现](#wsgi%e7%ae%80%e5%8d%95%e5%ae%9e%e7%8e%b0)
  - [总结](#%e6%80%bb%e7%bb%93)

<!-- /TOC -->

## WSGI 是什么
WSGI - Python Web Server Gateway Interface
是一种规范,用来规范**`Python web应用`**与服务器之间通信的标准

## WSGI Server是什么
wsgi server是自己做web服务器借用wsgi协议来调用application。

常用的WSGI服务器:
```
gunicorn
uwsgi
cheerypy
tornado
gevent
mod_wsgi
flask
....
```

## Web服务器

我们需要明确一点，nginx是无法直接跟flask application做通信，需要借用wsgi server。

flask本身也有个web服务器是werkzeug，so 才能启动服务并监听端口。

nginx、apache在这里只是启动了proxy的作用

那为什么不直接把uwsgi和gunicorn给暴露出去，**因为nginx的静态文件处理能力极强**


## WSGI工作机制

wsgi主要是两层，服务器端和 应用程序 

### 服务器端
>从底层解析http解析，然后调用应用程序，给应用程序提供(环境信息)和(回调函数)，这个回调函数是用来将应用程序设置的http header和status等信息传递给服务器方.

### 应用程序

> 用来生成返回的header,body和status,以便返回给服务器方。


## WSGI优缺点

### 优点
**多样的部署选择和组件之间的`高度解耦`**

由于上面提到的高度解耦特性，理论上，任何一个符合WSGI规范的App都可以部署在任何一个实现了WSGI规范的Server上，这给Python Web应用的部署带来了极大的灵活性。

Flask自带了一个基于Werkzeug的调试用服务器。

根据Flask的文档，在生产环境不应该使用内建的调试服务器，

而应该采取以下方式之一进行部署：
* **`GUNICORN  -目前这个是主流`**
* UWSGI

### 缺点
- 待补充



## HTTP基础知识

对于 web 应用程序来说，最基本的概念就是客户端发送请求（request），收到服务器端的响应（response）。


下面是简单的 HTTP 请求：
```
GET /Index.html HTTP/1.1\r\n
Connection: Keep-Alive\r\n
Accept: */*\r\n
User-Agent: Sample Application\r\n
Host: www.microsoft.com\r\n\r\n
```

内容包括了 method、 url、 protocol version 以及头部的信息。

而 HTTP 响应（不包括数据）可能是如下的内容：
```
HTTP/1.1 200 OK
Server: Microsoft-IIS/5.0\r\n
Content-Location: http://www.microsoft.com/default.htm\r\n
Date: Tue,25Jun200219:33:18GMT\r\n
Content-Type: text/html\r\n
Accept-Ranges: bytes\r\n
Last-Modified: Mon,24Jun200220:27:23GMT\r\n
Content-Length: 26812\r\n
```

实际生产中，python 程序是放在服务器的 http server（比如 apache， nginx 等）上的。

现在的问题是：
> 服务器程序怎么把接受到的请求传递给 python 呢?
> 怎么在网络的数据流和 python 的结构体之间转换呢？

这就是 wsgi 做的事情：
> 一套关于程序端和服务器端的规范，或者说统一的接口。



## 引出WSGI

先看一下面向 http 的 python 程序需要关心哪些内容

-  **请求**
    - 请求的方法 method
   - 请求的地址 url
	- 请求的内容
	- 请求的头部 header
	- 请求的环境信息

-  **响应**
	- 状态码 status_code
	- 响应的数据
	- 响应的头部

WSGI 的任务就是把上面的数据在 http server 和 python 程序之间简单友好地传递。

它是一个标准，需要 http server 和 python 程序都要遵守一定的规范，实现这个标准的约定内容，才能正常工作。


## WSGI 组件

这套规范将web组件分为三部分：server, framework/app, middleware对象

下面分别来介绍这三部分内容


### Web server side

Server 必须提供两样东西：**`environ, 和一个 start_response 函数`**

```
environ 是一个python dictionary，包含一些常见的东西(和CGI environment类似)。

start_response 是一个callable对象，有两个参数:

1.status —— string类型,包含一个标准的HTTP status 像"200 OK",
2.response_headers —— list类型,包含标准HTTP response headers。
```

案例：
```python
import os, sys


# application 是程序端的可调用对象
def run_with_cgi(application):
    # 准备 environ 参数，这是一个字典，里面的内容是一次 HTTP 请求的环境变量
    environ = dict(os.environ.items())
    environ['wsgi.input'] = sys.stdin
    environ['wsgi.errors'] = sys.stderr
    environ['wsgi.version'] = (1, 0)
    environ['wsgi.multithread'] = False
    environ['wsgi.multiprocess'] = True
    environ['wsgi.run_once'] = True
    environ['wsgi.url_scheme'] = 'http'

    headers_set = []
    headers_sent = []


# 把应答的结果输出到终端
def write(data):
    sys.stdout.write(data)
    sys.stdout.flush()


# 实现 start_response 函数，根据程序端传过来的 status 和 response_headers 参数，
# 设置状态和头部
def start_response(status, response_headers, exc_info=None):
    headers_set[:] = [status, response_headers]
    return write


# 调用客户端的可调用对象，把准备好的参数传递过去
result = application(environ, start_response)

# 处理得到的结果，这里简单地把结果输出到标准输出。
try:
    for data in result:
        if data:  # don't send headers until body appears
            write(data)
finally:
    if hasattr(result, 'close'):
        result.close()
```



### app side

WSGI 规定每个 python 程序（Application）必须是一个可调用的对象（实现了__call__ 函数的方法或者类），

**接受两个参数 environ（WSGI 的环境信息） 和 start_response（开始响应请求的函数），并且返回 iterable对象**。

几点说明：
```
environ 和 start_response 由 http server 提供并实现
environ 变量是包含了环境信息的字典
Application 内部在返回前调用 start_response
start_response也是一个 callable，接受两个必须的参数，status（HTTP状态）和 response_headers（响应消息的头）
可调用对象要返回一个值，这个值是可迭代的。
```

看一个例子：
```python
# 1. 可调用对象是一个函数
def application(environ,start_response):
    response_body = 'The request method was %s' % environ['REQUEST_METHOD']
    # HTTP response code and message
    status = '200 OK'
    # 应答的头部是一个列表，每对键值都必须是一个 tuple。
    response_headers = [('Content-Type','text/plain'),('Content-Length',str(len(response_body)))]
    # 调用服务器程序提供的 start_response，填入两个参数
    start_response(status,response_headers)
    # 返回必须是 iterable
    return[response_body]

# 2. 可调用对象是一个类
class AppClass(object):
    """这里的可调用对象就是 AppClass 这个类，调用它就能生成可以迭代的结果。
    使用方法类似于：
    for result in AppClass(env, start_response):
         do_somthing(result)
    """
    def __init__(self,environ,start_response):
        self.environ = environ
        self.start = start_response
 
    def __iter__(self):
        status = '200 OK'
        response_headers = [('Content-type','text/plain')]
        self.start(status,response_headers)
        yield"Hello world!\n"
 
# 3. 可调用对象是一个实例
class AppClass(object):
    """这里的可调用对象就是 AppClass 的实例，使用方法类似于：
    app = AppClass()
    for result in app(environ, start_response):
         do_somthing(result)
    """
 
    def __init__(self):
        pass
 
    def __call__(self,environ,start_response):
        status = '200 OK'
        response_headers = [('Content-type','text/plain')]
        self.start(status,response_headers)
        yield"Hello world!\n"
```

### Middleware
中间件，顾名思义在Web server和Web framework/app之间再插入一个环节。

看情况对数据做一些处理。

有些程序可能处于服务器端和程序端两者之间：
>对于服务器程序，它就是应用程序；
>而对于应用程序，它就是服务器程序。
>这就是中间层 middleware。

middleware 对服务器程序和应用是透明的，它像一个代理/管道一样，把接收到的请求进行一些处理，然后往后传递，一直传递到客户端程序，最后把程序的客户端处理的结果再返回。

middleware 做了两件事情：
Gunicorn对静态文件的支持不太好，所以生产环境下常用Nginx作为反向代理服务器。

> 被服务器程序（有可能是其他 middleware）调用，返回结果回去
> 调用应用程序（有可能是其他 middleware），把参数传递过去

middleware 的可能使用场景：

```
根据 url 把请求给到不同的客户端程序（url routing）
允许多个客户端程序/web 框架同时运行，就是把接到的同一个请求传递给多个程序。
负载均衡和远程处理：把请求在网络上传输
应答的过滤处理
```

那么简单地 middleware 实现是怎么样的呢？

下面的代码实现的是一个简单地 url routing 的 middleware：
```python
class Router(object):
    def __init__(self):
        self.path_info = {}
    def route(self,environ,start_response):
        application = self.path_info[environ['PATH_INFO']]
        return application(environ,start_response)
    def __call__(self,path):
        def wrapper(application):
            self.path_info[path] = application
        return wrapper
 
router = Router()
```
怎么在程序里面使用呢？

```python
#here is the application
@router('/hello')#调用 route 实例，把函数注册到 paht_info 字典
def hello(environ,start_response):
    status = '200 OK'
    output = 'Hello'
    response_headers = [('Content-type','text/plain'),
                        ('Content-Length',str(len(output)))]
    write = start_response(status,response_headers)
    return[output]
 
@router('/world')
def world(environ,start_response):
    status = '200 OK'
    output = 'World!'
    response_headers = [('Content-type','text/plain'),
                        ('Content-Length',str(len(output)))]
    write = start_response(status,response_headers)
    return[output]
 
#here run the application
result = router.route(environ,start_response)
for value in result:
    write(value)
```

## WSGI简单实现

wsgiref - WSGI utilities and Reference Implementation

Python自带的实现了WSGI协议的的wsgi server。
```
wsgi server可以理解为一个符合wsgi规范的web server
接收request请求，封装一系列环境变量，
按照wsgi规范调用注册的wsgi app，
最后将response返回给客户端。
```

简单例子：

```python
def application(environ, start_response):
    response_body = "Hello World"
    header = [('Content-Type', 'text/html')]
    status = "200 OK"
    start_response(status, header)
    print("environ http request method:" + environ["REQUEST_METHOD"])
    return [response_body.encode("utf-8")]


if __name__ == "__main__":
    from wsgiref.simple_server import make_server
    httpd = make_server("0.0.0.0", 8080, application)
    print("http run on" + str(httpd.server_port))
    httpd.serve_forever()
    
```



## 总结

WSGI把来自socket的数据包解析为http格式，然后进而变化为environ变量

这environ变量里面有wsgi本身的信息(比如 host， post，进程模式等)

还有client的header及body信息。

start_respnse是一个函调函数，必须要附带两个参数，一个是status(http状态)，response_headers(响应的header头信息)。

像flask、django、tornado都会暴露WSGI协议入口，我们只需要自己实现WSGI协议，wsgi server然后给flask传递environ，及start_response, 等到application返回值之后，我再socket send返回客户端。
