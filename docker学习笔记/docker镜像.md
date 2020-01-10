# docker镜像

<!-- TOC -->

- [docker镜像](#docker镜像)
    - [镜像是什么](#镜像是什么)
    - [UnionFS联合文件系统](#unionfs联合文件系统)
    - [Docker镜像加载原理](#docker镜像加载原理)
    - [分层的镜像](#分层的镜像)
    - [为什么要采用这种分层结构](#为什么要采用这种分层结构)
    - [特点](#特点)
    - [Docker镜像commit操作补充](#docker镜像commit操作补充)

<!-- /TOC -->

## 镜像是什么

镜像是一种轻量级、可执行的独立软件包，用来打包软件运行环境和基于运行环境开发的软件

它包含运行某个软件所需要的所有内容，包括代码、运行库、库、环境变量和配置文件


## UnionFS联合文件系统
是一种**`分层、轻量级并且高性能`**的文件系统，它支持对文件系统的修改作为一次提交来一层层的叠加，同时可以将不同目录挂载到同一个虚拟文件系统下。

**UnionFS文件系统 是docker镜像的基础**，镜像可以通过分层来进行继承，基于基础镜像（没有父镜像),可以制作各种具体的应用镜像

**`类似 花卷`**

特征：
一次同时加载多个文件系统，但从外面看起来，只能看到一个文件系统，联合加载会把各层文件系统叠加起来，这样最终的文件系统会包含所有底层的文件和目录



##Docker镜像加载原理

docker 的镜像实际上由一层一层的文件系统组成，这种层级的文件系统UnionFS

bootfs(boot file system)主要包含bootloader和kernel

bootloader主要是引导加载kernel, Linux刚启动时会加载bootfs文件系统

在Docker镜像的最底层就是bootfs. 这一层与我们典型的Linux/Unix系统是一样的，包含boot加载器和内核，当boot加载完成后整个内核就都在内存中，此时内存的使用权已由bootfs转交给内核，此时系统会卸载bootfs

rootfs(root file system), 在bootfs之上，包含就是典型的Linux系统中的/dev, /proc, /bin, /etc等标准目录和文件，rootfs就是各种不同的操作系统发行版，比如centos, ubuntu等


对于一个精简的OS,  rootfs可以很小，只需要包括最基本的命令，工具和程序库就可以了，因此底层直接用host的kernel，自己只需要提供rootfs就可以了。

因此对于不同的Linux发行版，bootfs基本是一致的，rootfs就会有差别，所以不同的发行版可以公用bootfs

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200110091127.png)

## 分层的镜像

以pull为例子，在下载的过程中可以看到docker的镜像好像是一层一层的在下载

**关于docker是如何安装的？会在后面的文章中介绍**

```
[root@laptop ~]# docker pull tomcat

Using default tag: latest
latest: Pulling from library/tomcat
9a0b0ce99936: Pull complete
db3b6004c61a: Pull complete
f8f075920295: Pull complete
6ef14aff1139: Pull complete
962785d3b7f9: Pull complete
631589572f9b: Pull complete
c55a0c6f4c7b: Pull complete
379605d88e88: Pull complete
e056aa10ded8: Pull complete
6349a1c98d85: Pull complete
Digest: sha256:77e41dbdf7854f03b9a933510e8852c99d836d42ae85cba4b3bc04e8710dc0f7
Status: Downloaded newer image for tomcat:latest
docker.io/library/tomcat:latest
```



## 为什么要采用这种分层结构


**最大的好处 就是`共享资源`**

比如：
有多个镜像都从相同的base镜像构建而来，那么宿主机只需要在磁盘保存一份base镜像
同时内存中也需加载一份base镜像，就可以为所有容器服务了。
而且镜像的每一层都可以被共享


## 特点
**Docker镜像都是只读的**

当容器启动时，一个新的可写层被加载到镜像的顶部
这一层通常被称为**容器层**，容器层之下的都叫**镜像层**


##Docker镜像commit操作补充

**`docker  commit提交容器副本使之成为一个新的镜像`**
**`docker  commit -m="提交的描述信息“ -a="作者" 容器ID  要创建的目标镜像名:[标签名]`**



1.docker数据默认都存在 /var/lib/docker 这里
```
[root@laptop ~]# docker info |grep "Docker Root Dir"
Docker Root Dir: /var/lib/docker
```


如果想修改可以按照下面步骤进行设置:

```
[root@laptop docker]# pwd
/var/lib/docker
[root@laptop docker]# ll
total 48
drwx------ 2 root root 4096 Oct 23 11:19 builder
drwx--x--x 4 root root 4096 Oct 23 11:19 buildkit
drwx------ 9 root root 4096 Nov 15 09:24 containers
drwx------ 3 root root 4096 Oct 23 11:19 image
drwxr-x--- 3 root root 4096 Oct 23 11:19 network
drwx------ 18 root root 4096 Nov 15 10:11 overlay2
drwx------ 4 root root 4096 Oct 23 11:19 plugins
drwx------ 2 root root 4096 Nov 15 08:23 runtimes
drwx------ 2 root root 4096 Oct 23 11:19 swarm
drwx------ 2 root root 4096 Nov 15 10:01 tmp
drwx------ 2 root root 4096 Oct 23 11:19 trust
drwx------ 2 root root 4096 Oct 23 11:19 volumes

[root@laptop ~]# cat /etc/docker/daemon.json
{
"registry-mirrors":["https://1664le6h.mirror.aliyuncs.com"],
"data-root":"/home/jian/prj/docker"
}

```

使用root用户重启服务：
```
[root@laptop ~]# systemctl daemon-reload
[root@laptop ~]# systemctl restart docker
[root@laptop ~]# docker info |grep "Docker Root Dir"
Docker Root Dir: /home/jian/prj/docker
```



例子：

1.从Hub上下载tomcat镜像到本地并运行

```
[root@laptop image]# docker pull tomcat

Using default tag: latest
latest: Pulling from library/tomcat


[root@laptop image]# docker images

REPOSITORY TAG IMAGE ID CREATED SIZE
tomcat latest 882487b8be1d 3 weeks ago 507MB
```


2.运行


```
 # -p 主机端口:docker容器端口
 # -i 表示以“交互模式”运行容器
 # -t 表示容器启动后会进入其命令行
 # -P 随机分配端口

[root@laptop image]# docker run -it  -p 8080:8080 tomcat 
[root@laptop image]# docker run -it -P tomcat

# -d 后台运行
[root@laptop image]#docker run --name tomcat  -p 8080:8080 -d 


[root@laptop ~]# docker ps
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
bcc3a519b844 tomcat "catalina.sh run" 25 seconds ago Up 24 seconds 0.0.0.0:32768->8080/tcp brave_bouman
```


3.故意删除上一步镜像产生tomcat容器文档

```
[root@laptop ~]# docker ps
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
0c26f645c079 tomcat "catalina.sh run" 10 seconds ago Up 9 seconds 8080/tcp serene_cannon

[root@laptop ~]# docker exec -it 0c26f645c079 /bin/bash
root@0c26f645c079:/usr/local/tomcat# cd webapps/
root@0c26f645c079:/usr/local/tomcat/webapps# ls -l
total 20
drwxr-xr-x 3 root root 4096 Oct 19 02:25 ROOT
drwxr-xr-x 15 root root 4096 Oct 19 02:25 docs
drwxr-xr-x 6 root root 4096 Oct 19 02:25 examples
drwxr-xr-x 5 root root 4096 Oct 19 02:25 host-manager
drwxr-xr-x 5 root root 4096 Oct 19 02:25 manager
root@0c26f645c079:/usr/local/tomcat/webapps# rm -rf docs/
root@0c26f645c079:/usr/local/tomcat/webapps# ls -l
total 16
drwxr-xr-x 3 root root 4096 Oct 19 02:25 ROOT
drwxr-xr-x 6 root root 4096 Oct 19 02:25 examples
drwxr-xr-x 5 root root 4096 Oct 19 02:25 host-manager
drwxr-xr-x 5 root root 4096 Oct 19 02:25 manager

```



4.也即当前的tomcat运行实例是一个没有文档内容的容器，
以它为模板commit一个没有doc的tomcat新镜像

```
[root@laptop ~]# docker commit -a "abc" -m="del tomcat docs" a0d94640ea19 abc/mytomcat:1.2
sha256:f73b412690d60b1160430ec18b50727c8c6877235e6aa53c4953a927bd605966

[root@laptop ~]# docker images
REPOSITORY TAG IMAGE ID CREATED SIZE
abc/mytomcat 1.2 f73b412690d6 10 seconds ago 507MB
```



4.启动新镜像和原来的对比

```
[root@laptop ~]# docker images
REPOSITORY TAG IMAGE ID CREATED SIZE
abc/mytomcat 1.2 f73b412690d6 4 minutes ago 507MB
tomcat latest 882487b8be1d 3 weeks ago 507MB

[root@laptop ~]# docker run -it -p 8080:8080 abc/mytomcat:1.2

```

然后访问浏览器 [地址](http://localhost:8080)查看效果