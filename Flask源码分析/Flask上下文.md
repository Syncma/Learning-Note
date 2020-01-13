# Flask上下文
<!-- TOC -->

- [Flask上下文](#flask%e4%b8%8a%e4%b8%8b%e6%96%87)
  - [上下文是什么](#%e4%b8%8a%e4%b8%8b%e6%96%87%e6%98%af%e4%bb%80%e4%b9%88)
  - [Flask 上下文作用](#flask-%e4%b8%8a%e4%b8%8b%e6%96%87%e4%bd%9c%e7%94%a8)
  - [Flask 上下文机制](#flask-%e4%b8%8a%e4%b8%8b%e6%96%87%e6%9c%ba%e5%88%b6)
    - [Local类](#local%e7%b1%bb)
    - [LocalStack类](#localstack%e7%b1%bb)
    - [LocalProxy类](#localproxy%e7%b1%bb)
    - [小总结](#%e5%b0%8f%e6%80%bb%e7%bb%93)
  - [问题1](#%e9%97%ae%e9%a2%981)
  - [问题2](#%e9%97%ae%e9%a2%982)
  - [问题3](#%e9%97%ae%e9%a2%983)
  - [问题4](#%e9%97%ae%e9%a2%984)

<!-- /TOC -->

## 上下文是什么

日常生活中的上下文：

>从一篇文章中抽取一段话，你阅读后，可能依旧无法理解这段话中想表达的内容，因为它引用了文章其他部分的观点，要理解这段话，需要先阅读理解这些观点。
>这些散落于文章的观点就是这段话的上下文。


对Flask框架来说就是：

>Flask从客户端收到请求的时候，视图函数如果要处理请求的话，可能就要访问一些对象。那么这些对象可以通过参数的形式传递进来，或者在函数中访问外部变量。

>这个外部变量要有特定的值才会有意义，也就是上下文

看一个例子：

```python
from flask import Flask， request

app = Flask(__name__)


@app.route('/')
def index():
    user_agent = request.headers.get('User-Agent')
    return 'Your brower is %s' % user_agent


if __name__ == '__main__':
    app.run()
    
```

这里的request变量就是请求上下文，也就是当请求被推送之后，request才有意义，接下来才可以用request



## Flask 上下文作用
程序中的上下文：一个函数通常涉及了外部变量 (或方法)，要正常使用这个函数，就需要先赋值给这些外部变量，这些外部变量值的集合就称为上下文

Flask 的视图函数需要知道前端请求的 url、参数以及数据库等应用信息才可以正常运行，要怎么做？

一个粗暴的方法是将这些信息通过传参的方式一层层传到到视图函数，太不优雅。

Flask 为此设计出了自己的上下文机制，当需要使用请求信息时，直接 from flask import request就可以获得当前请求的所有信息并且在多线程环境下是线程安全的，很酷。

实现这种效果的大致原理其实与 threading.local 实现原理相同，创建一个全局的字典对象，利用线程 id 作为 key，相应的数据作为 value，这样，不同的线程就可以获取专属于自己的数据。


## Flask 上下文机制
Flask 上下文定义在 globals.py 上
```python
from functools import partial

from werkzeug.local import LocalProxy
from werkzeug.local import LocalStack


_request_ctx_err_msg = """\
Working outside of request context.

This typically means that you attempted to use functionality that needed
an active HTTP request.  Consult the documentation on testing for
information about how to avoid this problem.\
"""
_app_ctx_err_msg = """\
Working outside of application context.

This typically means that you attempted to use functionality that needed
to interface with the current application object in some way. To solve
this, set up an application context with app.app_context().  See the
documentation for more information.\
"""


def _lookup_req_object(name):
    top = _request_ctx_stack.top
    if top is None:
        raise RuntimeError(_request_ctx_err_msg)
    return getattr(top, name)


def _lookup_app_object(name):
    top = _app_ctx_stack.top
    if top is None:
        raise RuntimeError(_app_ctx_err_msg)
    return getattr(top, name)


def _find_app():
    top = _app_ctx_stack.top
    if top is None:
        raise RuntimeError(_app_ctx_err_msg)
    return top.app


# context locals
_request_ctx_stack = LocalStack()
_app_ctx_stack = LocalStack()
current_app = LocalProxy(_find_app)
request = LocalProxy(partial(_lookup_req_object, "request"))
session = LocalProxy(partial(_lookup_req_object, "session"))
g = LocalProxy(partial(_lookup_app_object, "g"))
```

Flask 中看似有多个上下文，但其实都衍生于 _request_ctx_stack与 _app_ctx_stack

 **`_request_ctx_stack是请求上下文， _app_ctx_stack是应用上下文`**

从代码里面看到：
```
常用的 request 和 session 衍生于 _request_ctx_stack，
current_app 和 g 衍生于 _app_ctx_stack。
```

可以发现，这些上下文的实现都使用了 LocalStack 和 LocalProxy，这两个类的实现在 werkzeug 中

在实现这两个类之前，需要先理解 Local 类

### Local类
werkzeug/local.py
```python
class Local(object):
    __slots__ = ("__storage__", "__ident_func__")

    def __init__(self):
        object.__setattr__(self, "__storage__", {})
        object.__setattr__(self, "__ident_func__", get_ident)

    def __iter__(self):
        return iter(self.__storage__.items())

    def __call__(self, proxy):
        """Create a proxy for a name."""
        return LocalProxy(self, proxy)
        ...

```

看源码知道Local 类重写了 `__getattr__、 __setattr__和 __delattr__`，从而自定义了 Local 对象属性访问、设置与删除对应的操作

这些方法都通过 `__ident_func__()`方法获取当前线程 id 并以此作为 key 去操作当前线程对应的数据，Local 通过这种方式实现了多线程数据的隔离。


###  LocalStack类
LocalStack 类是基于 Local 类实现的栈结构
```python
class LocalStack(object):
    def __init__(self):
        self._local = Local()
        ...

```

LocalStack 类的代码简洁易懂，主要的逻辑就实例化 Local 类，获得 local 对象，在 local 对象中添加 stack，以 list 的形式来实现一个栈，

至此可知， _request_ctx_stack与 _app_ctx_stack这两个上下文就是一个线程安装的栈，线程所有的信息都会保存到相应的栈里，直到需要使用时，再出栈获取。

### LocalProxy类
LocalProxy 类是 Local 类的代理对象，它的作用就是将操作都转发到 Local 对象上。

```python
@implements_bool
class LocalProxy(object):
....
```

LocalProxy 类在` __init__()`方法中将 local 实例赋值给 `_LocalProxy__local`

并在后续的方法中通过` __local`的方式去操作它

因为 LocalProxy 类重写了 `__setattr__`方法，所以不能直接复制，此时要通过 `object.__setattr__()`进行赋值。


LocalProxy 类后面的逻辑其实都是一层代理，将真正的处理交个 local 对象。


### 小总结

```
所谓 Flask 上下文，其实就是基于 list 实现的栈，

这个 list 存放在 Local 类实例化的对象中，Local 类利用线程 id 作为字典的 key，线程具体的值作为字典的 values 来实现线程安全，使用的过程就是出栈入栈的过程，此外，在具体使用时，会通过 LocalProxy 类将操作都代理给 Local 类对象。
```


##  问题1
为何需要 werkzeug 库的 Local 类？
treading 标准库中已经提供了 local 对象，该对象实现的效果与 Local 类似，以线程 id 为字典的 key，将线程具体的值作为字典的 values 存储，简单使用如下。

```
 In [1]: import threading
 In [2]: local = threading.local()
 In [3]: local.name = 'tony'
 In [4]: local.name
 Out[4]: 'tony'
```

那为何 werkzeug 库要自己再实现一个功能类似的 Local 类呢？

主要原因是为了**`兼容协程`**

当用户通过 greenlet 库来构建协程时，因为多个协程可以在同一个线程中，threading.local 无法处理这种情况
而 Local 可以通过 getcurrent () 方法来获取协程的唯一标识。

werkzeug/local.py
```python
try:
    from greenlet import getcurrent as get_ident
except ImportError:
    try:
        from thread import get_ident
    except ImportError:
        from _thread import get_ident
```


## 问题2

为什么不构建一个上下文而是要将其分为请求上下文 (request context) 和应用上下文 (application context)？

**`为了灵活度`**

虽然在实际的 Web 项目中，每个请求只会对应一个请求上下文和应用上下文，但在 debug 或使用 flask shell 时，用户可以单独构建新的上下文，将一个上下文以请求上下文和应用上下文的形式分开，可以让用户单独创建其中一种上下文，这很方便用户在不同的情景使用不同的上下文。


## 问题3

为什么不直接使用 Local？而要通过 LocalStack 类将其封装成栈的操作？

总结而言，**通过 LocalStack 实现栈结构而不直接使用 Local 的目的是为了在多应用情景下让一个请求可以很简单的知道当前上下文是哪个。**

要理解这个回答，先要回顾一下 Flask 多应用开发的内容并将其与上下文的概念结合在一起理解。

Flask 多应用开发的简单例子如下。
```python
from werkzeug.wsgi import DispatcherMiddleware
from werkzeug.serving import run_simple
from flask import Flask

frontend = Flask('frontend')
backend = Flask('backend')

@frontend.route('/home')
def home():
    return 'frontend home'

@backend.route('/home')
def home():
    return 'backend home'

app = DispatcherMiddleware(frontend, {
    '/frontend': frontend,
    '/backend': backend
})

if __name__ == "__main__":
    run_simple('127.0.0.1', 5000, app)
```

利用 werkzeug 的 DispatcherMiddleware，让一个 Python 解释器可以同时运行多个独立的应用实例，其效果虽然跟使用蓝图时的效果类似

但要注意，此时是多个独立的 Flask 应用，具体而言，每个独立的 Flask 应用都创建了自己的上下文。

每个独立的 Flask 应用都是一个合法的 WSGI 应用，利用 DispatcherMiddleware，通过调度中间件的逻辑将多个 Flask 应用组合成一个大应用。

简单理解 Flask 多应用后，回顾一下 Flask 上下文的作用。

比如，要获得当前请求的 path 属性，可以通过如下方式。
```python
from flask import request
print(request.path)
```

Flask 在多应用的情况下，依旧可以通过 request.path 获得当前应用的信息，实现这个效果的前提就是，Flask 知道当前请求对应的上下文。

栈结构很好的实现了这个前提，每个请求，其相关的上下文就在栈顶，直接将栈顶上下文出栈就可以获得当前请求对应上下文中的信息了。

在上面 Flask 多应用的代码中，构建了 frontend 应用与 backend 应用，两个应用相互独立，分别负责前端逻辑与后端逻辑，通过 DispatcherMiddleware 将其整合在一起，这种情况下，appctx_stack 栈中就会有两个应用上下文。

访问 127.0.0.1:5000/backend/home时，backend 应用上下文入栈，成为栈顶。
想要获取当前请求中的信息时，直接出栈就可以获得与当前请求对应的上下文信息。

需要注意，请求上下文、应用上下文是具体的对象，而 requestctxstack (请求上下文栈) 与app ctxstack (应用上下文栈) 是数据结构

所谓栈就是一个 list，结合 Local 类的代码，上下文堆栈其结构大致为 {thread.get_ident():[]}，每个线程都有独立的一个栈。

此外，Flask 基于栈结构可以很容易实现内部重定向。
• 外部重定向：用户通过浏览器请求 URL-1 后，服务器返回 302 重定向请求，让其请求 URL-2，用户的浏览器会发起新的请求，请求新的 URL-2，获得相应的数据。

• 内部重定向：用户通过浏览器请求 URL-1 后，服务器内部之间将 ULR-2 对应的信息直接返回给用户。


**Flask 在内部通过多次入栈出栈的操作可以很方便的实现内部重定向。**


## 问题4

为什么不直接使用 LocalStack？而要通过 LocalProxy 类来代理操作？
这是因为 Flask 的上下文中保存的数据都是存放在栈里并且会动态变化的，
通过 LocalProxy 可以动态的访问相应的对象，从而避免造成数据访问异常。


怎么理解？看一个简单的例子，首先，直接操作 LocalStack，代码如下。

```python
from werkzeug.local import LocalStack

l_stack = LocalStack()
l_stack.push({'name': 'ayuliao'})
l_stack.push({'name': 'twotwo'})


def get_name():
    return l_stack.pop()


name = get_name()
print(f"name is {name['name']}")
print(f"name is {name['name']}")
```

运行上述代码，输出的结果如下。
```
 name is twotwo
name is twotwo
```

可以发现，结果相同。


利用 LocalProxy 代理操作，代码如下。
```
from werkzeug.local import LocalStack, LocalProxy

l_stack = LocalStack()
l_stack.push({'name': 'ayuliao'})
l_stack.push({'name': 'twotwo'})


def get_name():
    return l_stack.pop()


# 代理操作get_name
name2 = LocalProxy(get_name)
print(f"name is {name2['name']}")
print(f"name is {name2['name']}")
```

运行上述代码，输出的结果如下。
```
 name is twotwo
name is ayuliao
```


通过 LocalProxy 代理操作后，结果不同。

通过 LocalProxy 代理操作后，每一次获取值的操作其实都会调用 `__getitem__`，该方法是个匿名函数，x 就是 LocalProxy 实例本身，这里即为 name2，而 i 则为查询的属性，这里即为 name。

```python
class LocalProxy(object):
# ... 省略部分代码
__getitem__ = lambda x, i: x._get_current_object()[i]

# 结合 __init__与 _get_current_object()方法来看。
class LocalProxy(object):
    __slots__ = ("__local", "__dict__", "__name__", "__wrapped__")

    def __init__(self, local, name=None):
        object.__setattr__(self, "_LocalProxy__local", local)
        object.__setattr__(self, "__name__", name)
        if callable(local) and not hasattr(local, "__release_local__"):
            # "local" is a callable that is not an instance of Local or
            # LocalManager: mark it as a wrapped function.
            object.__setattr__(self, "__wrapped__", local)

    def _get_current_object(self):
        """Return the current object.  This is useful if you want the real
        object behind the proxy at a time for performance reasons or because
        you want to pass the object into a different context.
        """
        if not hasattr(self.__local, "__release_local__"):
            return self.__local()
        try:
            return getattr(self.__local, self.__name__)
        except AttributeError:
            raise RuntimeError("no object bound to %s" % self.__name__)
```

在 `__init__`方法中，将 getname 赋值给了` _LocalProxy__local`，因为 getname 不存在 `__release_local__`属性，此时使用 `_get_current_object()`方法，相当于再次执行 `get_name ()`，出栈后获得新的值。

通过上面的分析，明白了通过 LocalProxy 代理后，调用两次 name['name']获取的值不同的原因。

那为什么要这样做？看到 Flask 中 globals.py 的部分代码。
```
# context locals
_request_ctx_stack = LocalStack()
_app_ctx_stack = LocalStack()
current_app = LocalProxy(_find_app)
request = LocalProxy(partial(_lookup_req_object, "request"))
session = LocalProxy(partial(_lookup_req_object, "session"))
g = LocalProxy(partial(_lookup_app_object, "g"))
```

当前应用 current app 是通过 `LocalProxy(_find_app)`获得的，即每次调用 currentapp () 会执行出栈操作，获得与当前请求相对应的上下文信息。

如果 `current_app=_find_app()`，此时 current_app 就不会再变化了，在多应用多请求的情况下是不合理的，会抛出相应的异常。


