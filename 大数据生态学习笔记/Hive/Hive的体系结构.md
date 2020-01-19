# Hive的体系结构
<!-- TOC -->

- [Hive的体系结构](#hive%e7%9a%84%e4%bd%93%e7%b3%bb%e7%bb%93%e6%9e%84)
  - [体系图](#%e4%bd%93%e7%b3%bb%e5%9b%be)
  - [组件说明](#%e7%bb%84%e4%bb%b6%e8%af%b4%e6%98%8e)
    - [MetaStore](#metastore)
    - [HQL运行流程](#hql%e8%bf%90%e8%a1%8c%e6%b5%81%e7%a8%8b)
  - [Hive运行机制](#hive%e8%bf%90%e8%a1%8c%e6%9c%ba%e5%88%b6)

<!-- /TOC -->

## 体系图

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hive-3.png)


## 组件说明

### MetaStore
**Hive的元数据 Metastore**

- Hive将元数据存储在数据库中(metastore), 支持mysql, derby等数据库

- Hive中的元数据包括表的名字，表的列和分区以及属性，表的属性（是否为外部表等）， 表的数据所在目录等


###  HQL运行流程
一条HQL语句如何在Hive中进行查询？
解析器，编译器，优化器完成HQL查询语句从词法分析，语法分析、编译、优化以及查询计划(Plan)的生成。

生成的查询计划存储在HDFS中，并在随后由MapReduce调用执行

执行过程:
 > HQL->解析器（词法分析)->编译器(生成HQL的执行计划)->优化器(生成最佳的执行计划)->执行


## Hive运行机制
![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hive-2.png)

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hive.png)