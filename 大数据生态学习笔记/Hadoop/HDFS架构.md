# HDFS架构

<!-- TOC -->

- [HDFS架构](#hdfs%e6%9e%b6%e6%9e%84)
  - [架构图](#%e6%9e%b6%e6%9e%84%e5%9b%be)
  - [具体说明](#%e5%85%b7%e4%bd%93%e8%af%b4%e6%98%8e)

<!-- /TOC -->


## 架构图

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hdfs-1.png)


## 具体说明
说明：
1.NameNode(nn)：就是Master, 它是一个主管，管理者

```
1）管理HDFS的名称空间
2）配置副本策略
3）管理数据块(block)映射信息
4）处理客户端读写请求
```
2.DataNode: 就是Slave, NameNode下达命令，DataNode执行实际的操作
```
1）存储实际的数据块
2）执行数据块的读写操作
```
3.Client 就是客户端
```
1）文件切分，文件上传HDFS的时候，Client将文件切分成一个一个的block，然后进行上传
2）与NameNode交互，获取文件的位置信息
3）与DataNode交互，读取或者写入信息
4）Client提供一些命令来管理HDFS，比如NameNode格式化
5）Client可以通过一些命令来访问HDFS, 比如对HDFS增删改查操作
```

4.Secondary NameNode  并非NameNode的热备
```
当NameNode挂掉的时候， 它并不能马上替换NameNode并提供服务

1）辅助NameNode, 分担其工作量，比如定期合并Fsimage和Edits，并推送NameNode 
（什么是FsImage和Edits?)

FsImage 和 EditLog 是 HDFS 的核心数据，这些数据的意外丢失可能会导致整个 HDFS 服务不可用。

为了避免这个问题，可以配置 NameNode 使其支持 FsImage 和 EditLog 多副本同步，这样 FsImage 或 EditLog 的任何改变都会引起每个副本 FsImage 和 EditLog 的同步更新。

2）在紧急情况下，可辅助恢复NameNode

```