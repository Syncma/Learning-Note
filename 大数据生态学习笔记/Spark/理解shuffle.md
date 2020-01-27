# 理解shuffle

<!-- TOC -->

- [理解shuffle](#%e7%90%86%e8%a7%a3shuffle)
  - [什么是shuffle](#%e4%bb%80%e4%b9%88%e6%98%afshuffle)
  - [shuffle影响](#shuffle%e5%bd%b1%e5%93%8d)
  - [导致Shuffle的操作](#%e5%af%bc%e8%87%b4shuffle%e7%9a%84%e6%93%8d%e4%bd%9c)

<!-- /TOC -->



 ## 什么是shuffle

在 Spark 中，一个任务对应一个分区，通常不会跨分区操作数据。

但如果遇到 `reduceByKey` 等操作，Spark 必须从所有分区读取数据，并查找所有键的所有值，然后汇总在一起以计算每个键的最终结果 ，这称为 `Shuffle`

类似洗牌：

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200127094432.png)





##  shuffle影响

因为它通常会跨节点操作数据，这会涉及磁盘 I/O，网络 I/O，和数据序列化。

某些 Shuffle 操作还会消耗大量的堆内存，因为它们使用堆内存来临时存储需要网络传输的数据。



## 导致Shuffle的操作

由于 Shuffle 操作对性能的影响比较大，所以需要特别注意使用，以下操作都会导致 Shuffle：

* 涉及到重新分区操作**： 如 `repartition` 和 `coalesce`；

- **所有涉及到 ByKey 的操作**：如 `groupByKey` 和 `reduceByKey`，但 `countByKey` 除外；
- **联结操作**：如 `cogroup` 和 `join`。