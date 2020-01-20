# Spark介绍

<!-- TOC -->

- [Spark介绍](#spark%e4%bb%8b%e7%bb%8d)
  - [什么是Spark](#%e4%bb%80%e4%b9%88%e6%98%afspark)
  - [优势](#%e4%bc%98%e5%8a%bf)
  - [Spark特点](#spark%e7%89%b9%e7%82%b9)

<!-- /TOC -->
## 什么是Spark


Spark是一个针对大规模数据处理的快速通用引擎。
类似MapReduce，都进行数据的处理


## 优势

- 基于内存计算
 **DAG 有向无环图 -内存流水线优化**

- 抽象出分布式内存存储数据结构，弹性分布式数据集RDD

- 基于事件驱动，通过线程池复用线程提高性能

**`这些特点会在后续章节会详细说明`**
     
## Spark特点

```
基于Scala语言、Spark基于内存的计算

快：基于内存

易用：支持Scala、Java、Python

通用：Spark Core、Spark SQL、Spark Streaming、MLlib、Graphx

兼容性：完全兼容Hadoop
```
