# kafka消费者组
<!-- TOC -->

- [kafka消费者组](#kafka%e6%b6%88%e8%b4%b9%e8%80%85%e7%bb%84)
  - [消费者组](#%e6%b6%88%e8%b4%b9%e8%80%85%e7%bb%84)
  - [说明](#%e8%af%b4%e6%98%8e)
  - [例子](#%e4%be%8b%e5%ad%90)

<!-- /TOC -->
## 消费者组

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/kafka-group.png)



## 说明
消费者是以consumer group消费者的方式工作的，由一个或者多个消费者来组成一个组，共同消费一个topic

每个分区在同一个时间只能由group中的一个消费者读取，但是多个group可以同时消费这个partition

上图所示：

有一个由三个消费者组成的group, 有一个消费者读取主题中的两个分区，
另外两个分别读取一个分区

某个消费者读取某个分区，也可以叫做某个消费者是某个分区的拥有者

在这种情况下，消费者可以通过水平扩展的方式同时读取大量的消息
另外，如果一个消费者失败了，那么其他的group成员会自动负载均衡读取之前失败的消费者读取的分区


## 例子
```
[root@laptop config]# cat consumer.properties  |grep group.id
# consumer group id
group.id=test-consumer-group


[root@laptop bin]#  ./kafka-console-producer.sh --broker-list localhost:9092 --topic first
[root@laptop bin]#  ./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic first --consumer.config ../config/consumer.properties

```