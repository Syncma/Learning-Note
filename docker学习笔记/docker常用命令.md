# docker常用命令

<!-- TOC -->

- [docker常用命令](#docker常用命令)
    - [服务开启](#服务开启)
    - [帮助命令](#帮助命令)
    - [镜像命令](#镜像命令)
    - [容器命令](#容器命令)
        - [新建并启动容器](#新建并启动容器)
        - [列出当前所有正在运行的容器](#列出当前所有正在运行的容器)
        - [退出容器](#退出容器)
        - [启动容器](#启动容器)
        - [重启容器](#重启容器)
        - [停止容器](#停止容器)
        - [强制停止容器](#强制停止容器)
        - [删除已停止的容器](#删除已停止的容器)
        - [一次性删除多个容器](#一次性删除多个容器)
    - [启动守护式容器](#启动守护式容器)
        - [查看容器日志](#查看容器日志)
        - [查看容器内运行的进程](#查看容器内运行的进程)
        - [查看容器内部细节](#查看容器内部细节)
        - [命令行交互](#命令行交互)
        - [拷贝文件](#拷贝文件)

<!-- /TOC -->

## 服务开启

```
[root@laptop ~]# systemctl start docker.service 
```

使用普通用户执行docker命令会报错

```
[jian@laptop ~]$ docker run busybox echo "hello world"
docker: Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Post http://%2Fvar%2Frun%2Fdocker.sock/v1.40/containers/create: dial unix /var/run/docker.sock: connect: permission denied.
See 'docker run --help'.

解决的办法：
1.使用root用户

2.如果不想使用root， 可以执行下面步骤

1> 先使用root赋予普通用户权限：
[root@laptop ~]# setfacl -m user:jian:rw /var/run/docker.sock

-m  user:用户名:rw (rw表示可读可写）


或者也可以吧普通用户放到docker组

usermod -a -G docker $USER


2> 然后就可以使用普通用户进行docker操作了
```




## 帮助命令

```
docker version
docker info
docker --help
```

## 镜像命令

```
docker images  列出本地的镜像列表
[root@laptop ~]# docker images
REPOSITORY TAG IMAGE ID CREATED SIZE
hello-world latest fce289e99eb9 10 months ago 1.84kB

docker  search 某个xxx镜像名字
docker pull xxx  下载镜像
docker rmi  xxx  删除镜像

删除单个 docker rmi -f 镜像ID
删除多个  docker rmi -f 镜像名1:TAG 镜像2:TAG
删除全部  docker rmi -f $(docker images -qa)
```


## 容器命令

**`有镜像才能创建容器，这是前提`**

### 新建并启动容器
```
[root@laptop ~]# docker images
REPOSITORY TAG IMAGE ID CREATED SIZE
centos latest 0f3e07c0138f 6 weeks ago 220MB
[root@laptop ~]# docker run -it centos
[root@eff6100b601e /]#
```

### 列出当前所有正在运行的容器

```
[root@laptop ~]# docker ps
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
eff6100b601e centos "/bin/bash" 3 minutes ago Up 3 minutes peaceful_rubin
```

### 退出容器

两种退出方式：

* exit 容器停止退出
```
[root@eff6100b601e /]# exit 
exit
```

* 容器不停止退出   快捷键 **`Ctrl+P+Q`** 


### 启动容器

```
[root@laptop ~]# docker ps
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
e0e72d00e0e5 centos "/bin/bash" 2 minutes ago Up 2 minutes awesome_banach
[root@laptop ~]# docker start e0e72d00e0e5
e0e72d00e0e5
```

### 重启容器
```
[root@laptop ~]# docker restart e0e72d00e0e5
e0e72d00e0e5
```


### 停止容器
```
[root@laptop ~]# docker stop e0e72d00e0e5
e0e72d00e0e5
```

### 强制停止容器
```
[root@laptop ~]# docker kill be6bb678e724
be6bb678e724
```

### 删除已停止的容器
```
[root@laptop ~]# docker ps -l
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
be6bb678e724 centos "/bin/bash" About a minute ago Exited (137) 47 seconds ago modest_newton
[root@laptop ~]# docker rm be6bb678e724
be6bb678e724
```

### 一次性删除多个容器

```
[root@laptop ~]# docker rm -f $(docker ps -a -q)
c1b03eed297c
a725bb916b70
e0e72d00e0e5
eff6100b601e
5729d1fd4990


# 或者使用这个命令
[root@laptop ~]# docker ps -a -q |xargs docker rm

```



## 启动守护式容器

```
[root@laptop ~]# docker images
REPOSITORY TAG IMAGE ID CREATED SIZE
centos latest 0f3e07c0138f 6 weeks ago 220MB
[root@laptop ~]# docker run -d centos
032e471a76f72e4eae171ff0a3eb574144e15b6e6ce061bfffd3ef3f5496b2dc


[root@laptop ~]# docker ps
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
```

这里发现容器已经退出，原因在于：
**`Docker 容器后台运行，必须有一个前台进程`**

容器运行的命令如果不是那些一直挂起的命令（如top, tail)， 就是会自动退出的

**这是docker的机制问题**

>比如web容器 以nginx为例
>正常情况下，我们配置启动服务 只需要启动服务， 如service nginx start
>但是这样做，nginx为后台进程模式运行，这就导致docker前台没有运行的应用
>这样容器后台启动后，会立即自杀因为他觉得他没事可做了

所以，最佳的方案是将你要运行的程序以前台进程的形式运行

```
[root@laptop ~]# docker run -it centos
[root@e3e2159b1a89 /]#
```


### 查看容器日志

```
[root@laptop ~]# docker run -d centos /bin/sh -c "while true;do echo hello;sleep 2;done"

[root@laptop ~]# docker ps
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
1ba1e8801c18 centos "/bin/sh -c 'while t…" 18 seconds ago Up 17 seconds vigilant_matsumoto
e3e2159b1a89 centos "/bin/bash" 3 minutes ago Up 3 minutes gracious_hofstadter

[root@laptop ~]# docker logs -t -f 1ba1e8801c18 #容器ID
2019-11-15T01:06:18.309763627Z hello
2019-11-15T01:06:20.311834609Z hello

```

### 查看容器内运行的进程

```
[root@laptop ~]# docker top 1ba1e8801c18 #容器ID
UID PID PPID C STIME TTY TIME CMD
root 21677 21660 0 09:06 ? 00:00:00 /bin/sh -c while true;do echo hello;sleep 2;done
root 21981 21677 0 09:09 ? 00:00:00 /usr/bin/coreutils --coreutils-prog-shebang=sleep /usr/bin/sleep 2
```


### 查看容器内部细节

```
[root@laptop ~]# docker inspect 1ba1e8801c18 #容器ID

"Id": "1ba1e8801c185ae89f35a169181af040442d539987269ec4f54a5e581ee6966b",
"Created": "2019-11-15T01:06:17.947483399Z",
....
```

### 命令行交互

进入正在运行的容器并以命令行交互

```
[root@laptop ~]# docker ps
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
a7b5e40289de centos "/bin/bash" 16 seconds ago Up 15 seconds practical_tu
[root@laptop ~]# docker attach a7b5e40289de
[root@a7b5e40289de /]#

[root@laptop ~]# docker exec -t a7b5e40289de ls -l /tmp
total 8
-rwx------ 1 root root 1379 Sep 27 17:13 ks-script-0n44nrd1
-rwx------ 1 root root 671 Sep 27 17:13 ks-script-w6m6m_20
```

这两个命令的区别：
```
attach 直接进入容器启动命令的终端，不会启动新的进程
exec 是在容器中打开新的终端，并且可以启动新的进程
```


### 拷贝文件
从容器内拷贝文件到主机上

```
docker cp 容器ID:容器内路径  目的主机路径
```
