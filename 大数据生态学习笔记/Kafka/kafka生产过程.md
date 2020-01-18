# kafka生产过程
<!-- TOC -->

- [kafka生产过程](#kafka%e7%94%9f%e4%ba%a7%e8%bf%87%e7%a8%8b)
  - [写入方式](#%e5%86%99%e5%85%a5%e6%96%b9%e5%bc%8f)
    - [问题1](#%e9%97%ae%e9%a2%981)
    - [问题2](#%e9%97%ae%e9%a2%982)
  - [分区Partition](#%e5%88%86%e5%8c%bapartition)
  - [副本 Replication](#%e5%89%af%e6%9c%ac-replication)
  - [存储方式](#%e5%ad%98%e5%82%a8%e6%96%b9%e5%bc%8f)

<!-- /TOC -->

## 写入方式

Producer 采用推（push）模式将消息发不到broker, 每条消息被追加到分区(Partition)中，
属于顺序写磁盘

顺序写磁盘比随机写内存要高，保障Kafka吞吐率


### 问题1
**为什么要使用push的方法？ pull/push方法区别在哪里？customer是用pull还是push呢？**

在这方面，Kafka遵循了一种大部分消息系统共同的传统的设计：
**producer将消息推送到broker，consumer从broker拉取消息。**

一些消息系统比如Scribe和Apache Flume采用了push模式，将消息推送到下游的consumer。

这样做有好处也有坏处：
由broker决定消息推送的速率，对于不同消费速率的consumer就不太好处理了。

消息系统都致力于让consumer以最大的速率最快速的消费消息，但不幸的是，push模式下，当broker推送的速率远大于consumer消费的速率时，consumer恐怕就要崩溃了。

最终Kafka还是选取了传统的pull模式。

>Pull模式的另外一个好处是consumer可以自主决定是否批量的从broker拉取数据。

>Push模式必须在不知道下游consumer消费能力和消费策略的情况下决定是立即推送每条消息还是缓存之后批量推送。

如果为了避免consumer崩溃而采用较低的推送速率，将可能导致一次只推送较少的消息而造成浪费。

>Pull模式下，consumer就可以根据自己的消费能力去决定这些策略。

>Pull有个缺点是，如果broker没有可供消费的消息，将导致consumer不断在循环中轮询，直到新消息到达。

为了避免这点，Kafka有个参数可以让consumer阻塞知道新消息到达(当然也可以阻塞知道消息的数量达到某个特定的量这样就可以批量发)





### 问题2

我们知道内存的速度一定比磁盘快，这是肯定的

**那么`顺序写磁盘比随机写内存要高？`**

[参考1](https://cloud.tencent.com/developer/article/1448153)

[参考2](https://juejin.im/post/5b3af22bf265da62bd0dec2d)




## 分区Partition

消息发送都被发送到一个topic, 其本质就是一个目录，而topic是一些Partition Logs分区日志组成的

每个Partition消息都是有序的，生产的消息不断追加到Partition log上，每个消息都有唯一的offset值


![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/1_ipMuwhEg-LO6wBCy1jlkpg.png)


为什么要分区？

```
1.方便在集群中扩展，每个Partition可以通过调整以适应它所在的机器，
而一个topic又可以有多个Partition组成，因此整个集群就可以适应任意大小的数据了

2.可以提高并发，因为可以以Partition为单位读写了
```

分区的原则是什么?
```
1. 指定Partition, 直接使用
2. 未指定Partition 但指定key, 通过对key 的value进行hash出一个Partition
3. Partition和key都未指定，则使用轮询选出一个partition
```




## 副本 Replication

同一个partition 可能会有多个replication, 对应server.properties配置中的
**"default.replication.factor=N"**

在没有replication的情况下，一旦broker宕机，其上所有partition的数据都不可被消费

同时producer 也不能再将数据存在其上的patition

引入replication 之后，同一个partition可能会有多个replication
这时需要在replication之间选出一个leader
producer和consumer 只要跟这个leader交互，其他replication作为follower 从leader中复制数据


## 存储方式

物理上吧topic 分成一个或多个partition(server.properties中的num.partition=3配置），
每个partition物理上对应一个文件夹，存储partition所有的消息和索引文件

```
[root@laptop logs]# ll |grep first
drwxr-xr-x 2 root root   4096 Nov  8 09:41 first-0
[root@laptop logs]# cd first-0/
[root@laptop first-0]# ls
00000000000000000000.index  00000000000000000000.log  00000000000000000000.timeindex  leader-epoch-checkpoint
[root@laptop first-0]#

// xxx.log 就是实际数据
// 存储时间？  server.properties里面有配置
```


