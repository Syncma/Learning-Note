# Requests 源码阅读-Day2

<!-- TOC -->

- [Requests 源码阅读-Day2](#requests-%e6%ba%90%e7%a0%81%e9%98%85%e8%af%bb-day2)
  - [get方法](#get%e6%96%b9%e6%b3%95)
  - [request](#request)
  - [Session](#session)
    - [Session.__enter__](#sessionenter)
    - [Session.__init__](#sessioninit)
      - [问题1](#%e9%97%ae%e9%a2%981)
      - [问题2](#%e9%97%ae%e9%a2%982)

<!-- /TOC -->
## get方法

再来看这个文件：
tests/test_requests.py
```python
 def test_DIGEST_HTTP_200_OK_GET(self, httpbin):

        for authtype in self.digest_auth_algo:
            auth = HTTPDigestAuth('user', 'pass')
            url = httpbin('digest-auth', 'auth', 'user', 'pass', authtype, 'never')
            pytest.set_trace()

            r = requests.get(url, auth=auth)
            assert r.status_code == 200

            r = requests.get(url)
            assert r.status_code == 401
            print(r.headers['WWW-Authenticate'])

            s = requests.session()
            s.auth = HTTPDigestAuth('user', 'pass')
            r = s.get(url)
            assert r.status_code == 200

```

这里我们分析requests.get方法， 到requests模块中找到`__init__.py`文件
看到：
```python
from .api import request, get, head, post, patch, put, delete, options
```

<br>
OK, 不废话 直接找到api.py

```python
def get(url, params=None, **kwargs):
    r"""Sends a GET request.

    :param url: URL for the new :class:`Request` object.
    :param params: (optional) Dictionary, list of tuples or bytes to send
        in the query string for the :class:`Request`.
    :param \*\*kwargs: Optional arguments that ``request`` takes.
    :return: :class:`Response <Response>` object
    :rtype: requests.Response
    """

    kwargs.setdefault('allow_redirects', True)
    return request('get', url, params=params, **kwargs)
    
```
<br>

**该方法的作用是向url指定的地址发起GET请求。**

输入参数分别为：

>url：url全称叫统一资源定位符，即访问对象在互联网中的唯一地址。

>params：可选参数，字典类型，为请求提供查询参数，最后构造到url中。

>`**kwargs`：参数前加**在方法中会转换为字典类型，作为请求方法request的可选参数。

>kwargs.setdefault('allow_redirects', True)，设置默认键值对，若键值不存在，则插入值为"True"的键'allow_redirects'。

>返回请求方法request对象。


## request


再看这个request对象

```python
def request(method, url, **kwargs):
	    with sessions.Session() as session:
        return session.request(method=method, url=url, **kwargs)
```

<br>

请求方法request包含了许多输入参数

> with sessions.Session() as session，with语句的作用是确保**session对象无论是否正常运行都能确保正确退出，避免程序异常导致sockets接口无法正常关闭**。

> 最后返回session.request对象。


`那with是什么？`with是用来实现上下文管理的。

`那上下文管理是什么？`为了保证with对象无论是否正常运行都能确保正确退出。

with语句的原型如下：

```
with expression [as variable]:
    with-block
```

with语句中的[as variable]是可选的，如果指定了as variable说明符，则variable就是上下文管理器expression.`__enter__()`方法返回的对象。

with-block是执行语句，with-block执行完毕时，with语句会自动调用`expression.__exit__()`方法进行资源清理。

我们常见的读写文件建议使用with写法就是这个道理

例子：
```python
file = open("welcome.txt")
data = file.read()
print(data)

file.close()

```

使用with写法：
```python
with open("welcome.txt") as file:
    data = file.read()
    # do something
    
```

结合with语句，该部分代码的实现一目了然：

```python
session = sessions.Session().`__enter__`(self)  # 也即Session实例本身。
session.request(method=method, url=url, **kwargs) # 为with语句执行部分。
```

当执行部分session.request方法调用完成，

sessions.Session().`__exit__`(self, *args)方法被调用，

接着Session对象中的close(self)方法被执行，

完成Session对象资源的销毁，最后退出。

以上就是with语句的用途

其实with语句执行完后，requests.get方法也就执行完了，一次请求也即完成。



##  Session

上面说的sessions其实是导入了sessions.py这个文件里面的

```python
class Session(SessionRedirectMixin):
    """A Requests session.

    Provides cookie persistence, connection-pooling, and configuration.

    Basic Usage::

      >>> import requests
      >>> s = requests.Session()
      >>> s.get('https://httpbin.org/get')
      <Response [200]>

    Or as a context manager::

      >>> with requests.Session() as s:
      ...     s.get('https://httpbin.org/get')
      <Response [200]>
    """
    ...
```

Session是什么？ 这里要`好好研读下源码`


主要功能：

支持持久性的cookies，使用urllib3连接池功能，对参数进行配置，为request对象提供参数，拥有所有的请求方法等。

原来我们所有的设置操作，真真正正开始执行是在Session对象里。

同时Session继承了类SessionRedirectMixin，这个类实现了重定向的接口方法。

重定向的意思就是当我们通过url指定的路径向服务器请求资源时，发现该资源并不在url指定的路径上，这时服务器通过响应，给出新的资源地址，然后我们通过新的url再次发起请求。- `这里又涉及到了Mixin类的作用-这个会另外写一章节进行讲解`

接下去，我们来分析Session是如何被调用的。

前面提到过，Session调用时采用了with的方法，
然后我们看下源码中的with语句以及上下文管理器expression方法实现部分：

sessions.py

```python
class Session(SessionRedirectMixin):

    ...
    def __enter__(self):
        return self

    def __exit__(self, *args):
        self.close()
        ...
    
    ...
    def close(self):
        """Closes all adapters and as such the session"""
        for v in self.adapters.values():
            v.close()
```


### `Session.__enter__`

回到with语句中session获得上下文管理器sessions.Session()的`__enter__`(self)对象，

先会调用这个方法， 这个方法的返回值是 **<requests.sessions.Session object at 0x7fb690228080>**，也就是调用初始化方法`__init__`(self)

接下来对`__init__`方法分析


### `Session.__init__`

```python
def __init__(self):

        #: A case-insensitive dictionary of headers to be sent on each
        #: :class:`Request <Request>` sent from this
        #: :class:`Session <Session>`.
        self.headers = default_headers()

        #: Default Authentication tuple or object to attach to
        #: :class:`Request <Request>`.
        self.auth = None
		...
```

<br>

初始化方法主要实现了参数的默认设置，包括headers，auth，proxies，stream，verify，cookies，hooks等等。


首先我们看下header参数是怎么写的：

在发起一次请求时没有设置headers参数，那么header就会使用默认参数，由方法default_headers()来设置

utils.py
```python
def default_headers():
    """
    :rtype: requests.structures.CaseInsensitiveDict
    """
    return CaseInsensitiveDict({
        'User-Agent': default_user_agent(),
        'Accept-Encoding': ', '.join(('gzip', 'deflate')),
        'Accept': '*/*',
        'Connection': 'keep-alive',
    })
    
```

这时你会发现header默认参数中用户代理'User-Agent'将被设置为"python-requests"，

如果你正在写爬虫程序抓取某个网站的数据，那么建议你尽快修改用户代理，因为对方服务器可能很快就拒绝一个来之python的访问。

这里的CaseInsensitiveDict 方法是做什么用的，一直往里面分析。


它其实是**structures.py**里面的方法， 这个方法做什么事情呢？

`主要作用是：大小写不敏感的dict key-value`

```python
class CaseInsensitiveDict(MutableMapping):
    def __init__(self, data=None, **kwargs):
        self._store = OrderedDict()
        if data is None:
            data = {}
        self.update(data, **kwargs)
       ...
```
首先它继承了MutableMapping类，然后初始化做了一个赋值操作OrderedDict对象

两个问题：

`1.继承MutableMapping类的作用是什么？`

`2.OrderedDict对象是什么东西？`


慢慢来分析：

####  问题1

看到**structures.py**文件开头上面的import 信息，
```
from .compat import Mapping, MutableMapping
```


我们看到它其实是导入的是compat.py里面的MutableMapping方法

进入这个compat.py文件，看到首先进行了python版本判断：

```python

import chardet

import sys

# -------
# Pythons
# -------

# Syntax sugar.
_ver = sys.version_info

#: Python 2.x?
is_py2 = (_ver[0] == 2)

#: Python 3.x?
is_py3 = (_ver[0] == 3)
...
```

这里我本机是python3 版本 直接看python3的

```python
elif is_py3:
    from urllib.parse import urlparse, urlunparse, urljoin, urlsplit, urlencode, quote, unquote, quote_plus, unquote_plus, urldefrag
    from urllib.request import parse_http_list, getproxies, proxy_bypass, proxy_bypass_environment, getproxies_environment
    from http import cookiejar as cookielib
    from http.cookies import Morsel
    from io import StringIO
    # Keep OrderedDict for backwards compatibility.
    from collections import OrderedDict
    from collections.abc import Callable, Mapping, MutableMapping

    builtin_str = str
    str = str
    bytes = bytes
    basestring = (str, bytes)
    numeric_types = (int, float)
    integer_types = (int,)
    
```

可以看到其实调用的是collections.abc模块里面的MutableMapping方法，然后我们继续分析

这个collections.abc模块是系统自带的模块，根据python模块路径查找到这个模块

collections/abc.py
```python
from _collections_abc import *
from _collections_abc import __all__
```

调用的是`_collections_abc`模块里面的方法，继续分析这个也是系统自带的模块

叫`_collections_abc.py`, 进入这个模块找到MutableMappingf方法

```python

class MutableMapping(Mapping):

    __slots__ = ()
    """A MutableMapping is a generic container for associating
    key/value pairs.

    This class provides concrete generic implementations of all
    methods except for __getitem__, __setitem__, __delitem__,
    __iter__, and __len__.

    """

    @abstractmethod
    def __setitem__(self, key, value):
        raise KeyError

    @abstractmethod
    def __delitem__(self, key):
        raise KeyError

    __marker = object()

    def pop(self, key, default=__marker):
        '''D.pop(k[,d]) -> v, remove specified key and return the corresponding value.
          If key is not found, d is returned if given, otherwise KeyError is raised.
        '''
        try:
            value = self[key]
        except KeyError:
            if default is self.__marker:
                raise
            return default
        else:
            del self[key]
            return value

    def popitem(self):
        '''D.popitem() -> (k, v), remove and return some (key, value) pair
           as a 2-tuple; but raise KeyError if D is empty.
        '''
        try:
            key = next(iter(self))
        except StopIteration:
            raise KeyError
        value = self[key]
        del self[key]
        return key, value

    def clear(self):
        'D.clear() -> None.  Remove all items from D.'
        try:
            while True:
                self.popitem()
        except KeyError:
            pass

    def update(*args, **kwds):
        ''' D.update([E, ]**F) -> None.  Update D from mapping/iterable E and F.
            If E present and has a .keys() method, does:     for k in E: D[k] = E[k]
            If E present and lacks .keys() method, does:     for (k, v) in E: D[k] = v
            In either case, this is followed by: for k, v in F.items(): D[k] = v
        '''
        if not args:
            raise TypeError("descriptor 'update' of 'MutableMapping' object "
                            "needs an argument")
        self, *args = args
        if len(args) > 1:
            raise TypeError('update expected at most 1 arguments, got %d' %
                            len(args))
        if args:
            other = args[0]
            if isinstance(other, Mapping):
                for key in other:
                    self[key] = other[key]
            elif hasattr(other, "keys"):
                for key in other.keys():
                    self[key] = other[key]
            else:
                for key, value in other:
                    self[key] = value
        for key, value in kwds.items():
            self[key] = value

    def setdefault(self, key, default=None):
        'D.setdefault(k[,d]) -> D.get(k,d), also set D[k]=d if k not in D'
        try:
            return self[key]
        except KeyError:
            self[key] = default
        return default


MutableMapping.register(dict)

```

首先它继承了 Mapping 类，这个Mapping类是什么东西，

**`继续分析其实是继承了ABCMeta元类，这个ABC元类是做什么的？我会单独写一章介绍这个`**

通读代码发现其实它的作用就是对字典进行了一系列操作

重点是最后一句：

```python

MutableMapping.register(dict)

```

**`这里的作用是将"子类"注册为该抽象基类的”抽象子类"`**

例如：
```
from abc import ABC

class MyABC(ABC):
    pass

MyABC.register(tuple)

assert issubclass(tuple, MyABC)
assert isinstance((), MyABC)

```


####  问题2

OrderedDict对象是什么东西？

看到它其实是导入了collections 模块
```python
from collections import OrderedDict
```

这个collections模块也是系统自带的，到模块路径查看具体内容：

找到collections.`__init__.py`文件里面的OrderedDict方法：

```python
class OrderedDict(dict):
    'Dictionary that remembers insertion order'
    # An inherited dict maps keys to values.
    # The inherited dict provides __getitem__, __len__, __contains__, and get.
    # The remaining methods are order-aware.
    # Big-O running times for all methods are the same as regular dictionaries
    
    ....
```

它的作用就是做了个顺序的dict， 为啥要做个顺序dict， **`是为了解决啥问题呢`**？

用传统的dict 方法有什么不好的地方呢?

python中的字典是无序的，因为它是按照hash来存储的，但是OrderedDict，实现了对

字典对象中元素的排序，并且字典顺序保证是插入顺序


前面说的update方法 是继承compat里面的MutableMapping方法，继续深挖，
`_collections_abc.py`, 进入这个模块找到update方法

```python

from .compat import Mapping, MutableMapping


class CaseInsensitiveDict(MutableMapping):
....
def update(*args, **kwds):
        ''' D.update([E, ]**F) -> None.  Update D from mapping/iterable E and F.
            If E present and has a .keys() method, does:     for k in E: D[k] = E[k]
            If E present and lacks .keys() method, does:     for (k, v) in E: D[k] = v
            In either case, this is followed by: for k, v in F.items(): D[k] = v
        '''
        if not args:
            raise TypeError("descriptor 'update' of 'MutableMapping' object "
                            "needs an argument")
        self, *args = args
        if len(args) > 1:
            raise TypeError('update expected at most 1 arguments, got %d' %
                            len(args))
        if args:
            other = args[0]
            if isinstance(other, Mapping):
                for key in other:
                    self[key] = other[key]
            elif hasattr(other, "keys"):
                for key in other.keys():
                    self[key] = other[key]
            else:
                for key, value in other:
                    self[key] = value

        for key, value in kwds.items():
            self[key] = value
            
```

update的功能是对字典进行了一些key, value的赋值