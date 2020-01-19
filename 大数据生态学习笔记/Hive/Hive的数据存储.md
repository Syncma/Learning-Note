# Hive的数据存储

<!-- TOC -->

- [Hive的数据存储](#hive%e7%9a%84%e6%95%b0%e6%8d%ae%e5%ad%98%e5%82%a8)
	- [特点](#%e7%89%b9%e7%82%b9)
	- [表的种类](#%e8%a1%a8%e7%9a%84%e7%a7%8d%e7%b1%bb)
		- [Table 内部表](#table-%e5%86%85%e9%83%a8%e8%a1%a8)
		- [Partition 分区表](#partition-%e5%88%86%e5%8c%ba%e8%a1%a8)
		- [External Table 外部表](#external-table-%e5%a4%96%e9%83%a8%e8%a1%a8)
		- [Bucket Table 桶表](#bucket-table-%e6%a1%b6%e8%a1%a8)
		- [视图](#%e8%a7%86%e5%9b%be)

<!-- /TOC -->


## 特点
- 基于HDFS

- 没有专门的数据存储格式

- 存储结构主要包括： 数据库、文件、表、视图

- 可以直接加载文本文件（.txt文件等）

- 创建表时，指定Hive数据的列分隔符与行分割符

> create table t1(tid int, tname string, age int) row format delimited  fields terminated by ',';


## 表的种类
###  Table 内部表

1.与数据中的table 在概念上类似
2.每一个Table在hive都有一个相应的目录存储数据
3.所有的Table 数据（不包括External table)都保存在这个目录中
4.删除表时，元数据和数据都会被删除


### Partition 分区表
1.Partition 对应数据库中的Partition列的密集索引
2.在Hive中，表中的一个Partition对应表下的一个目录，所有的Partition的数据都存储在对应的目录中

> create table partition_table(tid int, sname string) partitioned by (gender string) xxxxx

有的时候查询我们并不希望扫描全表，所以分区能实现快速查询。


### External Table 外部表
1.指向已经在HDFS中存在的数据，也可以创建Partition
2.它和内部表在元数据的组织上是相同的，而实际数据的存储则有较大的差异
3.外部表只有一个过程，加载数据和创建表同时完成，并不会移动到数据仓库目录中，只是与外部数据建立一个连接。当删除
一个外部表，仅仅是删除该连接。

[参考地址](https://www.iteblog.com/archives/899.html)


最后归纳一下Hive中表与外部表的区别：
```
1、在导入数据到外部表，数据并没有移动到自己的数据仓库目录下，也就是说外部表中的数据并不是由它自己来管理的！而表则不一样；
2、在删除表的时候，Hive将会把属于表的元数据和数据全部删掉；而删除外部表的时候，Hive仅仅删除外部表的元数据，数据是不会删除的！

那么，应该如何选择使用哪种表呢？在大多数情况没有太多的区别，因此选择只是个人喜好的问题。但是作为一个经验，如果所有处理都需要由Hive完成，那么你应该创建表，否则使用外部表！
```


### Bucket Table 桶表
1.对数据进行哈希取值，然后放到不同文件中存储

分桶规则：对分桶字段值进行哈希，哈希值除以桶的个数求余，余数决定了该条记录在哪个桶中，也就是余数相同的在一个桶中。

优点：
1、提高join查询效率 
2、提高抽样效率


[参考内容](https://blog.csdn.net/qq_26937525/article/details/54880980)

另外一个要注意的问题是使用桶表的时候我们要开启桶表：
> set hive.enforce.bucketing = true;


### 视图
1.是一张虚表，是一个逻辑概念，可以跨越多张表
2.视图建立在已有表的基础上，视图赖以建立的这些表叫做基表
3.视图可以简化复杂的查询

和关系型数据库一样，Hive中也提供了视图的功能，注意Hive中视图的特性，和关系型数据库中的稍有区别：
	* 只有逻辑视图，没有物化视图；
	* 视图只能查询，不能Load/Insert/Update/Delete数据；
	* 视图在创建时候，只是保存了一份元数据，当查询视图的时候，才开始执行视图对应的那些子查询；


[Hive视图](https://www.cnblogs.com/zlslch/p/6105243.html)




