# Flask介绍

[toc]


## 什么是Flask

**它是python微型web框架**


## 应用场景

快速开发web
适用于做小网站以及web服务的API，开发大型网站无压力

## 特征

Flask只是一个内核，默认依赖于两个外部库： **`Jinja2 模板引擎和 Werkzeug WSGI 工具集`**

特点：
```
内置开发服务器和快速调试器
集成支持单元测试
RESTful可请求调度
Jinja2模板
支持安全cookie（客户端会话）
符合WSGI 1.0
基于Unicode
```


## 两大组件

### Jinja2
Jinja2 是一个 Python 的功能齐全的模板引擎。它有完整的 unicode 支持，一个可选 的集成沙箱执行环境，被广泛使用，以 BSD 许可证授权

[Jinja官网文档](http://docs.jinkan.org/docs/jinja2/)

### Werkzeug

Werkzeug是一个WSGI工具包，他可以作为一个Web框架的底层库

什么是WSGI?  

> 简单的说，WSGI 只是一种接口,它只适用于 Python 语言，其全称为 Web Server Gateway Interface，定义了 web服务器和 web应用之间的接口规范

[官网地址](https://www.palletsprojects.com/p/werkzeug/)
[文档手册](https://werkzeug.palletsprojects.com/en/0.15.x/)


####  Werkzeug优点

Werkzeug是一个符合WSGI规范的基础库，其中包含有很多Web开发的 常用功能，

比如：请求响应模型的类抽象，url路由，cookie实现，web调试界面等基础功能。

**`Flask是基于 Werkzeug的高层封装，添加了更完备的web开发功能`**


比bottle使用的socketserver  优势在哪？--- **socketServer 只提供基本的WSGI**




####  Werkzeug使用

```python
from werkzeug.wrappers import Request, Response

@Request.application
def application(request):
return Response("Hello, World!")


if __name__ == "__main__":
    from werkzeug.serving import run_simple
    run_simple("localhost", 5000, application)
```


## Flask安装

1.使用下面命令安装flask模块
```
pip install flask
```

2.简单例子

```
from flask import Flask
app = Flask(__name__)


@app.route('/')
def index():
    return 'hello world!'


if __name__ == '__main__':
    app.run()
```


运行结果：

```
* Serving Flask app "app" (lazy loading)
 * Environment: production
   WARNING: This is a development server. Do not use it in a production deployment.
   Use a production WSGI server instead.
 * Debug mode: off
 * Running on http://127.0.0.1:5000/ (Press CTRL+C to quit)

```

注意： **`服务默认端口5000`**




## 组件内部逻辑
会在后面源码分析会介绍

## 生产环境部署
```
gunicron + flask
```


## 与其他框架比较优缺点

Python Web框架分类
功能分类：
a:收发消息相关(socket)
b:根据不同的URL执行不同的函数（业务逻辑相关的）
c:实现动态网页（字符串的替换）

Web框架分类：
1、自己实现b，c，使用第三方的a（Django）
2、自己实现b，使用第三方的a，c（Flask）
3、自己实现a\b\c（Tornado）




1.大包大揽的Django

优点：完美的文档、全套的解决方案(Cache、Session、ORM...)、强大的URL路由配置、自助管理后台
缺点：系统紧耦合、自带的ORM不够强大、Template结构弱


2.力求精简的Web.py和Tornado

3.新生代的微框架Flask、Bottle



区别：
Bottle WSGI 是使用系统自带的socketserver.TCPServer
Flask 是使用它自带的werkzurg

Flask 强于 Bottle





## Flask很重要的几个概念

###路由

所谓路由，就是处理URL和函数之间关系的程序，Flask中也是对URL规则进行统一管理的，使用@app.route修饰器将一个函数注册为路由。

###  蓝图

 编程讲究的是功能模块化，从而使代码看起来更加的优雅和顺畅， 在Flask中，蓝图可以将各个应用组织成不同的组件，实现代码的模块化。

比如一个系统有两种角色，一个是普通用户user，另一个是管理员admin，那么他们所拥有的权限和功能有很大差异，若将其放在同一个文件下，代码量相对较大且不易维护，若进行版本控制时，也很容易出现冲突，这时可以创建蓝图加以区分。

### ORM框架 

