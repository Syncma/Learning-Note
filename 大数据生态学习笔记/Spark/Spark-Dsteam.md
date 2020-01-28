# Spark-Dsteam

<!-- TOC -->

- [Spark-Dsteam](#spark-dsteam)
  - [Dsteam介绍](#dsteam%e4%bb%8b%e7%bb%8d)
  - [其他内容](#%e5%85%b6%e4%bb%96%e5%86%85%e5%ae%b9)

<!-- /TOC -->
## Dsteam介绍

Spark Streaming 是将流式计算分解成一系列短小的批处理作业，也就是把输入数据按照batch size（如1秒）分成一段一段的数据（DStream），每一段数据都转换成Spark中的RDD，然后将对DStream的Transformation操作变为针对Spark中对 RDD的Transformation操作，将RDD经过操作变成中间结果保存在内存中。

Spark Streaming在内部的处理机制是：接收实时流的数据，并根据一定的时间间隔拆分成一批批的数据，然后通过Spark Engine处理这些批数据，最终得到处理后的一批批结果数据。

DStream作为Spark Streaming的基础抽象，它代表持续性的数据流。这些数据流既可以通过外部输入源赖获取，也可以通过现有的Dstream的 transformation操作来获得。

在内部实现上，DStream由一组时间序列上连续的RDD来表示。每个RDD都包含了自己特定时间间隔内的数据流



## 其他内容

* 待补充

  