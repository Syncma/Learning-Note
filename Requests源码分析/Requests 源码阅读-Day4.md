# Requests 源码阅读-Day4

<!-- TOC -->

- [Requests 源码阅读-Day4](#requests-%e6%ba%90%e7%a0%81%e9%98%85%e8%af%bb-day4)
  - [adapters](#adapters)
    - [self.mount](#selfmount)
    - [HTTPAdapter](#httpadapter)
    - [PoolManager](#poolmanager)
    - [总结](#%e6%80%bb%e7%bb%93)

<!-- /TOC -->

## adapters

继续往下看：
session.py
```python
class Session(SessionRedirectMixin):
	....
	    # Default connection adapters.
        self.adapters = OrderedDict()
        self.mount('https://', HTTPAdapter())
        self.mount('http://', HTTPAdapter())
```


调用了OrderedDict方法，这个方法在头声明过
```python
from collections import OrderedDict
```

这个在`Requests 源码阅读-Day2`里面讲过了，这里就不说了

主要讲讲后面的self.mount 和HTTPAdapter



### self.mount
sessions.py
```python
   def mount(self, prefix, adapter):
        """Registers a connection adapter to a prefix.

        Adapters are sorted in descending order by prefix length.
        """
        self.adapters[prefix] = adapter
        keys_to_move = [k for k in self.adapters if len(k) < len(prefix)]
        for key in keys_to_move:
            self.adapters[key] = self.adapters.pop(key)
```

这个方法的作用是**注册适配器，按照前缀长度降序排序**

### HTTPAdapter

这个方法在头部声明过：
```python
from .adapters import HTTPAdapter
```

所以直接到当前目录下的adapters.py里面找到HTTPAdapter方法

```python
class HTTPAdapter(BaseAdapter):
....

```

我读了下这里面的代码，大概了解这个适配器使用了强大的**urllib3**库，使用urllib3库里面的**PoolManager**方法，为requests提供了默认的HTTP和HTTPS交互方法。


### PoolManager

这个PoolManager要详细说说里面的逻辑。

urllib3/poolmanager.py
```python
class PoolManager(RequestMethods):
    proxy = None

    def __init__(self, num_pools=10, headers=None, **connection_pool_kw):
        RequestMethods.__init__(self, headers)
        self.connection_pool_kw = connection_pool_kw
        self.pools = RecentlyUsedContainer(num_pools,
                                           dispose_func=lambda p: p.close())

        # Locally set the pool classes and keys so other PoolManagers can
        # override them.
        self.pool_classes_by_scheme = pool_classes_by_scheme
        self.key_fn_by_scheme = key_fn_by_scheme.copy()
        ....
```

重点是里面的RecentlyUsedContainer 方法，这个方法可以看看
urllib3/_collections.py

```python
class RecentlyUsedContainer(MutableMapping):
    ContainerCls = OrderedDict

    def __init__(self, maxsize=10, dispose_func=None):
        self._maxsize = maxsize
        self.dispose_func = dispose_func

        self._container = self.ContainerCls()
        self.lock = RLock()
	....
```

提供了两个参数, maxsize和dispose_func，

这个maxsize很好理解，就是连接池里面最大的数量

这个dispose_func 就是每次从容器中取出一个

**`它是线程安全类似字典类型的连接池`**


### 总结

self.adapters = OrderedDict()将adapters指向一个新建的有序字典对象，用于存放传输适配器。传输适配器的作用是提供一种机制，让你可以为HTTP服务定义交互方法。

requests自带了一个传输适配器，就是源码中的HTTPAdapter

mount方法会注册一个传输适配器的特定实例到一个前缀上面。加载以后，任何使用该会话的 HTTP 请求，只要其 URL 是以给定的前缀开头，该传输适配器就会被使用到。

所以每当Session被实例化，就会有适配器附着在Session上，这里不管是HTTP还是HTTPS，用的都是同一个适配器HTTPAdapter。
