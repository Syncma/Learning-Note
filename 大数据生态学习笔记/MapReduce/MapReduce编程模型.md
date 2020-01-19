# MapReduce编程模型
<!-- TOC -->

- [MapReduce编程模型](#mapreduce%e7%bc%96%e7%a8%8b%e6%a8%a1%e5%9e%8b)
  - [模型图](#%e6%a8%a1%e5%9e%8b%e5%9b%be)
  - [组件说明](#%e7%bb%84%e4%bb%b6%e8%af%b4%e6%98%8e)
  - [总结](#%e6%80%bb%e7%bb%93)

<!-- /TOC -->


## 模型图

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/mapreduce-2.png)


## 组件说明

MapReduce是一种编程模型，是一种编程方法

* HDFS:  
Hadoop的分布式文件系统，为MapReduce提供数据源和Job信息存储。

* Client Node: 
执行MapReduce程序的进程，用来提交MapReduce Job。

* JobTracker Node: 
把完整的Job拆分成若干Task，负责调度协调所有Task，相当于Master的角色。

* TaskTracker Node: 
负责执行由JobTracker指派的Task，相当于Worker的角色。这其中的Task分为MapTask和ReduceTask。



## 总结

Map:  映射过程，将一组数据按照某种Map函数映射成新的数据
Reduce：归约过程，把若干结果进行汇总并输出


