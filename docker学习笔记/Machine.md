# Docker Machine 学习笔记

<!-- TOC -->

- [Docker Machine 学习笔记](#docker-machine-%e5%ad%a6%e4%b9%a0%e7%ac%94%e8%ae%b0)
  - [介绍](#%e4%bb%8b%e7%bb%8d)
  - [安装](#%e5%ae%89%e8%a3%85)

<!-- /TOC -->

## 介绍
Machine项目是Docker官方的开源项目，负责实现对Docker运行环境进行安装和管理，特别在管理多个Docker环境时，使用Machine要比手动管理高效得多。


[参考地址1](https://www.dongwm.com/post/docker-machine-and-swarm/)

[参考地址2](http://dockone.io/article/275)


简单来说：

>Docker Machine是一个简化Docker安装的命令行工具，通过一个简单的命令行即可在相应的平台上安装Docker，比如VirtualBox、 Digital Ocean、Microsoft Azure。


## 安装

```
[root@laptop ~]# curl -L https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-`uname -s`-`uname -m` >/usr/local/bin/docker-machine

[root@laptop docker]# chmod +x /usr/local/bin/docker-machine
[root@laptop tmp]# docker-machine -v
docker-machine version 0.16.2, build bd45ab13
```

Docker Machine 支持多种后端驱动，包括虚拟机、本地主机和云平台等。