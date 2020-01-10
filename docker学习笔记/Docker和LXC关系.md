# Docker和LXC关系
<!-- TOC -->

- [Docker和LXC关系](#docker%e5%92%8clxc%e5%85%b3%e7%b3%bb)
  - [LXC](#lxc)
    - [介绍](#%e4%bb%8b%e7%bb%8d)
    - [优缺点](#%e4%bc%98%e7%bc%ba%e7%82%b9)
  - [docker VS LXC](#docker-vs-lxc)

<!-- /TOC -->

## LXC

### 介绍
LXC为Linux Container的简写，是一种轻量级的虚拟化的手段

LXC 采用 Cgroup 系统来对容器进行资源管理，采用NameSpace来进行资源限制

[参考地址](https://www.cnblogs.com/createyuan/p/5248140.html)


### 优缺点

1.优点

>虚拟化开销小，一台物理机可以运行很多“小”虚拟机。

>通过 cgroup 的方法增减 CPU（中央处理器）/内存非常方便，调整速度很快。

>虚拟机运行速度和本地环境相同的速度基本相同。

2.缺点

>不能热迁移 。
>不能模拟不同体系结构、装不同操作系统。
>安全隔离差 。

## docker VS LXC


**docker底层使用了LXC来实现**, 在LXC的基础之上，docker提供了一系列更强大的功能。

主要提供了下面几个功能：
* 跨主机部署 -dockfile
* 自动构建
* 类似 git版本管理
* 组件重用
* 共享
* 生态链丰富-各种容器编排技术


**LXC 代表产品  ： Proxmox**[官问地址](https://www.proxmox.com/en/)

