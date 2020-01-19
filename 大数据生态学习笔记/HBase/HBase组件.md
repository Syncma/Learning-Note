# HBase组件
<!-- TOC -->

- [HBase组件](#hbase%e7%bb%84%e4%bb%b6)
  - [组件图](#%e7%bb%84%e4%bb%b6%e5%9b%be)
  - [组件介绍](#%e7%bb%84%e4%bb%b6%e4%bb%8b%e7%bb%8d)
  - [HRegionServer](#hregionserver)

<!-- /TOC -->


## 组件图

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hbase-art1.png)

## 组件介绍

HDFS、RegionServer、Master、ZooKeeper

Regionserver 向Master 、ZooKeeper 报告自己健康状态 自己负责的分区

HBase的主要进程： **Master, RegionServer**
HBase所依赖的两个外部服务：**ZooKeeper、 HDFS**

Zookeeper Quorum存储-ROOT-表地址、HMaster地址
HRegionServer把自己以Ephedral方式注册到Zookeeper中，
HMaster随时感知各个HRegionServer的健康状况
Zookeeper避免HMaster单点问题


## HRegionServer
![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hbase-art2.png)


HRegionServer结构：
* HLog：存储HBase的修改记录
* HRegion：根据rowkey（行键，类似id）分割的表的分片
* Store：对应HBase表中的一个列族，可存储多个字段
* HFile：真正的存储文件
* MemStore：保存当前的操作
* ZooKeeper：存放数据的元数据信息，负责维护RegionServer中保存的元数据信息
* DFS Client：存储数据信息到HDFS集群中

