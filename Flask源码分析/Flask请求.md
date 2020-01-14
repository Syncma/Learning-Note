# Flask请求

<!-- TOC -->

- [Flask请求](#flask%e8%af%b7%e6%b1%82)
  - [关于请求](#%e5%85%b3%e4%ba%8e%e8%af%b7%e6%b1%82)
    - [Request类](#request%e7%b1%bb)
    - [问题1](#%e9%97%ae%e9%a2%981)
    - [问题2](#%e9%97%ae%e9%a2%982)
    - [问题3](#%e9%97%ae%e9%a2%983)
    - [BaseRequest](#baserequest)
    - [cached_property](#cachedproperty)

<!-- /TOC -->
## 关于请求

要先了解Flask 与 WSGI 协议的关系

请求对 Flask 而言，主要就是 WSGI 的 environ 对象。
Flask 的请求由 Request 类实现。


还是根据例子看源码
```python
from flask import Flask

app = Flask(__name__)

@app.route('/')
def index():
    return 'hello world!'
    
if __name__ == '__main__':
    app.run()
```


直接定义Flask初始化代码的逻辑：flask/app.py
```python
class Flask(_PackageBoundObject):
	 #: The class that is used for request objects.  See :class:`~flask.Request`
    #: for more information.
    request_class = Request
...
```

### Request类
看到请求其实是调用了Request类，在头部看到这一行：
```python
from .wrappers import Request
```

那么我们直接去wrappers.py找下Request
```python
class Request(RequestBase, JSONMixin):
...

```
首先就是类继承，看到它继承了两个类 RequestBase和JSONMixin，
**`这两个类做了什么作用呢？为什么要继承这两个呢？`**

我们接着分析， 看到wrappers.py开头导入的模块：
```python
from werkzeug.wrappers import Request as RequestBase
from werkzeug.wrappers import Response as ResponseBase
from werkzeug.wrappers.json import JSONMixin as _JSONMixin
```

看到这个，我们了解到RequestBase其实是werkzeug.wrappers 里面的Request，
JSONMixin 是werkzeug.wrappers.json模块里面的JSONMixin类。

先看Request，直接定位Request方法：
```python
class Request(
    BaseRequest,
    AcceptMixin,
    ETagRequestMixin,
    UserAgentMixin,
    AuthorizationMixin,
    CommonRequestDescriptorsMixin,
):
    """Full featured request object implementing the following mixins:

    - :class:`AcceptMixin` for accept header parsing
    - :class:`ETagRequestMixin` for etag and cache control handling
    - :class:`UserAgentMixin` for user agent introspection
    - :class:`AuthorizationMixin` for http auth handling
    - :class:`CommonRequestDescriptorsMixin` for common headers
    """
```

一看这个有的人会蒙圈了，对于写了这么多年的python的我来说，第一次看也是蒙的，
继承了这么多类，然后什么都不写，它想要说明什么意思呢？

### 问题1
这里面看到很多Mixin类，Mixin机制的作用是什么？

看一个例子：
```python
class RunnableMixIn(object):
    def run(self):
        print('Running...')


class CarnivorousMixIn(object):
    def fly(self):
        print('Flying...')


class Animal(object):
    pass


class Mammal(Animal):
    pass


class Dog(Mammal, RunnableMixIn, CarnivorousMixIn):
    pass


Dog().fly()
```

运行结果：
```
Flying...
```

从上面的例子我们可以看到：

```
Dog继承了很多MixIn类，并且它本身没有fly方法，
但是继承的CarnivorousMixIn有fly方法，那么Dog也就有了fly方法
```


### 问题2
python类 里面只有注释也能执行？

看一个例子：
```python
class Test(object):
    """this is test class"""


Test()
```

运行了下，发现可以正常运行（python3.6+）

### 问题3
为什么要继承那么多Mixin类？

简单的说：
> Request 类之所以这样设计的目的是将**`功能分层`**
> 不同的功能交由不同的 Mixin 类去完成，从而构成清晰易读的结构。


简单解释一下 Mixin 机制:
> 某个类希望使用多个不同类中的属性或方法，就可以通过多继承的方式来弄


### BaseRequest
首先看继承的第一个类：
```python
class Request(
    BaseRequest,
    AcceptMixin,
    ETagRequestMixin,
    UserAgentMixin,
    AuthorizationMixin,
    CommonRequestDescriptorsMixin,
):
```
定位到BaseRequest类：
```python
class BaseRequest(object):
...
def __init__(self, environ, populate_request=True, shallow=False):
        self.environ = environ
        if populate_request and not shallow:
            self.environ["werkzeug.request"] = self
        self.shallow = shallow
```

`__init__`方法就是获取 environ 变量，可以在往下看到有application 方法

 这里涉及到WSGI相关知识，可以看我的WSGI笔记

来看下这里面的application方法

```python
@classmethod
    def application(cls, f):
    from ..exceptions import HTTPException

        def application(*args):
            request = cls(args[-2])
            with request:
                try:
                    resp = f(*args[:-2] + (request,))
                except HTTPException as e:
                    resp = e.get_response(args[-2])
                return resp(*args[-2:])

        return update_wrapper(application, f)
```

很明显知道该方法是一个装饰器，被装饰的方法就是所谓的视图函数。


### cached_property
还发现一个有趣的装饰器 cached_property，该装饰器的作用与 property一样，因为它继承自 property，但多了缓存功能，而缓存的实现方式就是通过最简单的字典来实现的

```python
class cached_property(property):
    def __init__(self, func, name=None, doc=None):
        self.__name__ = name or func.__name__
        self.__module__ = func.__module__
        self.__doc__ = doc or func.__doc__
        self.func = func

    def __set__(self, obj, value):
        obj.__dict__[self.__name__] = value

    def __get__(self, obj, type=None):
        if obj is None:
            return self
        value = obj.__dict__.get(self.__name__, _missing)
        if value is _missing:
            value = self.func(obj)
            obj.__dict__[self.__name__] = value
        return value
```

    


以方法名作为 dict 的 key，方法的返回值为 dict 的 value，非常简单。


举个例子来说明：
```python
class Demo(object):
    def __init__(self):
        self.num = 100

    @property
    def add(self):
        self.num += 50
        return self.num


demo = Demo()
print(demo.add)  # 150
print(demo.add)  # 200
```


使用cached_property装饰器：

1.安装：
```
pip install cached_property
```

2.使用：
```python
from cached_property import cached_property
# 或者使用werkzeug自带的， 原理都一样的
# from werkzeug.utils import cached_property

class Demo(object):
    def __init__(self):
        self.num = 100

    @cached_property
    def add(self):
        self.num += 50
        return self.num

demo = Demo()
print(demo.add)  # 150
print(demo.add)  # 150

```


BaseRequest 类方法太多，这里就不详细分析了。

接着看一下 AcceptMixin 类，代码如下。

werkzeug/wrappers/accept.py
```python
class AcceptMixin(object):
    """A mixin for classes with an :attr:`~BaseResponse.environ` attribute
    to get all the HTTP accept headers as
    :class:`~werkzeug.datastructures.Accept` objects (or subclasses
    thereof).
    """

    @cached_property
    def accept_mimetypes(self):
        """List of mimetypes this client supports as
        :class:`~werkzeug.datastructures.MIMEAccept` object.
        """
        return parse_accept_header(self.environ.get("HTTP_ACCEPT"), MIMEAccept)

    @cached_property
    def accept_charsets(self):
        """List of charsets this client supports as
        :class:`~werkzeug.datastructures.CharsetAccept` object.
        """
        return parse_accept_header(
            self.environ.get("HTTP_ACCEPT_CHARSET"), CharsetAccept
        )

    @cached_property
    def accept_encodings(self):
        """List of encodings this client accepts.  Encodings in a HTTP term
        are compression encodings such as gzip.  For charsets have a look at
        :attr:`accept_charset`.
        """
        return parse_accept_header(self.environ.get("HTTP_ACCEPT_ENCODING"))

    @cached_property
    def accept_languages(self):
        """List of languages this client accepts as
        :class:`~werkzeug.datastructures.LanguageAccept` object.

        .. versionchanged 0.5
           In previous versions this was a regular
           :class:`~werkzeug.datastructures.Accept` object.
        """
        return parse_accept_header(
            self.environ.get("HTTP_ACCEPT_LANGUAGE"), LanguageAccept
        )

```



AcceptMixin 类方法主要用于处理 HTTP 头，因为使用了 self.environ，所以依赖于 BaseRequest 类。

类继承的本质其实就是将多个属性与方法放在一起。

比如 A 类继承于 B 类与 C 类，那么在 A 类的内存空间中，就存在 B 类与 C 类间不重复的属性与方法。

AcceptMixin 可以直接使用 self.environ 也是这个原理。

还有很多的方法 这里就分析到这里了

