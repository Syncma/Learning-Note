# MapReduce介绍
<!-- TOC -->

- [MapReduce介绍](#mapreduce%e4%bb%8b%e7%bb%8d)
  - [大数据本质](#%e5%a4%a7%e6%95%b0%e6%8d%ae%e6%9c%ac%e8%b4%a8)
  - [MapReduce瓶颈](#mapreduce%e7%93%b6%e9%a2%88)
  - [MapReduce优缺点](#mapreduce%e4%bc%98%e7%bc%ba%e7%82%b9)
    - [优点](#%e4%bc%98%e7%82%b9)
    - [缺点](#%e7%bc%ba%e7%82%b9)

<!-- /TOC -->

## 大数据本质

* 数据的存储：分布式文件系统   HDFS   Hadoop Distributed File System
* 数据的计算：分布式计算 MapReduce

如何解决大数据的计算？分布式计算

1.什么是**PageRank**?(MapReduce的问题来源)
* 内容待补充


2.MapReduce基础编程模型

**`把一个大任务拆分成小任务，再进行汇总`**

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/mapreduce.png)



## MapReduce瓶颈

* 计算机性能: CPU、磁盘、内存、网络
* IO操作


## MapReduce优缺点

### 优点

1.MR易于编程，简单实现了一些接口可以完成一个分布式程序
2.良好的扩展性，可以简单的增加机器来扩展计算能力
3.高容错性
4.适合PB级以上海量数据的离线处理


### 缺点
1.不擅长实时计算
2.不擅长流式计算
3.不擅长DAG有向图计算
4.每个mapreduce作业的结果都会写入磁盘，造成大量的磁盘IO，性能低下


**`这些缺点 spark全部解决了，后面说到spark会详细说明`**

