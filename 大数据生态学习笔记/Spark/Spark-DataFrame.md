# Spark-DataFrame

<!-- TOC -->

- [Spark-DataFrame](#spark-dataframe)
  - [介绍](#%e4%bb%8b%e7%bb%8d)
  - [DataFrame创建](#dataframe%e5%88%9b%e5%bb%ba)
  - [DataFrame和RDD的区别](#dataframe%e5%92%8crdd%e7%9a%84%e5%8c%ba%e5%88%ab)
  - [RDD转换DataFrame](#rdd%e8%bd%ac%e6%8d%a2dataframe)

<!-- /TOC -->



## 介绍

为了支持结构化数据的处理，Spark SQL 提供了新的数据结构 DataFrame。DataFrame 是一个由具名列组成的数据集。它在概念上等同于关系数据库中的表或 R/Python 语言中的 `data frame`



## DataFrame创建

* 内容待补充



## DataFrame和RDD的区别

**1.DataFrame的推出，让Spark具备了处理大规模结构化数据的能力**

**不仅比原有的RDD转换方式更加简单易用，而且获得更高的计算性能**

**2.Spark能够轻松的从MySQL到DataFrame的转化，并且支持SQL查询**

**3.RDD是分布式 java对象的集合，但是对象内部结构对于RDD而言是不可知的**

**4.DataFrame是一种以RDD为基础的分布式数据集，提供了详细的结构信息**



![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200128093024.png)





## RDD转换DataFrame

* 内容待补充



