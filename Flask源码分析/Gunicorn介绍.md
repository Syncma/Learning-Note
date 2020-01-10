# Gunicorn介绍


<!-- TOC -->

- [Gunicorn介绍](#gunicorn%e4%bb%8b%e7%bb%8d)
  - [Gunicorn是什么](#gunicorn%e6%98%af%e4%bb%80%e4%b9%88)
  - [安装](#%e5%ae%89%e8%a3%85)
  - [Gunicorn源码结构](#gunicorn%e6%ba%90%e7%a0%81%e7%bb%93%e6%9e%84)
  - [gunicorn flask 压测对比](#gunicorn-flask-%e5%8e%8b%e6%b5%8b%e5%af%b9%e6%af%94)
    - [Flask启动压测](#flask%e5%90%af%e5%8a%a8%e5%8e%8b%e6%b5%8b)
    - [Gunicorn启动压测](#gunicorn%e5%90%af%e5%8a%a8%e5%8e%8b%e6%b5%8b)
    - [问题](#%e9%97%ae%e9%a2%98)
      - [gunicorn并发比flask好的原因](#gunicorn%e5%b9%b6%e5%8f%91%e6%af%94flask%e5%a5%bd%e7%9a%84%e5%8e%9f%e5%9b%a0)
      - [gunicorn和flask通信 流程](#gunicorn%e5%92%8cflask%e9%80%9a%e4%bf%a1-%e6%b5%81%e7%a8%8b)
  - [gunicorn 几种 worker 性能测试比较](#gunicorn-%e5%87%a0%e7%a7%8d-worker-%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e6%af%94%e8%be%83)
  - [gunicorn 配置](#gunicorn-%e9%85%8d%e7%bd%ae)
    - [workers模式](#workers%e6%a8%a1%e5%bc%8f)
    - [多线程模式](#%e5%a4%9a%e7%ba%bf%e7%a8%8b%e6%a8%a1%e5%bc%8f)
    - [伪线程 gevent (协程)](#%e4%bc%aa%e7%ba%bf%e7%a8%8b-gevent-%e5%8d%8f%e7%a8%8b)
    - [建议](#%e5%bb%ba%e8%ae%ae)
  - [wrk压测工具](#wrk%e5%8e%8b%e6%b5%8b%e5%b7%a5%e5%85%b7)
  - [其他备注](#%e5%85%b6%e4%bb%96%e5%a4%87%e6%b3%a8)

<!-- /TOC -->

## Gunicorn是什么

Gunicorn ‘Green Unicorn’ 是一个 UNIX 下的 WSGI HTTP 服务器，它是一个 移植自 Ruby 的 Unicorn 项目的 pre-fork worker 模型。它既支持 eventlet ， 也支持 greenlet


在管理 worker 上，使用了 pre-fork 模型，即一个 master 进程管理多个 worker 进程，所有请求和响应均由 Worker 处理。Master 进程是一个简单的 loop, 监听 worker 不同进程信号并且作出响应。比如接受到 TTIN 提升 worker 数量，TTOU 降低运行 Worker 数量。如果 worker 挂了，发出 CHLD, 则重启失败的 worker, 同步的 Worker 一次处理一个请求。


[通过优化 Gunicorn 配置提高性能](https://juejin.im/post/5ce8cab8e51d4577523f22f8)


## 安装

目前Gunicorn只能运行在Linux环境中，不支持windows平台

安装:
```
[jian@laptop ~]$  pip install gunicorn

[jian@laptop ~]$ pip show gunicorn
Name: gunicorn
Version: 19.7.1
Summary: WSGI HTTP Server for UNIX
Home-page: http://gunicorn.org
Author: Benoit Chesneau
Author-email: benoitc@e-engura.com
License: MIT
Location: /home/jian/.pyenv/versions/3.6.7/lib/python3.6/site-packages
Requires: 
Required-by: 


```

## Gunicorn源码结构

- 要分析源码 待补充

## gunicorn flask 压测对比

测试例子:  demo.py
```python
from flask import Flask
app = Flask(__name__)


@app.route('/')
def index():
    return 'hello world!'


if __name__ == '__main__':
    app.run()
```

### Flask启动压测

1. 直接运行demo.py，使用flask自带的WSGI
```python
[jian@laptop practics]$ python demo.py 
 * Serving Flask app "demo" (lazy loading)
 * Environment: production
   WARNING: This is a development server. Do not use it in a production deployment.
   Use a production WSGI server instead.
 * Debug mode: off
 * Running on http://127.0.0.1:5000/ (Press CTRL+C to quit)

```

2.使用ab工具进行压测

```
[jian@laptop practics]$ ab -n 500 -c 500 http://localhost:5000/
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Finished 500 requests


Server Software:        Werkzeug/0.16.0
Server Hostname:        localhost
Server Port:            5000

Document Path:          /
Document Length:        12 bytes

Concurrency Level:      500
Time taken for tests:   0.477 seconds
Complete requests:      500
Failed requests:        0
Total transferred:      83000 bytes
HTML transferred:       6000 bytes
Requests per second:    1049.28 [#/sec] (mean)
Time per request:       476.515 [ms] (mean)
Time per request:       0.953 [ms] (mean, across all concurrent requests)
Transfer rate:          170.10 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   4.6      0      11
Processing:    13  111  29.4    122     161
Waiting:        2  110  29.5    121     159
Total:         13  114  26.1    122     164

Percentage of the requests served within a certain time (ms)
  50%    122
  66%    124
  75%    125
  80%    128
  90%    136
  95%    145
  98%    153
  99%    160
 100%    164 (longest request)

```

上面结果可以得出，同时并发500请求，压测结果是这样：
```
Requests per second:    1049.28 [#/sec] (mean)
Time per request:       476.515 [ms] (mean)
Time per request:       0.953 [ms] (mean, across all concurrent 
```


### Gunicorn启动压测

1.使用gunicorn启动， 这里启动4个进程

```
# 其中： 
第一个 server 指的是 server.py 文件； 第二个指的是 flask 应用的名字，app = Flask(name)

-w 为开启n个进程

[jian@laptop practics]$ gunicorn -w 4 -b 0.0.0.0:8000 demo:app
[2020-01-10 21:08:56 +0800] [15563] [INFO] Starting gunicorn 19.7.1
[2020-01-10 21:08:56 +0800] [15563] [INFO] Listening at: http://0.0.0.0:8000 (15563)
[2020-01-10 21:08:56 +0800] [15563] [INFO] Using worker: sync
[2020-01-10 21:08:56 +0800] [15670] [INFO] Booting worker with pid: 15670
[2020-01-10 21:08:56 +0800] [15671] [INFO] Booting worker with pid: 15671
[2020-01-10 21:08:56 +0800] [15672] [INFO] Booting worker with pid: 15672
[2020-01-10 21:08:56 +0800] [15675] [INFO] Booting worker with pid: 15675

```

2.使用ab工具进行压测

```
[jian@laptop practics]$ ab -n 500 -c 500 http://localhost:8000/
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Finished 500 requests


Server Software:        gunicorn/19.7.1
Server Hostname:        localhost
Server Port:            8000

Document Path:          /
Document Length:        12 bytes

Concurrency Level:      500
Time taken for tests:   0.111 seconds
Complete requests:      500
Failed requests:        0
Total transferred:      86000 bytes
HTML transferred:       6000 bytes
Requests per second:    4522.80 [#/sec] (mean)
Time per request:       110.551 [ms] (mean)
Time per request:       0.221 [ms] (mean, across all concurrent requests)
Transfer rate:          759.69 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    4   6.9      0      16
Processing:     5   21   5.8     22      28
Waiting:        1   21   5.9     22      28
Total:         18   26   4.7     25      41

Percentage of the requests served within a certain time (ms)
  50%     25
  66%     27
  75%     28
  80%     28
  90%     31
  95%     37
  98%     40
  99%     40
 100%     41 (longest request)

```

上面结果可以得出，同时并发500请求，压测结果是这样：
```
Requests per second:    4522.80 [#/sec] (mean)
Time per request:       110.551 [ms] (mean)
Time per request:       0.221 [ms] (mean, across all concurrent 
```

可以明显看到Requests per second: 明显比flask自带的要高
而且Time per request 也少了很多


### 问题

####  gunicorn并发比flask好的原因

查看 flask 代码的时候可以看到这个 WebServer 的名称也叫做 run_simple 

too simple 的东西往往不太适合生产

```
from werkzeug.serving import run_simple
	run_simple('localhost', 5000, application, use_reloader=True)
```


1.单 Worker

只有一个进程在跑所有的请求，而由于实现的简陋性，内置 webserver 很容易卡死。

并且只有一个 Worker 在跑请求。在多核 CPU 下，仅仅占用一核。

当然，其实也可以多起几个进程。

2.缺乏 Worker 的管理

加入负载量上来了，Gunicorn 可以调节 Worker 的数量

flask内置的 Webserver 是不适合做这种事情的

一言以蔽之，太弱，几个请求就打满了


####  gunicorn和flask通信 流程

nginx<->gunicorn<->flask



## gunicorn 几种 worker 性能测试比较

1.Gunicorn目前自带支持几种工作方式:

sync (默认值)
eventlet
gevent
tornado


2.安装测试模块

```
[jian@laptop tmp]$ cat requirements.txt 
gunicorn==19.7.1
flask==1.1.1
flask-redis==0.4.0
gevent==1.2.2
tornado==4.5.3
eventlet==0.25.1

#这里要特别注意tornado版本必须是5.0以下，不然gunicorn 在启动会报错：
TypeError: __init__() got an unexpected keyword argument 'io_loop'

[jian@laptop tmp]$ pip install -r requirements.txt

```

3. 测试例子

测试环境: Fedora 29 x64

需要安装redis , 可以使用下面命令进行安装：
```
[root@laptop ~]# dnf install redis
```

开启redis服务：
```
[root@laptop ~]# /usr/bin/redis-server /etc/redis.conf 

```

测试程序：
```python
from flask import Flask
from flask_redis import FlaskRedis

REDIS_URL = "redis://:@localhost:6379/0"
app = Flask(__name__)
app.config.from_object(__name__)

redis = FlaskRedis(app, True)


@app.route('/')
def index():
    redis.incr("hit", 1)
    return redis.get("hit")


if __name__ == '__main__':
    app.run()

```

4.开始测试


分别使用四种方式开启服务

```
[jian@laptop practics]$ gunicorn -w 4 demo:app --worker-class sync
[jian@laptop practics]$ gunicorn -w 4 demo:app --worker-class gevent
[jian@laptop practics]$ gunicorn -w 4 demo:app --worker-class tornado
[jian@laptop practics]$ gunicorn -w 4 demo:app --worker-class eventlet

```


使用ab工具,并行500个客户端，发送50000次请求，压测命令：
```
[jian@laptop practics]$ ab -c 500 -t 30 -r http://localhost:8000/
```


5.测试结果

| Worker class | Time taken for tests | Complete requests | Failed requests | Requests per second | 用户平均请求等待时间 | 服务器平均处理时间 | 最小连接时间 | 平均连接时间 | 50%的连接时间 | 最大连接时间 |
| :----------- | -------------------: | :---------------: | :-------------: | :-----------------: | :------------------: | :----------------: | :----------: | :----------: | :-----------: | :----------: |
| sync         |              43.362s |       49719       |       157       |       1146.61       |      436.069ms       |      0.872ms       |     12ms     |     55ms     |     25ms      |   33574ms    |
| gevent       |              13.062s |       50000       |        0        |       3827.96       |      130.618ms       |      0.261ms       |     3ms      |    129ms     |     96ms      |    1477ms    |
| tornado      |              27.925s |       50000       |       17        |       1790.50       |      279.252ms       |      0.559ms       |     16ms     |    146ms     |    27850ms    |
| eventlet     |              12.601s |       50000       |        0        |       3967.88       |      126.012ms       |      0.252ms       |     9ms      |    125ms     |    1377ms     |



eventlet 和gevent两种方式效果最好，数据基本差不多.




## gunicorn 配置

### workers模式

每个worker都是一个加载python应用程序的UNIX进程
worker之间没有共享内存

建议workers 数量是 (2*CPU) + 1

### 多线程模式

gunicorn 还允许每个worker拥有多个线程

在这种模式下，每个worker都会加载一次，同一个worker生成的每个线程共享相同的内存空间

使用threads模式，每一次使用threads模式，worker类就会是gthread

```
gunicorn -w 5 --threads=2  main:app
```

等同于：
```
gunicorn -w 5 --thread=2 --worker-class=gthread main:app
```

最大的并发请求就是worker * 线程 ， 也就是10

建议最大并发数 是(2*CPU) +1


### 伪线程 gevent (协程)

```
gunicorn --worker-class=gevent --worker-connections=1000 -w 3 main:app
```

work-connections 是对gevent worker类的特殊设置

建议workers数量 仍然是 (2*CPU) + 1

在这种情况下，最大的并发请求数 是3000（3个worker * 1000连接/worker)


### 建议

IO 受限 							-建议使用gevent或者asyncio

CPU受限							-建议增加workers数量

不确定内存占用?			-建议使用gthread

不知道怎么选择？			-建议增加workers数量


## wrk压测工具

[git地址](https://github.com/wg/wrk)


格式：
```
wrk -d 20s -t 10 -c 200 url

-d : duration 测试持续时间
-t : threads 线程数
-c : connection 连接数

```

这里使用sync进程来进行开启服务：
```
[jian@laptop practics]$ gunicorn -w 4 demo:app --worker-class sync

也可以写成：
[jian@laptop practics]$ gunicorn -w 4 demo:app  -k sync

```

压测测试例子：

```
[jian@laptop practics]$ wrk -d 20s -t 10 -c 200 http://localhost:8000/
Running 20s test @ http://localhost:8000/
  10 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    44.89ms   80.33ms   1.69s    97.20%
    Req/Sec   419.02    207.58   818.00     61.69%
  74361 requests in 20.02s, 11.70MB read
  Socket errors: connect 0, read 0, write 0, timeout 10
Requests/sec:   3714.26
Transfer/sec:    598.49KB


```

需要注意：
1.一共产生了74361 个requests
2.一共产生了10个timeout
3.平均延迟访问：44.89ms



## 其他备注

Gunicorn对静态文件的支持不太好，所以生产环境下常用Nginx作为反向代理服务器。

生产环境都是Nginx + gunicorn + flask



