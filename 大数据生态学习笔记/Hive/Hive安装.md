# Hive安装
<!-- TOC -->

- [Hive安装](#hive安装)
    - [安装模式](#安装模式)
        - [嵌入模式](#嵌入模式)
            - [安装过程](#安装过程)
        - [本地模式](#本地模式)
            - [安装过程](#安装过程-1)
        - [远程模式](#远程模式)
            - [安装过程](#安装过程-2)

<!-- /TOC -->


## 安装模式
### 嵌入模式

```
	1. 元数据信息被存储在Hive自带的derby数据库中
	2. 只允许创建一个连接
	3. 多用于Demo
	4. 数据会存放在HDFS上
```

#### 安装过程
1.解压

```
[jian@laptop tools]$ tar xf apache-hive-2.3.6-bin.tar.gz -C ../
[jian@laptop tools]$ cd ..
[jian@laptop bigdata]$ mv apache-hive-2.3.6-bin/  hive-2.3.6
[jian@laptop bigdata]$
```


2.修改环境变量
```
//增加下面内容
[jian@laptop bigdata]$ cat ~/.bashrc
#Hive配置
export HIVE_HOME=/home/jian/prj/bigdata/hive-2.3.6
export PATH=$HIVE_HOME/bin:$PATH
```
执行source命令使配置文件生效：
```
[jian@laptop bigdata]$ source ~/.bashrc
```


3.修改配置文件
```
[jian@laptop conf]$ cp hive-env.sh.template hive-env.sh
[jian@laptop conf]$ pwd/home/jian/prj/bigdata/hive-2.3.6/conf
[jian@laptop conf]$ cat hive-env.sh
增加下面内容：HADOOP_HOME=/home/jian/prj/bigdata/hadoop-2.10.0
```

4.启动Hive在启动hive 之前 先开启Hadoop服务
```
[jian@laptop hadoop-2.10.0]$ sbin/start-all.sh
[jian@laptop hadoop-2.10.0]$ jps
19745 NodeManager
20114 Jps
19330 SecondaryNameNode
19572 ResourceManager
19060 DataNode
18846 NameNode

```

```
[jian@laptop conf]$ hive
xxxCaused by: org.apache.hadoop.hdfs.server.namenode.SafeModeException: Cannot create directory /tmp/hive. Name node is in safe mode.
```
启动报错解决：

```
//手动关闭安全模式[
jian@laptop hadoop-2.10.0]$ hadoop dfsadmin -safemode leave
DEPRECATED: Use of this script to execute hdfs command is deprecated.Instead use the hdfs command for it.
Safe mode is OFF
```

备注:

```

安全模式是HDFS所处的一种特殊状态，在这种状态下，文件系统只接受读数据请求，而不接受删除、修改等变更请求。在NameNode主节点启动时，HDFS首先进入安全模式，DataNode在启动的时候会向namenode汇报可用的block等状态，当整个系统达到安全标准时，HDFS自动离开安全模式。如果HDFS出于安全模式下，则文件block不能进行任何的副本复制操作，因此达到最小的副本数量要求是基于datanode启动时的状态来判定的，启动时不会再做任何复制（从而达到最小副本数量要求）

只有接收到的datanode上的块数量（datanodes  blocks）和实际的数量（total blocks）接近一致， 超过  datanodes blocks /  total blocks >= 99.9%  这个阀值，就表示 块数量一致，就会推出安全模式。达到99.9%的阀值之后，文件系统不会立即推出安全模式，而是会等待30秒之后才会退出。

在安全模式下不可以进行以下操作：
1）创建文件夹
2）上传文件
3）删除文件

注意：启动了namenode，未启动datanode，则文件系统处于安全模式，只可以查看文件系统上有哪些文件，不可以查看文件内容，因为datanode都还没启动，怎么可能可以查看文件内容
```

//手动开启安全模式
```
[jian@laptop hadoop-2.10.0]$ hadoop dfsadmin -safemode enter
```
再次执行：
```
[jian@laptop conf]$ hive
....
hive>
```


### 本地模式     
* 元数据信息被存储在MySQL数据库中
* MySQL数据库与Hive运行在同一台物理机器上
* 多用于开发与测试

#### 安装过程
1.登录mysql
```
[jian@laptop hive-2.3.6]$ mysql -u root -p
...
mysql> create database hive;
Query OK, 1 row affected (0.01 sec)
mysql> create user 'hive'@'localhost' identified by 'Qmmjj123#';
mysql> grant all privileges on *.* to 'hive'@'%';
ERROR 1410 (42000): You are not allowed to create a user with GRANT#In summary, now in MySQL 8.0 you cannot create a user from GRANT, you don't need to run FLUSH PRIVILEGES command
mysql> grant all privileges on *.* to 'hive'@'localhost';
Query OK, 0 rows affected (0.01 sec)
```


2.将JDBC驱动包拷贝到Hive lib 目录下mysql-connector-java-8.0.18.jar
```
[jian@laptop lib]$ pwd
/home/jian/prj/bigdata/hive-2.3.6/lib
```



3.修改配置文件
```
[jian@laptop conf]$ cp hive-default.xml.template hive-site.xml
[jian@laptop conf]$ pwd
/home/jian/prj/bigdata/hive-2.3.6/conf
[jian@laptop conf]$ cat hive-site.xml
修改下面内容：

<property>
<name>javax.jdo.option.ConnectionURL</name>
<value>jdbc:mysql://localhost/hive?createDatabaseIfNotExist=true</value><description>JDBC connect string for a JDBC metastore.To use SSL to encrypt/authenticate the connection, provide database-specific SSL flag in the connection URL.For example, jdbc:postgresql://myhost/db?ssl=true for postgres database.</description>
</property>

<property>
<name>javax.jdo.option.ConnectionDriverName</name><value>com.mysql.jdbc.Driver</value>
<description>Driver class name for a JDBC metastore</description>
</property>

<property><name>javax.jdo.option.ConnectionUserName</name>
<value>hive</value>
<description>Username to use against metastore database</description>
</property>


<property><name>javax.jdo.option.ConnectionPassword</name>
<value>Qmmjj123#</value>
<description>password to use against metastore database</description>
</property>

//添加：
<property>
<name>hive.metastore.local</name>
<value>true</value>
<description>controls whether to connect to remove metastore server or open a new metastore server in Hive Client JVM</description>
</property>

<property>
<name>hive.server2.logging.operation.log.location</name><value>/tmp/hive/operation_logs</value>
<description>Top level directory where operation logs are stored if logging functionality is enabled</description>
</property>

<property>
<name>hive.exec.local.scratchdir</name>
<value>/tmp/hive</value>
<description>Local scratch space for Hive jobs</description>
</property>

<property>
<name>hive.downloaded.resources.dir</name>
<value>/tmp/hive/resources</value>
<description>Temporary local directory for added resources in the remote file system.</description>
</property>

<property>
<name>hive.querylog.location</name>
<value>/tmp/hive/querylog</value>
<description>Location of Hive run time structured log file
</description>
</property>

[jian@laptop conf]$ cp hive-log4j2.properties.template hive-log4j2.properties

```


4.启动

```
[jian@laptop lib]$ hive
...
hive> show tables;
Loading class `com.mysql.jdbc.Driver'. This is deprecated. The new driver class is `com.mysql.cj.jdbc.Driver'. The driver is automatically registered via the SPI and manual loading of the driver class is generally unnecessary.FAILED: SemanticException org.apache.hadoop.hive.ql.metadata.HiveException: java.lang.RuntimeException: Unable to instantiate org.apache.hadoop.hive.ql.metadata.SessionHiveMetaStoreClient
```


出现上面错误，修改配置文件：
```
[jian@laptop conf]$ cat hive-site.xml
<property>
<name>javax.jdo.option.ConnectionDriverName</name><value>com.mysql.cj.jdbc.Driver</value>
<description>Driver class name for a JDBC metastore</description>
</property>
```

这个错误：

```
FAILED: SemanticException org.apache.hadoop.hive.ql.metadata.HiveException: java.lang.RuntimeException: Unable to instantiate org.apache.hadoop.hive.ql.metadata.SessionHiveMetaStoreClient
```

修改配置文件里面的
```
<name>hive.metastore.schema.verification</name>
<value>false</value>
```

原因：
> 因为没有正常启动Hive 的 Metastore Server服务进程。 

解决方法：
启动Hive 的 Metastore Server服务进程
执行如下命令：
```
[jian@laptop conf]$ hive --service metastore &
```

```
hive> show tables;

FAILED: SemanticException org.apache.hadoop.hive.ql.metadata.HiveException: org.apache.hadoop.hive.ql.metadata.HiveException: MetaException(message:Hive metastore database is not initialized. Please use schematool (e.g. ./schematool -initSchema -dbType ...) to create the schema. If needed, don't forget to include the option to auto-create the underlying database in your JDBC connection string (e.g. ?createDatabaseIfNotExist=true for mysql))
```


执行：
```
[jian@laptop conf]$ schematool -initSchema -dbType mysql
然后再次执行：

[jian@laptop lib]$ hive
hive> show tables;
OK
Time taken: 3.726 seconds


mysql> use hive;
mysql> show tables;
....
```

**可以发现这里多了很多的表，这些就是元数据。存放在mysql数据库当中。**


hive新建表：
hive> create table test2(id int, name string);
OK
Time taken: 0.73 seconds

mysql查看对应的元数据：
mysql> select * from COLUMNS_V2;
+-------+---------+-------------+-----------+-------------+| 
CD_ID | COMMENT | COLUMN_NAME | TYPE_NAME | INTEGER_IDX |
+-------+---------+-------------+-----------+-------------+| 
1 | NULL | id | int | 0 |
| 1 | NULL | name | string | 1 |
+-------+---------+-------------+-----------+-------------+
2 rows in set (0.00 sec)




这里的CD_ID指的是test2表的id，里面列对应列名，以及对应的类型说明等等。
这说明hive中的元数据确实是存在mysql当中的。

hive是基于hadoop的。
存放表的信息，即元数据是在mysql当中，但是表的实际数据是存放在hadoop的hdfs当中的。



浏览器访问：[访问这里](http://localhost:50070)

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hive-web1.png)



### 远程模式
* 元数据信息被存储在MySQL数据库中
* MySQL数据库与Hive运行在不同的操作系统上
*  多用于生产环境，允许多个连接

这种模式需要使用hive安装目录下提供的beeline+hiveserver2配合使用才可以。


其原理就是将metadata作为一个单独的服务进行启动。各种客户端通过beeline来连接，连接之前无需知道数据库的密码。

#### 安装过程
1.执行hiveserver2命令：
```
[jian@laptop hive-2.3.6]$ hiveserver2 start
2019-12-02 13:20:49: Starting HiveServer2SLF4J: Class path contains multiple SLF4J bindings.SLF4J: Found binding in [jar:file:/home/jian/prj/bigdata/hive-2.3.6/lib/log4j-slf4j-impl-2.6.2.jar!/org/slf4j/impl/StaticLoggerBinder.class]SLF4J: Found binding in [jar:file:/home/jian/prj/bigdata/hadoop-2.10.0/share/hadoop/common/lib/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]SLF4J: See http://www.slf4j.org/codes.html#multiple_bindings for an explanation.SLF4J: Actual binding is of type [org.apache.logging.slf4j.Log4jLoggerFactory]
```

启动后命令行会一直监听不退出，我们可以看到它监听了10000端口。
```
[jian@laptop conf]$ netstat -antp |grep 10000
(Not all processes could be identified, non-owned process infowill not be shown, you would have to be root to see it all.)
tcp 0 0 0.0.0.0:10000 0.0.0.0:* LISTEN 29979/java
```


2.新建窗口，执行beeline命令
基于 SQLLine CLI的JDBC客户端

```
beeline的命令在执行时，需要在前边加上!，help命令可以查看具体详细用法

[jian@laptop conf]$ beeline -u jdbc:hive2://

 jdbc:hive2://> show tables;
 

也可以这样执行：
[jian@laptop conf]$ beeline
Beeline version 2.3.6 by Apache Hive

beeline> !connect jdbc:hive2://
Connecting to jdbc:hive2://
Enter username for jdbc:hive2://: hive
Enter password for jdbc:hive2://: *********
 jdbc:hive2://> show tables;
```
