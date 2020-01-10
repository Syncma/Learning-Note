# Requests 源码阅读-Day3

<!-- TOC -->

- [Requests 源码阅读-Day3](#requests-%e6%ba%90%e7%a0%81%e9%98%85%e8%af%bb-day3)
  - [hooks](#hooks)
    - [hooks初始化](#hooks%e5%88%9d%e5%a7%8b%e5%8c%96)
    - [hooks使用](#hooks%e4%bd%bf%e7%94%a8)
  - [cookies](#cookies)

<!-- /TOC -->

## hooks

###  hooks初始化

继续前面的分析，来到了hook这部分

```python
class Session(SessionRedirectMixin):

    __attrs__ = [
        'headers',
        'cookies',
        'auth',
        'proxies',
        'hooks',
        'params',
        'verify',
        'cert',
        'prefetch',
        'adapters',
        'stream',
        'trust_env',
        'max_redirects',
    ]

    def __init__(self):

        #: A case-insensitive dictionary of headers to be sent on each
        #: :class:`Request <Request>` sent from this
        #: :class:`Session <Session>`.
        self.headers = default_headers()
        print("headers=", self.headers)

        #: Default Authentication tuple or object to attach to
        #: :class:`Request <Request>`.
        self.auth = None

        #: Dictionary mapping protocol or protocol and host to the URL of the proxy
        #: (e.g. {'http': 'foo.bar:3128', 'http://host.name': 'foo.bar:4012'}) to
        #: be used on each :class:`Request <Request>`.
        self.proxies = {}

        #: Event-handling hooks.
        self.hooks = default_hooks()
       ....

```


看到调用的是default_hook方法，然后到文件开头可以看到是导入了hooks.py

```python
from .hooks import default_hooks, dispatch_hook
```

到hooks.py这个文件查看：
```python
"""
requests.hooks
~~~~~~~~~~~~~~

This module provides the capabilities for the Requests hooks system.

Available hooks:

``response``:
    The response generated from a Request.
"""
HOOKS = ['response']

def default_hooks():
    return {event: [] for event in HOOKS}
```

<br>

就写了一句话，这是什么意思呢？

通过注释我们知道**hooks意为事件挂钩，可以用来操控部分请求过程或者信号事件处理。**

requests有一个钩子系统，在请求产生的响应response前，做一些想做的事情。

上述源代码中default_hooks()方法用了字典解析，最后返回{'response':[]}。

###  hooks使用

下面我们简单描述一下hooks是如何使用的。

首先我们需要传递一个字典{hook_name:callback_function}给参数hooks。

hook_name为钩子名，也就是 "response",callback_function为钩子方法，在目标事件发生时回调该方法。

callback_function会接受一个数据块作为它的第一个参数，

定义如下def callback_function(r, *args, **kwargs)。

从default_hooks()方法返回的hooks默认参数{'response':[]}

可知，键"response"所对应的值为一个列表

换句话说，对于一个钩子事件，可以有多个钩子方法。


下面我们写个例子演示一下。

```python
>>> def hooks1(r, *args, **kwargs):
...     print("hooks1 url=" + r.url)
... 
>>> def hooks2(r, *args, **kwargs):
...     print("hooks2 encoding=" + r.encoding)
... 
>>> hooks = dict(response=[hooks1,hooks2])
>>> requests.get("http://httpbin.org", hooks=hooks)
hooks1 url=http://httpbin.org/
hooks2 encoding=utf-8
<Response [200]>

```



## cookies

继续往下分析看看cookies初始化：

```python
        #: A CookieJar containing all currently outstanding cookies set on this
        #: session. By default it is a
        #: :class:`RequestsCookieJar <requests.cookies.RequestsCookieJar>`, but
        #: may be any other ``cookielib.CookieJar`` compatible object.
        self.cookies = cookiejar_from_dict({})
        
```

这里面的cookiejar_from_dict 到文件开头看是导入cookies.py：
```python
from .cookies import (cookiejar_from_dict, extract_cookies_to_jar,
```

 进入cookies.py查看到这个方法：
```python
def cookiejar_from_dict(cookie_dict, cookiejar=None, overwrite=True):
    """Returns a CookieJar from a key/value dictionary.

    :param cookie_dict: Dict of key/values to insert into CookieJar.
    :param cookiejar: (optional) A cookiejar to add the cookies to.
    :param overwrite: (optional) If False, will not replace cookies
        already in the jar with new ones.
    :rtype: CookieJar
    """
    if cookiejar is None:
        cookiejar = RequestsCookieJar()

    if cookie_dict is not None:
        names_from_jar = [cookie.name for cookie in cookiejar]
        for name in cookie_dict:
            if overwrite or (name not in names_from_jar):
                cookiejar.set_cookie(create_cookie(name, cookie_dict[name]))

    return cookiejar
```


这里面的RequestsCookieJar()是什么东西?

定位到该文件里面的RequestsCookieJar方法

```python
class RequestsCookieJar(cookielib.CookieJar, MutableMapping):
...
```

可以看到这又用到了多继承，继承了cookielib.CookieJar和MutableMapping类，

问题：

`1.继承这两个类是做什么用呢？`

`2.前面的cookiejar = RequestsCookieJar() 究竟执行了什么操作？`



分别进行解答：
1.继承这两个类是为了下面调用里面的方法，MutableMapping这个在前面的章节已经分析过了，这里就不在说明

2.RequestsCookieJar() 其实调用了cookielib.CookieJar里面的`__init__`方法

cookies.py
```python
from .compat import cookielib, urlparse, urlunparse, Morsel, MutableMapping
```

->调用了compat.py里面cookielib ， 继续分析调用了http模块里面的cookiejar
```python
...
elif is_py3:
    from urllib.parse import urlparse, urlunparse, urljoin, urlsplit, urlencode, quote, unquote, quote_plus, unquote_plus, urldefrag
    from urllib.request import parse_http_list, getproxies, proxy_bypass, proxy_bypass_environment, getproxies_environment
    from http import cookiejar as cookielib
    
```

然后到http模块里面查看这个cookiejar

http/cookiejar.py 里面
```python
class CookieJar:
    """Collection of HTTP cookies.

    You may not need to know about this class: try
    urllib.request.build_opener(HTTPCookieProcessor).open(url).
    """

    non_word_re = re.compile(r"\W")
    quote_re = re.compile(r"([\"\\])")
    strict_domain_re = re.compile(r"\.?[^.]*")
    domain_re = re.compile(r"[^.]*")
    dots_re = re.compile(r"^\.+")

    magic_re = re.compile(r"^\#LWP-Cookies-(\d+\.\d+)", re.ASCII)

    def __init__(self, policy=None):
    ....
```

通读了下 大概知道这个模块的作用是管理HTTP的cookie值，存储HTTP请求生成的cookie，向传出的HTTP请求添加cookie的对象


cookies初始化方法cookiejar_from_dict(cookie_dict, cookiejar=None, overwrite=True)的作用是将字典类型的cookies插入到cookiejar中，返回cookiejar。

整个cookie都存储在内存中，CookieJar实例销毁后cookie也将丢失。