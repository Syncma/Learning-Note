# docker底层原理

<!-- TOC -->

- [docker底层原理](#docker%e5%ba%95%e5%b1%82%e5%8e%9f%e7%90%86)
    - [CS结构](#cs%e7%bb%93%e6%9e%84)
    - [docker为啥比虚拟机快](#docker%e4%b8%ba%e5%95%a5%e6%af%94%e8%99%9a%e6%8b%9f%e6%9c%ba%e5%bf%ab)

<!-- /TOC -->
### CS结构
Docker是一个client-Server结构的系统

Docker守护进程运行在主机上，然后通过socket连接从客户端访问

守护进程从客户端接受命令并管理运行在主机上的容器。

容器就是一个运行环境，就是我们说的集装箱

```
[root@laptop ~]# ps -ef |grep docker

root     21696     1  0 22:40 ?        00:00:01 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
root     23160 20583  0 22:57 pts/0    00:00:00 grep --color docker

```

### docker为啥比虚拟机快


![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/docker1.png)