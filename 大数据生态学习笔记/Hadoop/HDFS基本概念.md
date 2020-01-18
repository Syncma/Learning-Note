# HDFS基本概念
<!-- TOC -->

- [HDFS基本概念](#hdfs%e5%9f%ba%e6%9c%ac%e6%a6%82%e5%bf%b5)
  - [块(block)](#%e5%9d%97block)
  - [NameNode](#namenode)
  - [DataNode](#datanode)
  - [Secondary NameNode](#secondary-namenode)

<!-- /TOC -->

## 块(block)

HDFS 的文件被分为块进行存储、块是文件存储处理的逻辑单位

**`这里要了解为什么hadoop block大小要设置在128M 左右`**

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200118102318.png)


如果设置大一点或者小一点有什么影响呢？

```
如果设置太小，会增加寻址时间，程序一直在找块的开始位置

如果设置太大，从磁盘传输数据的时间会明显大于定位这个块开始位置所需的时间
导致程序在处理这块数据时，会非常慢
```
总结： **`HDFS 块的大小设置主要取决于磁盘传输效率`**


## NameNode

**`管理节点、存放文件元数据`**

* 文件和数据块的映射表：文件名、文件目录结构、文件属性（生成时间，副本数、文件权限）
* 数据块与数据节点的映射表：每个文件的块列表和块所在的DataNode


## DataNode

 **`HDFS的工作节点、存放数据块`**
 
在本地文件系统存储文件块数据，以及块数据的校验和


## Secondary NameNode
**用来监控HDFS状态的辅助后台程序，每隔一段时间获取HDFS元数据的快照**