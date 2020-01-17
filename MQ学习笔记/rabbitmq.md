# rabbitmq
<!-- TOC -->

- [rabbitmq](#rabbitmq)
    - [介绍](#介绍)
        - [安装](#安装)
        - [服务开启](#服务开启)
        - [例子](#例子)
        - [web界面](#web界面)
    - [组件介绍](#组件介绍)
    - [消息流程分析](#消息流程分析)
        - [生产者发送消息](#生产者发送消息)
        - [消费者接收消息](#消费者接收消息)
    - [原理解析](#原理解析)

<!-- /TOC -->

## 介绍

rabbitmq，发现是用erlang写的

rabbimq是用来提供发送消息的服务，可以用在不同的应用程序之间进行通信

[官网地址](https://www.rabbitmq.com/)

### 安装
Fedora 平台安装：
```
[root@laptop ~]# dnf install rabbitmq-server
```

### 服务开启
```
[root@laptop ~]# systemctl start rabbitmq-server.service
```

### 例子

python使用rabbitmq服务，可以使用现成的类库pika


produce.py
```python
import pika

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

channel.queue_declare(queue='hello')

channel.basic_publish(exchange='', routing_key='hello', body='Hello World!')
print("[x] Sent 'Hello World!'")

connection.close()
```


receive.py
```python
import pika

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

channel.queue_declare(queue='hello')


def callback(ch, method, properties, body):
    print("[x] Received %r" % (body, ))


channel.basic_consume('hello', callback, auto_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()
```

### web界面

使用root用户开启web界面功能
```
[root@laptop ~]# rabbitmq-plugins enable rabbitmq_management
```
然后浏览器访问 [这里](http://localhost:15672)，默认账户和密码都是guest


## 组件介绍

![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/rabbitmq_example.png)




概念：
```
channel:通道，amqp支持一个tcp连接上启用多个mq通信通道，每个通道都可以被作为通信流。

producer：生产者，是消息产生的源头。

exchange：交换机，可以理解为具有路由表的路由规则。

queues：队列，装载消息的缓存容器。

consumer：消费者，连接到队列并取走消息的客户端。

核心思想：在RabbitMQ中，生产者从不直接将消息发送给队列。

事实上，有些生产者甚至不知道消息是否被送到某个队列中去了。生产者只负责将消息送给交换机，而交换机确切地知道什么消息应该送到哪。


RoutingKey：路由键，生产者将消息发给交换器的时候会指定一个路由键，用来指定路由规则
Binding：绑定，RabbitMQ通过绑定将交换器与队列关联起来，绑定时会指定BindingKey
Connection：连接，生产者或消费者和Broker之间的一条TCP连接

bind：绑定，实际上可以理解为交换机的路由规则。每个消息都有一个称为路由键的属性(routing key)，就是一个简单的字符串。一个绑定将【交换机，路由键，消息送达队列】三者绑定在一起，形成一条路由规则。

exchange type：交换机类型：

fanout：不处理路由键，转发到所有绑定的队列上
direct：处理路由键，必须完全匹配，即路由键字符串相同才会转发
topic：路由键模式匹配，此时队列需要绑定要一个模式上。


```

## 消息流程分析

### 生产者发送消息

```
生产者连接到 RabbitMQ Broker，建立一个连接（Connection），开启一个信道（Channel）
生产者声明一个交换器，并设置相关属性（交换机类型、持久化）
生产者声明一个队列，并设置相关属性（排他、持久化、自动删除）
生产者通过路由键将交换器和队列绑定起来
生产者发消息至RabbitMQ Broker（包含路由键、交换器信息）
相应的交换器根据接收到的路由键查找相匹配的队列
若找到队列，则把消息存入；若没有，则丢弃或者回退给生产者
关闭信道
关闭连接
```

### 消费者接收消息
```
消费者连接到 RabbitMQ Broker ，建立一个连接（Connection），开启一个信道（Channel）
消费者向 RabbitMQ Broker 请求消费对应队列中的消息
等待 RabbitMQ Broker 回应并投递相应队列中的消息，消费者接收消息
消费者确认（ack）接收到的消息
RabbitMQ 从队列中删除相应已经被确认的消息
关闭信道
关闭连接

```


## 原理解析

* 内容待补充