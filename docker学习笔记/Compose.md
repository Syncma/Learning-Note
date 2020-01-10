# Docker Compose 学习

<!-- TOC -->

- [Docker Compose 学习](#docker-compose-%e5%ad%a6%e4%b9%a0)
	- [编排技术](#%e7%bc%96%e6%8e%92%e6%8a%80%e6%9c%af)
	- [Compose 介绍](#compose-%e4%bb%8b%e7%bb%8d)
	- [Compose安装](#compose%e5%ae%89%e8%a3%85)
	- [Compose使用](#compose%e4%bd%bf%e7%94%a8)

<!-- /TOC -->


## 编排技术

Docker Compose 是 Docker 官方编排（Orchestration）项目之一，负责快速的部署分布式应用

把全部东西堆到一个容器里面是典型的虚拟机的使用方式，不是 Docker 的正确打开方式

正确的做法是让一个容器做一件事：
数据库、Nginx、Python 应用、缓存等等都是独立的容器，分别启动它们，这些容器组成了一个集群，需要某种方法把它们关联起来。

这个关联有一个非常专用、形象的称呼「编排」，我最早了解这个词是通过「Ansbile Playbooks」，而 Docker Compose 大家可以猜到就是负责实现对 Docker 容器集群编排的。


## Compose 介绍
Dockerfile 可以让用户管理一个单独的应用容器，而 Compose 则允许用户在一个模板 (YAML 格式) 中定义一组相关联的应用容器

Compose 非常适合构建开发和测试环境，但如果你想在生产中使用你的容器，应该选择 Kubernetes 来编排容器


Compose 项目由 Python 编写，实现上调用了 Docker 服务提供的 API 来对容器进行管理。
因此，只要所操作的平台支持 Docker API，就可以在其上利用 Compose 来进行编排管理。

## Compose安装

```
curl -L https://github.com/docker/compose/releases/download/1.25.0/docker-compose-`uname -s`-`uname -m`-o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
```

## Compose使用


Compose 中有两个重要的概念：

* 服务 (service)：一个应用的容器，实际上可以包括若干运行相同镜像的容器实例。
* 项目 (project)：由一组关联的应用容器组成的一个完整业务单元，在 docker-compose.yml 文件中定义。

Compose 的默认管理对象是项目，通过子命令对项目中的一组容器进行便捷地生命周期管理。

1.这里使用flask 框架来搭建一个简单的web应用

demo/app.py:
```python
from flask import Flask
from redis import Redis
app = Flask(__name__)
redis = Redis(host='localhost', port=6379)
@app.route('/')
def hello():
	count = redis.incr('hits')
	return 'Hello World! 该页面已被访问 {} 次。\n'.format(count)
if __name__ == "__main__":
	app.run(host="0.0.0.0", debug=True)
```


2.编写dockerfile文件

在app.py 外面一层

```
FROM python:3.6-alpine
ADD  demo/ /code
WORKDIR /code
RUN pip install redis flask
CMD ["python", "app.py"]
```

3.编写docker-compose.yml 文件

```
version: '3'
services:
web:
build: .
ports:
- "5000:5000"
redis:
image: "redis:alpine"
```


4.运行下面命令开启服务
```
docker-compose up
```

up 命令十分强大，它将尝试自动完成包括构建镜像，（重新）创建服务，启动服务，并关联服务相关容器的一系列操作。链接的服务都将会被自动启动，除非已经处于运行状态。

可以说，大部分时候都可以直接通过该命令来启动一个项目。

默认情况，docker-compose up 启动的容器都在前台，控制台将会同时打印所有容器的输出信息，可以很方便进行调试。

当通过 Ctrl-C 停止命令时，所有容器将会停止。

如果使用 docker-compose up -d，将会在后台启动并运行所有的容器。一般推荐生产环境下使用该选项。


5.此时访问本地 5000 端口，每次刷新页面，计数就会加 1。
