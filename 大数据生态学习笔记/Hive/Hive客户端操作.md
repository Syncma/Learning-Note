# Hive客户端操作

<!-- TOC -->

- [Hive客户端操作](#hive%e5%ae%a2%e6%88%b7%e7%ab%af%e6%93%8d%e4%bd%9c)
  - [操作步骤](#%e6%93%8d%e4%bd%9c%e6%ad%a5%e9%aa%a4)

<!-- /TOC -->

## 操作步骤
```
[jian@laptop tmp]$ beeline -u jdbc:hive2://
0: jdbc:hive2://> create table test
. . . . . . . . > (id int, name string,
. . . . . . . . > age int, tel string)
. . . . . . . . > ROW FORMAT DELIMITED
. . . . . . . . > FIELDS TERMINATED BY '\t'
. . . . . . . . > STORED AS TEXTFILE;


0: jdbc:hive2://> desc test;
19/12/02 15:03:32 [HiveServer2-Background-Pool: Thread-32]: WARN conf.HiveConf: HiveConf of name hive.metastore.local does not exist
OK
+-----------+------------+----------+
| col_name | data_type | comment |
+-----------+------------+----------+
| id | int | |
| name | string | |
| age | int | |
| tel | string | |
+-----------+------------+----------+
4 rows selected (0.322 seconds)

```

出现“HiveConf of name hive.metastore.local”，删除配置文件hive-site.xml里面关于hive.metastore.local


```
hive> insert into table test2 values('1', 'jacky');
WARNING: Hive-on-MR is deprecated in Hive 2 and may not be available in the future versions. Consider using a different execution engine (i.e. spark, tez) or using Hive 1.X releases.
Query ID = jian_20191202141555_885890ed-25b6-4719-a7b5-98ce07a8b320
Total jobs = 3
Launching Job 1 out of 3
Number of reduce tasks is set to 0 since there's no reduce operator
Starting Job = job_1575251487179_0001, Tracking URL = http://laptop:8088/proxy/application_1575251487179_0001/
Kill Command = /home/jian/prj/bigdata/hadoop-2.10.0/bin/hadoop job -kill job_1575251487179_0001
Hadoop job information for Stage-1: number of mappers: 1; number of reducers: 0
2019-12-02 14:16:03,075 Stage-1 map = 0%, reduce = 0%
2019-12-02 14:16:07,225 Stage-1 map = 100%, reduce = 0%, Cumulative CPU 2.6 sec
MapReduce Total cumulative CPU time: 2 seconds 600 msec
Ended Job = job_1575251487179_0001
Stage-4 is selected by condition resolver.
Stage-3 is filtered out by condition resolver.
Stage-5 is filtered out by condition resolver.
Moving data to directory hdfs://localhost:9000/user/hive/warehouse/test2/.hive-staging_hive_2019-12-02_14-15-56_006_1521573675172182819-1/-ext-10000
Loading data to table default.test2
MapReduce Jobs Launched:
Stage-Stage-1: Map: 1 Cumulative CPU: 2.6 sec HDFS Read: 4144 HDFS Write: 77 SUCCESS
Total MapReduce CPU Time Spent: 2 seconds 600 msec
OK
Time taken: 12.89 seconds
```

我们看到使用传统的insert 语句进行数据插入， 非常耗时。
**`因为hive把sql->mapreduce`** 

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hive-mr.png)


这里推荐方法：
**使用load 命令**

```
local:可选，表示从本地文件系统中加载，而非hdfs
overwrite:可选，先删除原来数据，然后再加载

partition:这里是指将inpath中的所有数据加载到那个分区，并不会判断待加载的数据中每一条记录属于哪个分区。

有LOCAL表示从本地文件系统加载（文件会被拷贝到HDFS中）
无LOCAL表示从HDFS中加载数据（注意：文件直接被移动！！！而不是拷贝！！！ 并且。。文件名都不带改的。。）
OVERWRITE 表示是否覆盖表中数据（或指定分区的数据）（没有OVERWRITE 会直接APPEND，而不会滤重!）
```


1.准备测试数据：
```
###字段以tab间隔###

[jian@laptop tmp]$ cat sqldata.txt
1 tom 20 138111111
2 jacky 30 138111112
3 test 40 138111113
```

2.导入数据
```
hive> load data local inpath '/tmp/sqldata.txt' into table test2;
Loading data to table default.test2
OK
Time taken: 0.438 seconds


0: jdbc:hive2://> select * from test2;
OK
+-----------+-------------+
| test2.id | test2.name |
+-----------+-------------+
| NULL | NULL |
| NULL | NULL |
| NULL | NULL |
+-----------+-------------+
```


发现数据并没有写进去？？？？

原因查看数据的并没有以\t为间隔


数据修正完成后，再次执行：
```
hive> load data local inpath '/tmp/sqldata.txt' overwrite into table test;
Loading data to table default.test
OK
Time taken: 0.391 seconds
hive> select * from test;
OK
1 tom 20 138111111
2 jacky 30 138111112
3 test 40 138111113
Time taken: 0.153 seconds, Fetched: 3 row(s)
```
**`数据成功写入 并且消耗的时间很短`**

在vim中，默认情况下，没法区分空格和缩进，所以我们需要配置，使其能够区分

```
cat ~/.vimrc
" Show spaces and tabs; to turn off for copying, use `:set nolist`
set listchars=tab:→\ ,space:·,trail:·,nbsp:·
```

数据导出命令：
```
insert overwrite local directory '/home/hadoop/temp'  select * from test
```