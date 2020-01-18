# kafka架构
<!-- TOC -->

- [kafka架构](#kafka%e6%9e%b6%e6%9e%84)
  - [架构图](#%e6%9e%b6%e6%9e%84%e5%9b%be)
  - [名词解释](#%e5%90%8d%e8%af%8d%e8%a7%a3%e9%87%8a)
    - [Producer](#producer)
    - [Consumer](#consumer)
    - [Topic](#topic)
    - [Broker](#broker)
    - [Consumer Group](#consumer-group)
      - [分区介绍](#%e5%88%86%e5%8c%ba%e4%bb%8b%e7%bb%8d)
      - [为什么要设计分区](#%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e8%ae%be%e8%ae%a1%e5%88%86%e5%8c%ba)
    - [Offset](#offset)
    - [Topic &amp; Partition 关系](#topic-amp-partition-%e5%85%b3%e7%b3%bb)
  - [总结](#%e6%80%bb%e7%bb%93)

<!-- /TOC -->


## 架构图

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/kafka.png)

备注：
> **`同一个组里面多个消费者不能够同时消费同一个分区的数据`**
>一个分区只能被同一个消费者群组里面的一个消费者读取，但可以被不同消费者群组中所组成的多个消费者共同读取
>多个消费者群组中消费者共同读取同一个主题时，彼此之间互不影响。


## 名词解释

### Producer
消息生产者，就是向 kafka broker 发消息的客户端

### Consumer 
消息消费者，向 kafka broker 读取消息的客户端

consumer从broker拉取(pull)数据并进行处理



### Topic 
可以理解为一个队列，一个 Topic 又分为一个或多个分区

每条发布到Kafka集群的消息都有一个类别，这个**类别被称为Topic**

物理上不同Topic的消息分开存储，逻辑上一个Topic的消息虽然保存于一个或多个broker上但用户只需指定消息的Topic即可生产或消费数据而不必关心数据存于何处


### Broker 
一台 kafka 服务器就是一个 broker
一个集群由多个 broker 组成
一个 broker 可以容纳多个 topic
	

### Consumer Group
这是 kafka 用来实现一个 topic 消息的广播（发给所有的 consumer）和单播（发给任意一个 consumer）的手段。一个 topic 可以有多个 Consumer Group

每个Consumer属于一个特定的Consumer Group（可为每个Consumer指定group name，若不指定group name则属于默认的group）


###Partition
#### 分区介绍
为了实现扩展性，一个非常大的 topic 可以分布到多个 broker上，每个 partition 是一个有序的队列。

partition 中的每条消息都会被分配一个有序的id（offset）。

将消息发给 consumer，kafka 只保证按一个 partition 中的消息的顺序，不保证一个 topic 的整体（多个 partition 间）的顺序。

Parition是物理上的概念，每个Topic包含一个或多个Partition.

#### 为什么要设计分区

>分区对于 Kafka 集群的好处是：实现负载均衡。
>分区对于消费者来说，可以提高并发度，提高效率。

### Offset
kafka 的存储文件都是按照 offset.kafka 来命名，用 offset 做名字的好处是方便查找。

例如你想找位于 2049 的位置，只要找到 2048.kafka 的文件即可。
当然 the first offset 就是 00000000000.kafka。


消费者通过检查消息的偏移量 (offset) 来区分读取过的消息。
偏移量是一个不断递增的数值，在创建消息时，Kafka 会把它添加到其中，在给定的分区里，每个消息的偏移量都是唯一的。
消费者把每个分区最后读取的偏移量保存在 Zookeeper 或 Kafka 上，如果消费者关闭或者重启，它还可以重新获取该偏移量，以保证读取状态不会丢失。


### `Topic & Partition 关系`

**这个要重点掌握**

Topic在逻辑上可以被认为是一个queue，每条消费都必须指定它的Topic，可以简单理解为必须指明把这条消息放进哪个queue里。

为了使得Kafka的吞吐率可以线性提高，物理上把Topic分成一个或多个Partition，每个Partition在物理上对应一个文件夹，该文件夹下存储这个Partition的所有消息和索引文件


Kafka 的消息通过 Topics(主题) 进行分类，
一个主题可以被分为若干个 Partitions(分区)，
一个分区就是一个提交日志 (commit log)。

消息以追加的方式写入分区，然后以先入先出的顺序读取。

Kafka 通过分区来实现数据的冗余和伸缩性，分区可以分布在不同的服务器上，这意味着一个 Topic 可以横跨多个服务器，以提供比单个服务器更强大的性能。

由于一个 Topic 包含多个分区，因此无法在整个 Topic 范围内保证消息的顺序性，但可以保证消息在单个分区内的顺序性。

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/kafka-topic.png)




## 总结

kafka流行主要使因为：

1.数据的聚合，就是可以把产生数据的系统就做生产者，有多少生产者无所谓，只要建立对应的topic，把数据发送到这个topic中就可以，

2.高并发性，并支持分布式部署，在一个繁忙的系统中，产生的日志或者其它数据是非常庞大的，要对这些数据进行处理，必须要一个高吞吐量、低延迟的处理系统

Kafka正好满足这个需求，每个topic可以建立一个或多个分区，每个分区你可以简单理解为一个公路上的多个车道，每个车就是数据，因为车道多所以它可以加速数据的传输。

