# pytest 学习笔记

<!-- TOC -->

- [pytest 学习笔记](#pytest-%e5%ad%a6%e4%b9%a0%e7%ac%94%e8%ae%b0)
  - [pytest是什么](#pytest%e6%98%af%e4%bb%80%e4%b9%88)
  - [pytest使用](#pytest%e4%bd%bf%e7%94%a8)
  - [简单例子](#%e7%ae%80%e5%8d%95%e4%be%8b%e5%ad%90)
  - [常用的pytest第三方插件](#%e5%b8%b8%e7%94%a8%e7%9a%84pytest%e7%ac%ac%e4%b8%89%e6%96%b9%e6%8f%92%e4%bb%b6)
    - [pytest-sugar](#pytest-sugar)
    - [pytest-assume](#pytest-assume)
    - [pytest-ordering](#pytest-ordering)
  - [pytest参数化](#pytest%e5%8f%82%e6%95%b0%e5%8c%96)
  - [pytest执行级别](#pytest%e6%89%a7%e8%a1%8c%e7%ba%a7%e5%88%ab)
    - [setup_module 模块级别](#setupmodule-%e6%a8%a1%e5%9d%97%e7%ba%a7%e5%88%ab)
    - [setup_class 类级别](#setupclass-%e7%b1%bb%e7%ba%a7%e5%88%ab)
    - [setup_function 函数级别](#setupfunction-%e5%87%bd%e6%95%b0%e7%ba%a7%e5%88%ab)
  - [fixture](#fixture)
    - [fixture函数作用](#fixture%e5%87%bd%e6%95%b0%e4%bd%9c%e7%94%a8)
    - [fixture 参数](#fixture-%e5%8f%82%e6%95%b0)
      - [函数级别](#%e5%87%bd%e6%95%b0%e7%ba%a7%e5%88%ab)
      - [模块级别](#%e6%a8%a1%e5%9d%97%e7%ba%a7%e5%88%ab)
  - [webtest](#webtest)

<!-- /TOC -->

## pytest是什么

pytest是python的一种单元测试框架


##pytest安装

```
pip install pytest
```

## pytest使用


pytest测试样例非常简单，只需要按照下面的规则：

>测试文件以test_开头（以_test结尾也可以-建议是**test开头**）
>测试类以Test开头，并且不能带有` __init__ `方法
>测试函数以test_开头
>断言使用基本的assert即可
>fixture的文件名必须是conftest.py


## 简单例子

例子1: demo.py
```python
def add(a, b):
    return a + b
```

test_demo.py
```python
from demo import add

def test_add():
    assert add(1, 2) == 3
    assert add(1, 0) == 1
    assert add(1, -1) == 0
```
<br>

执行下面命令进行测试：
```
 pytest test_demo.py 
```
<br>

例子2: 使用类来组成多个用例的
```python
import pytest

# content of test_class.py
class TestClass:
    def test_one(self):
        x = "this"
        assert 'h' in x

    def test_two(self):
        x = "hello"
        assert hasattr(x, 'check')
```

## 常用的pytest第三方插件


插件安装方式pip install 插件名称


### pytest-sugar 

这个插件是显示执行的进度条


### pytest-assume 
这个插件功能是一次性执行完所有测试用例，

不会因为某个测试用例失败而终止后面的用例

test_demo.py
```python
import pytest
from demo import add


def test_add():
    pytest.assume(add(1, 2) == 3)
    pytest.assume(add(1, 0) == 2)
    pytest.assume(add(1, -1) == 0)
```

### pytest-ordering 
设置测试用例执行顺序



## pytest参数化

1.通过@pytest.mark.parametrize装饰器传参

test_demo.py
```python
import pytest
from demo import add

# 参数
@pytest.mark.parametrize('a, b, re', [
    (3, 4, 7),
    (1, -1, 0),
    (1, 1.2, 2.2),
])
# 测试加法
def test_add(a, b, re):
    pytest.assume(add(a, b) == re)
```



2.通过读取文件的形式参数化

test.json:

```python
[[3, 4, 7], [1, -1, 0], [1, 1.2, 2.2]]
```

test_demo.py

```python
import json
import pytest
from demo import add

# 读取json文件
def load_json():
    with open('test.json', 'r') as f:
        re = json.load(f)
        return re

# 把读到文件的数据通过参数传递
@pytest.mark.parametrize('a, b, re', load_json())
# 测试加法运算
def test_add(a, b, re):
    pytest.assume(add(a, b) == re)
```

备注：

> 1.pytest.mark.parametrize的第一个参数代表测试方法接收的参数个数，第二个参数为接收的数据内容。
> 
>2.如果是通过读文件的形式传参需要注意读出来的文件内容是否和要传的参数类型一致。

<br>
## pytest执行级别

### setup_module 模块级别

如果在单个模块中有多个测试函数和测试类，则可以选择实现以下方法（只会执行一次）

test_demo.py
```python
def setup_module():
    """ 模块执行前准备工作 """

class Test1():
    """ 测试类1 """

class Test2():
    """ 测试类2 """

def teardown_module():
    """ 模块执行后的结束工作 """
```


### setup_class 类级别

在调用类的所有测试方法之前和之后调用以下方法（只会执行一次）
```python
class TestDemo():
    @classmethod
    def setup_class(cls):
        """ 类方法执行前准备工作 """

    def test_add(self):
        """ 测试方法 """

    @classmethod
    def teardown_class(cls):
        """ 类方法执行后结束工作 """
```

### setup_function 函数级别

在执行每个函数执行，调用以下方法（只会执行一次）
```python
def setup_function():
    print("setup_function：每个用例开始前都会执行")


def teardown_function():
    print("teardown_function：每个用例结束后都会执行")


def test_one():
    print("正在执行----test_one")
    x = "this"
    assert 'h' in x


def test_two():
    print("正在执行----test_two")
    x = "hello"
    assert hasattr(x, 'check')


def test_three():
    print("正在执行----test_three")
    a = "hello"
    b = "hello world"
    assert a in b
```



## fixture


### fixture函数作用
完成setup和teardown操作，处理数据库、文件等资源的打开和关闭

完成大部分测试用例需要完成的通用操作，例如login、设置config参数、环境变量等

准备测试数据，将数据提前写入到数据库，或者通过params返回给test用例等

有独立的命名，可以按照测试的用途来激活，比如用于functions/modules/class/session

### fixture 参数

scope参数

scope=function：每个test都运行，默认是function的scope

scope=class：每个class的所有test只运行一次

scope=module：每个module的所有test只运行一次

scope=session：每个session只运行一次


#### 函数级别

test_demo.py
```python
import pytest

# 定义fixture函数
@pytest.fixture()
def fixture_func():
    print("\n这是一个fixture函数,用来完成一些提前准备")

def test_1(fixture_func):
    print("测试用例1")
    assert 1 == 1

def test_2(fixture_func):
    print("测试用例2")
    assert 2 == 2
```

使用`pytest  -s  test_demo.py `运行，才可以打印出print语句


#### 模块级别

```python
import pytest


@pytest.fixture(scope="module")
def fixture_func():
    print("\n这是一个fixture函数,用来完成一些提前准备")

def test_1(fixture_func):
    print("测试用例1")
    assert 1 == 1

def test_2(fixture_func):
    print("测试用例2")
    assert 2 == 2
```

####session级别fixture

1.先创建一个conftest.py文件，输入以下代码

说明：当有测试用例调用pytest.fixture函数时，

**pytest会自动去找conftest.py文件里找pytest.fixture函数，不需要import。**

```python
import pytest
@pytest.fixture(scope="session")
def fixture_func():
    print("\n这是一个fixture函数,用来完成一些提前准备")
```



2.创建第一个测试用例文件 test_demo.py

```python
def test_1(fixture_func):
    print('测试用例1')
    assert 1 == 1

def test_2(fixture_func):
    print('测试用例2')
    assert 2 == 2
```

3.再创建一个测试文件test_demo2.py

```python
def test_1(fixture_func):
    print('测试用例3')
    assert 3 == 3

def test_2(fixture_func):
    print('测试用例4')
    assert 4 == 4
```

执行命令：
```
pytest -s test_*
```



总结：

装饰器@pytest.fixture用于声明一个方法是fixture方法，如果测试用例的参数列表中包含fixture对象，那么测试用例运行之前会先调用fixture方法。

fixture可以指定其方法范围，由参数scope决定。

@pytest.fixture(scope="function"),function级别的fixture方法在每一个测试用例运行前被调用，待测试用例结束之后再销毁。

如果不给定scope参数，默认情况下就是“function”。

@pytest.fixture(scope="session"),session级别的fixture方法在每一次会话中只运行一次，也就是在所有测试用例之前运行一次，且被所有测试用例共享。

conftest.py可以认为是pytest中的配置文件，单独管理一些预置的操作，与fixture方法配合，pytest在运行测试用例之前会事先调用conftest.py中预置的fixture方法，然后供所有测试用例使用。


## webtest

```python
from webtest import TestApp


def application(environ, start_response):
    body = b"hello"
    headers = [('Content-Type', 'text/html; charset=utf8'),
               ('Content-Length', str(len(body)))]
    start_response('200 OK', headers)
    return [body]


def test_qiniu():
    app = TestApp(application)
    resp = app.get("/")
    assert resp.status == '200 OK'

```
