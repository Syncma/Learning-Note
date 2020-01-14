# Flask响应

<!-- TOC -->

- [Flask响应](#flask%e5%93%8d%e5%ba%94)
  - [基本使用](#%e5%9f%ba%e6%9c%ac%e4%bd%bf%e7%94%a8)
  - [Response](#response)
    - [wsgi_app](#wsgiapp)
    - [full_dispatch_request](#fulldispatchrequest)
    - [finalize_request](#finalizerequest)
    - [make_response](#makeresponse)
    - [response_class](#responseclass)
    - [Response](#response-1)
    - [BaseResponse](#baseresponse)
    - [Headers](#headers)
    - [自定义 Response](#%e8%87%aa%e5%ae%9a%e4%b9%89-response)
    - [复习__call__](#%e5%a4%8d%e4%b9%a0call)
      - [问题1](#%e9%97%ae%e9%a2%981)
      - [问题2](#%e9%97%ae%e9%a2%982)
      - [__init__和__call__区别](#init%e5%92%8ccall%e5%8c%ba%e5%88%ab)

<!-- /TOC -->

## 基本使用

看一个例子：
```python
from flask import Flask

app = Flask(__name__)

@app.route('/')
def index():
    return 'hello world!', 200, {'Content-Type': 'application/json'}


if __name__ == '__main__':
    app.run()
    
```

在 index 方法中，返回了了 http 状态、body 以及 header 等，该方法返回的其实是一个 tuple

思考：

>这里究竟发送了什么？才让一个 tuple 变为一个 Response


## Response

开始源码分析吧， 

简单而言，app.run () 会启动一个满足 WSGI 协议的 web 服务，它会监听指定的端口，将 HTTP 请求解析为 WSGI 格式的数据，然后将 environ, start_response 传递给 Flask () 实例对象。

**`类对象作为方法被调用`，需要看到`__call__() `方法**

```python
class Flask(_PackageBoundObject):
...
    def __call__(self, environ, start_response):
        """The WSGI server calls the Flask application object as the
        WSGI application. This calls :meth:`wsgi_app` which can be
        wrapped to applying middleware."""
        return self.wsgi_app(environ, start_response)

    def __repr__(self):
        return "<%s %r>" % (self.__class__.__name__, self.name)

```

### wsgi_app

看到这里是调用`self.wsgi_app`方法 开启WSGI服务

进行深入看看这个方法：
```python
def wsgi_app(self, environ, start_response):
		ctx = self.request_context(environ)
        error = None
        try:
            try:
                ctx.push()
                response = self.full_dispatch_request()
            except Exception as e:
                error = e
                response = self.handle_exception(e)
            except:  # noqa: B001
                error = sys.exc_info()[1]
                raise
            return response(environ, start_response)
        finally:
            if self.should_ignore_error(error):
                error = None
            ctx.auto_pop(error)
```
该方法会找到当前请求路由对应的方法，调用该方法，获得返回 (即 response)，如果请求路由不存在，则进行错误处理，返回 500 错误。

### full_dispatch_request
这个重点看看full_dispatch_request () 方法：
```python
    def full_dispatch_request(self):
        """Dispatches the request and on top of that performs request
        pre and postprocessing as well as HTTP exception catching and
        error handling.

        .. versionadded:: 0.7
        """
        self.try_trigger_before_first_request_functions()
        try:
            request_started.send(self)
            rv = self.preprocess_request()
            if rv is None:
                rv = self.dispatch_request()
        except Exception as e:
            rv = self.handle_user_exception(e)
        return self.finalize_request(rv)
```

在 `full_dispatch_request ()` 方法中会调用 `self.finalize_request () `方法对返回数据进行处理，response 对象的构建就在该方法中实现


### finalize_request
继续定位`finalize_request`：
```python
def finalize_request(self, rv, from_error_handler=False):
        response = self.make_response(rv)
        try:
            response = self.process_response(response)
            request_finished.send(self, response=response)
        except Exception:
            if not from_error_handler:
                raise
            self.logger.exception(
                "Request finalizing failed with an error while handling an error"
            )
        return response

```

finalizerequest () 方法调用 makeresponse () 方法将视图函数返回的 tuple 转为了 response 对象

随后再通过 process_response () 方法进行了相关的 hooks 处理，具体而言就是执行 ctx._after_request_functions变量中存放的方法。


### make_response
这里重点看一下 make_response () 方法，其源码如下。

```
def make_response(self, rv):
...
```

make_response () 方法写的非常直观，将传入的内容按不同的情况进行处理，最终通过 response_class () 将其其转为 response 对象。

会发现有些直接通过 jsonify () 方法就将内容返回了，其实 jsonify () 方法内部也使用了 response_class () 将内容转为 response 对象。

### response_class

response_class 其实就是 Response 类， 可以看到代码里面有这么一段
```
response_class = Response
```

### Response
接着看看Response 类：
```python
class Response(ResponseBase, JSONMixin):
...
```

我们看到这里继承了ResponseBase和JSONMixin类，先看下ResponseBase

```
from werkzeug.wrappers import Response as ResponseBase
```

OK, 定位到Response方法中看看里面的逻辑：
```python
class Response(
    BaseResponse,
    ETagResponseMixin,
    ResponseStreamMixin,
    CommonResponseDescriptorsMixin,
    WWWAuthenticateMixin,
):
    """Full featured response object implementing the following mixins:

    - :class:`ETagResponseMixin` for etag and cache control handling
    - :class:`ResponseStreamMixin` to add support for the `stream` property
    - :class:`CommonResponseDescriptorsMixin` for various HTTP descriptors
    - :class:`WWWAuthenticateMixin` for HTTP authentication support
    """
```

跟前面的Request一样，通过 Mixin 机制，让具体逻辑都在 BaseResponse 中实现

### BaseResponse
简单看一下 BaseResponse 类的逻辑
```python
class BaseResponse(object):
....
```

在 BaseResponse () 中，定义了 Response 返回的默认属性，此外还提供了很多方法

这里看一下 Headers 类的实现，该类用于定义 Response 中 header 的细节

```python
	....
    if isinstance(headers, Headers):
            self.headers = headers
        elif not headers:
            self.headers = Headers()
        else:
            self.headers = Headers(headers)
```

### Headers
进一步定位Headers类：
```python
@native_itermethods(["keys", "values", "items"])
class Headers(object):
    def __init__(self, defaults=None):
        self._list = []
        if defaults is not None:
            if isinstance(defaults, (list, Headers)):
                self._list.extend(defaults)
            else:
                self.extend(defaults)
....
```

Headers 类通过 list 的形式构建出一个类似与 dict 的对象 (操作方面像 dict，key-value)，这要做的目的是为了保证 headers 中元素的顺序

通过 list 来保证顺序，此外 Headers 类还运行使用相同的 key 存储不同的 values，同样通过 list 来实现，这里看一下它的 get () 方法。

```python
def get(self, key, default=None, type=None, as_bytes=False):
		try:
	        rv = self.__getitem__(key, _get_mode=True)
        except KeyError:
            return default
        if as_bytes:
            rv = rv.encode("latin1")
        if type is None:
            return rv
        try:
            return type(rv)
        except ValueError:
            return default

```

使用者可以通过如下方式使用
```
>>> d = Headers([('Content-Length', '42')])
>>> d.get('Content-Length', type=int)
42
```

get () 会调用` __getitem__()`方法去获取具体的值

而 `__getitem__()`方法的主要逻辑就是遍历 _list

Headers 所有的属性都以元组的形式存放在 _list中


### 自定义 Response
如果想自定义 Response，直接继承 Flask 中的 Response 则可。

```
from flask import Flask, Response
class MyResponse(Response):
     pass
app = Flask(__name__)
app.response_class = MyResp
```

### 复习`__call__`
这里简单复习下`__call__`方法

对象后面加括号，触发执行。

>构造方法的执行是由创建对象触发的，即：对象 = 类名() 
而对于` __call__ `方法的执行是由对象后加括号触发的，即：对象() 或者 类()()

```python
class Foo(object):
    def __init__(self):
        print('init')
        print("init self=", self)

    def __call__(self):
        print("call self=", self)
        print('call')


obj = Foo()
print("init==",obj)  # 执行__init__方法

print("call==",obj())  # 执行__call__方法

```

运行结果：
```
init
init self= <__main__.Foo object at 0x7fdda7583128>
init== <__main__.Foo object at 0x7fdda7583128>
call self= <__main__.Foo object at 0x7fdda7583128>
call
call== None
```

#### 问题1
`__init__`方法有没有返回值？
```python
class Foo(object):
    def __init__(self):
        print('init')
        print("init self=", self)
        return "Hello" # 这里给个返回值

obj = Foo()
print("init==",obj)
```

运行结果：
```
TypeError: __init__() should return None, not 'str'
```

所以知道`__init__`只能返回None


#### 问题2

关于` __call__` 方法，不得不先提到一个概念，就是可调用对象（callable）

我们平时自定义的函数、内置函数和类都属于可调用对象，**但凡是可以把一对括号()应用到某个对象身上都可称之为可调用对象**

判断对象是否为可调用对象可以用函数 callable

如果在类中实现了 `__call__ `方法，那么实例对象也将成为一个可调用对象


例子：
```python
class Foo(object):
    def __init__(self):
        print('init')
        print("init self=", self)

    def __call__(self):
        print("call")
        print("call self=", self)

obj = Foo()
print("init==",obj)

print(callable(obj)) # True

print(obj())

```

返回结果：
```
init
init self= <__main__.Foo object at 0x7f7e84d350f0>
init== <__main__.Foo object at 0x7f7e84d350f0>
True
call
call self= <__main__.Foo object at 0x7f7e84d350f0>
None
```

我们发现obj是实例对象，同时还是可调用对象，那么就可以像函数一样调用它

实例对象也可以像函数一样作为可调用对象来用

那么，这个特点在什么场景用得上呢？这个要结合类的特性来说，类可以记录数据（属性），而函数不行（闭包某种意义上也可行）


利用这种特性可以实现基于类的装饰器，在类里面记录状态

```python
class Counter(object):
    def __init__(self, func):
        self.func = func
        self.count = 0

    def __call__(self, *args, **kwargs):
        self.count += 1
        return self.func(*args, **kwargs)

@Counter
def foo():
    pass

for i in range(10):
    foo()

print(foo.count)  # 10

```



#### `__init__`和`__call__`区别

例子：
```python
class Foo(object):
    #def __init__(self):
    #    print('init')
    #    print("init self=", self)

    def __call__(self):
        print("call")
        print("call self=", self)

obj = Foo()
print("init==",obj)

print(callable(obj)) # True

print(obj())
```

运行结果：
```
init== <__main__.Foo object at 0x7ff507a9b0b8>
True
call
call self= <__main__.Foo object at 0x7ff507a9b0b8>
None
```

