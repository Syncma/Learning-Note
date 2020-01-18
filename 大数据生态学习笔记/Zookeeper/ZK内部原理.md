# ZK内部原理

<!-- TOC -->

- [ZK内部原理](#zk%e5%86%85%e9%83%a8%e5%8e%9f%e7%90%86)
  - [选举机制](#%e9%80%89%e4%b8%be%e6%9c%ba%e5%88%b6)
  - [ZK节点类型](#zk%e8%8a%82%e7%82%b9%e7%b1%bb%e5%9e%8b)

<!-- /TOC -->
## 选举机制

1）半数机制，集群中半数以上机器存活，集群可用； 所以ZK适合安装奇数台服务器

2）ZK 虽然在配置文件中没有指定Master和Slave,但是ZK工作时， 是有一个节点为Leader, 其他为Follower

Leader是通过内部的选举机制临时产生的
![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200118214415.png)



## ZK节点类型

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/zk-leixing.png)