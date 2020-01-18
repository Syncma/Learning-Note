# ZK常见问题

<!-- TOC -->

- [ZK常见问题](#zk%e5%b8%b8%e8%a7%81%e9%97%ae%e9%a2%98)
  - [问题1](#%e9%97%ae%e9%a2%981)
  - [其他问题](#%e5%85%b6%e4%bb%96%e9%97%ae%e9%a2%98)

<!-- /TOC -->

## 问题1
ZK创建目录：

```
[jian@laptop zookeeper-3.4.14]$ bin/zkCli.sh

Connecting to localhost:2181
xxxx
[zk: localhost:2181(CONNECTED) 9] create /hbase null
Created /hbase
[zk: localhost:2181(CONNECTED) 10] create /hbase/master null
Created /hbase/master
 
//但是/hbase/master目录 不是自动创建的
//只能创建一级目录的节点，多级时，必须一级一级创建
//不支持递归创建，必须先创建父节点
//节点不能以 / 结尾，会直接报错

[zk: 127.0.0.1:2181(CONNECTED) 60] create /zk/test2/ null
Command failed: java.lang.IllegalArgumentException: Path must not end with / character
//节点不为空不能删除
[zk: 127.0.0.1:2181(CONNECTED) 58] delete /zk
Node not empty: /zk
```

删除时，须先清空节点下的内容，才能删除节点

```
delete /zk/test1
delete /zk/test2
delete /zk
```


## 其他问题

* 待补充