# stat结构体

<!-- TOC -->

- [stat结构体](#stat%e7%bb%93%e6%9e%84%e4%bd%93)

<!-- /TOC -->

```
[zk: localhost:2181(CONNECTED) 12] stat /
#创建节点的事务zxid, 每次修改zk状态都会收到一个zxid形式的时间戳，也就是zk事务ID
事务ID是ZK中所有修改总的次序，每个修改都有唯一的zxid, 如果zxid1小于zxid2，那么
zxid1是在zxid2之前发生的
cZxid = 0x0  
#znode被创建的毫秒数(从1970年开始）
ctime = Thu Jan 01 08:00:00 CST 1970
#最后更新的事务zxid
mZxid = 0x0
#最后修改的毫秒数
mtime = Thu Jan 01 08:00:00 CST 1970
#最后更新的子节点
pZxid = 0x6
#Znode子节点变化号，znode子节点修改次数
cversion = 0
#Znode数据变化号
dataVersion = 0
#Znode访问控制列表的变化号
aclVersion = 0
#如果是临时节点，这个是znode拥有者的session id
如果不是临时节点，则是0
ephemeralOwner = 0x0
#Znode的数据长度
dataLength = 0
#Znode子节点数量
numChildren = 2

```
