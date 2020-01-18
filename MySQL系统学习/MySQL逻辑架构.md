# MySQL逻辑架构
<!-- TOC -->

- [MySQL逻辑架构](#mysql逻辑架构)
    - [说明](#说明)
    - [安装过程](#安装过程)
    - [逻辑架构](#逻辑架构)
        - [连接器](#连接器)
            - [原理图](#原理图)
            - [说明](#说明-1)
    - [长连接和短连接](#长连接和短连接)
    - [查询过程](#查询过程)
    - [查询缓存](#查询缓存)
    - [分析器](#分析器)
    - [优化器](#优化器)
    - [执行器](#执行器)
    - [基本命令](#基本命令)
        - [新建数据库](#新建数据库)

<!-- /TOC -->
## 说明

这里使用 **`MariaDB`** 作为数据库
 MariaDB数据库管理系统是MySQL的一个分支，主要由开源社区在维护

[下载地址](https://downloads.mariadb.org/)


## 安装过程
windows 平台直接下载msi格式 最新稳定版本是10.4  
备注：**windows安装完成需要重新**

Fedora 平台 -----**`默认测试环境`**

1.使用root用户执行下面的命令进行安装：
```
dnf install mariadb-server mariadb-client
```

2.服务开启，使用root用户执行
```
systemctl start mariadb #启动服务
systemctl enable mariadb #设置开机启动
systemctl restart mariadb #重新启动
systemctl stop mariadb.service #停止MariaDB
```

3.为 MariaDB 配置远程访问权限
```
select User, host from mysql.user;

GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '123456' WITH GRANT OPTION;

FLUSH PRIVILEGES;
```


## 逻辑架构


![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/1075436-20190215135724690-741976355.png)


### 连接器

#### 原理图

![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/1075436-20190215135942438-1134796733.png)



#### 说明

1. 连接mysql
```
[jian@laptop ~]$ mysql -u root -p
Enter password: 
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MariaDB connection id is 8
Server version: 10.3.18-MariaDB MariaDB Server

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
```

2.查询链接状态

![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/processlist.png)



## 长连接和短连接
1、什么是长链接？
数据库里面，长连接是连接成功后，如果客户端持续有请求，则一直使用同一个链接。

2、什么是短连接？
短连接则是指每次执行完很少的几次查询就断开连接，下次查询重新建立一个

3、尽量使用长链接
建立连接的过程通常是比较复杂的，所以使用中尽量减少建立的动作，也就是使用长连接


## 查询过程
用户总是希望MySQL能够获得更高的查询性能，最好的办法是弄清楚MySQL是如何优化和执行查询的。

一旦理解了这一点，就会发现：

很多的查询优化工作实际上就是遵循一些原则让MySQL的优化器能够按照预想的合理方式运行而已。

当向MySQL发送一个请求的时候，MySQL到底做了些什么呢？

![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/查询过程.png)


## 查询缓存
![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/1075436-20190215153911680-2025779852.png)


MySQL拿到一个查询请求后，会先到查询缓存看看，之前是不是执行过这条语句，如果有，就直接返回给客户端

如果语句不在查询缓存中，就会继续后面的执行阶段。

执行完成后，执行结果会被存入查询缓存中，

如果查询命中缓存MySQL不需要执行后面的复杂操作，就可以直接返回结果，这个效率会很高

简单的说就是MySQL会优先检查这个查询是否命中查询缓存中的数据

可以使用下面命令查看：

![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/mysql-cache.png)


这些参数介绍：
```
query_cache_limit: MySQL能够缓存的最大结果,如果超出,则增加 Qcache_not_cached的值,并删除查询结果
query_cache_min_res_unit: 分配内存块时的最小单位大小
query_cache_size: 缓存使用的总内存空间大小,单位是字节,这个值必须是1024的整数倍,否则MySQL实际分配可能跟这个数值不同(感觉这个应该跟文件系统的blcok大小有关)
query_cache_type: 是否打开缓存 OFF: 关闭 ON: 总是打开
query_cache_wlock_invalidate: 如果某个数据表被锁住,是否仍然从缓存中返回数据,默认是OFF,表示仍然可以返回
```


## 分析器

![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/1075436-20190215162821053-414836620.png)

如果你的语句不对，就会收到“You have an error in your SQL syntax”的错误提醒，


## 优化器

1、在表里面有多个索引的时候，决定使用哪个索引

2、多表关联(ioin)的时候，决定各个表的链接顺序

## 执行器

![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/1075436-20190215164402749-1383016574.png)

查询执行的最后一个阶段是将结果返回给客户端。

即使查询不需要返回结果给客户端，MySQL仍然会返回这个查询的一些信息，如查询影响到的行数。

## 基本命令


### 新建数据库

1.默认create database db_name命令，生成的数据库是latin1编码

```

MariaDB [(none)]> create database test;
Query OK, 1 row affected (0.001 sec)


//查看所有数据库情况：
MariaDB [test]> SELECT SCHEMA_NAME 'database', default_character_set_name 'charset', DEFAULT_COLLATION_NAME 'collation' FROM information_schema.SCHEMATA;
+--------------------+---------+-------------------+
| database           | charset | collation         |
+--------------------+---------+-------------------+
| information_schema | utf8    | utf8_general_ci   |
| test               | latin1  | latin1_swedish_ci |
| mysql              | latin1  | latin1_swedish_ci |
| performance_schema | utf8    | utf8_general_ci   |
+--------------------+---------+-------------------+
4 rows in set (0.001 sec)


//查看单个数据库：
MariaDB [(none)]> use test;
Database changed


MariaDB [test]> show variables like "collation_database";
+--------------------+-------------------+
| Variable_name      | Value             |
+--------------------+-------------------+
| collation_database | latin1_swedish_ci |
+--------------------+-------------------+
1 row in set (0.001 sec)

MariaDB [test]> show variables like "character_set_database";
+------------------------+--------+
| Variable_name          | Value  |
+------------------------+--------+
| character_set_database | latin1 |
+------------------------+--------+
1 row in set (0.001 sec)

```

如果想要生成utf8mb4编码,可以使用下面命令

```
create database test default character set utf8mb4 collate utf8mb4_unicode_ci;
```