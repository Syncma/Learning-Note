# kafka和flume的区别

<!-- TOC -->

- [kafka和flume的区别](#kafka%e5%92%8cflume%e7%9a%84%e5%8c%ba%e5%88%ab)
  - [区别1](#%e5%8c%ba%e5%88%ab1)
  - [区别2](#%e5%8c%ba%e5%88%ab2)
  - [区别3](#%e5%8c%ba%e5%88%ab3)
  - [区别4](#%e5%8c%ba%e5%88%ab4)
  - [区别5](#%e5%8c%ba%e5%88%ab5)
  - [区别6](#%e5%8c%ba%e5%88%ab6)
  - [总结](#%e6%80%bb%e7%bb%93)

<!-- /TOC -->


## 区别1
kafka和flume都是日志系统

* kafka是分布式消息中间件，自带存储，提供push和pull存取数据功能。
* flume分为agent（数据采集器）[source channel sink]。



## 区别2
kafka做**日志缓存**应该是更为合适的，但是 flume的**数据采集部分**做的很好，可以定制很多数据源，减少开发量。

所以比较流行**`flume+kafka`**模式，如果为了利用flume写hdfs的能力，也可以采用kafka+flume的方式。

**采集层 主要可以使用Flume, Kafka两种技术。**

## 区别3
Flume：Flume 是管道流方式，提供了很多的默认实现，让用户通过参数部署，及扩展API.

Kafka：Kafka是一个可持久化的分布式的消息队列。

Kafka 是一个非常通用的系统。你可以有许多生产者和很多的消费者共享多个主题Topics。

相比之下,Flume是一个专用工具被设计为旨在往HDFS,HBase发送数据。
它对HDFS有特殊的优化，并且集成了Hadoop的安全特性。


所以:
> Cloudera 建议如果数据被多个系统消费的话，使用kafka
> 如果数据被设计给Hadoop使用，使用Flume。


## 区别4
fume可以使用**`拦截器`**实时处理数据。
这些对数据屏蔽或者过量是很有用的

Kafka需要外部的流处理系统才能做到。


## 区别5
Kafka和Flume都是可靠的系统,通过适当的配置能保证零数据丢失。然而，Flume不支持副本事件。

于是，如果Flume代理的一个节点奔溃了，即使使用了可靠的文件管道方式，你也将丢失这些事件直到你恢复这些磁盘。

**如果你需要一个高可靠行的管道，那么使用Kafka是个更好的选择。**


## 区别6
Flume和Kafka可以很好地结合起来使用。

如果你的设计需要从Kafka到Hadoop的流数据，使用Flume代理并配置Kafka的Source读取数据也是可行的

你没有必要实现自己的消费者。你可以直接利用Flume与HDFS及HBase的结合的所有好处。

你可以使用Cloudera Manager对消费者的监控，并且你甚至可以添加拦截器进行一些流处理。

Flume和Kafka可以结合起来使用，通常会使用Flume + Kafka的方式。

## 总结

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/flume-kafka.png)