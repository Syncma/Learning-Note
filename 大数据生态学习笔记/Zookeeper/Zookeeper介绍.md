# Zookeeper介绍
<!-- TOC -->

- [Zookeeper介绍](#zookeeper介绍)
    - [介绍](#介绍)
    - [工作机制](#工作机制)
    - [特点](#特点)
    - [数据结构](#数据结构)
    - [应用场景](#应用场景)
        - [统一命名服务](#统一命名服务)
        - [统一配置管理](#统一配置管理)
        - [统一集群管理](#统一集群管理)
        - [服务器动态上下线](#服务器动态上下线)
        - [软负载均衡](#软负载均衡)
    - [ZK安装](#zk安装)
    - [ZK 配置文件说明](#zk-配置文件说明)

<!-- /TOC -->

## 介绍

ZK 是一个开源的分布式的，为分布式应用**提供协调服务**的Apache项目


## 工作机制


![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/zook1.png)

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/zook2.png)


总结：

>  **ZK = 文件系统+ 通知机制**

## 特点

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/zk333.png)



## 数据结构

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/zk-info.png)


## 应用场景

### 统一命名服务

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200118213459.png)

### 统一配置管理

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200118213535.png)

### 统一集群管理

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200118213619.png)

### 服务器动态上下线

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200118213650.png)


### 软负载均衡

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200118213745.png)




## ZK安装


[下载地址](https://mirrors.tuna.tsinghua.edu.cn/apache/zookeeper/)

1.从上面网站下载zk :  zookeeper-3.4.14.tar.gz
```
[jian@laptop bigdata]$ tar xf zookeeper-3.4.14.tar.gz
```

2.安装JDK

```
[jian@laptop bigdata]$ java --version
openjdk 11-ea 2018-09-25
OpenJDK Runtime Environment (build 11-ea+28)
OpenJDK 64-Bit Server VM (build 11-ea+28, mixed mode, sharing)
```

3.修改配置文件

```
[jian@laptop zookeeper-3.4.14]$ pwd
/home/jian/prj/bigdata/zookeeper-3.4.14
[jian@laptop zookeeper-3.4.14]$ mkdir zkData
[jian@laptop zookeeper-3.4.14]$ cd zkData/
[jian@laptop zkData]$ pwd
/home/jian/prj/bigdata/zookeeper-3.4.14/zkData


[jian@laptop conf]$ pwd
/home/jian/prj/bigdata/zookeeper-3.4.14/conf
[jian@laptop conf]$ mv zoo_sample.cfg  zoo.cfg
[jian@laptop conf]$ vi zoo_cfg
[jian@laptop conf]$ cat zoo_cfg
# The number of milliseconds of each tick
tickTime=2000
# The number of ticks that the initial
# synchronization phase can take
initLimit=10
# The number of ticks that can pass between
# sending a request and getting an acknowledgement
syncLimit=5
# the directory where the snapshot is stored.
# do not use /tmp for storage, /tmp here is just
# example sakes.
dataDir=/home/jian/prj/bigdata/zookeeper-3.4.14/zkData
# the port at which the clients will connect
clientPort=2181
# the maximum number of client connections.
# increase this if you need to handle more clients
#maxClientCnxns=60
#
# Be sure to read the maintenance section of the
# administrator guide before turning on autopurge.
#
# http://zookeeper.apache.org/doc/current/zookeeperAdmin.html#sc_maintenance
#
# The number of snapshots to retain in dataDir
#autopurge.snapRetainCount=3
# Purge task interval in hours
# Set to "0" to disable auto purge feature
#autopurge.purgeInterval=1

```

4.启动

```
[jian@laptop zookeeper-3.4.14]$ bin/zkServer.sh start
ZooKeeper JMX enabled by default
Using config: /home/jian/prj/bigdata/zookeeper-3.4.14/bin/../conf/zoo.cfg
Starting zookeeper ... STARTED
```

查看进程是否启动：
```
[jian@laptop zookeeper-3.4.14]$ jps
6162 QuorumPeerMain
6255 Jps
```

查看状态：
```
[jian@laptop zookeeper-3.4.14]$ bin/zkServer.sh status
ZooKeeper JMX enabled by default
Using config: /home/jian/prj/bigdata/zookeeper-3.4.14/bin/../conf/zoo.cfg
Mode: standalone
```

启动客户端
```
[jian@laptop zookeeper-3.4.14]$ bin/zkCli.sh
Connecting to localhost:2181

退出客户端： 输入quit
```
停止ZK:
```
[jian@laptop zookeeper-3.4.14]$ bin/zkServer.sh stop
ZooKeeper JMX enabled by default
Using config: /home/jian/prj/bigdata/zookeeper-3.4.14/bin/../conf/zoo.cfg
Stopping zookeeper ... STOPPED
```



## ZK 配置文件说明

```

# The number of milliseconds of each tick
tickTime=2000
# The number of ticks that the initial
# synchronization phase can take
initLimit=10
# The number of ticks that can pass between
# sending a request and getting an acknowledgement
syncLimit=5
# the directory where the snapshot is stored.
# do not use /tmp for storage, /tmp here is just
# example sakes.
dataDir=/home/jian/prj/bigdata/zookeeper-3.4.14/zkData
# the port at which the clients will connect
clientPort=2181
# the maximum number of client connections.
# increase this if you need to handle more clients
#maxClientCnxns=60
#
# Be sure to read the maintenance section of the
# administrator guide before turning on autopurge.
#
# http://zookeeper.apache.org/doc/current/zookeeperAdmin.html#sc_maintenance
#
# The number of snapshots to retain in dataDir
#autopurge.snapRetainCount=3
# Purge task interval in hours
# Set to "0" to disable auto purge feature
#autopurge.purgeInterval=1

```

```
tickTime=2000 通信心跳数， 与客户端心跳时间，单位毫秒
initLimit = 10 :  LF (Leader->Follower)初始通信时限
syncLimit=5 ： LF同步通信时限
```






