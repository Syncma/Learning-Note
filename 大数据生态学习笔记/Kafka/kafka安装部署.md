# kafka安装部署
<!-- TOC -->

- [kafka安装部署](#kafka%e5%ae%89%e8%a3%85%e9%83%a8%e7%bd%b2)
  - [部署环境](#%e9%83%a8%e7%bd%b2%e7%8e%af%e5%a2%83)
  - [ZK&amp;kafka关系](#zkampkafka%e5%85%b3%e7%b3%bb)
  - [配置](#%e9%85%8d%e7%bd%ae)
  - [测试](#%e6%b5%8b%e8%af%95)

<!-- /TOC -->

## 部署环境
环境: Fedora 29 x64

Kafka是使用Java开发的应用程序，所以它可以运行在windows、MacOS和Linux等多种操作系统上。 

运行Zookeeper和Kafka需要Java运行时版本，所以在安装Zookeeper和Kafka之前，需要先安装Java环境。
```
[jian@laptop logs]$ java -version
openjdk version "11-ea" 2018-09-25
OpenJDK Runtime Environment (build 11-ea+28)
OpenJDK 64-Bit Server VM (build 11-ea+28, mixed mode, sharing)
```



## ZK&kafka关系
![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/2018-08-19-kafka-setup-linux-1.PNG)



**Kafka使用Zookeeper保存集群的元数据信息和消费者信息，**


## 配置

这里使用kafka自带的ZK:


1.[官网下载](https://kafka.apache.org/downloads)，选择“Binary”



2.安装
```
[jian@laptop kafka]# tar xf kafka_2.12-2.3.1.tgz
[jian@laptop kafka]# mv kafka_2.12-2.3.1/ kafka
[jian@laptop kafka]$ pwd
/home/jian/prj/bigdata/kafka

[jian@laptop kafka]# ll
total 52
drwxr-xr-x 3 root root  4096 Oct 18 08:12 bin
drwxr-xr-x 2 root root  4096 Oct 18 08:12 config
drwxr-xr-x 2 root root  4096 Nov  7 22:22 libs
-rw-r--r-- 1 root root 32216 Oct 18 08:10 LICENSE
-rw-r--r-- 1 root root  337 Oct 18 08:10 NOTICE
drwxr-xr-x 2 root root  4096 Oct 18 08:12 site-docs

[jian@laptop kafka]# mkdir logs  //这一步不需要创建，系统会自动创建logs目录
```
修改配置文件：
```

[jian@laptop config]# pwd
/home/jian/prj/bigdata/kafka/config

[jian@laptop config]# vi server.properties
broker.id=0#默认为0，在多节点下，每个节点的broker.id要不同 >
listeners = PLAINTEXT://127.0.0.1:9092
log.dirs=/home/jian/prj/bigdata/kafka/logs/
zookeeper.connect=localhost:2181  //在多节点下，要指定其他节点的zookooper.connect，以逗号分隔
```

开启服务：
```
[jian@laptop kafka]$ bin/zookeeper-server-start.sh -daemon  config/zookeeper.properties

[jian@laptop kafka]$ bin/kafka-server-start.sh -daemon config/server.properties
[jian@laptop kafka]$ pwd
/home/jian/prj/bigdata/kafka
```

检测服务是否开启：
```
[jian@laptop kafka]$ jps
19522 Kafka
13837 QuorumPeerMain
19630 Jps

[jian@laptop kafka]# netstat -tunlp |grep 2181
tcp6       0      0 :::2181                 :::*                    LISTEN      23405/java
```


## 测试

1.创建topic:

```
[jian@laptop kafka]# bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic testTopic
Created topic testTopic.
```

如果创建Topic出现下面的错误：
```
[jian@laptop bin]$ ./kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic testTopic
Error while executing topic command : Replication factor: 1 larger than available brokers: 0.
[2019-12-05 12:44:31,697] ERROR org.apache.kafka.common.errors.InvalidReplicationFactorException: Replication factor: 1 larger than available brokers: 0.
(kafka.admin.TopicCommand$)

//解决办法：查看logs/server.log日志看看有什么报错信息
```



2.查看当前环境下所有的topic:

```
[jian@laptop bin]$ pwd
/home/jian/prj/bigdata/kafka/bin

[jian@laptop bin]$ ./kafka-topics.sh --list --zookeeper localhost:2181
testTopic
```


3.运行一个producer:

```
[jian@laptop bin]# ./kafka-console-producer.sh --broker-list localhost:9092 --topic testTopic
>My name is jack
>
```

4.启动另外一个终端运行一个consumer:

```
[jian@laptop bin]# ./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic testTopic --from-beginning
//加上from-begining表示从一开始就获取，如果不加就取最新的
My name is jack
```

producer 与 consumer要启在两个控制台，在producer控制台下输入，consumer控制台就会输出


5.logs目录 记录kafka日志 以及数据

```
[jian@laptop logs]$ pwd
/home/jian/prj/bigdata/kafka/logs

[jian@laptop logs]$ ll
total 636
-rw-rw-r-- 1 jian jian 0 Dec 5 14:13 cleaner-offset-checkpoint
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-0
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-1
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-10
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-11
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-12
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-13
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-14
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-15
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-16
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-17
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-18
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-19
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-2
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-20
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-21
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-22
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-23
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-24
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-25
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-26
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-27
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-28
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-29
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-3
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-30
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-31
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-32
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-33
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-34
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-35
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-36
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-37
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-38
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-39
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-4
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-40
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-41
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-42
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-43
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-44
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-45
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-46
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-47
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-48
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-49
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-5
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-6
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-7
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-8
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 __consumer_offsets-9
-rw-rw-r-- 1 jian jian 13250 Dec 5 14:13 controller.log
-rw-rw-r-- 1 jian jian 0 Dec 5 14:13 kafka-authorizer.log
-rw-rw-r-- 1 jian jian 0 Dec 5 14:13 kafka-request.log
-rw-rw-r-- 1 jian jian 9931 Dec 5 14:13 kafkaServer-gc.log.0.current
-rw-rw-r-- 1 jian jian 142718 Dec 5 14:13 kafkaServer.out
-rw-rw-r-- 1 jian jian 172 Dec 5 14:13 log-cleaner.log
-rw-rw-r-- 1 jian jian 4 Dec 5 14:14 log-start-offset-checkpoint
-rw-rw-r-- 1 jian jian 54 Dec 5 14:13 meta.properties
-rw-rw-r-- 1 jian jian 1209 Dec 5 14:14 recovery-point-offset-checkpoint
-rw-rw-r-- 1 jian jian 1209 Dec 5 14:14 replication-offset-checkpoint
-rw-rw-r-- 1 jian jian 142718 Dec 5 14:13 server.log
-rw-rw-r-- 1 jian jian 105222 Dec 5 14:13 state-change.log
drwxrwxr-x 2 jian jian 4096 Dec 5 14:13 testTopic-0
```

**这里留个坑，思考**：

**> `为什么产生了50个offsets目录？`** 


系统默认生成topic： **`__consumer_offsets`** 

```
jian@laptop kafka]$ bin/kafka-topics.sh --zookeeper localhost:2181 --topic __consumer_offsets --describe
Topic:__consumer_offsets    PartitionCount:50    ReplicationFactor:1    Configs:segment.bytes=104857600,cleanup.policy=compact,compression.type=producer
Topic: __consumer_offsets    Partition: 0    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 1    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 2    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 3    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 4    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 5    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 6    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 7    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 8    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 9    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 10    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 11    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 12    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 13    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 14    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 15    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 16    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 17    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 18    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 19    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 20    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 21    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 22    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 23    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 24    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 25    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 26    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 27    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 28    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 29    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 30    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 31    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 32    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 33    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 34    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 35    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 36    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 37    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 38    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 39    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 40    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 41    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 42    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 43    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 44    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 45    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 46    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 47    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 48    Leader: 0    Replicas: 0    Isr: 0
Topic: __consumer_offsets    Partition: 49    Leader: 0    Replicas: 0    Isr: 0

```


6.再次查看所有的topic:

```
[jian@laptop bin]# ./kafka-topics.sh --list --zookeeper localhost:2181
__consumer_offsets   // 这是系统自动生成的
testTopic
```

7.查看某个topic的详情：
```
[jian@laptop bin]# ./kafka-topics.sh  --zookeeper localhost:2181 --describe --topic testTopic
Topic:testTopic PartitionCount:1 ReplicationFactor:1 Configs:
Topic: testTopic Partition: 0 Leader: 0 Replicas: 0 Isr: 0

//这几个字段会后面篇章详细说明
Isr: 作用是为了leader挂了 做选举用的
Replicas: 副本？
Leader:  跟Partition有关系
```


8.删除topic:
```
[jian@laptop bin]# ./kafka-topics.sh --delete --topic testTopic
Exception in thread "main" java.lang.IllegalArgumentException: Only one of --bootstrap-server or --zookeeper must be specified
at kafka.admin.TopicCommand$TopicCommandOptions.checkArgs(TopicCommand.scala:630)
at kafka.admin.TopicCommand$.main(TopicCommand.scala:50)
at kafka.admin.TopicCommand.main(TopicCommand.scala)
```

正确删除应该使用这个命令：
```
[jian@laptop bin]# ./kafka-topics.sh --delete --bootstrap-server localhost:9092 --topic testTopic
```

超过了限制就会报错：
```
[root@laptop bin]# ./kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 5 --partitions 1 --topic second
Error while executing topic command : Replication factor: 5 larger than available brokers: 1.
[2019-11-08 09:45:11,783] ERROR org.apache.kafka.common.errors.InvalidReplicationFactorException: Replication factor: 5 larger than available brokers: 1.
(kafka.admin.TopicCommand$)


--topic 定义topic 名
--replication-factor 定义副本数
--partitions 定义分区数
```