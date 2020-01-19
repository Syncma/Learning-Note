# HBase表结构模型

<!-- TOC -->

- [HBase表结构模型](#hbase%e8%a1%a8%e7%bb%93%e6%9e%84%e6%a8%a1%e5%9e%8b)
  - [关系型数据库](#%e5%85%b3%e7%b3%bb%e5%9e%8b%e6%95%b0%e6%8d%ae%e5%ba%93)
  - [BigTable](#bigtable)
  - [HBase](#hbase)
  - [HBase表结构模型](#hbase%e8%a1%a8%e7%bb%93%e6%9e%84%e6%a8%a1%e5%9e%8b-1)
    - [重要概念](#%e9%87%8d%e8%a6%81%e6%a6%82%e5%bf%b5)
      - [Rowkey](#rowkey)
      - [Column Family 列族](#column-family-%e5%88%97%e6%97%8f)
      - [Timestamp时间戳](#timestamp%e6%97%b6%e9%97%b4%e6%88%b3)
      - [Cell](#cell)
    - [HBase VS 关系型数据库区别](#hbase-vs-%e5%85%b3%e7%b3%bb%e5%9e%8b%e6%95%b0%e6%8d%ae%e5%ba%93%e5%8c%ba%e5%88%ab)
      - [RDBMS表](#rdbms%e8%a1%a8)
      - [HBase表](#hbase%e8%a1%a8)
      - [总结](#%e6%80%bb%e7%bb%93)
    - [HBase特点](#hbase%e7%89%b9%e7%82%b9)

<!-- /TOC -->


## 关系型数据库

关系型数据库(Oracle、MySQL、SQL Server)的特点
   1、什么是关系型数据库？基于关系模型（基于二维表）所提出的一种数据库
   
   2、**ER（Entity-Relationalship）模型**：通过增加外键来减少数据的冗余
   
   3、举例：学生-系


![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/normal-table.png)



## BigTable

BigTable简单来说就是把所有的数据保存到一张表中
采用冗余的特点，这样带来的好处 是提高效率

## HBase
因为有了bigtable的思想：NoSQL：HBase数据库

HBase基于Hadoop的HDFS的

下面是描述HBase的表结构图：

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/big-table.png)




## HBase表结构模型

通过上面的图 我们发现：

> HBase以表的形式存储数据。表有行和列组成。列划分为若干个列族(Column family)

### 重要概念

#### Rowkey 

数据唯一标识，按字典排序，就是检索记录的主键

1.访问Hbase table有三种方式：
```
1.通过单个row key 访问
2.通过row key 的range
3.全表扫描
```

2.row key 可以是任意字符串， 按照字典排序存储

3.不支持条件查询和Orderby 查询


#### Column Family 列族
多个列的集合，最多不超过3个
就是表中的列
列名都是以列族为前缀

这里要注意：
>一张表列簇不能超过5个
>每个列簇中的列数没有限制
>列只有插入数据后存在


#### Timestamp时间戳
 支持多个版本数据同时存在


#### Cell
**`HBase中通过row和columns确定的为一个存贮单元称为cell`**

每个 cell都保存着同一份数据的多个版本。
版本通过时间戳来索引。
时间戳的类型是 64位整型。
时间戳可以由hbase(在数据写入时自动 )赋值，此时时间戳是精确到毫秒的当前系统时间。
时间戳也可以由客户显式赋值。如果应用程序要避免数据版本冲突，就必须自己生成具有唯一性的时间戳。

每个 cell中，不同版本的数据按照时间倒序排序，即最新的数据排在最前面。

为了避免数据存在过多版本造成的的管理 (包括存贮和索引)负担，
hbase提供了两种数据版本回收方式。

一是保存数据的最后n个版本，
二是保存最近一段时间内的版本（比如最近七天）。
用户可以针对每个列族进行设置。
 

由`{row key, column( =<family> + <label>), version}` 唯一确定的单元。
cell中的数据是没有类型的，全部是字节码形式存贮。



### HBase VS 关系型数据库区别

HBase: 列动态增加，数据自动切分，高并发读取、不支持条件查询
关系型数据库： 复杂查询，其他都不支持


#### RDBMS表

| Primary Key | Column1 | Column2 |
| :---------- | ------: | :-----: |
| 记录1       |      XX |   XX    |
| 记录2       |      XX |   XX    |
| 记录3       |      XX |   XX    |



#### HBase表
| ROW Key |       CF1 |    CF2    |
| :------ | --------: | :-------: |
| 记录1   | 列1...列n | 列1...列n |
| 记录2   | 列1...列n | 列1...列n |
| 记录3   | 列1...列n | 列1...列n |


#### 总结
| 不同点   |        HBase |     RDBMS      |
| :------- | -----------: | :------------: |
| 数据类型 |       字符串 | 丰富的数据类型 |
| 数据操作 | 简单增删改查 |  丰富SQL支持   |
| 存储模式 |       列存储 |     行存储     |
| 数据保护 |         保留 |      替换      |
| 可伸缩性 |           好 |       差       |



### HBase特点

1.容量大
HBase单表可以有百亿行，百万列，数据矩阵横向和纵向两个维度所支持的数据量级都非常具有弹性

2.面向列
动态增加列存储，Hbase是面向列的存储和权限控制，并支持独立检索。
列式存储，其数据在表中是按照某列来存储的，这样在查询只需要少数几个字段的时候，能大大减少读取的数据量

3.多版本
HBase每一列的数据存储有多个Version

4.稀疏性
为空的列并不占用存储空间，表可以设计的非常稀疏

5.扩展性
底层依赖于HDFS

6.高可靠性
WAL机制保证了数据写入时不会因集群异常而导致写入数据丢失
Replication机制保证了在集群出现严重的问题时，数据不会发生丢失或损坏
而且HBase 底层使用HDFS, HDFS本身也有备份

7.高性能
底层的LSM数据结构和RowKey有序排列等架构上的独特设计，使得HBase具有非常高的写入性能
Region切分、主键索引和缓存机制使得HBase在海量数据中具有一定的随机读取性能，
该性能针对Rowkey的查询能达到毫秒级别