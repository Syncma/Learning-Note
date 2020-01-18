# Flume安装配置

<!-- TOC -->

- [Flume安装配置](#flume%e5%ae%89%e8%a3%85%e9%85%8d%e7%bd%ae)
  - [下载地址](#%e4%b8%8b%e8%bd%bd%e5%9c%b0%e5%9d%80)
  - [安装配置](#%e5%ae%89%e8%a3%85%e9%85%8d%e7%bd%ae)

<!-- /TOC -->


## 下载地址

[下载地址](https://mirrors.tuna.tsinghua.edu.cn/apache/flume)


## 安装配置
1.解压
```
[jian@laptop bigdata]$ tar xf apache-flume-1.9.0-bin.tar.gz
[jian@laptop conf]$ mv flume-env.sh.template flume-env.sh
```

2.配置环境变量
```
[jian@laptop conf]$ cat flume-env.sh |grep JAVA_HOME
export JAVA_HOME=/home/jian/prj/bigdata/jdk1.8.0_231
```
```
[jian@laptop conf]$ cat ~/.bashrc
#Flume配置
export FLUME_HOME=/home/jian/prj/bigdata/flume
export PATH=$FLUME_HOME/bin:$PATH

//执行source使配置文件生效
[jian@laptop conf]$ source ~/.bashrc
```
3.查看版本
```
[jian@laptop conf]$ flume-ng version
Flume 1.9.0
Source code repository: https://git-wip-us.apache.org/repos/asf/flume.git
Revision: d4fcab4f501d41597bc616921329a4339f73585e
Compiled by fszabo on Mon Dec 17 20:45:25 CET 2018
From source with checksum 35db629a3bda49d23e9b3690c80737f9
```

4. 修改配置文件
```
[jian@laptop bigdata]$ mv apache-flume-1.9.0-bin flume
[jian@laptop bigdata]$ cd flume/
[jian@laptop flume]$ mkdir job
[jian@laptop flume]$ cd job
[jian@laptop job]$ cat netcat-flume-logger.conf
# Name the components on this agent
a1.sources = r1
a1.sinks = k1
a1.channels = c1
# Describe/configure the source
a1.sources.r1.type = netcat
a1.sources.r1.bind = localhost
a1.sources.r1.port = 44444
# Describe the sink
a1.sinks.k1.type = logger
# Use a channel which buffers events in memory
a1.channels.c1.type = memory
a1.channels.c1.capacity = 1000
a1.channels.c1.transactionCapacity = 100
# Bind the source and sink to the channel
a1.sources.r1.channels = c1
a1.sinks.k1.channel = c1
```

[官网文档](https://flume.apache.org/releases/content/1.9.0/FlumeUserGuide.html)


5.测试
```
[jian@laptop flume]$ bin/flume-ng agent --conf conf/ --conf-file job/netcat-flume-logger.conf --name a1 -Dflume.root.logger=INFO,console
```
```
[jian@laptop job]$ nc localhost 44444
hello world
OK
```
[更多例子](https://www.cnblogs.com/qingyunzong/p/8995554.html)