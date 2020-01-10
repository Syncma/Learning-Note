# docker架构
<!-- TOC -->

- [docker架构](#docker%e6%9e%b6%e6%9e%84)
  - [架构总述](#%e6%9e%b6%e6%9e%84%e6%80%bb%e8%bf%b0)
  - [组件介绍](#%e7%bb%84%e4%bb%b6%e4%bb%8b%e7%bb%8d)

<!-- /TOC -->

## 架构总述
Docker 架构Docker 使用客户端-服务器 (C/S) 架构模式，使用远程API来管理和创建Docker容器。

Docker 容器通过 Docker 镜像来创建。



![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/vm3.png)


容器与镜像的关系类似于面向对象编程中的对象与类。

| Docker | 面向对象 |
| :----- | -------: |
| 容器   |     对象 |
| 镜像   |       类 |


## 组件介绍

| 组件                   |                                                        介绍 |
| :--------------------- | ----------------------------------------------------------: |
| Docker 容器(Container) |                              容器是独立运行的一个或一组应用 |
| Docker 镜像(Images)    |                     Docker 镜像是用于创建 Docker 容器的模板 |
| Docker 客户端(Client)  | Docker 客户端通过命令行或者其他工具与 Docker 的守护进程通信 |
| Docker 主机(Host)      |        一个物理或者虚拟的机器用于执行 Docker 守护进程和容器 |
| Docker 仓库(Registry)  |     Docker 仓库用来保存镜像，可以理解为代码控制中的代码仓库 |
| Docker Machine         |              Docker Machine是一个简化Docker安装的命令行工具 |
