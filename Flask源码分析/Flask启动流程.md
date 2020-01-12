# 第1讲-Flask 启动流程

<!-- TOC -->

- [第1讲-Flask 启动流程](#%e7%ac%ac1%e8%ae%b2-flask-%e5%90%af%e5%8a%a8%e6%b5%81%e7%a8%8b)
  - [启动流程](#%e5%90%af%e5%8a%a8%e6%b5%81%e7%a8%8b)
  - [app初始化](#app%e5%88%9d%e5%a7%8b%e5%8c%96)
    - [问题-__name__](#%e9%97%ae%e9%a2%98-name)
    - [问题-_PackageBoundObject](#%e9%97%ae%e9%a2%98-packageboundobject)
    - [问题-_PackageBoundObject作用](#%e9%97%ae%e9%a2%98-packageboundobject%e4%bd%9c%e7%94%a8)
      - [import_name](#importname)
      - [root_path](#rootpath)
      - [template_folder](#templatefolder)
    - [其他问题](#%e5%85%b6%e4%bb%96%e9%97%ae%e9%a2%98)
      - [static_url_path和static_folder](#staticurlpath%e5%92%8cstaticfolder)
      - [host_matching和static_host](#hostmatching%e5%92%8cstatichost)
      - [subdomain_matching](#subdomainmatching)
      - [instance_path和instance_relative_config](#instancepath%e5%92%8cinstancerelativeconfig)
  - [路由](#%e8%b7%af%e7%94%b1)
  - [run](#run)

<!-- /TOC -->

## 启动流程

从例子出发

```python
from flask import Flask

app = Flask(__name__)


@app.route('/')
def index():
    return 'hello world!'


if __name__ == '__main__':
    app.run()

```


## app初始化
首先通过 `Flask(__name__)`实例化了 Flask 类，对应的`__init__() `方法如下

```python
class Flask(_PackageBoundObject):
...
def __init__(
        self,
        import_name,
        static_url_path=None,
        static_folder="static",
        static_host=None,
        host_matching=False,
        subdomain_matching=False,
        template_folder="templates",
        instance_path=None,
        instance_relative_config=False,
        root_path=None,
    ):
    #对每次请求，创建一个处理通道。
    self.config = self.make_config()
    self.view_functions = {}
    self.error_handler_spec = {}
    self.before_request_funcs = {}
    self.before_first_request_funcs = []
    self.after_request_funcs = {}
    self.teardown_request_funcs = {}
    self.teardown_appcontext_funcs = []
    self.url_value_preprocessors = {}
    self.url_default_functions = {}
    self.url_map = Map()
    self.blueprints = {}
    self._blueprint_order = []
    self.extensions = {}
    ...
```
`__init__() `方法中有大量的注释，注释中解释了这些变量的用途，但仅从变量名就知道，它们用于存储每次请求对应的信息，相当于一个处理通道。

### 问题-`__name__`
为什要使用`Flask(__name__)`实例化了 Flask 类， 这里的`__name__`是什么意思？

学过python的马上会联想到这个：

`if __name__ == "__main__"`

我们知道：
>1.python文件的后缀为.py
>
>2..py文件既可以用来直接执行，就像一个小程序一样，也可以用来作为模块被导入
>
>3.在python中导入模块一般使用的是import


so...开始解释`if __name__ == "__main__"`:


首先解释一下if，顾名思义，if就是如果的意思，在句子开始处加上if，就说明，这个句子是一个条件语句。

接着是 `__name__`，`__name__`作为模块的内置属性，简单点说呢，就是.py文件的调用方式。


看一个例子： demo.py
```python
print("Doing...")
print("__name__=", __name__)

def hello():
    print("Hello")



if __name__ == "__main__":
    print("__name__=", __name__)
    hello()

```

1.首先执行导入模块的方式：
```
[jian@laptop tmp]$ python
Python 3.6.7 (default, Mar 21 2019, 20:23:57) 
[GCC 8.3.1 20190223 (Red Hat 8.3.1-2)] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> import demo
Doing...
__name__= demo
```

看了`__name__` 它的值是demo， 也就是文件名


2.使用脚本运行的方式：
```
[jian@laptop tmp]$ python demo.py 
Doing...
__name__= __main__
__name__= __main__
Hello

```
看了`__name__` 它的值是`__main__`， 也就是文件名

3.这里的`__main__`是什么呢？

可以查看[官方文档](https://docs.python.org/3/library/__main__.html)

简单说就是**`顶层代码执行的作用域的名称`**


4.再看一个例子
```
[jian@laptop tmp]$ tree a/
a/
├── b
│   ├── c.py
│   └── __init__.py
└── __init__.py

1 directory, 3 files

[jian@laptop tmp]$ cat a/__init__.py 
print("a init name=", __name__)
[jian@laptop tmp]$ cat a/b/__init__.py 
print("b init name", __name__)
[jian@laptop tmp]$ cat a/b/c.py 
print("This is c")
print("c init name", __name__)
```

当一个.py文件（模块）被其他.py文件（模块）导入时，我们在命令行执行
```
[jian@laptop tmp]$ python -c "import a.b.c"
a init name= a
b init name a.b
This is c
c init name a.b.c

```

由此可见，`__name__`**还可以清晰地反映一个模块在包中的层次**。

再回到上面说的内容，最后是`__main__`，.py文件有两种使用方式：作为模块被调用和直接使用。

如果它等于"`__main__`"就表示是直接执行。


总结：

```
在if __name__ == "__main__"：
之后的语句作为模块被调用的时候，语句之后的代码不执行；

直接使用的时候，语句之后的代码执行。
通常，此语句用于模块测试中使用。
```



### 问题-_PackageBoundObject

```
class Flask(_PackageBoundObject):
```

Flask继承了_PackageBoundObject， 这个东西是什么？

定位到这个类看到：
```python
class _PackageBoundObject(object):
    #: The name of the package or module that this app belongs to. Do not
    #: change this once it is set by the constructor.
    import_name = None

    #: Location of the template files to be added to the template lookup.
    #: ``None`` if templates should not be added.
    template_folder = None

    #: Absolute path to the package on the filesystem. Used to look up
    #: resources contained in the package.
    root_path = None

    def __init__(self, import_name, template_folder=None, root_path=None):
        self.import_name = import_name
        self.template_folder = template_folder
    ...
```

再回到flask 这边往下看代码：

```
class Flask(_PackageBoundObject):
...
        _PackageBoundObject.__init__(
            self, import_name, template_folder=template_folder, root_path=root_path
        )

        self.static_url_path = static_url_path
        self.static_folder = static_folder
```

看到这个就知道了 其实`_PackageBoundObject`类也只是初始化import_name和template_folder这个属性，方便文件路径处理。


### 问题-`_PackageBoundObject`作用

研读里面的源码发现:

```python
class _PackageBoundObject(object):
    #: The name of the package or module that this app belongs to. Do not
    #: change this once it is set by the constructor.
    import_name = None

    #: Location of the template files to be added to the template lookup.
    #: ``None`` if templates should not be added.
    template_folder = None

    #: Absolute path to the package on the filesystem. Used to look up
    #: resources contained in the package.
    root_path = None
    
	def __init__(self, import_name, template_folder=None, root_path=None):
        self.import_name = import_name
        self.template_folder = template_folder

        if root_path is None:
            root_path = get_root_path(self.import_name)

        self.root_path = root_path
        self._static_folder = None
        self._static_url_path = None

```

这里面有三个实例：
>import_name

>template_folder 

>root_path

下面分别来介绍这三个东西是什么，以及分别起了什么作用?

#### import_name 
这里就是属于flask app包名或是模块名


#### root_path 

看源码知道，root_path 属性的值是使用import_name 属性作为参数,
调用get_root_path方法得到的.

定位get_root_path 方法， 一行行分析：
```python
def get_root_path(import_name):
    # Module already imported and has a file attribute.  Use that first.
    mod = sys.modules.get(import_name)
    if mod is not None and hasattr(mod, "__file__"):
        return os.path.dirname(os.path.abspath(mod.__file__))
        ....
```

首先这句话是什么意思？ 来看一个例子：

```python
[jian@laptop tmp]$ pwd
/tmp
[jian@laptop tmp]$ cat test.py 
import sys
import os


print("__name__=", __name__)

mod = sys.modules.get(__name__)
mod_dict = sys.modules.get(__name__).__dict__

print("mod=",mod)
print("mod dict=", mod_dict)


is_have_file = hasattr(sys.modules.get(__name__), '__file__')
print("is_have_file=", is_have_file)

filename = mod_dict.get("__file__")
print("filename=", filename)


filepath = os.path.abspath(mod.__file__)
print("filepath=", filepath)


filedirname = os.path.dirname(filename)
print("filedirname=", filedirname)

```


运行结果：
```
__name__= __main__
mod= <module '__main__' from 'test.py'>
mod dict= {'__name__': '__main__', '__doc__': None, '__package__': None, '__loader__': <_frozen_importlib_external.SourceFileLoader object at 0x7ff7895b3320>, '__spec__': None, '__annotations__': {}, '__builtins__': <module 'builtins' (built-in)>, '__file__': 'test.py', '__cached__': None, 'sys': <module 'sys' (built-in)>, 'os': <module 'os' from '/home/jian/.pyenv/versions/3.6.7/lib/python3.6/os.py'>, 'mod': <module '__main__' from 'test.py'>, 'mod_dict': {...}}
is_have_file= True
filename= test.py
filepath= /tmp/test.py
filedirname= 

```

通过结果看到：
> get_root_path函数用于根据模块名得到模块路径。

> 如果模块已经加载过，也即通过`sys.modules`列表能找到该模块名，并且该模块有`__file__`属性，则直接返回`__file__`即可得到模块路径


再继续往下研究：
```python
    # Next attempt: check the loader.
    loader = pkgutil.get_loader(import_name)

    # current working directory.
    if loader is None or import_name == "__main__":
        return os.getcwd()

if hasattr(loader, "get_filename"):
        filepath = loader.get_filename(import_name)
    else:
        # Fall back to imports.
        __import__(import_name)
        mod = sys.modules[import_name]
        filepath = getattr(mod, "__file__", None)

        if filepath is None:
            raise RuntimeError(
                "No root path can be found for the provided "
                'module "%s".  This can happen because the '
                "module came from an import hook that does "
                "not provide file name information or because "
                "it's a namespace package.  In this case "
                "the root path needs to be explicitly "
                "provided." % import_name
            )
    return os.path.dirname(os.path.abspath(filepath))
```

这句话又是什么意思呢？ 看一个例子：
```python
[jian@laptop tmp]$ cat test.py 
import pkgutil
import sys


print("__name__=", __name__)


loader = pkgutil.get_loader(__name__)
print("loader=", loader)
print("loader dict=", loader.__dict__)

import_name = __name__
__import__(import_name)
mod = sys.modules[import_name]
filepath = getattr(mod, "__file__", None)

print("filepath=", filepath)

```

使用脚本方式调用：
```
[jian@laptop tmp]$ python test.py 
__name__= __main__
loader= <_frozen_importlib_external.SourceFileLoader object at 0x7fa9b15f0320>
loader dict= {'name': '__main__', 'path': 'test.py'}
filepath= test.py
```

使用模块方式调用：
```
[jian@laptop tmp]$ python -c "import test"
__name__= test
loader= <_frozen_importlib_external.SourceFileLoader object at 0x7f667f040128>
loader dict= {'name': 'test', 'path': '/tmp/test.py'}
filepath= /tmp/test.py
```


通过例子我们知道先通过模块名找到模块加载器，如果加载器为空或者模块是通过执行python模块的方式来访问模块，则直接返回当前进程工作路径作为模块路径

如果加载器方法get_filename，则调用该方法获取模块名；否则导入该模块，并且通过sys.modules获取模块及其属性`__file__`，如果`__file__`属性值未空，则直接抛出异常， 最后根据前面确定的文件路径确定根目录


#### template_folder 
从字面意思就知道 它是模板文件夹， 默认文件夹templates

看一个例子：

```
[jian@laptop demo]$ tree .
.
├── app.py
├── static
│   ├── css
│   ├── images
│   └── js
│       └── main.js
└── templates
    └── index.html

5 directories, 3 files

```

```python
[jian@laptop demo]$ cat app.py 
from flask import Flask, render_template

# 可以不写static_folder,默认就是static目录
#app = Flask(__name__, static_folder='static')
app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

if __name__ == "__main__":
    app.run(debug=True)

```

```
[jian@laptop demo]$ cat templates/index.html 
<html>
	<head>
	<!--使用这种方式没有生效?
	 <script type="text/javascript" src="../static/js/main.js></script>
	-->
	<script type="text/javascript" src="{{url_for('static', filename='js/main.js')}}"> </script>
	</head>

	<body>
		<h1>I am here , hello world </h1>
	</body>
</html>
```

```
[jian@laptop demo]$ cat static/js/main.js 
window.onload = function() {
	alert("Hello");
}

```

服务开启：
```
[jian@laptop demo]$ python app.py 
 * Serving Flask app "app" (lazy loading)
 * Environment: production
   WARNING: This is a development server. Do not use it in a production deployment.
   Use a production WSGI server instead.
 * Debug mode: on
 * Running on http://127.0.0.1:5000/ (Press CTRL+C to quit)
 * Restarting with stat
 * Debugger is active!
 * Debugger PIN: 229-100-610

```

使用浏览器访问http://localhost:5000 查看结果



### 其他问题

#### static_url_path和static_folder
static_url_path 和static_folder 区别，这两个有什么区别？

```
class Flask(_PackageBoundObject):
...
    def __init__(
        self,
        import_name,
        static_url_path=None,
        static_folder="static",
        static_host=None,
        host_matching=False,
        subdomain_matching=False,
        template_folder="templates",
        instance_path=None,
        instance_relative_config=False,
        root_path=None,
    ):
```

修改前面的例子app.py
```python
from flask import Flask, render_template

# 可以不写static_folder,默认就是static目录
#app = Flask(__name__, static_folder='static')
app = Flask(__name__, static_url_path="/stat")



@app.route('/')
def index():
    return render_template('index.html')



if __name__ == "__main__":
    app.run(debug=True)

```

然后运行服务，分别进行访问下面两个文件

```
http://localhost:5000/stat/js/main.js
http://localhost:5000/static/js/main.js
```

查看日志：
```
127.0.0.1 - - [12/Jan/2020 15:35:16] "GET /stat/js/main.js HTTP/1.1" 200 -
127.0.0.1 - - [12/Jan/2020 15:36:30] "GET /stat/js/main.js HTTP/1.1" 304 -

127.0.0.1 - - [12/Jan/2020 15:34:59] "GET /static/js/main.js HTTP/1.1" 404 -

```

对于"/stat/js/main.js" 进行几次访问，发现一开始是200， 然后就是304， 然后又是200
这里的304的意思是， 客户端有缓存情况下，服务端这边给的一种响应

>第一次访问 200

>按F5刷新（第二次访问） 304

>按Ctrl+F5强制刷新 200


从上面的代码中可以看到：

**`static_url_path和static_folder 同时存在时`**

**`static_url_path代替static_folder 指明了静态资源的路径`**


#### host_matching和static_host
host_matching和static_host区别？

修改前面的例子：
```
from flask import Flask, render_template

# host_matching 和 static_host 组合更改静态资源的访问地址(主机:端口)
# 结合 static_url_path 指定文件
app = Flask(__name__, host_matching=True, static_host="localhost:8888",static_url_path="/stat")

@app.route('/', host="localhost:8888")
def index():
    print("url_map=", app.url_map)
    print("host_matching=", app.url_map.host_matching)
    return render_template('index.html')

if __name__ == "__main__":
    app.run(debug=True, port=8888)

```

然后直接浏览器访问http://localhost:8888 就可以访问index.html

服务器日志：
```
url_map= Map([<Rule 'localhost:8888|/' (OPTIONS, HEAD, GET) -> index>,
 <Rule 'localhost:8888|/stat/<filename>' (OPTIONS, HEAD, GET) -> static>])
host_matching= True

```

#### subdomain_matching

在匹配路由时，考虑相对于server_name的子域。默认为false

本地增加域名为后面做测试：
```
[root@laptop ~]# cat /etc/hosts
127.0.0.1 dev.com
127.0.0.1 test.dev.com
```

修改前面的例子: app.py
```python
from flask import Flask, render_template


app = Flask(__name__)
app.config["SERVER_NAME"] = "dev.com:8888"


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/', subdomain='test')
def demo_home():
    return "This is test"



if __name__ == "__main__":
    app.run(debug=True)

```

服务开启，浏览器分别访问下面两个地址：

```
http://dev.com:8888
http://test.dev.com:8888
```

服务器日志：
```
http://dev.com:8888/
Map([<Rule 'test|/' (GET, HEAD, OPTIONS) -> demo_home>,
 <Rule '/' (GET, HEAD, OPTIONS) -> index>,
 <Rule '/static/<filename>' (GET, HEAD, OPTIONS) -> static>])
127.0.0.1 - - [12/Jan/2020 16:12:35] "GET / HTTP/1.1" 200 -

http://test.dev.com:8888/
Map([<Rule 'test|/' (GET, HEAD, OPTIONS) -> demo_home>,
 <Rule '/' (GET, HEAD, OPTIONS) -> index>,
 <Rule '/static/<filename>' (GET, HEAD, OPTIONS) -> static>])
127.0.0.1 - - [12/Jan/2020 16:12:46] "GET / HTTP/1.1" 200 -

```


增加subdomain_matching配置

修改前面的例子, app.py
```python
[jian@laptop demo]$ cat app.py 
from flask import Flask, render_template, request


app = Flask(__name__,subdomain_matching=True)
app.config["SERVER_NAME"] = "dev.com:8888"


@app.route('/')
def index():
    print(request.url)
    print(app.url_map)
    return render_template('index.html')


@app.route('/', subdomain='test')
def demo_home():
    print(request.url)
    print(app.url_map)
    return "This is test"



if __name__ == "__main__":
    app.run(debug=True)

```

再进行测试：
```
http://dev.com:8888/
Map([<Rule 'test|/' (OPTIONS, HEAD, GET) -> demo_home>,
 <Rule '/' (OPTIONS, HEAD, GET) -> index>,
 <Rule '/static/<filename>' (OPTIONS, HEAD, GET) -> static>])
127.0.0.1 - - [12/Jan/2020 16:15:02] "GET / HTTP/1.1" 200 -

http://test.dev.com:8888/
Map([<Rule 'test|/' (OPTIONS, HEAD, GET) -> demo_home>,
 <Rule '/' (OPTIONS, HEAD, GET) -> index>,
 <Rule '/static/<filename>' (OPTIONS, HEAD, GET) -> static>])

```

**`实验表明加不加subdomain_matching 都是一样的结果？？？`**


#### instance_path和instance_relative_config

 instance_relative_config 和 instance_path
两者配合从版本控制外加载配置信息


看一个例子：
```python
[jian@laptop demo]$ cat app.py 
from flask import Flask, render_template

app = Flask(__name__, instance_path="/tmp", instance_relative_config=True)
app.config.from_pyfile("config.py")

@app.route('/')
def index():
    print(app.instance_path)
    return render_template('index.html')

if __name__ == "__main__":
    app.run()

```


```python
[jian@laptop tmp]$ cat /tmp/config.py 
DEBUG = True
```

服务开启：
```
[jian@laptop demo]$ python app.py 
 * Serving Flask app "app" (lazy loading)
 * Environment: production
   WARNING: This is a development server. Do not use it in a production deployment.
   Use a production WSGI server instead.
 * Debug mode: on
 * Running on http://127.0.0.1:5000/ (Press CTRL+C to quit)
 * Restarting with stat
 * Debugger is active!
 * Debugger PIN: 229-100-610

```


默认是没有开启debug， 但是载入了其他配置文件，就变成了debug=True



## 路由
Flask 实例化后，接着利用 @app.route('/')装饰器的方式将 hello () 方法映射成了路由

找到router这个方法：

```python
    def route(self, rule, **options):
      def decorator(f):
            endpoint = options.pop("endpoint", None)
            self.add_url_rule(rule, endpoint, f, **options)
            return f
        return decorator
```

可以发现 route () 方法就是一个简单的装饰器，**具体处理逻辑在 `addurlrule ()` 方法中**。

定位到**addurlrule**这个方法中查看：

```python
@setupmethod
def add_url_rule(self, rule, endpoint=None, view_func=None,
    provide_automatic_options=None, **options):
    # ... 省略其他代码细节
    rule = self.url_rule_class(rule, methods=methods, **options)
    rule.provide_automatic_options = provide_automatic_options
    # 将路由rule添加到url_map中
    self.url_map.add(rule)
    if view_func is not None:
        old_func = self.view_functions.get(endpoint)
    # 每个方法的endpoint必须不同
    if old_func is not None and old_func != view_func:
        raise AssertionError('View function mapping is overwriting an '
        'existing endpoint function: %s' % endpoint)
        # 将rule对应的endpoint与view_func通过view_functions字典对应上
    self.view_functions[endpoint] = view_func

```

从 addurlrule () 方法可以看出， @app.route('/')的主要作用就是将路由保存到 urlmap 中，将装饰的方法保存到 viewfunctions 中。

需要注意的是，每个方法的 endpoint 必须不同，否则会抛出 AssertionError。


## run
最后调用了 app.run () 方法运行 Flask 应用，对应代码如下。

```
  def run(self, host=None, port=None, debug=None,
    load_dotenv=True, **options):
    # ... 省略
    from werkzeug.serving import run_simple
    try:
    # 利用run_simple方法启动web服务
        run_simple(host, port, self, **options)
    finally:
        self._got_first_request = False

```

run () 方法进一步调用 werkzeug.serving 下的 run_simple () 方法启动 web 服务，其中 self 就是 Flask () 的 application。

继续往下面查看：
```python
def run_simple(....):
....

 def inner():
        try:
            fd = int(os.environ["WERKZEUG_SERVER_FD"])
        except (LookupError, ValueError):
            fd = None
        srv = make_server(
            hostname,
            port,
            application,
            threaded,
            processes,
            request_handler,
            passthrough_errors,
            ssl_context,
            fd=fd,
        )
        ...

```

看到run_simple其实是调用了make_server方法，进入make_server方法继续往下查看：
```python
def make_server(...):
...
if threaded and processes > 1:
        raise ValueError("cannot have a multithreaded and multi process server.")
    elif threaded:
        return ThreadedWSGIServer(
            host, port, app, request_handler, passthrough_errors, ssl_context, fd=fd
        )
    elif processes > 1:
        return ForkingWSGIServer(
            host,
            port,
            app,
            processes,
            request_handler,
            passthrough_errors,
            ssl_context,
            fd=fd,
        )
    else:
        return BaseWSGIServer(
            host, port, app, request_handler, passthrough_errors, ssl_context, fd=fd
        )
```


看到make_server其实是调用了BaseWSGIServer方法，进入BaseWSGIServer方法继续往下查看：

```python
class BaseWSGIServer(HTTPServer, object):

    """Simple single-threaded, single-process WSGI server."""

    multithread = False
    multiprocess = False
    request_queue_size = LISTEN_QUEUE

    def __init__(
        self,
        host,
        port,
        app,
        handler=None,
        passthrough_errors=False,
        ssl_context=None,
        fd=None,
    ):
	if handler is None:
       handler = WSGIRequestHandler
    ....

```

总结：
> runsimple() -> makeserver() -> BaseWSGIServer() -> WSGIRequestHandler


继续往下研究：

WSGIRequestHandler 类从名称就可以知，它主要用于处理满足 WSGI 协议的请求，该类中的 execute () 方法部分代码如下。

```python

        def execute(app):
            application_iter = app(environ, start_response)
            try:
                for data in application_iter:
                    write(data)
                if not headers_sent:
                    write(b"")
            finally:
                if hasattr(application_iter, "close"):
                    application_iter.close()
                application_iter = None
```

简单而言，app.run () 会启动一个满足 WSGI 协议的 web 服务，它会监听指定的端口，将 HTTP 请求解析为 WSGI 格式的数据，然后将 environ, start_response 传递给 Flask () 实例对象。

类对象作为方法被调用，需要看到`__call__() `方法

```python
 
def __call__(self, environ, start_response):
    return self.wsgi_app(environ, start_response)

def wsgi_app(self, environ, start_response):
    # 请求上下文
    ctx = self.request_context(environ)
    error = None
    try:
        try:
            ctx.push()
            # 正确的请求处理路径，会通过路由找到对应的处理函数
            # 通过路由找到相应的处理函数
            response = self.full_dispatch_request()
        except Exception as e:
            # 错误处理
            error = e
            response = self.handle_exception(e)
        except:
            error = sys.exc_info()[1]
            raise
            return response(environ, start_response)
    finally:
        if self.should_ignore_error(error):
            error = None
            # 无论是否发生异常都需要从请求上下文中error pop出
            ctx.auto_pop(error)
```

主要逻辑在 wsgiapp () 方法中，一开始进行了请求上下文的处理 

随后通过 full_dispatch_request() 方法找到当前请求路由对应的方法，调用该方法，获得返回，如果请求路由不存在，则进行错误处理，返回 500 错误。

full_dispatch_request () 方法代码如下。

```python
def full_dispatch_request(self):
    self.try_trigger_before_first_request_functions()
    try:
        request_started.send(self)
        rv = self.preprocess_request()
        if rv is None:
        # 调用该路由对应的处理函数并将处理函数的结果返回。
            rv = self.dispatch_request()
    except Exception as e:
        rv = self.handle_user_exception(e)
    return self.finalize_request(rv)
```

full_dispatch_request () 方法中最关键的逻辑在于 dispatch_request () 方法，该方法会将调用对应路由的处理函数并获得该函数的结果。


此外还有 

```
try_trigger_before_first_request_functions()、
preprocess_request()、
finalize_request () 方法
...
```

这些方法会执行我们自定义的各种钩子函数(用什么用处？)，这些钩子函数会存储在 before_request_funcs()、before_first_requestfuncs()、after_request_funcs () 方法中。


最后，需要注意:

```
app.run () 方法仅在开发环境中会被使用，通过上面的分析，已经知道 app.run () 背后就是使用 werkzeug 构建了一个简单的 web 服务，

但这个 web 服务并不牢靠，生产环境通常利用 gunicron，通过配置的形式，指定 WSGI app 所在文件来启动 Flask 应用。
```