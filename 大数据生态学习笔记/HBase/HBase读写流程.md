# HBase读写流程

<!-- TOC -->

- [HBase读写流程](#hbase读写流程)
    - [写流程](#写流程)
        - [流程图](#流程图)
        - [流程讲解](#流程讲解)
    - [读流程](#读流程)
        - [流程图](#流程图-1)
        - [流程讲解](#流程讲解-1)

<!-- /TOC -->


## 写流程



### 流程图
![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/169ce5b07ff84fe3.png)


### 流程讲解

客户端Client发送写数据请求，通过ZooKeeper获取到表的元数据信息，客户端通过RPC通信查找到对应的RegionServer，进而找到Region，同时在HLog中记录写操作，通过HLog把数据信息写入到memstore（16KB），memstore存满后溢写到storeFile中，最后HDFS客户端统一存储到HFile。


## 读流程

### 流程图

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/169b3693fdc1cb5a.png)

### 流程讲解


客户端Client访问ZooKeeper，返回-ROOT-表元数据位置，根据元数据位置去查找对应的RegionServer，同时根据-ROOT-查找到.META表，再根据.META表的元数据查找到Region，返回Region的业务元数据给客户端。

客户端Client从Region中的Store中读取数据，若在写数据缓存memstore（存放用户最近写入的数据）中读到对应数据，则直接返回数据信息到客户端，若memstore中不存在对应数据，则去读数据缓存blockcache中查找，若blockcache中仍未找到，则去对应的HFile查找数据信息并存入blockcache，进而通过blockcache返回数据信息到客户端。