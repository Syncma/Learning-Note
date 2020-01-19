# HiveVSHbase区别
<!-- TOC -->

- [HiveVSHbase区别](#hivevshbase%e5%8c%ba%e5%88%ab)
	- [hbase与hive的联系](#hbase%e4%b8%8ehive%e7%9a%84%e8%81%94%e7%b3%bb)
	- [Hive](#hive)
	- [HBase](#hbase)
	- [性能比较](#%e6%80%a7%e8%83%bd%e6%af%94%e8%be%83)
	- [各自适用的场景](#%e5%90%84%e8%87%aa%e9%80%82%e7%94%a8%e7%9a%84%e5%9c%ba%e6%99%af)

<!-- /TOC -->

## hbase与hive的联系

1.hive适合处理离线的数据
2.hbase适合处理实时的数据的查询
两者合并起来使用可以达到‘＋’的效果

## Hive
**hive适合用于网络日志等数据量大的静态数据查询**
HIVE是hadoop的数据仓库，依赖于HDFS和mapreduce
类似于SQL操作
把MAPREDUCE的程序作为插件来支持HIVE的数据分析
作用于全表扫描使用（HIVE+HADOOP）
hive的操作是基于整个数据表的、
所以查询起来常常是以小时来计
不支持常规的更新语句，插入，更新，删除

## HBase
**hbase适合大数据的实时查询**
是一个数据库系统，面向列的数据库查询，有自己的查询 语句
支持横向扩展，减少成本
由自己的查询方式，不用依赖于MAPREDUCE
索引访问使用（HBASE+HADOOP）

## 性能比较

HBASE相对于HIVE是比较高效的多的
HIVE需要使用到HDFS存储，要用到MAPREDUCE计算框架
HBASE需要使用HDFS存放文件，HBASE负责组织文件
HIVE需要借助MAPREDUCE来完成HIVE的命令执行


相同
> hbase与HIVE都是架构在HADOOP之上的，都是用HADOOP作为底层存储


HBASE优点
* 列的动态增加，并且列为空就不存储数据，节约存储空间
* 支持高并发读写操作


HBASE缺点
* 不支持条件查询，只支持按照ROWKEY查询
* 不支持MASTER的故障切换，当MASTER宕机，整个系统就瘫痪掉了
* 只保存字符类型
* 没有表与表之间的关系


## 各自适用的场景

先放结论：Hbase和Hive在大数据架构中处在不同位置，

Hbase主要解决实时数据查询问题，
Hive主要解决数据处理和计算问题，
一般是配合使用。


一、区别：

	1. Hbase： Hadoop database 的简称，也就是基于Hadoop数据库，是一种NoSQL数据库，主要适用于海量明细数据（十亿、百亿）的随机实时查询，如日志明细、交易清单、轨迹行为等。
	
	2. Hive：Hive是Hadoop数据仓库，严格来说，不是数据库，主要是让开发人员能够通过SQL来计算和处理HDFS上的结构化数据，适用于离线的批量数据计算。

	* 通过元数据来描述Hdfs上的结构化文本数据，通俗点来说，就是定义一张表来描述HDFS上的结构化文本，包括各列数据名称，数据类型是什么等，方便我们处理数据，
	* 当前很多SQL ON Hadoop的计算引擎均用的是hive的元数据，如Spark SQL、Impala等；
	* 基于第一点，通过SQL来处理和计算HDFS的数据，Hive会将SQL翻译为Mapreduce来处理数据；


二、关系

在大数据架构中，Hive和HBase是协作关系，数据流一般如下图：
	1. 通过ETL工具将数据源抽取到HDFS存储；
	2. 通过Hive清洗、处理和计算原始数据；
	3. HIve清洗处理后的结果，如果是面向海量数据随机查询场景的可存入Hbase
	4. 数据应用从HBase查询数据；

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hive-hbase.png)



更为细致的区别如下：
```
1.Hive中的表是纯逻辑表，就只是表的定义等，即表的元数据。Hive本身不存储数据，它完全依赖HDFS和MapReduce。这样就可以将结构化的数据文件映射为为一张数据库表，并提供完整的SQL查询功能，并将SQL语句最终转换为MapReduce任务进行运行。而HBase表是物理表，适合存放非结构化的数据。

2.Hive是基于MapReduce来处理数据,而MapReduce处理数据是基于行的模式；HBase处理数据是基于列的而不是基于行的模式，适合海量数据的随机访问。

3.HBase的表是疏松的存储的，因此用户可以给行定义各种不同的列；而Hive表是稠密型，即定义多少列，每一行有存储固定列数的数据。

4.Hive使用Hadoop来分析处理数据，而Hadoop系统是批处理系统，因此不能保证处理的低迟延问题；而HBase是近实时系统，支持实时查询。

5.Hive不提供row-level的更新，它适用于大量append-only数据集（如日志）的批任务处理。而基于HBase的查询，支持和row-level的更新。

6.Hive提供完整的SQL实现，通常被用来做一些基于历史数据的挖掘、分析。而HBase不适用与有join，多级索引，表关系复杂的应用场景。

```

