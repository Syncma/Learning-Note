#  HDFS读写流程
<!-- TOC -->

- [HDFS读写流程](#hdfs%e8%af%bb%e5%86%99%e6%b5%81%e7%a8%8b)
  - [写过程](#%e5%86%99%e8%bf%87%e7%a8%8b)
  - [读过程](#%e8%af%bb%e8%bf%87%e7%a8%8b)

<!-- /TOC -->

## 写过程



![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hdfs-read.png)


写流程
- 客户端向NameNode发起写数据请求
- 分块写入DataNode节点，DataNode自动完成副本备份
- DataNode向NameNode汇报存储完成，NameNode通知客户端

## 读过程

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hdfs-write.png)




读流程：
- 客户端向NameNode发起读数据请求
- NameNode找出距离最近的DataNode节点信息
- 客户端从DataNode分块下载文件
