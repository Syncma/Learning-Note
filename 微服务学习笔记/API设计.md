# API 设计

<!-- TOC -->

- [API 设计](#api-%e8%ae%be%e8%ae%a1)
  - [API架构风格](#api%e6%9e%b6%e6%9e%84%e9%a3%8e%e6%a0%bc)
  - [媒体类型](#%e5%aa%92%e4%bd%93%e7%b1%bb%e5%9e%8b)
  - [常见的开发组合](#%e5%b8%b8%e8%a7%81%e7%9a%84%e5%bc%80%e5%8f%91%e7%bb%84%e5%90%88)
  - [REST VS RPC](#rest-vs-rpc)
    - [REST 优点](#rest-%e4%bc%98%e7%82%b9)
    - [RPC优点](#rpc%e4%bc%98%e7%82%b9)
    - [如何选择](#%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9)
  - [SOAP](#soap)
    - [python例子](#python%e4%be%8b%e5%ad%90)
  - [RPC](#rpc)
    - [RPC类型](#rpc%e7%b1%bb%e5%9e%8b)
      - [XML-RPC](#xml-rpc)
      - [JSON-RPC](#json-rpc)
      - [Messagepack-RPC](#messagepack-rpc)
  - [RPC 调用详细流程图](#rpc-%e8%b0%83%e7%94%a8%e8%af%a6%e7%bb%86%e6%b5%81%e7%a8%8b%e5%9b%be)
  - [总结](#%e6%80%bb%e7%bb%93)

<!-- /TOC -->

## API架构风格

**常用的API风格主要有`SOAP、RPC、REST`**

## 媒体类型

**常用的媒体类型`JSON、XML、Protobuf`**


## 常见的开发组合

SOAP + XML

RPC + Protobuf

REST + JSON


## REST VS RPC

### REST 优点

>1. 轻量级，简单易用，维护性和扩展性都比较好

>2. REST 相对更规范，更标准，更通用，无论哪种语言都支持 HTTP 协议，可以对接外部很多系统，只要满足 HTTP 调用即可，更适合对外，RPC 会有语言限制，不同语言的 RPC 调用起来很麻烦
>
>3. JSON 格式可读性更强，开发调试都很方便
>4. 在开发过程中，如果严格按照 REST 规范来写 API，API 看起来更清晰，更容易被大家理解


### RPC优点

>1. RPC+Protobuf 采用的是 TCP 做传输协议，REST 直接使用 HTTP 做应用层协议，这种区别导致 REST 在调用性能上会比 RPC+Protobuf 低
>2. RPC 不像 REST 那样，每一个操作都要抽象成对资源的增删改查，在实际开发中，有很多操作很难抽象成资源，比如登录操作。所以在实际开发中并不能严格按照 REST 规范来写 API，RPC 就不存在这个问题
> 3. RPC 屏蔽网络细节、易用，和本地调用类似
>这里的易用指的是调用方式上的易用性。在做 RPC 开发时，开发过程很烦琐，需要先写一个 DSL 描述文件，然后用代码生成器生成各种语言代码，当描述文件有更改时，必须重新定义和编译，维护性差。


### 如何选择

**`内部系统调用用RPC、对外使用REST`**




## SOAP

SOAP（原为Simple Object Access Protocol的首字母缩写，即简单对象访问协议）是交换数据的一种协议规范，使用在计算机网络Web服务（web service）中，交换带结构信息。

SOAP为了简化网页服务器（Web Server）从XML数据库中提取数据时，节省去格式化页面时间，以及不同应用程序之间按照HTTP通信协议，遵从XML格式执行资料互换，使其抽象于语言实现、平台和硬件。


### python例子

针对Python的WebService开发，开发者讨论最多的库是soaplib，但从其官网可知，其最新版本“soaplib-2.0.0-beta2”从2011年3月发布后就不再进行更新了。

通过阅读soaplib的官方文档，可知其不再维护后已经转向了一个新的项目：rpclib（官方地址：http://github.com/arskom/rpclib）

进行后续开发，但在rpclib的readme中，介绍了rpclib已经更名为spyne，并将持续进行更新，那就选用spyne进行开发了。


python3 +

1.安装：
```
pip install spyne==2.13.4a1

```

[官网地址](http://spyne.io/)

2.服务端代码：
```python
from spyne import Application, rpc, ServiceBase, Iterable, Integer, Unicode
from spyne.protocol.soap import Soap11
from spyne.server.wsgi import WsgiApplication
from spyne.protocol.json import JsonDocument

import logging
from wsgiref.simple_server import make_server


# 定义服务
class HelloWorldService(ServiceBase):
    @rpc(Unicode, Integer, _returns=Iterable(Unicode))
    def say_hello(self, name, times):
        for i in range(times):
            yield u'Hello, %s' % name


# 服务注册
soap_app = Application([HelloWorldService],
                       tns='example.soap',
                       in_protocol=Soap11(validator='lxml'),
                       out_protocol=JsonDocument())

# WSGI
wsgi_app = WsgiApplication(soap_app)

if __name__ == "__main__":

    logging.basicConfig(level=logging.DEBUG)
    logging.getLogger("spyne.protocol.xml").setLevel(logging.DEBUG)

    logging.info("Listening to http://localhost:8000")
    logging.info("wsdl is at: http://localhost:8000/?wsdl")

    server = make_server("localhost", 8000, wsgi_app)
    server.serve_forever()

```

运行服务端代码，然后在浏览器重访问http://127.0.0.1:8000/SOAP/?wsdl ，如果正常的话，则能看到该服务的描述信息，包括各个方法的输入参数、返回值，以及实体类的信息


3.客户端代码

1>安装客户端模块：

```
pip install zep

```

2> 代码：
```python
from zeep import Client

url = 'http://localhost:8000/?wsdl'
client = Client(url)

try:
    result = client.service.say_hello('tony', 10)
    print("result=", result)

except Exception as e:
    print("ERROR=", e)

```



##  RPC

RPC（Remote Procedure Call Protocol）——远程过程调用协议，它是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络技术的协议。

> “远程调用”意思就是：
> 
> 被调用方法的具体实现不在程序运行本地，而是在别的某个地方（分布到各个服务器），但是用起来像是在本地

RPC协议假定某些传输协议的存在，如TCP或UDP，为通信程序之间携带信息数据。

在OSI网络通信模型中，RPC跨越了传输层和应用层。

RPC使得开发包括网络分布式多程序在内的应用程序更加容易。

总结:
服务提供的两大流派.传统意义以方法调用为导向通称RPC。

为了企业SOA,若干厂商联合推出webservice,制定了wsdl接口定义,传输soap.当互联网时代,臃肿SOA被简化为http+xml/json.但是简化出现各种混乱。

以资源为导向,任何操作无非是对资源的增删改查，于是统一的REST出现了.

进化的顺序: **RPC -> SOAP -> RESTful**


### RPC类型

#### XML-RPC 
XML-RPC:XML Remote Procedure Call，即XML远程方法调用,利用http+xml封装进行RPC调用。

基于http协议传输、XML作为信息编码格式。

一个xml-rpc消息就是一个请求体为xml的http-post请求，服务端执行后也以xml格式编码返回。

这个标准面前已经演变为下面的SOAP协议。可以理解SOAP是XML-RPC的高级版本。

测试环境:  **python3.6+**

服务端代码：
```python
from xmlrpc.server import SimpleXMLRPCServer


def respon_string(str):
    return "get string :%s" % str


if __name__ == '__main__':
    s = SimpleXMLRPCServer(('0.0.0.0', 8080))
    s.register_function(respon_string, "get_string")
    s.serve_forever()

```

客户端代码：
```python
import xmlrpc.client

proxy = xmlrpc.client.ServerProxy('http://localhost:8080')
print(proxy.get_string('hello'))
```


####  JSON-RPC

JSON-RPC:JSON Remote Procedure Call，即JSON远程方法调用 。

类似于XML-RPC，不同之处是使用JSON作为信息交换格式

[JsonrpcServer官方地址](https://jsonrpcserver.readthedocs.io/en/latest/index.html)

模块安装：
```python
pip install jsonrpcserver
pip install jsonrpcclient
```


服务端代码：
```python
from jsonrpcserver import method, serve


@method
def ping():
    return "pong"


if __name__ == "__main__":
    serve()  # 默认端口5000
```


客户端代码：
```python
from jsonrpcclient import request

response = request("http://localhost:5000", "ping")

# print(response.text)
# {"jsonrpc": "2.0", "result": "pong", "id": 1}

print(response.data.result)
#  "pong"
```



####  Messagepack-RPC
 Messagepack 是一个基于二进制的高效对象序列化库。 它支持在多种语言(如 JSON)之间交换结构化对象。 但与 JSON 不同的是，它非常快速和小巧。

模块安装：
```
pip install msgpack
```

服务端代码：
```python
import msgpackrpc


class SumServer(object):
    def sum(self, x, y):
        return x + y

server = msgpackrpc.Server(SumServer())
server.listen(msgpackrpc.Address("localhost", 18800))
server.start()
```


客户端代码：
```python
import msgpackrpc

client = msgpackrpc.Client(msgpackrpc.Address("localhost", 18800))
result = client.call('sum', 1, 2)
print(result)
```


##  RPC 调用详细流程图

python自带RPC 的实现过程


![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/RPC.jpg)


##  总结
RPC 主要用于公司内部的服务调用，性能消耗低，传输效率高，实现复杂。
HTTP 主要用于对外的异构环境，浏览器接口调用，App 接口调用，第三方接口调用等。

RPC 使用场景（大型的网站，内部子系统较多、接口非常多的情况下适合使用 RPC）：
* 长链接。不必每次通信都要像 HTTP 一样去 3 次握手，减少了网络开销。
*  注册发布机制。RPC 框架一般都有注册中心，有丰富的监控管理；发布、下线接口、动态扩展等，对调用方来说是无感知、统一化的操作。
*  安全性，没有暴露资源操作。
*  微服务支持。就是最近流行的服务化架构、服务化治理，RPC 框架是一个强力的支撑。

