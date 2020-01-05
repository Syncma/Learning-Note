# Requests 源码阅读-Day5


## Get方法

Session对象实例化后指向session，接着调用了其内部方法request：

从例子分析：

```python
[jian@laptop requests]$ cat test.py 
import requests
url = "http://www.baidu.com"
resp = requests.get(url)
```

然后到`requests/__init__.py` 看到是导入api.py里面的get方法

```
from .api import request, get, head, post, patch, put, delete, options
```

继续到api.py找到get方法

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


从request方法的注释中可以看出，该方法的作用是构造Request对象，准备并发送它，最后返回Response对象。

## Request

我们查看一下Request：
```python
def request(method, url, **kwargs):
    with sessions.Session() as session:
        return session.request(method=method, url=url, **kwargs)
```


**为什么要这样做呢？**

Session对象具有保留参数的功能，支持持久性的cookies以及urllib3的连接池功能

当我们向同一主机发送多个请求的时候，底层的TCP连接将会被重用，从而带来显著的性能提升，同时也为我们节省了很多工作量，不必为每次请求都去设置参数。

但是，不是每一次请求都需要保持长连接，保留参数，因为这会带来资源释放失败的风险，所以在我们常规方法中，引入了with 语句确保Session对象的正常退出。

通过代码演示一下两者的区别：

```python
>>> s = requests.session()
>>> s.get("http://httpbin.org/cookies/set/sessioncookie/123456789")
<Response [200]>
>>> r = s.get("http://httpbin.org/cookies")
>>> print(r.text)
{
    "cookies": {
        "sessioncookie": "123456789"
    }
}

>>> requests.get("http://httpbin.org/cookies/set/sessioncookie/123456789")
<Response [200]>
>>> r = requests.get("http://httpbin.org/cookies")
>>> print(r.text)
{
    "cookies": {}
}

```

着重看下session.request方法:
sessions.py
```python
class Session(SessionRedirectMixin):
    ...
    def request(self, method, url, ...):
        ...
        # Create the Request.
        req = Request(
            method=method.upper(),
            url=url,
            headers=headers,
            files=files,
            data=data or {},
            json=json,
            params=params or {},
            auth=auth,
            cookies=cookies,
            hooks=hooks,
        )
        prep = self.prepare_request(req)
        ...
... 
```

再继续研究Request， 其实是导入models.py里面的Request
```
from .models import Request, PreparedRequest, DEFAULT_REDIRECT_LIMIT
```

继续往里面看
models.py

```python
class Request(RequestHooksMixin):
    """A user-created :class:`Request <Request>` object.
    Used to prepare a :class:`PreparedRequest <PreparedRequest>`, which is sent to the server.
    ...
    """

    def __init__(self,
            method=None, url=None, headers=None, files=None, data=None,
            params=None, auth=None, cookies=None, hooks=None, json=None):
	...
```

类Request继承了类RequestHooksMixin，类RequestHooksMixin提供了hooks事件注册与注销的接口方法。

初始化方法实现了hooks事件注册，然后又是一波参数设置。

类Request对象是为后面的类PreparedRequest对象创建做准备。

Request对象实例构造完成后，继续调用了prepare_request方法：

sessions.py
```python
class Session(SessionRedirectMixin):
    """A Requests session.
    Provides cookie persistence, connection-pooling, and configuration.
    ...
    """
    
    ...
    def prepare_request(self, request):
        """Constructs a :class:`PreparedRequest <PreparedRequest>` for
        transmission and returns it. The :class:`PreparedRequest` has settings
        merged from the :class:`Request <Request>` instance and those of the
        :class:`Session`.
        ...
```

prepare_request(self, request)方法的作用是构造用于传输的PreparedRequest对象并返回它。

**`那么PreparedRequest对象是如何构建的？`**

它是由Request实例对象与Session对象中的数据（如cookies，stream，verify，proxies等等）合并而来。

为什么参数分别放在Request对象与Session对象中呢？

猜测与Session的参数持久化与连接池等有关，可以充分利用之前请求时存储的参数。

PreparedRequest对象构造完成后，最后再通过send方法将其发送出去：
<br>

sessions.py
```python
class Session(SessionRedirectMixin):
    """A Requests session.
    Provides cookie persistence, connection-pooling, and configuration.
    ...
    """
    ...
    def request(self, method, url, ...):
        """Constructs a :class:`Request <Request>`, prepares it and sends it.
        Returns :class:`Response <Response>` object.
        ...
        :rtype: requests.Response
        """
        ...
        resp = self.send(prep, **send_kwargs)
        return resp
        ...
    
        ...
        def send(self, request, **kwargs):
        """Send a given PreparedRequest.

        :rtype: requests.Response
        """
        ...
        # Get the appropriate adapter to use
        adapter = self.get_adapter(url=request.url)
				
        ...
        
        # Send the request
        r = adapter.send(request, **kwargs)
        ...
        
        return r
```

send方法：
```python
def send(self, request, **kwargs):
        """Send a given PreparedRequest.

        :rtype: requests.Response
        """
        # Set defaults that the hooks can utilize to ensure they always have
        # the correct parameters to reproduce the previous request.
        kwargs.setdefault('stream', self.stream)
        kwargs.setdefault('verify', self.verify)
        kwargs.setdefault('cert', self.cert)
        kwargs.setdefault('proxies', self.proxies)

        # It's possible that users might accidentally send a Request object.
        # Guard against that specific failure case.
        if isinstance(request, Request):
            raise ValueError('You can only send PreparedRequests.')
           ...
           ...
            
```

send方法接收PreparedRequest对象，然后根据该对象的url参数获取对应的传输适配器。

这个传输适配器就是我们前面讲的HTTPAdapter，它的底层由强大的urllib3库实现，为requests提供了的HTTP和HTTPS全部接口方法，包括send等。

当send方法将请求发送给服务器后，等待服务器的响应，最后返回Response对象。

这里又涉及到`NotImplementedError`以及生成器的一些用法

到这里，一个完整的HTTP请求完成了。

总之：

> sessions.py是最大的难点，因为里面包含了很多东西

