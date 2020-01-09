# Requests 源码阅读-Day1


[toc]


## Requests介绍

requests是一个Python的网络请求库，和urllib、httplib之流相比起来最大的优点就是好用，requests官方标榜的就是“我们的库是给人用的哦”。

此外，requests还支持https验证并且是线程安全的

[官方文档](https://requests.readthedocs.io/en/master/)



## 分析环境

Requests版本： V2.20.0

python版本: 3.6.7

[git下载地址](https://github.com/psf/requests)



## 源码分析


首先git clone源码到本地，然后使用**cloc工具**查看文件格式

>如果系统没有cloc工具，可以进行安装。
>由于我本地是fedora系统，可以使用dnf install cloc命令进行安装


使用下面命令查看项目关于python文件的统计情况：
```
[jian@laptop requests]$ cloc --include-lang=Python .
```

| Language | files | blank | comment | code  |
| :------- | ----: | :---: | :-----: | :---: |
| Python   |    35 | 1951  |  1990   | 5937  |

<br>

**`可以看到总共python文件有35个，代码量是5937行`**

### 单元测试

```python
[jian@laptop requests]$ pwd
/home/jian/prj/github/requests

[jian@laptop requests]$ python

Python 3.6.7 (default, Mar 21 2019, 20:23:57) 
[GCC 8.3.1 20190223 (Red Hat 8.3.1-2)] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> import requests
>>> r = requests.get("http://www.baidu.com")
>>> r.status_code
200
```

<br>
打开README.md文件看看requests当前都支持哪些功能

```
International Domains and URLs      #国际化域名和URLS
Keep-Alive & Connection Pooling		#keep—Alive&连接池
Sessions with Cookie Persistence    #持久性cookie的会话
Browser-style SSL Verification		#浏览器式SSL认证
Basic & Digest Authentication       #基本/摘要认证
Familiar  dict–like Cookies			#key/value cookies
Automatic Decompression of Content  #自动内容解压缩
Automatic Content Decoding			#自动内容解码
Automatic Connection Pooling        #自动连接池
Unicode Response Bodies				#Unicode响应体
Multi-part File Uploads             #文件分块上传
SOCKS Proxy Support					#HTTP(S)代理支持
Connection Timeouts                 #连接超时
Streaming Downloads					#数据流下载
Automatic honoring of .netrc        #netrc支持
Chunked HTTP Requests				#Chunked请求

```

<br>
看源码看的是思想，要明白作者的设计思路到底是什么。

**`比如requests，看完了你应当问问自己，cookie为什么要封装而不是直接用？request为什么要有两个形态？设计session是为了解决什么问题？`**

只有理解了设计思路，再去看具体的细节实现才能有收获，否则你看到的就是满屏的raise、isinstanceof，这样去看代码恐怕是浪费时间了。

so...开始开干吧


#### test_requests.py

源码目录下有一个tests文件夹，这里面以test开头的测试文件是专门用于测试requests接口，使用的是pytest方式，pytest我会单独写一章节介绍具体的内容


首先选第一个方法分析，找到`test_DIGEST_HTTP_200_OK_GET`方法

tests/test_requests.py

```python
    def test_DIGEST_HTTP_200_OK_GET(self, httpbin):

        for authtype in self.digest_auth_algo:
            auth = HTTPDigestAuth('user', 'pass')
            url = httpbin('digest-auth', 'auth', 'user', 'pass', authtype, 'never')

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


上面这段代码主要做了下面的工作：

>1.传递一个httpbin参数，这个`httpbin`是什么东西？

>2.遍历不同的摘要认证算法，`self.digest_auth_algo`是什么？

>3.摘要认证变量auth及url变量设置, `HTTPDigestAuth`是什么？

>4.对url发起get请求，200表示请求成功，401表示未经授权。

>这个测试是为了验证auth的必要性

>5.新建了一个会话对象s，同时也设置了auth变量，跟前面不同的是这个请求是由会话对象s发起的

<br>

##### digest_auth_algo

首先回答前面说的self.digest_auto_algo的问题：

tests/test_requests.py
```python
class TestRequests:

    digest_auth_algo = ('MD5', 'SHA-256', 'SHA-512')
    ....
```

摘要访问认证，它是一种协议规定的web服务器用来同网页浏览器进行认证信息协商的方法。

浏览器在向服务器发送请求的过程中需要传递认证信息auth

auth经过摘要算法加密形成秘文，最后发送给服务器。

服务器验证成功后返回“200”告知浏览器可以继续访问

若验证失败则返回"401"告诉浏览器禁止访问。

当前该摘要算法分别选用了"MD5","SHA-256","SHA-512"。
<br>

##### HTTPDigestAuth

tests/test_requests.py 
```python
from requests.auth import HTTPDigestAuth, _basic_auth_str
```

看到是导入requests.auth模块里面的HTTPDigestAuth方法
好 我们去查看下这个东西是什么？

requests/auth.py
```python
class HTTPDigestAuth(AuthBase):
    """Attaches HTTP Digest Authentication to the given Request object."""

    def __init__(self, username, password):
        self.username = username
        self.password = password
        # Keep state in per-thread local storage
        self._thread_local = threading.local()

    ....
    def __call__(self, r):
        # Initialize per-thread state, if needed
        self.init_per_thread_state()
        # If we have a saved nonce, skip the 401
        if self._thread_local.last_nonce:
            r.headers['Authorization'] = self.build_digest_header(r.method, r.url)
        try:
            self._thread_local.pos = r.body.tell()
        except AttributeError:
            # In the case of HTTPDigestAuth being reused and the body of
            # the previous request was a file-like object, pos has the
            # file position of the previous body. Ensure it's set to
            # None.
            self._thread_local.pos = None
        r.register_hook('response', self.handle_401)
        r.register_hook('response', self.handle_redirect)
        self._thread_local.num_401_calls = 1

        return r

	...

```

<br>

HTTPDigestAuth：为http请求对象提供摘要认证。

实例化对象auth时需要传入认证所需的username及password。

**threading.local(**)在这里的作用是保存一个全局变量，但是这个全局变量只能在当前线程才能访问，每一个线程都有单独的内存空间来保存这个变量，它们在逻辑上是隔离的，其他线程都无法访问。

我们可以通过实例演示一下摘要认证：

```python
[jian@laptop requests]$ python
Python 3.6.7 (default, Mar 21 2019, 20:23:57) 
[GCC 8.3.1 20190223 (Red Hat 8.3.1-2)] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> import requests
>>> from requests.auth import HTTPDigestAuth
>>> r = requests.get('http://httpbin.org/digest-auth/auth/user/pass',auth=HTTPDigestAuth
... ('user','pass'))

>>> r.status_code
200
```

<br>

##### httpbin

终于要解决这个东西了，这个东西是啥呢？
<br>

1.设置断点：
tests/conftest.py
```
@pytest.fixture
def httpbin(httpbin):
    pytest.set_trace() # 设置断点
    return prepare_url(httpbin)
```
<br>

在tests目录新创建个文件test_xx.py
tests/test_xx.py
```python
from requests.auth import HTTPDigestAuth
import requests
import pytest


class TestRequests:

    digest_auth_algo = ('MD5', 'SHA-256', 'SHA-512')

    def test_DIGEST_HTTP_200_OK_GET(self, httpbin):

        for authtype in self.digest_auth_algo:
            auth = HTTPDigestAuth('user', 'pass')
            url = httpbin('digest-auth', 'auth', 'user', 'pass', authtype,
                          'never')
            pytest.set_trace() # 设置断点

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
<br>

然后在终端执行

```
[jian@laptop requests]$ pwd
/home/jian/prj/github/requests
[jian@laptop requests]$ pytest  tests/test_xx.py 
xxx
  @pytest.fixture
  def httpbin(httpbin):
E       fixture 'httpbin' not found
```

怎么报错呢？ -说httpbin 这个fixture没有找到
<br>

查资料是说缺少`pytest-httpbin`模块

pip安装起来：
```
pip install pytest-httpbin 
```
<br>

然后再次执行 `pytest  tests/test_xx.py `命令

```
[jian@laptop requests]$ pytest  tests/test_xx.py 
Test session starts (platform: linux, Python 3.6.7, pytest 3.2.1, pytest-sugar 0.9.2)
rootdir: /home/jian/prj/github/requests, inifile: pytest.ini
plugins: sugar-0.9.2, pep8-1.0.6, mock-1.6.2, httpbin-1.0.0, flakes-2.0.0, env-0.6.2, cov-2.5.1, assume-2.2.0, allure-adaptor-1.7.10, celery-4.4.0

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PDB set_trace (IO-capturing turned off) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
> /home/jian/prj/github/requests/tests/conftest.py(20)httpbin()
-> return prepare_url(httpbin)
(Pdb) httpbin
<pytest_httpbin.serve.Server object at 0x7f5e78e44438>
(Pdb) c

>>>>>>>>>>>>>>>>>>> PDB set_trace (IO-capturing turned off) >>>>>>>>>>>>>>>>>>>>
> /home/jian/prj/github/requests/tests/test_xx.py(18)test_DIGEST_HTTP_200_OK_GET()
-> r = requests.get(url, auth=auth)
(Pdb) httpbin
<function prepare_url.<locals>.inner at 0x7f5e78ea3c80>
(Pdb) url
'http://127.0.0.1:44131/digest-auth/auth/user/pass/MD5/never'
(Pdb) 

```

<br>
在调试窗口PDB set_trace中可以看到，首先被调用的是的conftest.py中的httpbin()方法

我们在（pdb）中输入httpbin变量，结果
返回了<pytest_httpbin.serve.Server object at 0x7f5e78e44438>

然后继续调用方法test_DIGEST_HTTP_200_OK_GET()，输入httpbin变量，结果返回了<function prepare_url.<locals>.inner at 0x7f5e78ea3c80>

经过调试后，httpbin的面貌渐渐变得清晰了


test_DIGEST_HTTP_200_OK_GET(self, httpbin)中的httpbin对象为<function prepare_url.<locals>.inner at 0x7f5e78ea3c80>


也就是源码中prepare_url(value)方法里的inner(*suffix)方法。
<br>

也就是这个文件:
tests/conftest.py
```
def prepare_url(value):
    # Issue #1483: Make sure the URL always has a trailing slash
    httpbin_url = value.url.rstrip('/') + '/'

    def inner(*suffix):
        return urljoin(httpbin_url, '/'.join(suffix))

    return inner
```



这里使用了函数`闭包`，有什么作用？ -`保持程序上一次运行后的状态然后继续执行`
<br>

httpbin(httpbin)方法中参数httpbin对象为<pytest_httpbin.serve.Server object at 0x7f5e78e44438>

pytest_httpbin是pytest的一个插件，那肯定跟pytest调用有关系了

然后Server是什么东东？我们来查看下它的源码：
<br>

pytest_httpbin/serve.py
```python
class Server(object):
    """
    HTTP server running a WSGI application in its own thread.
    """

    port_envvar = 'HTTPBIN_HTTP_PORT'

    def __init__(self, host='127.0.0.1', port=0, application=None, **kwargs):
        self.app = application
        if self.port_envvar in os.environ:
            port = int(os.environ[self.port_envvar])
        self._server = make_server(
            host,
            port,
            self.app,
            handler_class=Handler,
            **kwargs
        )
```

<br>

原来这是一个本地的WSGI服务器，专门用于pytest进行网络测试，这样的好处在于我们无需连接外部网络环境，在本地就能实现一系列的网络测试工作。

WSGI全称是Web Server Gateway Interface,它其实是一个标准，介于web应用与web服务器之间。

只要我们遵循WSGI接口标准设计web应用，就无需在意TCP连接，HTTP请求等等底层的实现，全权交由web服务器即可。

上述代码实现的逻辑已经比较清晰了，httpbin对象被实例化的时候调用__init__(self, host='127.0.0.1',port=0, application=None, **kwargs)。

提到fixture方法httpbin(httpbin)中的参数httpbin是一个Server对象，但是这个对象是在什么时候创建的？原来这个httpbin也是一个fixture方法，存在于pytest-httpbin插件中。

pytest-httpbin/plugin.py
```python
from __future__ import absolute_import
import pytest
from httpbin import app as httpbin_app
from . import serve, certs

@pytest.fixture(scope='session')
def httpbin(request):
    server = serve.Server(application=httpbin_app)
    server.start()
    request.addfinalizer(server.stop)
    return server
   
```
<br>

这是一个"session"级别的fixture方法，首先实例化Server对象为server，传入application参数"httpbin_app"，application参数我们在前面提到过，它指向我们的web应用程序。

这里的httpbin_app是pytest-httpbin下app模块的别称，该模块是专门用于http测试而编写的web应用程序，这里就不扩展了。

然后server继续调用start()方法，启动线程，开启WSGI服务器，最后返回server。


总结：

![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/httpbin.png)


