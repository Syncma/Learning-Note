# docker介绍

<!-- TOC -->

- [docker介绍](#docker%e4%bb%8b%e7%bb%8d)
  - [Docker是什么](#docker%e6%98%af%e4%bb%80%e4%b9%88)
  - [Docker能干嘛](#docker%e8%83%bd%e5%b9%b2%e5%98%9b)
    - [之前的虚拟机技术](#%e4%b9%8b%e5%89%8d%e7%9a%84%e8%99%9a%e6%8b%9f%e6%9c%ba%e6%8a%80%e6%9c%af)
    - [容器虚拟化技术](#%e5%ae%b9%e5%99%a8%e8%99%9a%e6%8b%9f%e5%8c%96%e6%8a%80%e6%9c%af)
    - [区别](#%e5%8c%ba%e5%88%ab)

<!-- /TOC -->
## Docker是什么

**`Docker是基于Go语言实现的云开源项目`**

对应用程序的封装、分发、部署、运行等生命周期的管理，使用户的APP（或者一个web应用)以及运行环境做到“**一次封装，到处运行**“

解决了运行环境和配置问题软件容器，方便做持续继承并有助于整体发布的容器虚拟化技术


## Docker能干嘛

### 之前的虚拟机技术

虚拟机的缺点： 
* 资源占用多
*  冗余步骤多 
*  启动慢


![传统虚拟机](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/vm1.png)



### 容器虚拟化技术

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/vm2.png)



### 区别


>1.传统虚拟化技术是虚拟出一套硬件后，在其上运行一个完整操作系统，在该系统上再运行所需应用进程

>2.而容器内的应用进程直接运行于宿主的内核，容器内没有自己的内核，而且也没有进行硬件虚拟
因此容器要比传统虚拟机更为轻便

>3.每个容器之间互相隔离，每个容器有自己的文件系统，容器之间进程不会相互影响，能区分计算资源