# spark安装
<!-- TOC -->

- [spark安装](#spark%e5%ae%89%e8%a3%85)
  - [Spark部署模式](#spark%e9%83%a8%e7%bd%b2%e6%a8%a1%e5%bc%8f)
    - [单机模式](#%e5%8d%95%e6%9c%ba%e6%a8%a1%e5%bc%8f)
    - [Standalone模式](#standalone%e6%a8%a1%e5%bc%8f)
    - [Yarn模式](#yarn%e6%a8%a1%e5%bc%8f)
    - [Mesos模式](#mesos%e6%a8%a1%e5%bc%8f)

<!-- /TOC -->

## Spark部署模式
Spark部署模式主要是四种：

* **`Local模式`**（单机模式，是本文讲的方式，仅供熟悉Spark和scala入门用）
* **`Standalone模式`**（使用Spark自带的简单集群管理器,计算数据不是特别庞大）
* **`YARN模式`**（使用YARN作为集群管理器，配合hadoop集群使用）-通用最广
* **`Mesos模式`**（使用Mesos作为集群管理器，配合docker）-性能匹配最好



### 单机模式

1.解压
```
[jian@laptop tools]$ gunzip spark-2.4.4-bin-hadoop2.7.tgz
[jian@laptop tools]$ tar xf spark-2.4.4-bin-hadoop2.7.tar -C ../
[jian@laptop tools]$ pwd
/home/jian/prj/bigdata/tools
[jian@laptop bigdata]$ mv spark-2.4.4-bin-hadoop2.7/ spark
[jian@laptop bigdata]$ pwd
/home/jian/prj/bigdata
```

2.编辑配置文件
```
[jian@laptop conf]$ cp spark-env.sh.template spark-env.sh
[jian@laptop conf]$ pwd
/home/jian/prj/bigdata/spark/conf
[jian@laptop conf]$ cat spark-env.sh
在末尾新增下面内容：
SPARK_LOCAL_IP=localhost
```

3.启动

测试：
```
[jian@laptop spark]$ bin/run-example SparkPi 2>&1 |grep "Pi is roughly"
Pi is roughly 3.144115720578603
```
到bin目录执行：
```
[jian@laptop bin]$ ./spark-shell
xxxx
Using Scala version 2.11.12 (Java HotSpot(TM) 64-Bit Server VM, Java 1.8.0_231)
Type in expressions to have them evaluated.
Type :help for more information.
scala>
```


4.浏览器可以访问 [访问这里](http://localhost:4040)

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/spark-web.png)






测试小例子：

在/tmp目录下准备README.md文件，然后到spark命令行输入下面命令：
```
[jian@laptop tmp]$ cat README.md
1111
1111
1111
1111
1111
1111
1111
1111
1111
1111
```
```
scala> var lines = sc.textFile("/tmp/README.md")
lines: org.apache.spark.rdd.RDD[String] = /tmp/README.md MapPartitionsRDD[3] at textFile at <console>:24
scala> lines.count()
res1: Long = 10
```




### Standalone模式

分布式集群模式:

**Master-Worker架构，Master负责调度，Worker负责具体Task的执行**

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200120142911.png)


1.编辑配置文件
```
[jian@laptop conf]$ cat spark-env.sh
在末尾新增下面内容：
SPARK_MASTER_HOST=localhost
```
```
[jian@laptop conf]$ cp slaves.template slaves
默认已经是localhost
# A Spark Worker will be started on each of the machines listed below.
localhost
```

2.服务启动
```
[jian@laptop spark]$ sbin/start-all.sh
查看进程：
[jian@laptop spark]$ jps
24608 Master
24858 Jps
24767 Worker

```
可以在浏览器输入     [访问这里](http://localhost:8080)

3.测试例子
```
[jian@laptop spark]$ bin/spark-submit --class org.apache.spark.examples.SparkPi --master spark://localhost:7077 ./examples/jars/spark-examples_2.11-2.4.4.jar 100
xxx
Pi is roughly 3.1415339141533916
xxxx
```

启动spark-shell:
```
[jian@laptop spark]$ bin/spark-shell --master spark://localhost:7077
```

4.浏览器可以访问 [访问这里](http://localhost:7070)
![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/spark-web2.png)


### Yarn模式

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200120143511.png)

Spark 客户端可以直接连接 Yarn，不需要额外构建Spark集群

有 yarn-client 和 yarn-cluster 两种模式，主要区别在于：Driver 程序的运行节点不同

* yarn-client：Driver程序运行在客户端，适用于交互、调试，希望立即看到app的输出
*  yarn-cluster：Driver程序运行在由 RM（ResourceManager）启动的 AM（AplicationMaster）上, 适用于生产环境。

Yarn的ResourceManager就相对于Spark Standalone模式下的Master

我们启动spark集群是要用到standalone，现在有yarn了，就不用spark集群了

[参考内容](https://blog.csdn.net/huojiao2006/article/details/80563112)


1.修改配置文件

```
[jian@laptop spark]$ cat conf/spark-env.sh
增加：
YARN_CONF_DIR=/home/jian/prj/bigdata/hadoop-2.10.0/etc/hadoop
```

2.服务开启

首先开启Hadoop, yarn服务
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

#关闭安全模式
[jian@laptop hadoop-2.10.0]$ hadoop dfsadmin -safemode leave
DEPRECATED: Use of this script to execute hdfs command is deprecated.
Instead use the hdfs command for it.
Safe mode is OFF

#关闭standlone模式的服务
[jian@laptop spark]$ sbin/stop-all.sh


#测试


[jian@laptop spark]$ bin/spark-shell --master yarn



[jian@laptop spark]$ bin/spark-submit --class org.apache.spark.examples.SparkPi --master yarn --deploy-mode client ./examples/jars/spark-examples_2.11-2.4.4.jar 100
Failing this attempt.Diagnostics: [2019-12-03 16:02:41.720]Container [pid=13426,containerID=container_1575360080350_0001_02_000001] is running beyond virtual memory limits. Current usage: 360.9 MB of 1 GB physical memory used; 2.5 GB of 2.1 GB virtual memory used. Killing container.
Dump of the process-tree for container_1575360080350_0001_02_000001 :
xxx
19/12/03 13:02:40 ERROR client.TransportClient: Failed to send RPC RPC 8185245102213421039 to /192.168.43.19:51962: java.nio.channels.ClosedChannelException
java.nio.channels.ClosedChannelException
at io.netty.channel.AbstractChannel$AbstractUnsafe.write(...)(Unknown Source)
19/12/03 13:02:40 INFO storage.BlockManagerMasterEndpoint: Trying to remove executor 1 from BlockManagerMaster.
19/12/03 13:02:40 WARN cluster.YarnSchedulerBackend$YarnSchedulerEndpoint: Attempted to get executor loss reason for executor id 1 at RPC address 192.168.43.19:51970, but got no response. Marking as slave lost.
java.io.IOException: Failed to send RPC RPC 8185245102213421039 to /192.168.43.19:51962: java.nio.channels.ClosedChannelException
at org.apache.spark.network.client.TransportClient$RpcChannelListener.handleFailure(TransportClient.java:362)
at org.apache.spark.network.client.TransportClient$StdChannelListener.operationComplete(TransportClient.java:339)
at io.netty.util.concurrent.DefaultPromise.notifyListener0(DefaultPromise.java:507)
at io.netty.util.concurrent.DefaultPromise.notifyListenersNow(DefaultPromise.java:481)
at io.netty.util.concurrent.DefaultPromise.notifyListeners(DefaultPromise.java:420)
at io.netty.util.concurrent.DefaultPromise.tryFailure(DefaultPromise.java:122)
at io.netty.channel.AbstractChannel$AbstractUnsafe.safeSetFailure(AbstractChannel.java:987)
at io.netty.channel.AbstractChannel$AbstractUnsafe.write(AbstractChannel.java:869)
at io.netty.channel.DefaultChannelPipeline$HeadContext.write(DefaultChannelPipeline.java:1316)
at io.netty.channel.AbstractChannelHandlerContext.invokeWrite0(AbstractChannelHandlerContext.java:738)
at io.netty.channel.AbstractChannelHandlerContext.invokeWrite(AbstractChannelHandlerContext.java:730)
at io.netty.channel.AbstractChannelHandlerContext.access$1900(AbstractChannelHandlerContext.java:38)
at io.netty.channel.AbstractChannelHandlerContext$AbstractWriteTask.write(AbstractChannelHandlerContext.java:1081)
at io.netty.channel.AbstractChannelHandlerContext$WriteAndFlushTask.write(AbstractChannelHandlerContext.java:1128)
at io.netty.channel.AbstractChannelHandlerContext$AbstractWriteTask.run(AbstractChannelHandlerContext.java:1070)
at io.netty.util.concurrent.AbstractEventExecutor.safeExecute(AbstractEventExecutor.java:163)
at io.netty.util.concurrent.SingleThreadEventExecutor.runAllTasks(SingleThreadEventExecutor.java:403)
at io.netty.channel.nio.NioEventLoop.run(NioEventLoop.java:463)
at io.netty.util.concurrent.SingleThreadEventExecutor$5.run(SingleThreadEventExecutor.java:858)
at io.netty.util.concurrent.DefaultThreadFactory$DefaultRunnableDecorator.run(DefaultThreadFactory.java:138)
at java.lang.Thread.run(Thread.java:748)
Caused by: java.nio.channels.ClosedChannelException
at io.netty.channel.AbstractChannel$AbstractUnsafe.write(...)(Unknown Source)
xxxx
```

从log可以看出，这是RPC通信的问题，RPC无法建立连接。

报错信息解决：

修改yarn配置文件：
```
[jian@laptop hadoop]$ vi yarn-site.xml
[jian@laptop hadoop]$ pwd
/home/jian/prj/bigdata/hadoop-2.10.0/etc/hadoop
新增下面内容：
<!--错误忽略-->
<property>
<name>yarn.nodemanager.pmem-check-enabled</name>
<value>false</value>
</property>
<property>
<name>yarn.nodemanger.vmem-check-enabled</name>
<value>false</value>
</property>
```
重启服务：
```
[jian@laptop hadoop-2.10.0]$ sbin/stop-all.sh
[jian@laptop hadoop-2.10.0]$ sbin/start-all.sh
[jian@laptop hadoop-2.10.0]$ jps
19125 Jps
18329 SecondaryNameNode
17849 NameNode
18571 ResourceManager
18748 NodeManager
18062 DataNode

---未解决-----
```


### Mesos模式

Spark客户端直接连接 Mesos；不需要额外构建 Spark 集群。
国内应用比较少，更多的是运用yarn调度。

运行模式对比：

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200120143607.png)



