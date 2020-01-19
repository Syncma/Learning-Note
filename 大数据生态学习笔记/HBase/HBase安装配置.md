# HBase安装配置

<!-- TOC -->

- [HBase安装配置](#hbase安装配置)
    - [环境](#环境)
    - [源码安装](#源码安装)
        - [单机模式](#单机模式)
            - [配置过程](#配置过程)
        - [伪分布式模式](#伪分布式模式)
            - [安装过程](#安装过程)
        - [完全分布式模式](#完全分布式模式)
            - [配置过程](#配置过程-1)
            - [备注](#备注)
    - [docker安装](#docker安装)

<!-- /TOC -->

## 环境
Fedora 29 x64

这里使用 **`hbase-1.4.12-bin.tar.gz`** 稳定版本
[下载地址](https://mirrors.tuna.tsinghua.edu.cn/apache/hbase/stable/)

[安装文档](https://github.com/apachecn/hbase-doc-zh/blob/master/docs/1.md)


## 源码安装
1.解压

```
[jian@laptop tools]$ tar xf hbase-1.4.12-bin.tar.gz  -C ../
[jian@laptop hbase-1.4.12]$ pwd
/home/jian/prj/bigdata/hbase-1.4.12

[jian@laptop hbase-1.4.12]$ mkdir logs  //创建日志目录
```


2.修改配置文件：
```
[jian@laptop conf]$ echo $JAVA_HOME
/home/jian/prj/bigdata/jdk1.8.0_231

[jian@laptop conf]$ vi hbase-env.sh
# export JAVA_HOME=/usr/java/jdk1.8.0/
export JAVA_HOME=/home/jian/prj/bigdata/jdk1.8.0_231
#HBase使用自带的ZK, 不需要单独的ZK
export HBASE_MANAGES_ZK=true

[jian@laptop conf]$ pwd
/home/jian/prj/bigdata/hbase-2.2.2/conf
```

[jian@laptop conf]$ cat ~/.bashrc
增加下面的内容：
```
#HBase配置
export HBASE_HOME=/home/jian/prj/bigdata/hbase-1.4.12
export PATH=$HBASE_HOME/bin:$PATH
```

执行source命令使配置文件生效：
```
[jian@laptop conf]$ source ~/.bashrc
```




### 单机模式

* Hbase不使用HDFS,仅使用本地文件系统
* ZooKeeper与Hbase运行在同一个JVM中


#### 配置过程

1.修改配置文件：

建议修改 ${HBase-Dir}/conf/hbase-site.xml 文件，因为即使你修改了hbase-default.xml文件，也会被hbase-site.xml中的配置所覆盖。

也就是说，最终是以 hbase-site.xml 中的配置为准的


```
[jian@laptop hbase-2.2.2]$ cat conf/hbase-site.xml
<configuration>
    <property>
        <name>hbase.rootdir</name>
        <value>file:///tmp/hbase-jian/hbase</value>
    </property>
    <property>
        <name>hbase.zookeeper.property.dataDir</name>
        <value>/hbase-jian/zookeeper</value>
    </property>
    <property>
        <name>hbase.unsafe.stream.capability.enforce</name>
        <value>false</value>
        <description>
            Controls whether HBase will check for stream capabilities (hflush/hsync).
            Disable this if you intend to run on LocalFileSystem, denoted by a rootdir
            with the 'file://' scheme, but be mindful of the NOTE below.
            WARNING: Setting this to false blinds you to potential data loss and
            inconsistent system state in the event of process and/or node failures. If
            HBase is complaining of an inability to use hsync or hflush it's most
            likely not a false positive.
        </description>
    </property>
</configuration>
```

2.启动：

单机模式下不需要HDFS，因此不需要事先启动Hadoop，直接启动HBase即可。

终端下输入命令：./start-hbase.sh
```
[jian@laptop hbase-2.2.2]$ bin/start-hbase.sh
```
3.查看进程：

```
[jian@laptop hbase-2.2.2]$ jps
11024 DataNode
11714 NodeManager
23091 HMaster
10744 NameNode
23691 Jps
11292 SecondaryNameNode
23279 HRegionServer
11535 ResourceManager
```

```
[jian@laptop hbase-1.4.12]$ hbase shell
SLF4J: Class path contains multiple SLF4J bindings.
SLF4J: Found binding in [jar:file:/home/jian/prj/bigdata/hbase-1.4.12/lib/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: Found binding in [jar:file:/home/jian/prj/bigdata/hadoop-2.10.0/share/hadoop/common/lib/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: See http://www.slf4j.org/codes.html#multiple_bindings for an explanation.
SLF4J: Actual binding is of type [org.slf4j.impl.Log4jLoggerFactory]
HBase Shell
Use "help" to get list of supported commands.
Use "exit" to quit this interactive shell.
Version 1.4.12, r6ae4a77408ad35d6a7a4e5cebfd401fc4b72b5ec, Sun Nov 24 13:25:41 CST 2019
hbase(main):001:0> status
1 active master, 0 backup masters, 1 servers, 0 dead, 2.0000 average load
```

4.服务关闭：
```
[jian@laptop hbase-2.2.2]$ bin/stop-hbase.sh
```

### 伪分布式模式

* 所有进程运行在同一个节点上,不同进程运行在不同的JVM当中
* 比较适合实验测试

伪分布式模式下，HBase的所有组件还是运行在同一台主机，不同的是，每个组件独立运行在不同的JVM

伪分布模式是一个运行在单台机器上的分布式模式。

此模式下，HBase所有的守护进程将运行在同一个节点之上，而且需要依赖HDFS
因此在此之前必须保证**HDFS**已经成功运行

#### 安装过程

1.修改配置文件
```
[jian@laptop conf]$ cat hbase-site.xml
<?xml version="1.0"?>
<?xml-stylesheet type="text/xsl" href="configuration.xsl"?>
<!--
/**
*
* Licensed to the Apache Software Foundation (ASF) under one
* or more contributor license agreements. See the NOTICE file
* distributed with this work for additional information
* regarding copyright ownership. The ASF licenses this file
* to you under the Apache License, Version 2.0 (the
* "License"); you may not use this file except in compliance
* with the License. You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/
-->
<configuration>
    <property>
        <name>hbase.rootdir</name>
        <value>hdfs://localhost:9000/hbase</value>
    </property>
    <property>
        <name>hbase.cluster.distributed</name>
        <value>true</value>
    </property>
    <property>
        <name>hbase.zookeeper.property.dataDir</name>
        <value>/tmp/hbase-jian/zookeeper</value>
    </property>
    <property>
        <name>hbase.unsafe.stream.capability.enforce</name>
        <value>false</value>
    </property>
</configuration>
```
2.启动HDFS

继续执行：
```
[jian@laptop hadoop-2.10.0]$ sbin/start-all.sh
This script is Deprecated. Instead use start-dfs.sh and start-yarn.sh
Starting namenodes on [localhost]
localhost: starting namenode, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/hadoop-jian-namenode-laptop.out
localhost: starting datanode, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/hadoop-jian-datanode-laptop.out
Starting secondary namenodes [0.0.0.0]
0.0.0.0: starting secondarynamenode, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/hadoop-jian-secondarynamenode-laptop.out
starting yarn daemons
starting resourcemanager, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/yarn-jian-resourcemanager-laptop.out
localhost: starting nodemanager, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/yarn-jian-nodemanager-laptop.out
```



3.启动HBase
```
[jian@laptop hbase-2.2.2]$ bin/start-hbase.sh
```

4.查看进程
```
[jian@laptop logs]$ jps
26272 HRegionServer
11024 DataNode
11714 NodeManager
26613 Jps
25990 HQuorumPeer
10744 NameNode
26092 HMaster
11292 SecondaryNameNode
11535 ResourceManager
```

5.进入 hbase shell模式

```
[jian@laptop logs]$ hbase shell
SLF4J: Class path contains multiple SLF4J bindings.
SLF4J: Found binding in [jar:file:/home/jian/prj/bigdata/hadoop-2.10.0/share/hadoop/common/lib/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: Found binding in [jar:file:/home/jian/prj/bigdata/hbase-2.2.2/lib/client-facing-thirdparty/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: See http://www.slf4j.org/codes.html#multiple_bindings for an explanation.
SLF4J: Actual binding is of type [org.slf4j.impl.Log4jLoggerFactory]
HBase Shell
Use "help" to get list of supported commands.
Use "exit" to quit this interactive shell.
For Reference, please visit: http://hbase.apache.org/2.0/book.html#shell
Version 2.2.2, re6513a76c91cceda95dad7af246ac81d46fa2589, Sat Oct 19 10:10:12 UTC 2019
Took 0.0036 seconds
hbase(main):001:0> list
TABLE
0 row(s)
Took 0.3299 seconds
=> []
```

### 完全分布式模式
* 进程运行在多个服务器集群中
* 分布式依赖于HDFS系统，因此布署Hbase之前一定要有一个正常工作的HDFS集群
* 不使用Hbase自带的ZK



#### 配置过程
1.开启zk

```
[jian@laptop zookeeper-3.4.14]$ bin/zkServer.sh start
ZooKeeper JMX enabled by default
Using config: /home/jian/prj/bigdata/zookeeper-3.4.14/bin/../conf/zoo.cfg
Starting zookeeper ... STARTED
```

```
[jian@laptop conf]$ vi hbase-env.sh
#HBase不使用自带的ZK, 需要单独的ZK
export HBASE_MANAGES_ZK=false
```

2.修改配置文件：

```
<configuration>
    <property>
        <name>hbase.rootdir</name>
        <value>hdfs://localhost:9000/hbase</value>
    </property>
    <property>
        <name>hbase.cluster.distributed</name>
        <value>true</value>
    </property>
    <property>
        <name>hbase.zookeeper.quorum</name>
        <value>localhost</value>
    </property>
    <property>
        <name>zookeeper.znode.parent</name>
        <value>/hbase-test</value>
    </property>
    <property>
        <name>hbase.zookeeper.property.dataDir</name>
        <value>/home/jian/prj/bigdata/zookeeper-3.4.14/zkData</value>
    </property>
    <property>
        <name>hbase.unsafe.stream.capability.enforce</name>
        <value>false</value>
    </property>
</configuration>
```



#### 备注
hbase 会自动在ZK上创建/hbase目录，所以不需要手动创建的



3.开启Hadoop和ZK服务：
```
[jian@laptop bin]$ jps
28240 Jps
11024 DataNode
11714 NodeManager
10744 NameNode
11292 SecondaryNameNode
28207 QuorumPeerMain
11535 ResourceManager
```

4.Hbase开启：
```
[jian@laptop hbase-1.4.12]$ bin/start-hbase.sh
```

5.查看进程：
```
[jian@laptop bin]$ jps
11024 DataNode
11714 NodeManager
30771 HRegionServer
31125 Jps
10744 NameNode
30585 HMaster
11292 SecondaryNameNode
28207 QuorumPeerMain
11535 ResourceManager
```
```
[jian@laptop hbase-1.4.12]$ hbase shell
SLF4J: Class path contains multiple SLF4J bindings.
SLF4J: Found binding in [jar:file:/home/jian/prj/bigdata/hbase-1.4.12/lib/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: Found binding in [jar:file:/home/jian/prj/bigdata/hadoop-2.10.0/share/hadoop/common/lib/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: See http://www.slf4j.org/codes.html#multiple_bindings for an explanation.
SLF4J: Actual binding is of type [org.slf4j.impl.Log4jLoggerFactory]
HBase Shell
Use "help" to get list of supported commands.
Use "exit" to quit this interactive shell.
Version 1.4.12, r6ae4a77408ad35d6a7a4e5cebfd401fc4b72b5ec, Sun Nov 24 13:25:41 CST 2019
hbase(main):001:0> list
TABLE
0 row(s) in 0.1490 seconds
=> []
```

6.ZK查看：
```
[jian@laptop zookeeper-3.4.14]$ bin/zkCli.sh
Connecting to localhost:2181
xxxx
[zk: localhost:2181(CONNECTED) 2] get /hbase-test
cZxid = 0x1f5
ctime = Sun Dec 01 19:57:14 CST 2019
mZxid = 0x1f5
mtime = Sun Dec 01 19:57:14 CST 2019
pZxid = 0x256
cversion = 24
dataVersion = 0
aclVersion = 0
ephemeralOwner = 0x0
dataLength = 0
numChildren = 18
```

7.Hadoop查看：

```
[jian@laptop bin]$ hdfs dfs -fs hdfs://localhost:9000/ -ls /
Found 1 items
drwxr-xr-x - jian supergroup 0 2019-12-01 20:01 /hbase
[jian@laptop bin]$ hdfs dfs -fs hdfs://localhost:9000/ -ls /hbase
Found 7 items
drwxr-xr-x - jian supergroup 0 2019-12-01 20:01 /hbase/.tmp
drwxr-xr-x - jian supergroup 0 2019-12-01 20:01 /hbase/MasterProcWALs
drwxr-xr-x - jian supergroup 0 2019-12-01 20:01 /hbase/WALs
drwxr-xr-x - jian supergroup 0 2019-12-01 19:47 /hbase/data
-rw-r--r-- 1 jian supergroup 42 2019-12-01 19:47 /hbase/hbase.id
-rw-r--r-- 1 jian supergroup 7 2019-12-01 19:47 /hbase/hbase.version
drwxr-xr-x - jian supergroup 0 2019-12-01 20:00 /hbase/oldWALs
```

浏览器可以输入进行查看： [访问这里](http://localhost:16010)

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/Hbase-web.png)


## docker安装
* 内容待补充
